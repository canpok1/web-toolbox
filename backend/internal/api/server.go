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

func NewServer(redis redis.Client) *Server {
	return &Server{redis: redis}
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdReveal(ctx echo.Context, roundId uuid.UUID) error {
	return ctx.String(http.StatusNotImplemented, fmt.Sprintf("PostApiPlanningPokerRoundsRoundIdReveal: %s", roundId))
}

func (s *Server) PostApiPlanningPokerRoundsRoundIdVotes(ctx echo.Context, roundId uuid.UUID) error {
	var req SendVoteRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	voteId, err := uuid.NewUUID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}

	res := SendVoteResponse{
		VoteId: &voteId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessions(ctx echo.Context) error {
	var req CreateSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	hostId, err := uuid.NewUUID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}

	sessionId, err := uuid.NewUUID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}

	res := CreateSessionResponse{
		HostId:    &hostId,
		SessionId: &sessionId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) GetApiPlanningPokerSessionsSessionId(ctx echo.Context, sessionId uuid.UUID) error {
	return ctx.String(http.StatusNotImplemented, fmt.Sprintf("GetApiPlanningPokerSessionsSessionId: %s", sessionId))
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdEnd(ctx echo.Context, sessionId uuid.UUID) error {
	return ctx.String(http.StatusNotImplemented, fmt.Sprintf("PostApiPlanningPokerSessionsSessionIdEnd: %s", sessionId))
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdParticipants(ctx echo.Context, sessionId uuid.UUID) error {
	var req JoinSessionRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	participantId, err := uuid.NewUUID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}

	res := JoinSessionResponse{
		ParticipantId: &participantId,
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (s *Server) PostApiPlanningPokerSessionsSessionIdRounds(ctx echo.Context, sessionId uuid.UUID) error {
	roundId, err := uuid.NewUUID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to generate uuid")
	}
	res := StartRoundResponse{RoundId: &roundId}
	return ctx.JSON(http.StatusCreated, res)
}
