package api

import (
	"encoding/json"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/redis"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
	redis redis.Client
}

func NewServer(redis redis.Client) *Server {
	return &Server{redis: redis}
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(w http.ResponseWriter, r *http.Request, roundId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) PostApiPlanningPokerSessions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(w http.ResponseWriter, r *http.Request, sessionId openapi_types.UUID) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "not implemented"})
}
