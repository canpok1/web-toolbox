package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/google/uuid"
)

var validScaleTypeMap = map[ScaleType]struct{}{
	Fibonacci:  {},
	TShirt:     {},
	PowerOfTwo: {},
	Custom:     {},
}

func (s *Server) ValidatePostApiPlanningPokerSessions(body *CreateSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required")
	}
	if body.HostName == "" {
		return fmt.Errorf("hostName is required")
	}
	if body.ScaleType == "" {
		return fmt.Errorf("scaleType is required")
	}
	if _, exists := validScaleTypeMap[ScaleType(body.ScaleType)]; !exists {
		return fmt.Errorf("invalid scaleType: %s", body.ScaleType)
	}
	if body.ScaleType == Custom && len(*body.CustomScale) == 0 {
		return fmt.Errorf("customScale is required when scaleType is custom")
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessions(body *CreateSessionRequest) (*CreateSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	hostId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate host uuid: %v", err)
	}
	hostIdValue := hostId.String()

	sessionId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session uuid: %v", err)
	}
	sessionIdValue := sessionId.String()

	// セッション情報の保存
	session := redis.Session{
		HostId:      hostIdValue,
		ScaleType:   string(body.ScaleType),
		CustomScale: []string{},
		Status:      "waiting",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if body.CustomScale != nil {
		session.CustomScale = *body.CustomScale
	}

	ctx := context.Background()
	if err = s.redis.CreateSession(ctx, sessionIdValue, session); err != nil {
		return nil, fmt.Errorf("failed to save session to redis: %v", err)
	}

	err = s.redis.CreateParticipant(ctx, hostIdValue, redis.Participant{
		SessionId: sessionIdValue,
		Name:      body.HostName,
		IsHost:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save participant to redis: %v", err)
	}

	err = s.redis.AddParticipantToSession(ctx, sessionIdValue, hostIdValue)
	if err != nil {
		return nil, fmt.Errorf("failed to add participant list to redis: %v", err)
	}

	// レスポンスの作成
	res := CreateSessionResponse{
		HostId:    hostId,
		SessionId: sessionId,
	}
	return &res, nil
}

func (s *Server) ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionID uuid.UUID, body *JoinSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (sessionID: %s)", sessionID.String())
	}
	if body.Name == "" {
		return fmt.Errorf("name is required (sessionID: %s)", sessionID.String())
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionID uuid.UUID, body *JoinSessionRequest) (*JoinSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required (sessionID: %s)", sessionID.String())
	}

	participantId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate participant uuid (sessionID: %s): %v", sessionID.String(), err)
	}

	ctx := context.Background()

	// セッションの存在チェック
	session, err := s.redis.GetSession(ctx, sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis (sessionID: %s): %v", sessionID.String(), err)
	}
	if session == nil {
		return nil, fmt.Errorf("session is not found (sessionID: %s)", sessionID.String())
	}

	// 参加者登録
	participant := redis.Participant{
		SessionId: sessionID.String(),
		Name:      body.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.redis.CreateParticipant(ctx, participantId.String(), participant); err != nil {
		return nil, fmt.Errorf("failed to create participant (sessionID: %s): %v", sessionID.String(), err)
	}
	if err := s.redis.AddParticipantToSession(ctx, sessionID.String(), participantId.String()); err != nil {
		return nil, fmt.Errorf("failed to add participant to session (sessionID: %s, participantID: %s): %v", sessionID.String(), participantId.String(), err)
	}

	res := JoinSessionResponse{
		ParticipantId: participantId,
	}
	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerRoundsRoundId(ctx context.Context, roundId uuid.UUID, participantId *uuid.UUID) (*GetRoundResponse, error) {
	redisRound, err := s.redis.GetRound(ctx, roundId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundId, err)
	}
	if redisRound == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundId)
	}

	sessionUUID, err := uuid.Parse(redisRound.SessionId)
	if err != nil {
		log.Printf("invalid session ID format found in round data: roundID=%s, sessionID=%s, err=%v", roundId, redisRound.SessionId, err)
		return nil, fmt.Errorf("internal server error: invalid session ID format")
	}

	apiRound := Round{
		RoundId:   roundId,
		SessionId: sessionUUID,
		Status:    RoundStatus(redisRound.Status),
		CreatedAt: redisRound.CreatedAt,
		UpdatedAt: redisRound.UpdatedAt,
		Votes:     []Vote{},
	}

	voteIDs, err := s.redis.GetVotesInRound(ctx, roundId.String())
	if err != nil {
		// 投票リスト取得エラーはログに残すが、ラウンド情報自体は返す（投票情報なしで）
		log.Printf("failed to get votes in round, returning round data without votes: roundID=%s, err=%v", roundId, err)
		// エラーを返さずに処理を続行。apiRound.Votes は空のまま
	}

	if len(voteIDs) > 0 {
		apiVotes := make([]Vote, 0, len(voteIDs))
		for _, voteID := range voteIDs {
			redisVote, err := s.redis.GetVote(ctx, voteID)
			if err != nil {
				log.Printf("failed to get vote details, skipping vote: roundID=%s, voteID=%s, err=%v", roundId, voteID, err)
				continue // 個別の投票取得エラーはスキップ
			}
			if redisVote == nil {
				log.Printf("vote not found, but ID was listed in round, skipping vote: roundID=%s, voteID=%s", roundId, voteID)
				continue // データ不整合の可能性、スキップ
			}

			participantUUID, err := uuid.Parse(redisVote.ParticipantId)
			if err != nil {
				log.Printf("failed to parse participantId for vote, skipping vote: roundID=%s, voteID=%s, participantID=%s, err=%v", roundId, voteID, redisVote.ParticipantId, err)
				continue // ParticipantId の形式不正、スキップ
			}

			redisParticipant, err := s.redis.GetParticipant(ctx, redisVote.ParticipantId)
			if err != nil {
				log.Printf("failed to get participant details for vote, skipping participant name: roundID=%s, voteID=%s, participantID=%s, err=%v", roundId, voteID, redisVote.ParticipantId, err)
				continue // 参加者情報が取れない場合は投票情報を含めない
			}
			if redisParticipant == nil {
				log.Printf("participant not found for vote, skipping participant name: roundID=%s, voteID=%s, participantID=%s", roundId, voteID, redisVote.ParticipantId)
				continue // 参加者情報が取れない場合は投票情報を含めない
			}

			apiVote := Vote{
				ParticipantId:   participantUUID,
				ParticipantName: redisParticipant.Name,
			}

			// 公開時か自身のもののみ投票結果をセット
			if apiRound.Status == Revealed || (participantId != nil && *participantId == participantUUID) {
				if redisVote.Value != "" {
					valueCopy := redisVote.Value // ポインタ用にコピー
					apiVote.Value = &valueCopy
				}
			}
			apiVotes = append(apiVotes, apiVote)
		}
		if len(apiVotes) > 0 {
			apiRound.Votes = apiVotes
		}
	}

	res := GetRoundResponse{
		Round: apiRound,
	}

	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdReveal(ctx context.Context, roundId uuid.UUID) (*RevealRoundResponse, error) {
	// Retrieve the round from Redis
	round, err := s.redis.GetRound(ctx, roundId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundId, err)
	}
	if round == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundId)
	}

	// Update the round status to "revealed"
	round.Status = "revealed"
	round.UpdatedAt = time.Now()
	if err := s.redis.UpdateRound(ctx, roundId.String(), *round); err != nil {
		return nil, fmt.Errorf("failed to update round in redis: roundID=%s, err=%v", roundId, err)
	}

	res := RevealRoundResponse{}

	return &res, nil
}

