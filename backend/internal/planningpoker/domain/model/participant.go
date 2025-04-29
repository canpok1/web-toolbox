package model

import (
	"context"
	"time"
)

type Participant struct {
	SessionId string    `json:"sessionId"`
	Name      string    `json:"name"`
	IsHost    bool      `json:"isHost"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ParticipantRepository interface {
	CreateParticipant(ctx context.Context, participantId string, participant Participant) error
	GetParticipant(ctx context.Context, participantId string) (*Participant, error)
	UpdateParticipant(ctx context.Context, participantId string, participant Participant) error
	AddParticipantToSession(ctx context.Context, sessionId, participantId string) error
	GetParticipantsInSession(ctx context.Context, sessionId string) ([]string, error)
}
