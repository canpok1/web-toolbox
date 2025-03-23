package api

import (
	"context"
	"fmt"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/google/uuid"
)

func (s *Server) ValidatePostApiPlanningPokerSessions(body *CreateSessionRequest) error {
	if body == nil {
		return fmt.Errorf("request body is required")
	}
	if body.SessionName == "" {
		return fmt.Errorf("sessionName is required")
	}
	if body.HostName == "" {
		return fmt.Errorf("hostName is required")
	}
	if body.ScaleType == "" {
		return fmt.Errorf("scaleType is required")
	}
	if body.ScaleType != Fibonacci && body.ScaleType != TShirt && body.ScaleType != Custom {
		return fmt.Errorf("invalid scaleType: %s", body.ScaleType)
	}
	if body.ScaleType == Custom && len(*body.CustomScale) == 0 {
		return fmt.Errorf("customScale is required when scaleType is custom")
	}

	return nil
}

func (s *Server) HandlePostApiPlanningPokerSessions(body *CreateSessionRequest) (*CreateSessionResponse, error) {
	if body == nil {
		return nil, fmt.Errorf("request body is required")
	}

	hostId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate host uuid: %v", err)
	}

	sessionId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session uuid: %v", err)
	}

	// セッション情報の保存
	session := redis.Session{
		SessionName: body.SessionName,
		HostId:      hostId.String(),
		ScaleType:   string(body.ScaleType),
		CustomScale: []string{},
		Status:      "waiting",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if body.CustomScale != nil {
		session.CustomScale = *body.CustomScale
	}

	ctx := context.Background()
	if err = s.redis.CreateSession(ctx, sessionId.String(), session); err != nil {
		return nil, fmt.Errorf("failed to save session to redis: %v", err)
	}

	// レスポンスの作成
	res := CreateSessionResponse{
		HostId:    hostId,
		SessionId: sessionId,
	}
	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionID uuid.UUID, body *JoinSessionRequest) (*JoinSessionResponse, error) {
	// TODO POST /api/planning-poker/sessions/{sessionId}/participants の処理を実装
	participantId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate participant uuid: %v", err)
	}

	res := JoinSessionResponse{
		ParticipantId: participantId,
	}
	return &res, nil
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdReveal(roundId uuid.UUID) (*RevealRoundResponse, error) {
	// TODO POST /api/planning-poker/rounds/{roundId}/reveal の処理を実装
	return nil, fmt.Errorf("HandlePostApiPlanningPokerRoundsRoundIdReveal: %s is not implemented yet", roundId)
}

func (s *Server) HandlePostApiPlanningPokerRoundsRoundIdVotes(roundId uuid.UUID, body *SendVoteRequest) (*SendVoteResponse, error) {
	// TODO POST /api/planning-poker/rounds/{roundId}/votes の処理を実装
	voteId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate vote uuid: %v", err)
	}

	res := SendVoteResponse{
		VoteId: voteId,
	}
	return &res, nil
}

func (s *Server) HandleGetApiPlanningPokerSessionsSessionId(sessionID uuid.UUID) (*GetSessionResponse, error) {
	// TODO POST /api/planning-poker/sessions/{sessionId}/participants の処理を実装
	return nil, fmt.Errorf("GetApiPlanningPokerSessionsSessionId: %s is not implemented yet", sessionID)
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdEnd(sessionID uuid.UUID) (*EndSessionResponse, error) {
	// TODO POST /api/planning-poker/sessions/{sessionId}/end の処理を実装
	return nil, fmt.Errorf("PostApiPlanningPokerSessionsSessionIdEnd: %s is not implemented yet", sessionID)
}

func (s *Server) HandlePostApiPlanningPokerSessionsSessionIdRounds(sessionID uuid.UUID) (*StartRoundResponse, error) {
	// TODO POST /api/planning-poker/sessions/{sessionId}/rounds の処理を実装
	roundId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate round uuid: %v", err)
	}
	res := StartRoundResponse{RoundId: roundId}
	return &res, nil
}
