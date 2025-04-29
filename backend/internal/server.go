package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker"
	"github.com/labstack/echo/v4"
)

type Server struct {
	planningpokerServer *planningpoker.Server
}

var _ api.ServerInterface = &Server{}

func NewServer(planningpokerServer *planningpoker.Server) *Server {
	return &Server{planningpokerServer: planningpokerServer}
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req api.CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to bind request body: %v", err)})
	}
	if err := s.planningpokerServer.ValidatePostApiPlanningPokerSessions(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: err.Error()})
	}

	res, err := s.planningpokerServer.HandlePostApiPlanningPokerSessions(&req)
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
	if err := s.planningpokerServer.ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.planningpokerServer.HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetApiPlanningPokerRoundsRoundId(ctx echo.Context, roundId string, params api.GetApiPlanningPokerRoundsRoundIdParams) error {
	res, err := s.planningpokerServer.HandleGetApiPlanningPokerRoundsRoundId(context.Background(), roundId, params.ParticipantId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId string) error {
	res, err := s.planningpokerServer.HandlePostApiPlanningPokerRoundsRoundIdReveal(context.Background(), roundId)
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
	if err := s.planningpokerServer.ValidatePostApiPlanningPokerRoundsRoundIdVotes(roundId, &req); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.ErrorResponse{Message: fmt.Sprintf("failed to validate request: %v", err)})
	}

	res, err := s.planningpokerServer.HandlePostApiPlanningPokerRoundsRoundIdVotes(context.Background(), roundId, &req)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId string) error {
	res, err := s.planningpokerServer.HandleGetApiPlanningPokerSessionsSessionId(sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId string) error {
	res, err := s.planningpokerServer.HandlePostApiPlanningPokerSessionsSessionIdEnd(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId string) error {
	res, err := s.planningpokerServer.HandlePostApiPlanningPokerSessionsSessionIdRounds(context.Background(), sessionId)
	if err != nil {
		log.Printf("failed to handle request: %v", err)
		return ctx.JSON(http.StatusInternalServerError, api.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, res)
}

func (s *Server) GetApiPlanningPokerWsSessionId(ctx echo.Context, sessionID string) error {
	return s.planningpokerServer.HandleGetApiPlanningPokerWsSessionId(ctx, sessionID)
}
