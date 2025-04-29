package internal

import (
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
	return s.planningpokerServer.PostApiPlanningPokerSessions(ctx)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostApiPlanningPokerSessionsSessionIdParticipants(ctx, sessionId)
}

func (s *Server) GetApiPlanningPokerRoundsRoundId(ctx echo.Context, roundId string, params api.GetApiPlanningPokerRoundsRoundIdParams) error {
	return s.planningpokerServer.GetApiPlanningPokerRoundsRoundId(ctx, roundId, params)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId string) error {
	return s.planningpokerServer.PostApiPlanningPokerRoundsRoundIdReveal(ctx, roundId)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId string) error {
	return s.planningpokerServer.PostApiPlanningPokerRoundsRoundIdVotes(ctx, roundId)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.GetApiPlanningPokerSessionsSessionId(ctx, sessionId)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostApiPlanningPokerSessionsSessionIdEnd(ctx, sessionId)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostApiPlanningPokerSessionsSessionIdRounds(ctx, sessionId)
}

func (s *Server) GetApiPlanningPokerWsSessionId(ctx echo.Context, sessionID string) error {
	return s.planningpokerServer.HandleGetApiPlanningPokerWsSessionId(ctx, sessionID)
}