func (s *Server) ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId uuid.UUID, body *SendVoteRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (roundID: %s)", roundId)
	}
	if body.ParticipantId == uuid.Nil {
		return fmt.Errorf("participantId is required (roundID: %s)", roundId)
	}
	if body.Value == "" {
		return fmt.Errorf("value is required (roundID: %s)", roundId)
	}
	return nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdVotes(ctx context.Context, roundId uuid.UUID, body *SendVoteRequest) (*SendVoteResponse, error) {
	// Validate request body
	if body == nil {
		return nil, fmt.Errorf("request body is required (roundID: %s)", roundId)
	}

	// Retrieve the round from Redis
	round, err := s.redis.GetRound(ctx, roundId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundId, err)
	}
	if round == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundId)
	}

	// Check if the round is in the "voting" state
	if round.Status != "voting" {
		return nil, fmt.Errorf("round is not in voting state: roundID=%s", roundId)
	}

	// Check if the participant exists
	participant, err := s.redis.GetParticipant(ctx, body.ParticipantId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get participant from redis: roundID=%s, participantID=%s, err=%v", roundId, body.ParticipantId, err)
	}
	if participant == nil {
		return nil, fmt.Errorf(
			"participant not found: roundID=%s, participantID=%s",
			roundId,
			body.ParticipantId,
		)
	}

	// Check if the participant has already voted in this round
	voteId, err := s.redis.GetVoteIdByRoundIdAndParticipantId(ctx, roundId.String(), body.ParticipantId.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get vote id from redis: roundID=%s, participantID=%s, err=%v", roundId.String(), body.ParticipantId.String(), err)
	}

	var vote redis.Vote
	if voteId == nil {
		// Create a new vote
		newVoteId, err := uuid.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate vote uuid: roundID=%s, err=%v", roundId.String(), err)
		}

		vote = redis.Vote{
			RoundId:       roundId.String(),
			ParticipantId: body.ParticipantId.String(),
			Value:         body.Value,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// Save the vote to Redis
		if err := s.redis.CreateVote(ctx, newVoteId.String(), vote); err != nil {
			return nil, fmt.Errorf("failed to create vote in redis: roundID=%s, voteID=%s, err=%v", roundId.String(), newVoteId.String(), err)
		}

		// Add the vote to the round's vote list
		if err := s.redis.AddVoteToRound(ctx, roundId.String(), newVoteId.String()); err != nil {
			return nil, fmt.Errorf("failed to add vote to round in redis: roundID=%s, voteID=%s, err=%v", roundId.String(), newVoteId.String(), err)
		}

		res := SendVoteResponse{VoteId: newVoteId}
		return &res, nil
	} else {
		// Update the existing vote
		vote, err := s.redis.GetVote(ctx, *voteId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vote from redis: roundID=%s, voteID=%s, err=%v", roundId.String(), *voteId, err)
		}
		vote.Value = body.Value
		vote.UpdatedAt = time.Now()

		if err := s.redis.UpdateVote(ctx, *voteId, *vote); err != nil {
			return nil, fmt.Errorf("failed to update vote in redis: roundID=%s, voteID=%s, err=%v", roundId.String(), *voteId, err)
		}
		res := SendVoteResponse{VoteId: uuid.MustParse(*voteId)}
		return &res, nil
	}
}

