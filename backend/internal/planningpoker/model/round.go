package model

import (
	"context"
	"time"
)

type Round struct {
	SessionId string    `json:"sessionId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RoundRepository interface {
	CreateRound(ctx context.Context, roundId string, round Round) error
	GetRound(ctx context.Context, roundId string) (*Round, error)
	UpdateRound(ctx context.Context, roundId string, round Round) error
}
