//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// PostApiPlanningPokerRoundsRoundIdVotesJSONBody defines parameters for PostApiPlanningPokerRoundsRoundIdVotes.
type PostApiPlanningPokerRoundsRoundIdVotesJSONBody struct {
	ParticipantId *openapi_types.UUID `json:"participantId,omitempty"`
	Value         *string             `json:"value,omitempty"`
}

// PostApiPlanningPokerSessionsJSONBody defines parameters for PostApiPlanningPokerSessions.
type PostApiPlanningPokerSessionsJSONBody struct {
	CustomScale *[]string `json:"customScale,omitempty"`
	HostName    *string   `json:"hostName,omitempty"`
	ScaleType   *string   `json:"scaleType,omitempty"`
	SessionName *string   `json:"sessionName,omitempty"`
}

// PostApiPlanningPokerSessionsSessionIdParticipantsJSONBody defines parameters for PostApiPlanningPokerSessionsSessionIdParticipants.
type PostApiPlanningPokerSessionsSessionIdParticipantsJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// PostApiPlanningPokerRoundsRoundIdVotesJSONRequestBody defines body for PostApiPlanningPokerRoundsRoundIdVotes for application/json ContentType.
type PostApiPlanningPokerRoundsRoundIdVotesJSONRequestBody PostApiPlanningPokerRoundsRoundIdVotesJSONBody

// PostApiPlanningPokerSessionsJSONRequestBody defines body for PostApiPlanningPokerSessions for application/json ContentType.
type PostApiPlanningPokerSessionsJSONRequestBody PostApiPlanningPokerSessionsJSONBody

// PostApiPlanningPokerSessionsSessionIdParticipantsJSONRequestBody defines body for PostApiPlanningPokerSessionsSessionIdParticipants for application/json ContentType.
type PostApiPlanningPokerSessionsSessionIdParticipantsJSONRequestBody PostApiPlanningPokerSessionsSessionIdParticipantsJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ラウンドを終了する
	// (POST /api/planning-poker/rounds/{roundId}/reveal)
	PostApiPlanningPokerRoundsRoundIdReveal(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID)
	// 投票を送信する
	// (POST /api/planning-poker/rounds/{roundId}/votes)
	PostApiPlanningPokerRoundsRoundIdVotes(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID)
	// セッションを作成する
	// (POST /api/planning-poker/sessions)
	PostApiPlanningPokerSessions(w http.ResponseWriter, r *http.Request)
	// セッションを取得する
	// (GET /api/planning-poker/sessions/{sessionId})
	GetApiPlanningPokerSessionsSessionId(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID)
	// セッションを終了する
	// (POST /api/planning-poker/sessions/{sessionId}/end)
	PostApiPlanningPokerSessionsSessionIdEnd(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID)
	// セッションに参加する
	// (POST /api/planning-poker/sessions/{sessionId}/participants)
	PostApiPlanningPokerSessionsSessionIdParticipants(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID)
	// ラウンドを開始する
	// (POST /api/planning-poker/sessions/{sessionId}/rounds)
	PostApiPlanningPokerSessionsSessionIdRounds(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostApiPlanningPokerRoundsRoundIdReveal operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerRoundsRoundIdReveal(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "roundId" -------------
	var roundId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roundId", r.PathValue("roundId"), &roundId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "roundId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerRoundsRoundIdReveal(w, r, roundId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiPlanningPokerRoundsRoundIdVotes operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerRoundsRoundIdVotes(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "roundId" -------------
	var roundId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "roundId", r.PathValue("roundId"), &roundId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "roundId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerRoundsRoundIdVotes(w, r, roundId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiPlanningPokerSessions operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerSessions(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerSessions(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetApiPlanningPokerSessionsSessionId operation middleware
func (siw *ServerInterfaceWrapper) GetApiPlanningPokerSessionsSessionId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", r.PathValue("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sessionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiPlanningPokerSessionsSessionId(w, r, sessionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiPlanningPokerSessionsSessionIdEnd operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdEnd(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", r.PathValue("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sessionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerSessionsSessionIdEnd(w, r, sessionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiPlanningPokerSessionsSessionIdParticipants operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdParticipants(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", r.PathValue("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sessionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerSessionsSessionIdParticipants(w, r, sessionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiPlanningPokerSessionsSessionIdRounds operation middleware
func (siw *ServerInterfaceWrapper) PostApiPlanningPokerSessionsSessionIdRounds(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "sessionId" -------------
	var sessionId openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "sessionId", r.PathValue("sessionId"), &sessionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sessionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiPlanningPokerSessionsSessionIdRounds(w, r, sessionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/rounds/{roundId}/reveal", wrapper.PostApiPlanningPokerRoundsRoundIdReveal)
	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/rounds/{roundId}/votes", wrapper.PostApiPlanningPokerRoundsRoundIdVotes)
	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/sessions", wrapper.PostApiPlanningPokerSessions)
	m.HandleFunc("GET "+options.BaseURL+"/api/planning-poker/sessions/{sessionId}", wrapper.GetApiPlanningPokerSessionsSessionId)
	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/sessions/{sessionId}/end", wrapper.PostApiPlanningPokerSessionsSessionIdEnd)
	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/sessions/{sessionId}/participants", wrapper.PostApiPlanningPokerSessionsSessionIdParticipants)
	m.HandleFunc("POST "+options.BaseURL+"/api/planning-poker/sessions/{sessionId}/rounds", wrapper.PostApiPlanningPokerSessionsSessionIdRounds)

	return m
}
