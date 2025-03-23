package api

import (
	"fmt"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Server struct {
	redis redis.Client
}

func NewServer(redis redis.Client) *Server {
	return &Server{redis: redis}
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId uuid.UUID) error {
	// TODO: 実装をここに記述
	return ctx.JSON(http.StatusNotImplemented, ErrorResponse{Message: fmt.Sprintf("PostApiPlanningPokerRoundsRoundIdReveal: %s is not implemented yet", roundId)})
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId uuid.UUID) error {
	var req SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	voteId, err := uuid.NewUUID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("failed to generate vote uuid: %v", err)})
	}

	res := SendVoteResponse{
		VoteId: &voteId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	hostId, err := uuid.NewUUID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("failed to generate host uuid: %v", err)})
	}

	sessionId, err := uuid.NewUUID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("failed to generate session uuid: %v", err)})
	}

	res := CreateSessionResponse{
		HostId:    &hostId,
		SessionId: &sessionId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId uuid.UUID) error {
	// TODO: 実装をここに記述
	return ctx.JSON(http.StatusNotImplemented, ErrorResponse{Message: fmt.Sprintf("GetApiPlanningPokerSessionsSessionId: %s is not implemented yet", sessionId)})
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId uuid.UUID) error {
	// TODO: 実装をここに記述
	return ctx.JSON(http.StatusNotImplemented, ErrorResponse{Message: fmt.Sprintf("PostApiPlanningPokerSessionsSessionIdEnd: %s is not implemented yet", sessionId)})
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId uuid.UUID) error {
	var req JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	participantId, err := uuid.NewUUID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("failed to generate participant uuid: %v", err)})
	}

	res := JoinSessionResponse{
		ParticipantId: &participantId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId uuid.UUID) error {
	roundId, err := uuid.NewUUID()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: fmt.Sprintf("failed to generate round uuid: %v", err)})
	}
	res := StartRoundResponse{RoundId: &roundId}
	return ctx.JSON(http.StatusCreated, res)
}
