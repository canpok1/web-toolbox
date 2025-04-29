package planningpoker

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain/model"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/infra"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var validScaleTypeMap = map[api.ScaleType]struct{}{
	api.Fibonacci:  {},
	api.TShirt:     {},
	api.PowerOfTwo: {},
	api.Custom:     {},
}

type Server struct {
	redis infra.RedisClient
	wsHub infra.WebSocketHub
}

func NewServer(redisClient infra.RedisClient, wsHub infra.WebSocketHub) *Server {
	return &Server{redis: redisClient, wsHub: wsHub}
}

func (s *Server) PostSessions(ctx echo.Context) error {
	var req api.CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostSessions(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
	}

	res, err := s.HandlePostSessions(&req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostSessionsSessionIdParticipants(ctx echo.Context, sessionId string) error {
	var req api.JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostSessionsSessionIdParticipants(sessionId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostSessionsSessionIdParticipants(sessionId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetRoundsRoundId(ctx echo.Context, roundId string, params api.GetApiPlanningPokerRoundsRoundIdParams) error {
	res, err := s.HandleGetRoundsRoundId(context.Background(), roundId, params.ParticipantId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostRoundsRoundIdReveal(ctx echo.Context, roundId string) error {
	res, err := s.HandlePostRoundsRoundIdReveal(context.Background(), roundId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostRoundsRoundIdVotes(ctx echo.Context, roundId string) error {
	var req api.SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostRoundsRoundIdVotes(roundId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostRoundsRoundIdVotes(context.Background(), roundId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetSessionsSessionId(ctx echo.Context, sessionId string) error {
	res, err := s.HandleGetSessionsSessionId(sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostSessionsSessionIdEnd(ctx echo.Context, sessionId string) error {
	res, err := s.HandlePostSessionsSessionIdEnd(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostSessionsSessionIdRounds(ctx echo.Context, sessionId string) error {
	res, err := s.HandlePostSessionsSessionIdRounds(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) ValidatePostSessions(body *api.CreateSessionRequest) error {
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

func (s *Server) HandlePostSessions(body *api.CreateSessionRequest) (*api.CreateSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	usecase := domain.NewCreateSessionUsecase(s.redis, s.redis)
	result, err := usecase.Create(string(body.ScaleType), body.CustomScale, body.HostName)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	res := api.CreateSessionResponse{
		HostId:    result.HostId,
		SessionId: result.SessionId,
	}
	return &res, nil
}

func (s *Server) ValidatePostSessionsSessionIdParticipants(sessionID string, body *api.JoinSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required (sessionID: %s)", sessionID)
	}
	if body.Name == "" {
		return fmt.Errorf("name is required (sessionID: %s)", sessionID)
	}

	return nil
}

func (s *Server) HandlePostSessionsSessionIdParticipants(sessionID string, body *api.JoinSessionRequest) (*api.JoinSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required (sessionID: %s)", sessionID)
	}

	usecase := domain.NewCreateParticipantUsecase(s.redis, s.redis)
	result, err := usecase.Create(sessionID, body.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create participant: %w", err)
	}

	s.wsHub.BroadcastParticipantJoined(sessionID, result.ParticipantId, body.Name)

	res := api.JoinSessionResponse{
		ParticipantId: result.ParticipantId,
	}
	return &res, nil
}

func (s *Server) HandleGetRoundsRoundId(ctx context.Context, roundId string, participantId *string) (*api.GetRoundResponse, error) {
	usecase := domain.NewReadRoundUsecase(s.redis, s.redis, s.redis)
	result, err := usecase.Read(ctx, roundId, participantId)
	if err != nil {
		return nil, fmt.Errorf("failed to read round: %w", err)
	}

	res := api.GetRoundResponse{
		Round: result.Round,
	}
	return &res, nil
}

func (s *Server) HandlePostRoundsRoundIdReveal(ctx context.Context, roundId string) (*api.RevealRoundResponse, error) {
	usecase := domain.NewUpdateRoundStatusUsecase(s.redis)
	result, err := usecase.Update(ctx, roundId, "revealed")
	if err != nil {
		return nil, fmt.Errorf("failed to update round status: %w", err)
	}

	s.wsHub.BroadcastVotesRevealed(result.Round.SessionId, roundId)

	res := api.RevealRoundResponse{}
	return &res, nil
}

func (s *Server) ValidatePostRoundsRoundIdVotes(roundId string, body *api.SendVoteRequest) error {
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

func (s *Server) HandlePostRoundsRoundIdVotes(ctx context.Context, roundId string, body *api.SendVoteRequest) (*api.SendVoteResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required (roundID: %s)", roundId)
	}

	usecase := domain.NewUpsertVoteUsecase(s.redis, s.redis, s.redis)
	result, err := usecase.Upsert(ctx, roundId, body.ParticipantId, body.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert vote: %w", err)
	}

	s.wsHub.BroadcastVoteSubmitted(result.Round.SessionId, body.ParticipantId)

	res := api.SendVoteResponse{VoteId: result.VoteID}
	return &res, nil
}

func (s *Server) HandleGetSessionsSessionId(sessionID string) (*api.GetSessionResponse, error) {
	ctx := context.Background()

	usecase := domain.NewReadSessionUsecase(s.redis, s.redis)
	result, err := usecase.Read(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to read session: %w", err)
	}

	var participants []api.SessionParticipant
	for _, participant := range result.Participants {
		participants = append(participants, api.SessionParticipant{
			ParticipantId: participant.ParticipantId,
			Name:          participant.Name,
		})
	}

	res := api.GetSessionResponse{
		Session: api.Session{
			SessionId:      sessionID,
			HostId:         result.Session.HostId,
			ScaleType:      api.ScaleType(result.Session.ScaleType),
			Status:         result.Session.Status,
			Scales:         result.Scales,
			CurrentRoundId: nil,
			Participants:   participants,
			CreatedAt:      result.Session.CreatedAt,
			UpdatedAt:      result.Session.UpdatedAt,
		},
	}
	if result.Session.CurrentRoundId != "" {
		res.Session.CurrentRoundId = &result.Session.CurrentRoundId
	}

	return &res, nil
}

func (s *Server) HandlePostSessionsSessionIdEnd(ctx context.Context, sessionID string) (*api.EndSessionResponse, error) {
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

func (s *Server) HandlePostSessionsSessionIdRounds(ctx context.Context, sessionID string) (*api.StartRoundResponse, error) {
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

	round := model.Round{
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

func (s *Server) HandleGetWsSessionId(ctx echo.Context, sessionID string) error {
	return s.wsHub.HandleWebSocket(ctx, sessionID)
}
