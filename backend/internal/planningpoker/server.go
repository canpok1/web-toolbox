package planningpoker

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/redis"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/websocket"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var scaleListMap = map[api.ScaleType][]string{
	api.Fibonacci:  {"0", "1", "2", "3", "5", "8", "13", "21", "34", "55", "89", "?"},
	api.TShirt:     {"XS", "S", "M", "L", "XL", "XXL", "?"},
	api.PowerOfTwo: {"1", "2", "4", "8", "16", "32", "64", "128", "256", "512", "1024", "?"},
	api.Custom:     {},
}

var scaleOrder = []string{
	"0", "1", "2", "3", "4", "5", "8", "13", "16", "21",
	"32", "34", "55", "64", "89", "128", "256", "512", "1024",
	"XS", "S", "M", "L", "XL", "XXL", "?",
}

var validScaleTypeMap = map[api.ScaleType]struct{}{
	api.Fibonacci:  {},
	api.TShirt:     {},
	api.PowerOfTwo: {},
	api.Custom:     {},
}

type Server struct {
	redis redis.Client
	wsHub websocket.WebSocketHub
}

func NewServer(redisClient redis.Client, wsHub websocket.WebSocketHub) *Server {
	return &Server{redis: redisClient, wsHub: wsHub}
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req api.CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerSessions(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
	}

	res, err := s.HandlePostApiPlanningPokerSessions(&req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId string) error {
	var req api.JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetApiPlanningPokerRoundsRoundId(ctx echo.Context, roundId string, params api.GetApiPlanningPokerRoundsRoundIdParams) error {
	res, err := s.HandleGetApiPlanningPokerRoundsRoundId(context.Background(), roundId, params.ParticipantId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId string) error {
	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdReveal(context.Background(), roundId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId string) error {
	var req api.SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdVotes(context.Background(), roundId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId string) error {
	res, err := s.HandleGetApiPlanningPokerSessionsSessionId(sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId string) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdEnd(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId string) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdRounds(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) ValidatePostApiPlanningPokerSessions(body *api.CreateSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required")
	}
	if body.HostName == "" {
		return fmt.Errorf("hostName is required")
	}
	if body.ScaleType == "" {
		return fmt.Errorf("scaleType is required")
	}
	if _, exists := validScaleTypeMap[api.ScaleType(body.ScaleType)]; !exists {
		return fmt.Errorf("invalid scaleType: %s", body.ScaleType)
	}
	if body.ScaleType == api.Custom && len(*body.CustomScale) == 0 {
		return fmt.Errorf("customScale is required when scaleType is custom")
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessions(body *api.CreateSessionRequest) (*api.CreateSessionResponse, error) {
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
	res := api.CreateSessionResponse{
		HostId:    hostIdValue,
		SessionId: sessionIdValue,
	}
	return &res, nil
}

func (s *Server) ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionID string, body *api.JoinSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (sessionID: %s)", sessionID)
	}
	if body.Name == "" {
		return fmt.Errorf("name is required (sessionID: %s)", sessionID)
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionID string, body *api.JoinSessionRequest) (*api.JoinSessionResponse, error) {
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

	s.wsHub.BroadcastParticipantJoined(sessionID, participantId.String(), body.Name)

	res := api.JoinSessionResponse{
		ParticipantId: participantId.String(),
	}
	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerRoundsRoundId(ctx context.Context, roundId string, participantId *string) (*api.GetRoundResponse, error) {
	redisRound, err := s.redis.GetRound(ctx, roundId)
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundId, err)
	}
	if redisRound == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundId)
	}

	apiRound := api.Round{
		RoundId:   roundId,
		SessionId: redisRound.SessionId,
		Status:    api.RoundStatus(redisRound.Status),
		CreatedAt: redisRound.CreatedAt,
		UpdatedAt: redisRound.UpdatedAt,
		Votes:     []api.Vote{},
		// Summary は後で設定
	}

	voteIDs, err := s.redis.GetVotesInRound(ctx, roundId)
	if err != nil {
		// 投票リスト取得エラーはログに残すが、ラウンド情報自体は返す（投票情報なしで）
		log.Printf("failed to get votes in round, returning round data without votes: roundID=%s, err=%v", roundId, err)
		// エラーを返さずに処理を続行。apiRound.Votes は空のまま
	}

	numericVotes := []float32{} // 数値として扱える投票値を格納するスライス
	voteCountMap := map[string]api.VoteCount{}

	if len(voteIDs) > 0 {
		apiVotes := make([]api.Vote, 0, len(voteIDs))
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

			apiVote := api.Vote{
				ParticipantId:   redisVote.ParticipantId,
				ParticipantName: redisParticipant.Name,
			}

			// 公開時か自身のもののみ投票結果をセット
			isRevealed := apiRound.Status == api.Revealed
			isOwnVote := participantId != nil && *participantId == apiVote.ParticipantId

			if isRevealed || isOwnVote {
				if redisVote.Value != "" {
					valueCopy := redisVote.Value // ポインタ用にコピー
					apiVote.Value = &valueCopy

					// revealed 状態の場合、数値変換を試みて集計用スライスに追加
					if isRevealed {
						participant := api.SessionParticipant{
							ParticipantId: redisVote.ParticipantId,
							Name:          redisParticipant.Name,
						}
						if voteCount, exist := voteCountMap[redisVote.Value]; exist {
							voteCountMap[redisVote.Value] = api.VoteCount{
								Value:        voteCount.Value,
								Count:        voteCount.Count + 1,
								Participants: append(voteCount.Participants, participant),
							}
						} else {
							voteCountMap[redisVote.Value] = api.VoteCount{
								Value:        redisVote.Value,
								Count:        1,
								Participants: []api.SessionParticipant{participant},
							}
						}

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

	// revealed 状態の場合、サマリーを生成
	if apiRound.Status == api.Revealed {
		summary := api.RoundSummary{
			VoteCounts: []api.VoteCount{},
		}

		for _, scale := range scaleOrder {
			if voteCount, exist := voteCountMap[scale]; exist {
				summary.VoteCounts = append(summary.VoteCounts, voteCount)
			}
		}

		if len(numericVotes) > 0 {
			sort.Slice(numericVotes, func(i, j int) bool {
				return numericVotes[i] < numericVotes[j]
			})

			var sum float32
			for _, v := range numericVotes {
				sum += v
			}
			average := sum / float32(len(numericVotes))
			summary.Average = &average

			var median float32
			n := len(numericVotes)
			if n%2 == 1 {
				// 奇数個の場合、中央の要素
				median = numericVotes[n/2]
			} else {
				// 偶数個の場合、中央の2つの要素の平均
				median = (numericVotes[n/2-1] + numericVotes[n/2]) / 2.0
			}
			summary.Median = &median

			max := numericVotes[len(numericVotes)-1]
			summary.Max = &max

			min := numericVotes[0]
			summary.Min = &min
		}

		apiRound.Summary = &summary
	}

	res := api.GetRoundResponse{
		Round: apiRound,
	}

	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdReveal(ctx context.Context, roundId string) (*api.RevealRoundResponse, error) {
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

	s.wsHub.BroadcastVotesRevealed(round.SessionId, roundId)

	res := api.RevealRoundResponse{}
	return &res, nil
}

func (s *Server) ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId string, body *api.SendVoteRequest) error {
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

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdVotes(ctx context.Context, roundId string, body *api.SendVoteRequest) (*api.SendVoteResponse, error) {
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

	var res api.SendVoteResponse
	if voteId == nil {
		// Create a new vote
		newVoteId, err := uuid.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate vote uuid: roundID=%s, err=%v", roundId, err)
		}

		vote := redis.Vote{
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

		res = api.SendVoteResponse{VoteId: newVoteId.String()}
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
		res = api.SendVoteResponse{VoteId: *voteId}
	}

	s.wsHub.BroadcastVoteSubmitted(round.SessionId, body.ParticipantId)

	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerSessionsSessionId(sessionID string) (*api.GetSessionResponse, error) {
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

	participants := []api.SessionParticipant{}
	for _, participantID := range participantIDs {
		participant, err := s.redis.GetParticipant(ctx, participantID)
		if err != nil {
			return nil, fmt.Errorf("failed to get participant from redis: sessionID=%s, participantID=%s, err=%w", sessionID, participantID, err)
		}
		participants = append(participants, api.SessionParticipant{
			Name:          participant.Name,
			ParticipantId: participantID,
		})
	}

	var scales []string
	if session.ScaleType == "custom" {
		scales = session.CustomScale
	} else {
		scales = scaleListMap[api.ScaleType(session.ScaleType)]
	}

	// Convert the redis.Session to GetSessionResponse
	res := api.GetSessionResponse{
		Session: api.Session{
			SessionId:      sessionID,
			HostId:         session.HostId,
			ScaleType:      api.ScaleType(session.ScaleType),
			Status:         session.Status,
			Scales:         scales,
			CurrentRoundId: nil,
			Participants:   participants,
			CreatedAt:      session.CreatedAt,
			UpdatedAt:      session.UpdatedAt,
		},
	}
	if session.CurrentRoundId != "" {
		res.Session.CurrentRoundId = &session.CurrentRoundId
	}

	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdEnd(ctx context.Context, sessionID string) (*api.EndSessionResponse, error) {
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

	return &api.EndSessionResponse{}, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdRounds(ctx context.Context, sessionID string) (*api.StartRoundResponse, error) {
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
	roundIdValue := roundId.String()

	round := redis.Round{
		SessionId: sessionID,
		Status:    "voting",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the round to Redis
	if err := s.redis.CreateRound(ctx, roundIdValue, round); err != nil {
		return nil, fmt.Errorf("failed to create round in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	// Update the session's currentRoundId
	session.CurrentRoundId = roundIdValue
	session.Status = "inProgress"
	session.UpdatedAt = time.Now()
	if err := s.redis.UpdateSession(ctx, sessionID, *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	s.wsHub.BroadcastRoundStarted(sessionID, roundIdValue)

	res := api.StartRoundResponse{RoundId: roundIdValue}
	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerWsSessionId(ctx echo.Context, sessionID string) error {
	return s.wsHub.HandleWebSocket(ctx, sessionID)
}
