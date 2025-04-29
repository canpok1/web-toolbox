package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain/model"
	"github.com/google/uuid"
)

type UpsertVoteUsecase struct {
	roundRepo       model.RoundRepository
	participantRepo model.ParticipantRepository
	voteRepo        model.VoteRepository
}

func NewUpsertVoteUsecase(rRepo model.RoundRepository, pRepo model.ParticipantRepository, vRepo model.VoteRepository) *UpsertVoteUsecase {
	return &UpsertVoteUsecase{
		roundRepo:       rRepo,
		participantRepo: pRepo,
		voteRepo:        vRepo,
	}
}

type UpsertVoteResult struct {
	VoteID string
	Vote   model.Vote
	Round  model.Round
}

func (r *UpsertVoteUsecase) Upsert(ctx context.Context, roundID string, participantID string, value string) (*UpsertVoteResult, error) {
	round, err := r.roundRepo.GetRound(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundID, err)
	}
	if round == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundID)
	}

	if round.Status != "voting" {
		return nil, fmt.Errorf("round is not in voting state: roundID=%s", roundID)
	}

	participant, err := r.participantRepo.GetParticipant(ctx, participantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get participant from redis: roundID=%s, participantID=%s, err=%v", roundID, participantID, err)
	}
	if participant == nil {
		return nil, fmt.Errorf(
			"participant not found: roundID=%s, participantID=%s",
			roundID,
			participantID,
		)
	}

	voteID, err := r.voteRepo.GetVoteIdByRoundIdAndParticipantId(ctx, roundID, participantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vote id from redis: roundID=%s, participantID=%s, err=%v", roundID, participantID, err)
	}

	if voteID == nil {
		// Create a new vote
		newVoteId, err := uuid.NewUUID()
		if err != nil {
			return nil, fmt.Errorf("failed to generate vote uuid: roundID=%s, err=%v", roundID, err)
		}

		vote := model.Vote{
			RoundId:       roundID,
			ParticipantId: participantID,
			Value:         value,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		// Save the vote to Redis
		if err := r.voteRepo.CreateVote(ctx, newVoteId.String(), vote); err != nil {
			return nil, fmt.Errorf("failed to create vote in redis: roundID=%s, voteID=%s, err=%v", roundID, newVoteId.String(), err)
		}

		// Add the vote to the round's vote list
		if err := r.voteRepo.AddVoteToRound(ctx, roundID, newVoteId.String()); err != nil {
			return nil, fmt.Errorf("failed to add vote to round in redis: roundID=%s, voteID=%s, err=%v", roundID, newVoteId.String(), err)
		}

		return &UpsertVoteResult{
			VoteID: newVoteId.String(),
			Vote:   vote,
			Round:  *round,
		}, nil
	}

	// Update the existing vote
	vote, err := r.voteRepo.GetVote(ctx, *voteID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vote from redis: roundID=%s, voteID=%s, err=%v", roundID, *voteID, err)
	}
	vote.Value = value
	vote.UpdatedAt = time.Now()

	if err := r.voteRepo.UpdateVote(ctx, *voteID, *vote); err != nil {
		return nil, fmt.Errorf("failed to update vote in redis: roundID=%s, voteID=%s, err=%v", roundID, *voteID, err)
	}

	return &UpsertVoteResult{
		VoteID: *voteID,
		Vote:   *vote,
		Round:  *round,
	}, nil
}
