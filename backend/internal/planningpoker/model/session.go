package model

import (
	"context"
	"time"
)

type Session struct {
	HostId         string    `json:"hostId"`
	ScaleType      string    `json:"scaleType"`
	CustomScale    []string  `json:"customScale"`
	CurrentRoundId string    `json:"currentRoundId"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type SessionRepository interface {
	CreateSession(ctx context.Context, sessionId string, session Session) error
	GetSession(ctx context.Context, sessionId string) (*Session, error)
	UpdateSession(ctx context.Context, sessionId string, session Session) error
}
