package api

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"
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

func (s *Server) ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionID string, body *JoinSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (sessionID: %s)", sessionID)
	}
	if body.Name == "" {
		return fmt.Errorf("name is required (sessionID: %s)", sessionID)
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionID string, body *JoinSessionRequest) (*JoinSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required (sessionID: %s)", sessionID)
	}

	participantId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate participant uuid (sessionID: %s): %v", sessionID, err)
	}

	ctx := context.Background()

	// セッションの存在チェック
	session, err := s.redis.GetSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis (sessionID: %s): %v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session is not found (sessionID: %s)", sessionID)
	}

	// 参加者登録
	participant := redis.Participant{
		SessionId: sessionID,
		Name:      body.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.redis.CreateParticipant(ctx, participantId.String(), participant); err != nil {
		return nil, fmt.Errorf("failed to create participant (sessionID: %s): %v", sessionID, err)
	}
	if err := s.redis.AddParticipantToSession(ctx, sessionID, participantId.String()); err != nil {
		return nil, fmt.Errorf("failed to add participant to session (sessionID: %s, participantID: %s): %v", sessionID, participantId.String(), err)
	}

	res := JoinSessionResponse{
		ParticipantId: participantId,
	}
	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerRoundsRoundId(ctx context.Context, roundId string, participantId *string) (*GetRoundResponse, error) {
	redisRound, err := s.redis.GetRound(ctx, roundId)
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
		RoundId:   uuid.MustParse(roundId),
		SessionId: sessionUUID,
		Status:    RoundStatus(redisRound.Status),
		CreatedAt: redisRound.CreatedAt,
		UpdatedAt: redisRound.UpdatedAt,
		Votes:     []Vote{},
		// Summary は後で設定
	}

	voteIDs, err := s.redis.GetVotesInRound(ctx, roundId)
	if err != nil {
		// 投票リスト取得エラーはログに残すが、ラウンド情報自体は返す（投票情報なしで）
		log.Printf("failed to get votes in round, returning round data without votes: roundID=%s, err=%v", roundId, err)
		// エラーを返さずに処理を続行。apiRound.Votes は空のまま
	}

	numericVotes := []float32{} // 数値として扱える投票値を格納するスライス

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
			// GetParticipant でエラーが発生した場合や見つからなかった場合は、その投票を除外する
			if err != nil {
				log.Printf("failed to get participant details for vote, skipping vote: roundID=%s, voteID=%s, participantID=%s, err=%v", roundId, voteID, redisVote.ParticipantId, err)
				continue
			}
			if redisParticipant == nil {
				log.Printf("participant not found for vote, skipping vote: roundID=%s, voteID=%s, participantID=%s", roundId, voteID, redisVote.ParticipantId)
				continue
			}

			apiVote := Vote{
				ParticipantId:   participantUUID,
				ParticipantName: redisParticipant.Name,
			}

			// 公開時か自身のもののみ投票結果をセット
			isRevealed := apiRound.Status == Revealed
			isOwnVote := participantId != nil && *participantId == participantUUID.String()

			if isRevealed || isOwnVote {
				if redisVote.Value != "" {
					valueCopy := redisVote.Value // ポインタ用にコピー
					apiVote.Value = &valueCopy

					// revealed 状態の場合、数値変換を試みて集計用スライスに追加
					if isRevealed {
						numVal, err := strconv.ParseFloat(redisVote.Value, 32)
						if err == nil {
							numericVotes = append(numericVotes, float32(numVal))
						} else {
							// 数値に変換できない値はログに残しても良い（例: '?', 'coffee'）
							log.Printf("Vote value is not numeric, skipping for summary calculation: roundID=%s, voteID=%s, value=%s", roundId, voteID, redisVote.Value)
						}
					}
				}
			}
			apiVotes = append(apiVotes, apiVote)
		}
		apiRound.Votes = apiVotes
	}

	// revealed 状態かつ数値の投票が1つ以上ある場合、サマリーを生成
	if apiRound.Status == Revealed && len(numericVotes) > 0 {
		sort.Slice(numericVotes, func(i, j int) bool {
			return numericVotes[i] < numericVotes[j]
		})

		var sum float32
		for _, v := range numericVotes {
			sum += v
		}
		average := sum / float32(len(numericVotes))

		var median float32
		n := len(numericVotes)
		if n%2 == 1 {
			// 奇数個の場合、中央の要素
			median = numericVotes[n/2]
		} else {
			// 偶数個の場合、中央の2つの要素の平均
			median = (numericVotes[n/2-1] + numericVotes[n/2]) / 2.0
		}

		max := numericVotes[len(numericVotes)-1]
		min := numericVotes[0]

		apiRound.Summary = &RoundSummary{
			Average: average,
			Median:  median,
			Max:     max,
			Min:     min,
		}
	}

	res := GetRoundResponse{
		Round: apiRound,
	}

	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdReveal(ctx context.Context, roundId string) (*RevealRoundResponse, error) {
	// Retrieve the round from Redis
	round, err := s.redis.GetRound(ctx, roundId)
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundId, err)
	}
	if round == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundId)
	}

	// Update the round status to "revealed"
	round.Status = "revealed"
	round.UpdatedAt = time.Now()
	if err := s.redis.UpdateRound(ctx, roundId, *round); err != nil {
		return nil, fmt.Errorf("failed to update round in redis: roundID=%s, err=%v", roundId, err)
	}

	res := RevealRoundResponse{}

	return &res, nil
}

