package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/api/planningpoker"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime/types"
)

type Server struct {
	redis redis.Client
	wsHub planningpoker.WebSocketHub
}

var _ ServerInterface = &Server{}

func NewServer(redisClient redis.Client, wsHub planningpoker.WebSocketHub) *Server {
	return &Server{redis: redisClient, wsHub: wsHub}
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerSessions(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	res, err := s.HandlePostApiPlanningPokerSessions(&req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId uuid.UUID) error {
	var req JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetApiPlanningPokerRoundsRoundId(ctx echo.Context, roundId types.UUID) error {
	res, err := s.HandleGetApiPlanningPokerRoundsRoundId(context.Background(), roundId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdReveal(context.Background(), roundId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId uuid.UUID) error {
	var req SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.HandlePostApiPlanningPokerRoundsRoundIdVotes(context.Background(), roundId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandleGetApiPlanningPokerSessionsSessionId(sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdEnd(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId uuid.UUID) error {
	res, err := s.HandlePostApiPlanningPokerSessionsSessionIdRounds(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerWs(ctx echo.Context) error {
	return s.wsHub.HandleWebSocket(ctx)
}
