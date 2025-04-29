package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain/model"
)

type UpdateRoundStatusUsecase struct {
	roundRepo model.RoundRepository
}

func NewUpdateRoundStatusUsecase(rRepo model.RoundRepository) *UpdateRoundStatusUsecase {
	return &UpdateRoundStatusUsecase{
		roundRepo: rRepo,
	}
}

type UpdateRoundStatusResult struct {
	Round model.Round
}

func (r *UpdateRoundStatusUsecase) Update(ctx context.Context, roundID string, status string) (*UpdateRoundStatusResult, error) {
	round, err := r.roundRepo.GetRound(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundID, err)
	}
	if round == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundID)
	}

	round.Status = status
	round.UpdatedAt = time.Now()
	if err := r.roundRepo.UpdateRound(ctx, roundID, *round); err != nil {
		return nil, fmt.Errorf("failed to update round in redis: roundID=%s, err=%v", roundID, err)
	}

	return &UpdateRoundStatusResult{
		Round: *round,
	}, nil
}
