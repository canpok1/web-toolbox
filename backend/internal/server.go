package internal

import (
	"net/http" // Required for http.StatusOK

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker"
	"github.com/canpok1/web-toolbox/backend/internal/talkroulette" // New import
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

// GetApiTalkRouletteThemes は api.ServerInterface を実装します。
func (s *Server) GetApiTalkRouletteThemes(ctx echo.Context, params api.GetApiTalkRouletteThemesParams) error {
	var maxCountPtr *int
	if params.MaxCount != nil {
		mc := int(*params.MaxCount) // int32 を int に変換
		maxCountPtr = &mc
	}

	// talkroulette.TalkRouletteTheme と api.TalkRouletteTheme に互換性があることを前提とします。
	// 互換性がない場合は、手動でのマッピングが必要になります。
	selectedThemesLogic, errLogic := talkroulette.GetTalkRouletteThemesLogic(params.Genre, maxCountPtr)
	if errLogic != nil {
		// エラー処理、例: 500エラーを返す
		// この特定のロジック関数はまだエラーを返しませんが、良い習慣です。
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "テーマの処理中にエラーが発生しました"})
	}

	// []talkroulette.TalkRouletteTheme を []api.TalkRouletteTheme に変換します。
	// OpenAPI の 'required' フィールドに基づき、api.TalkRouletteTheme のフィールドは非ポインタです。
	apiThemes := make([]api.TalkRouletteTheme, len(selectedThemesLogic))
	for i, t := range selectedThemesLogic {
		apiThemes[i] = api.TalkRouletteTheme{
			Id:    t.ID,    // 直接の文字列代入
			Genre: t.Genre, // 直接の文字列代入
			Theme: t.Theme, // 直接の文字列代入
		}
	}

	// OpenAPI 仕様に基づき、TalkRouletteThemeResponse.Themes は `[]api.TalkRouletteTheme` です。
	response := api.TalkRouletteThemeResponse{
		Themes: apiThemes,
	}

	return ctx.JSON(http.StatusOK, response)
}
