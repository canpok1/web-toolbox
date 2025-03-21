package api

import (
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) PostApiPlanningPokerSessions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
