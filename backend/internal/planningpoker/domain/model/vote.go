package model

import (
	"context"
	"time"
)

type Vote struct {
	RoundId       string    `json:"roundId"`
	ParticipantId string    `json:"participantId"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type VoteRepository interface {
	CreateVote(ctx context.Context, voteId string, vote Vote) error
	GetVote(ctx context.Context, voteId string) (*Vote, error)
	GetVoteIdByRoundIdAndParticipantId(ctx context.Context, roundId, participantId string) (*string, error)
	UpdateVote(ctx context.Context, voteId string, vote Vote) error
	AddVoteToRound(ctx context.Context, roundId, voteId string) error
	GetVotesInRound(ctx context.Context, roundId string) ([]string, error)
}
