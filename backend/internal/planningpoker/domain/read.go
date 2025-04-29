package domain

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/domain/model"
)

type ReadRoundUsecase struct {
	roundRepo       model.RoundRepository
	participantRepo model.ParticipantRepository
	voteRepo        model.VoteRepository
}

func NewReadRoundUsecase(rRepo model.RoundRepository, pRepo model.ParticipantRepository, vRepo model.VoteRepository) *ReadRoundUsecase {
	return &ReadRoundUsecase{
		roundRepo:       rRepo,
		participantRepo: pRepo,
		voteRepo:        vRepo,
	}
}

type ReadRoundResult struct {
	Round api.Round
}

func (r *ReadRoundUsecase) Read(ctx context.Context, roundID string, participantID *string) (*ReadRoundResult, error) {
	redisRound, err := r.roundRepo.GetRound(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get round from redis: roundID=%s, err=%v", roundID, err)
	}
	if redisRound == nil {
		return nil, fmt.Errorf("round not found: roundID=%s", roundID)
	}

	apiRound := api.Round{
		RoundId:   roundID,
		SessionId: redisRound.SessionId,
		Status:    api.RoundStatus(redisRound.Status),
		CreatedAt: redisRound.CreatedAt,
		UpdatedAt: redisRound.UpdatedAt,
		Votes:     []api.Vote{},
		// Summary は後で設定
	}

	voteIDs, err := r.voteRepo.GetVotesInRound(ctx, roundID)
	if err != nil {
		// 投票リスト取得エラーはログに残すが、ラウンド情報自体は返す（投票情報なしで）
		log.Printf("failed to get votes in round, returning round data without votes: roundID=%s, err=%v", roundID, err)
		// エラーを返さずに処理を続行。apiRound.Votes は空のまま
	}

	numericVotes := []float32{} // 数値として扱える投票値を格納するスライス
	voteCountMap := map[string]api.VoteCount{}

	if len(voteIDs) > 0 {
		apiVotes := make([]api.Vote, 0, len(voteIDs))
		for _, voteID := range voteIDs {
			redisVote, err := r.voteRepo.GetVote(ctx, voteID)
			if err != nil {
				log.Printf("failed to get vote details, skipping vote: roundID=%s, voteID=%s, err=%v", roundID, voteID, err)
				continue // 個別の投票取得エラーはスキップ
			}
			if redisVote == nil {
				log.Printf("vote not found, but ID was listed in round, skipping vote: roundID=%s, voteID=%s", roundID, voteID)
				continue // データ不整合の可能性、スキップ
			}

			redisParticipant, err := r.participantRepo.GetParticipant(ctx, redisVote.ParticipantId)
			// GetParticipant でエラーが発生した場合や見つからなかった場合は、その投票を除外する
			if err != nil {
				log.Printf("failed to get participant details for vote, skipping vote: roundID=%s, voteID=%s, participantID=%s, err=%v", roundID, voteID, redisVote.ParticipantId, err)
				continue
			}
			if redisParticipant == nil {
				log.Printf("participant not found for vote, skipping vote: roundID=%s, voteID=%s, participantID=%s", roundID, voteID, redisVote.ParticipantId)
				continue
			}

			apiVote := api.Vote{
				ParticipantId:   redisVote.ParticipantId,
				ParticipantName: redisParticipant.Name,
			}

			// 公開時か自身のもののみ投票結果をセット
			isRevealed := apiRound.Status == api.Revealed
			isOwnVote := participantID != nil && *participantID == apiVote.ParticipantId

			if isRevealed || isOwnVote {
				if redisVote.Value != "" {
					valueCopy := redisVote.Value // ポインタ用にコピー
					apiVote.Value = &valueCopy

					// revealed 状態の場合、数値変換を試みて集計用スライスに追加
					if isRevealed {
						participant := api.SessionParticipant{
							ParticipantId: redisVote.ParticipantId,
							Name:          redisParticipant.Name,
						}
						if voteCount, exist := voteCountMap[redisVote.Value]; exist {
							voteCountMap[redisVote.Value] = api.VoteCount{
								Value:        voteCount.Value,
								Count:        voteCount.Count + 1,
								Participants: append(voteCount.Participants, participant),
							}
						} else {
							voteCountMap[redisVote.Value] = api.VoteCount{
								Value:        redisVote.Value,
								Count:        1,
								Participants: []api.SessionParticipant{participant},
							}
						}

						numVal, err := strconv.ParseFloat(redisVote.Value, 32)
						if err == nil {
							numericVotes = append(numericVotes, float32(numVal))
						} else {
							// 数値に変換できない値はログに残しても良い（例: '?', 'coffee'）
							log.Printf("Vote value is not numeric, skipping for summary calculation: roundID=%s, voteID=%s, value=%s", roundID, voteID, redisVote.Value)
						}
					}
				}
			}
			apiVotes = append(apiVotes, apiVote)
		}
		apiRound.Votes = apiVotes
	}

	// revealed 状態の場合、サマリーを生成
	if apiRound.Status == api.Revealed {
		summary := api.RoundSummary{
			VoteCounts: []api.VoteCount{},
		}

		for _, scale := range model.ScaleOrder {
			if voteCount, exist := voteCountMap[scale]; exist {
				summary.VoteCounts = append(summary.VoteCounts, voteCount)
			}
		}

		if len(numericVotes) > 0 {
			sort.Slice(numericVotes, func(i, j int) bool {
				return numericVotes[i] < numericVotes[j]
			})

			var sum float32
			for _, v := range numericVotes {
				sum += v
			}
			average := sum / float32(len(numericVotes))
			summary.Average = &average

			var median float32
			n := len(numericVotes)
			if n%2 == 1 {
				// 奇数個の場合、中央の要素
				median = numericVotes[n/2]
			} else {
				// 偶数個の場合、中央の2つの要素の平均
				median = (numericVotes[n/2-1] + numericVotes[n/2]) / 2.0
			}
			summary.Median = &median

			max := numericVotes[len(numericVotes)-1]
			summary.Max = &max

			min := numericVotes[0]
			summary.Min = &min
		}

		apiRound.Summary = &summary
	}

	return &ReadRoundResult{
		Round: apiRound,
	}, nil
}