func (s *Server) ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId string, body *SendVoteRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (roundID: %s)", roundId)
	}
	if body.ParticipantId == "" {
		return fmt.Errorf("participantId is required (roundID: %s)", roundId)
	}
	if body.Value == "" {
		return fmt.Errorf("value is required (roundID: %s)", roundId)
	}
	return nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdVotes(ctx context.Context, roundId string, body *SendVoteRequest) (*SendVoteResponse, error) {
	// Validate request body
	if body == nil {
		return nil, fmt.Errorf("request body is required (roundID: %s)", roundId)
	}

	// Retrieve the round from Redis
	round, err := s.redis.GetRound(ctx, roundId)
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
	participant, err := s.redis.GetParticipant(ctx, body.ParticipantId)
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
	voteId, err := s.redis.GetVoteIdByRoundIdAndParticipantId(ctx, roundId, body.ParticipantId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vote id from redis: roundID=%s, participantID=%s, err=%v", roundId, body.ParticipantId, err)
	}

	var vote redis.Vote
	if voteId == nil {
		// Create a new vote
		newVoteId, err := uuid.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate vote uuid: roundID=%s, err=%v", roundId, err)
		}

		vote = redis.Vote{
			RoundId:       roundId,
			ParticipantId: body.ParticipantId,
			Value:         body.Value,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// Save the vote to Redis
		if err := s.redis.CreateVote(ctx, newVoteId.String(), vote); err != nil {
			return nil, fmt.Errorf("failed to create vote in redis: roundID=%s, voteID=%s, err=%v", roundId, newVoteId.String(), err)
		}

		// Add the vote to the round's vote list
		if err := s.redis.AddVoteToRound(ctx, roundId, newVoteId.String()); err != nil {
			return nil, fmt.Errorf("failed to add vote to round in redis: roundID=%s, voteID=%s, err=%v", roundId, newVoteId.String(), err)
		}

		res := SendVoteResponse{VoteId: newVoteId}
		return &res, nil
	} else {
		// Update the existing vote
		vote, err := s.redis.GetVote(ctx, *voteId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vote from redis: roundID=%s, voteID=%s, err=%v", roundId, *voteId, err)
		}
		vote.Value = body.Value
		vote.UpdatedAt = time.Now()

		if err := s.redis.UpdateVote(ctx, *voteId, *vote); err != nil {
			return nil, fmt.Errorf("failed to update vote in redis: roundID=%s, voteID=%s, err=%v", roundId, *voteId, err)
		}
		res := SendVoteResponse{VoteId: uuid.MustParse(*voteId)}
		return &res, nil
	}
}

func (s *Server) HandleGetApiPlanningPokerSessionsSessionId(sessionID string) (*GetSessionResponse, error) {
	ctx := context.Background()

	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis (sessionID: %s): %v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found (sessionID: %s)", sessionID)
	}

	participantIDs, err := s.redis.GetParticipantsInSession(ctx, sessionID)
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
			SessionId:      uuid.MustParse(sessionID),
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

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdEnd(ctx context.Context, sessionID string) (*EndSessionResponse, error) {
	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID)
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
	if err := s.redis.UpdateSession(ctx, sessionID, *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, err=%v", sessionID, err)
	}

	return &EndSessionResponse{}, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdRounds(ctx context.Context, sessionID string) (*StartRoundResponse, error) {
	// Retrieve the session from Redis
	session, err := s.redis.GetSession(ctx, sessionID)
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
		SessionId: sessionID,
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
	if err := s.redis.UpdateSession(ctx, sessionID, *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	res := StartRoundResponse{RoundId: roundId}
	return &res, nil
}
