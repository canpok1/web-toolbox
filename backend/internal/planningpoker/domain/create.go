package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain/model"
	"github.com/google/uuid"
)

type CreateSessionUsecase struct {
	sessionRepo     model.SessionRepository
	participantRepo model.ParticipantRepository
}

func NewCreateSessionUsecase(sRepo model.SessionRepository, pRepo model.ParticipantRepository) *CreateSessionUsecase {
	return &CreateSessionUsecase{
		sessionRepo:     sRepo,
		participantRepo: pRepo,
	}
}

type CreateSessionResult struct {
	SessionId string
	HostId    string
}

func (r *CreateSessionUsecase) Create(scaleType string, customScale *[]string, hostName string) (*CreateSessionResult, error) {
	hostId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate host uuid: %v", err)
	}
	hostIdValue := hostId.String()

	sessionId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate session uuid: %v", err)
	}
	sessionIdValue := sessionId.String()

	session := model.Session{
		HostId:      hostIdValue,
		ScaleType:   scaleType,
		CustomScale: []string{},
		Status:      "waiting",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if customScale != nil {
		session.CustomScale = *customScale
	}

	ctx := context.Background()
	if err = r.sessionRepo.CreateSession(ctx, sessionIdValue, session); err != nil {
		return nil, fmt.Errorf("failed to save session to redis: %v", err)
	}

	err = r.participantRepo.CreateParticipant(ctx, hostIdValue, model.Participant{
		SessionId: sessionIdValue,
		Name:      hostName,
		IsHost:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to save participant to redis: %v", err)
	}

	err = r.participantRepo.AddParticipantToSession(ctx, sessionIdValue, hostIdValue)
	if err != nil {
		return nil, fmt.Errorf("failed to add participant list to redis: %v", err)
	}

	return &CreateSessionResult{
		SessionId: sessionIdValue,
		HostId:    hostIdValue,
	}, nil
}

type CreateParticipantUsecase struct {
	sessionRepo     model.SessionRepository
	participantRepo model.ParticipantRepository
}

func NewCreateParticipantUsecase(sRepo model.SessionRepository, pRepo model.ParticipantRepository) *CreateParticipantUsecase {
	return &CreateParticipantUsecase{
		sessionRepo:     sRepo,
		participantRepo: pRepo,
	}
}

type CreateParticipantResult struct {
	ParticipantId string
}

func (r *CreateParticipantUsecase) Create(sessionID string, participantName string) (*CreateParticipantResult, error) {
	participantId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate participant uuid (sessionID: %s): %v", sessionID, err)
	}

	ctx := context.Background()

	session, err := r.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis (sessionID: %s): %v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session is not found (sessionID: %s)", sessionID)
	}

	participant := model.Participant{
		SessionId: sessionID,
		Name:      participantName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := r.participantRepo.CreateParticipant(ctx, participantId.String(), participant); err != nil {
		return nil, fmt.Errorf("failed to create participant (sessionID: %s): %v", sessionID, err)
	}
	if err := r.participantRepo.AddParticipantToSession(ctx, sessionID, participantId.String()); err != nil {
		return nil, fmt.Errorf("failed to add participant to session (sessionID: %s, participantID: %s): %v", sessionID, participantId.String(), err)
	}

	return &CreateParticipantResult{
		ParticipantId: participantId.String(),
	}, nil
}

type CreateRoundUsecase struct {
	sessionRepo model.SessionRepository
	roundRepo   model.RoundRepository
}

func NewCreateRoundUsecase(sRepo model.SessionRepository, rRepo model.RoundRepository) *CreateRoundUsecase {
	return &CreateRoundUsecase{sessionRepo: sRepo, roundRepo: rRepo}
}

type CreateRoundResult struct {
	RoundId string
}

func (r *CreateRoundUsecase) Create(ctx context.Context, sessionID string) (*CreateRoundResult, error) {
	session, err := r.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session from redis: sessionID=%s, err=%v", sessionID, err)
	}
	if session == nil {
		return nil, fmt.Errorf("session not found: sessionID=%s", sessionID)
	}

	// Create a new round
	roundId, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate round uuid: sessionID=%s, err=%v", sessionID, err)
	}
	roundIdValue := roundId.String()

	round := model.Round{
		SessionId: sessionID,
		Status:    "voting",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the round to Redis
	if err := r.roundRepo.CreateRound(ctx, roundIdValue, round); err != nil {
		return nil, fmt.Errorf("failed to create round in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	// Update the session's currentRoundId
	session.CurrentRoundId = roundIdValue
	session.Status = "inProgress"
	session.UpdatedAt = time.Now()
	if err := r.sessionRepo.UpdateSession(ctx, sessionID, *session); err != nil {
		return nil, fmt.Errorf("failed to update session in redis: sessionID=%s, roundID=%s, err=%v", sessionID, roundId, err)
	}

	return &CreateRoundResult{
		RoundId: roundIdValue,
	}, nil
}