func (s *Server) HandleGetApiPlanningPokerSessionsSessionId(sessionID uuid.UUID) (*GetSessionResponse, error) {
	ctx := context.Background()

	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis (sessionID: %s): %v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found (sessionID: %s)", sessionID)
	}

	participantIDs, err := s.redis.GetParticipantsInSession(ctx, sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get participants in session (sessionID: %s)", sessionID)
	}

	participants := []SessionParticipant{}
	for _, participantID := range participantIDs {
		participant, err := s.redis.GetParticipant(ctx, participantID)
		if err != nil {
			return nil, fmt.Errorf("failed to get participant from redis: sessionID=%s, participantID=%s, err=%w", sessionID, participantID, err)
		}
		participants = append(participants, SessionParticipant{
			Name:          participant.Name,
			ParticipantId: participantID,
		})
	}

	// Convert the redis.Session to GetSessionResponse
	res := GetSessionResponse{
		Session: Session{
			SessionId:      sessionID,
			HostId:         uuid.MustParse(session.HostId),
			ScaleType:      ScaleType(session.ScaleType),
			Status:         session.Status,
			CustomScale:    session.CustomScale,
			CurrentRoundId: nil,
			Participants:   participants,
			CreatedAt:      session.CreatedAt,
			UpdatedAt:      session.UpdatedAt,
		},
	}
	if session.CurrentRoundId != "" {
		currendRoundId, err := uuid.Parse(session.CurrentRoundId)
		if err != nil {
			log.Printf("failed to parse CurrentRoundId: %v", err)
			return nil, fmt.Errorf("failed to parse CurrentRoundId: %w", err)
		}
		res.Session.CurrentRoundId = &currendRoundId
	}

	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdEnd(ctx context.Context, sessionID uuid.UUID) (*EndSessionResponse, error) {
	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis: sessionID=%s, err=%v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found: sessionID=%s", sessionID)
	}

	// Update the session status to "finished"
	session.Status = "finished"
	session.UpdatedAt = time.Now()

	// Save the updated session back to Redis
	if err := s.redis.UpdateSession(ctx, sessionID.String(), *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, err=%v", sessionID, err)
	}

	return &EndSessionResponse{}, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdRounds(ctx context.Context, sessionID uuid.UUID) (*StartRoundResponse, error) {
	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis: sessionID=%s, err=%v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found: sessionID=%s", sessionID)
	}

	// Create a new round
	roundId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate round uuid: sessionID=%s, err=%v", sessionID, err)
	}

	round := redis.Round{
		SessionId: sessionID.String(),
		Status:    "voting",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the round to Redis
	if err := s.redis.CreateRound(ctx, roundId.String(), round); err != nil {
		return nil, fmt.Errorf("failed to create round in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	// Update the session's currentRoundId
	session.CurrentRoundId = roundId.String()
	session.Status = "inProgress"
	session.UpdatedAt = time.Now()
	if err := s.redis.UpdateSession(ctx, sessionID.String(), *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	res := StartRoundResponse{RoundId: roundId}
	return &res, nil
}
