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

func NewServer(redisClient redis.Client) *Server {
	return &Server{redis: redisClient}
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerSessions(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId uuid.UUID) error {
	var req JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdReveal(roundId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId uuid.UUID) error {
	var req SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdVotes(roundId, &req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandleGetApiPlanningPokerSessionsSessionId(sessionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdEnd(sessionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdRounds(sessionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}
