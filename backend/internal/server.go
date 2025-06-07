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
	return s.planningpokerServer.PostSessions(ctx)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostSessionsSessionIdParticipants(ctx, sessionId)
}

func (s *Server) GetApiPlanningPokerRoundsRoundId(ctx echo.Context, roundId string, params api.GetApiPlanningPokerRoundsRoundIdParams) error {
	return s.planningpokerServer.GetRoundsRoundId(ctx, roundId, params)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId string) error {
	return s.planningpokerServer.PostRoundsRoundIdReveal(ctx, roundId)
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId string) error {
	return s.planningpokerServer.PostRoundsRoundIdVotes(ctx, roundId)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.GetSessionsSessionId(ctx, sessionId)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostSessionsSessionIdEnd(ctx, sessionId)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId string) error {
	return s.planningpokerServer.PostSessionsSessionIdRounds(ctx, sessionId)
}

func (s *Server) GetApiPlanningPokerWsSessionId(ctx echo.Context, sessionID string) error {
	return s.planningpokerServer.HandleGetWsSessionId(ctx, sessionID)
}

// GetApiTalkRouletteThemes implements api.ServerInterface.
func (s *Server) GetApiTalkRouletteThemes(ctx echo.Context, params api.GetApiTalkRouletteThemesParams) error {
	// TODO: Implement
	return echo.ErrNotImplemented
}
