// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CreateSessionRequest defines model for CreateSessionRequest.
type CreateSessionRequest struct {
	CustomScale *[]string `json:"customScale,omitempty"`
	HostName    *string   `json:"hostName,omitempty"`
	ScaleType   *string   `json:"scaleType,omitempty"`
	SessionName *string   `json:"sessionName,omitempty"`
}

// CreateSessionResponse defines model for CreateSessionResponse.
type CreateSessionResponse struct {
	HostId    *openapi_types.UUID `json:"hostId,omitempty"`
	SessionId *openapi_types.UUID `json:"sessionId,omitempty"`
}

// EndSessionResponse defines model for EndSessionResponse.
type EndSessionResponse = map[string]interface{}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// GetSessionResponse defines model for GetSessionResponse.
type GetSessionResponse struct {
	CreatedAt      *time.Time          `json:"createdAt,omitempty"`
	CurrentRoundId *openapi_types.UUID `json:"currentRoundId,omitempty"`
	CustomScale    *[]string           `json:"customScale,omitempty"`
	HostId         *openapi_types.UUID `json:"hostId,omitempty"`
	ScaleType      *string             `json:"scaleType,omitempty"`
	SessionId      *openapi_types.UUID `json:"sessionId,omitempty"`
	SessionName    *string             `json:"sessionName,omitempty"`
	Status         *string             `json:"status,omitempty"`
	UpdatedAt      *time.Time          `json:"updatedAt,omitempty"`
}

// JoinSessionRequest defines model for JoinSessionRequest.
type JoinSessionRequest struct {
	Name *string `json:"name,omitempty"`
}

// JoinSessionResponse defines model for JoinSessionResponse.
type JoinSessionResponse struct {
	ParticipantId *openapi_types.UUID `json:"participantId,omitempty"`
}

// RevealRoundResponse defines model for RevealRoundResponse.
type RevealRoundResponse = map[string]interface{}

// SendVoteRequest defines model for SendVoteRequest.
type SendVoteRequest struct {
	ParticipantId *openapi_types.UUID `json:"participantId,omitempty"`
	Value         *string             `json:"value,omitempty"`
}

// SendVoteResponse defines model for SendVoteResponse.
type SendVoteResponse struct {
	VoteId *openapi_types.UUID `json:"voteId,omitempty"`
}

// StartRoundResponse defines model for StartRoundResponse.
type StartRoundResponse struct {
	RoundId *openapi_types.UUID `json:"roundId,omitempty"`
}

// PostApiPlanningPokerRoundsRoundIdVotesJSONRequestBody defines body for PostApiPlanningPokerRoundsRoundIdVotes for application/json ContentType.
type PostApiPlanningPokerRoundsRoundIdVotesJSONRequestBody = SendVoteRequest

// PostApiPlanningPokerSessionsJSONRequestBody defines body for PostApiPlanningPokerSessions for application/json ContentType.
type PostApiPlanningPokerSessionsJSONRequestBody = CreateSessionRequest

// PostApiPlanningPokerSessionsSessionIdParticipantsJSONRequestBody defines body for PostApiPlanningPokerSessionsSessionIdParticipants for application/json ContentType.
type PostApiPlanningPokerSessionsSessionIdParticipantsJSONRequestBody = JoinSessionRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ラウンドを終了する
	// (POST /api/planning-poker/rounds/{roundId}/reveal)
	PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId openapi_types.UUID) error
	// 投票を送信する
	// (POST /api/planning-poker/rounds/{roundId}/votes)
	PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId openapi_types.UUID) error
	// セッションを作成する
	// (POST /api/planning-poker/sessions)
	PostApiPlanningPokerSessions(ctx echo.Context) error
	// セッションを取得する
	// (GET /api/planning-poker/sessions/{sessionId})
	GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId openapi_types.UUID) error
	// セッションを終了する
	// (POST /api/planning-poker/sessions/{sessionId}/end)
	PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId openapi_types.UUID) error
	// セッションに参加する
	// (POST /api/planning-poker/sessions/{sessionId}/participants)
	PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId openapi_types.UUID) error
	// ラウンドを開始する
	// (POST /api/planning-poker/sessions/{sessionId}/rounds)
	PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostApiPlanningPokerRoundsRoundIdReveal converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roundId" -------------
	var roundId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roundId", ctx.Param("roundId"), &roundId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roundId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerRoundsRoundIdReveal(ctx, roundId)
	return err
}

// PostApiPlanningPokerRoundsRoundIdVotes converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "roundId" -------------
	var roundId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roundId", ctx.Param("roundId"), &roundId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter roundId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerRoundsRoundIdVotes(ctx, roundId)
	return err
}

// PostApiPlanningPokerSessions converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerSessions(ctx)
	return err
}

// GetApiPlanningPokerSessionsSessionId converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiPlanningPokerSessionsSessionId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", ctx.Param("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sessionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetApiPlanningPokerSessionsSessionId(ctx, sessionId)
	return err
}

// PostApiPlanningPokerSessionsSessionIdEnd converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", ctx.Param("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sessionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerSessionsSessionIdEnd(ctx, sessionId)
	return err
}

// PostApiPlanningPokerSessionsSessionIdParticipants converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", ctx.Param("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sessionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerSessionsSessionIdParticipants(ctx, sessionId)
	return err
}

// PostApiPlanningPokerSessionsSessionIdRounds converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", ctx.Param("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sessionId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostApiPlanningPokerSessionsSessionIdRounds(ctx, sessionId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/api/planning-poker/rounds/:roundId/reveal", wrapper.PostApiPlanningPokerRoundsRoundIdReveal)
	router.POST(baseURL+"/api/planning-poker/rounds/:roundId/votes", wrapper.PostApiPlanningPokerRoundsRoundIdVotes)
	router.POST(baseURL+"/api/planning-poker/sessions", wrapper.PostApiPlanningPokerSessions)
	router.GET(baseURL+"/api/planning-poker/sessions/:sessionId", wrapper.GetApiPlanningPokerSessionsSessionId)
	router.POST(baseURL+"/api/planning-poker/sessions/:sessionId/end", wrapper.PostApiPlanningPokerSessionsSessionIdEnd)
	router.POST(baseURL+"/api/planning-poker/sessions/:sessionId/participants", wrapper.PostApiPlanningPokerSessionsSessionIdParticipants)
	router.POST(baseURL+"/api/planning-poker/sessions/:sessionId/rounds", wrapper.PostApiPlanningPokerSessionsSessionIdRounds)

}
