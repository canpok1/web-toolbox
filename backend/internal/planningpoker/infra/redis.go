package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/planningpoker/model"
	redislib "github.com/redis/go-redis/v9"
)

// Client is an interface for interacting with Redis.
type RedisClient interface {
	model.SessionRepository
	model.RoundRepository
	model.ParticipantRepository
	model.VoteRepository
	Close() error
}

// client is a wrapper around the redislib.Client.
type redisClient struct {
	client            *redislib.Client
	defaultExpiration time.Duration
}

// NewClient creates a new Redis client.
func NewRedisClient(addr, password string, db int, defaultExpiration time.Duration) (RedisClient, error) {
	rdb := redislib.NewClient(&redislib.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s: %w", addr, err)
	}

	return &redisClient{client: rdb, defaultExpiration: defaultExpiration}, nil
}

// Close closes the Redis connection.
func (c *redisClient) Close() error {
	return c.client.Close()
}

// --- Session ---

// CreateSession creates a new session in Redis.
func (c *redisClient) CreateSession(ctx context.Context, sessionId string, session model.Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// GetSession retrieves a session from Redis.
func (c *redisClient) GetSession(ctx context.Context, sessionId string) (*model.Session, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil { // redislib.Nil を使用
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get session with key %s: %w", key, err)
	}
	var session model.Session
	err = json.Unmarshal([]byte(data), &session)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}
	return &session, nil
}

// UpdateSession updates a session in Redis.
func (c *redisClient) UpdateSession(ctx context.Context, sessionId string, session model.Session) error {
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// --- Round ---

// CreateRound creates a new round in Redis.
func (c *redisClient) CreateRound(ctx context.Context, roundId string, round model.Round) error {
	round.CreatedAt = time.Now()
	round.UpdatedAt = time.Now()
	data, err := json.Marshal(round)
	if err != nil {
		return fmt.Errorf("failed to marshal round: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// GetRound retrieves a round from Redis.
func (c *redisClient) GetRound(ctx context.Context, roundId string) (*model.Round, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil { // redislib.Nil を使用
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get round with key %s: %w", key, err)
	}
	var round model.Round
	err = json.Unmarshal([]byte(data), &round)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal round: %w", err)
	}
	return &round, nil
}

// UpdateRound updates a round in Redis.
func (c *redisClient) UpdateRound(ctx context.Context, roundId string, round model.Round) error {
	round.UpdatedAt = time.Now()
	data, err := json.Marshal(round)
	if err != nil {
		return fmt.Errorf("failed to marshal round: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// --- Participant ---

// CreateParticipant creates a new participant in Redis.
func (c *redisClient) CreateParticipant(ctx context.Context, participantId string, participant model.Participant) error {
	participant.CreatedAt = time.Now()
	participant.UpdatedAt = time.Now()
	data, err := json.Marshal(participant)
	if err != nil {
		return fmt.Errorf("failed to marshal participant: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// GetParticipant retrieves a participant from Redis.
func (c *redisClient) GetParticipant(ctx context.Context, participantId string) (*model.Participant, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil {
		return nil, fmt.Errorf("participant %s not found", participantId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get participant %s: %w", participantId, err)
	}
	var participant model.Participant
	err = json.Unmarshal([]byte(data), &participant)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal participant: %w", err)
	}
	return &participant, nil
}

// GetVoteIdByRoundIdAndParticipantId retrieves a voteId from Redis by roundId and participantId.
func (c *redisClient) GetVoteIdByRoundIdAndParticipantId(ctx context.Context, roundId, participantId string) (*string, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s:votes", roundId)
	voteIds, err := c.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get vote ids with key %s: %w", key, err)
	}
	for _, voteId := range voteIds {
		vote, err := c.GetVote(ctx, voteId)
		if err != nil {
			return nil, fmt.Errorf("failed to get vote with voteId %s: %w", voteId, err)
		}
		if vote.ParticipantId == participantId {
			return &voteId, nil
		}
	}
	return nil, nil
}

// UpdateParticipant updates a participant in Redis.
func (c *redisClient) UpdateParticipant(ctx context.Context, participantId string, participant model.Participant) error {
	participant.UpdatedAt = time.Now()
	data, err := json.Marshal(participant)
	if err != nil {
		return fmt.Errorf("failed to marshal participant: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// --- Vote ---

// CreateVote creates a new vote in Redis.
func (c *redisClient) CreateVote(ctx context.Context, voteId string, vote model.Vote) error {
	vote.CreatedAt = time.Now()
	vote.UpdatedAt = time.Now()
	data, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// GetVote retrieves a vote from Redis.
func (c *redisClient) GetVote(ctx context.Context, voteId string) (*model.Vote, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil {
		return nil, fmt.Errorf("vote %s not found", voteId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get vote %s: %w", voteId, err)
	}
	var vote model.Vote
	err = json.Unmarshal([]byte(data), &vote)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal vote: %w", err)
	}
	return &vote, nil
}

// UpdateVote updates a vote in Redis.
func (c *redisClient) UpdateVote(ctx context.Context, voteId string, vote model.Vote) error {
	vote.UpdatedAt = time.Now()
	data, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	return c.client.Set(ctx, key, data, c.defaultExpiration).Err()
}

// --- Session Participants ---

// AddParticipantToSession adds a participant to a session's participant list.
func (c *redisClient) AddParticipantToSession(ctx context.Context, sessionId, participantId string) error {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s:participants", sessionId)
	if err := c.client.SAdd(ctx, key, participantId).Err(); err != nil {
		return err
	}
	if err := setExpireIfNotSet(ctx, c.client, key, c.defaultExpiration); err != nil {
		return err
	}
	return nil
}

// GetParticipantsInSession retrieves all participants in a session.
func (c *redisClient) GetParticipantsInSession(ctx context.Context, sessionId string) ([]string, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s:participants", sessionId)
	return c.client.SMembers(ctx, key).Result()
}

// --- Round Votes ---

// AddVoteToRound adds a vote to a round's vote list.
func (c *redisClient) AddVoteToRound(ctx context.Context, roundId, voteId string) error {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s:votes", roundId)
	if err := c.client.SAdd(ctx, key, voteId).Err(); err != nil {
		return err
	}
	if err := setExpireIfNotSet(ctx, c.client, key, c.defaultExpiration); err != nil {
		return err
	}
	return nil
}

// GetVotesInRound retrieves all votes in a round.
func (c *redisClient) GetVotesInRound(ctx context.Context, roundId string) ([]string, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s:votes", roundId)
	return c.client.SMembers(ctx, key).Result()
}

func setExpireIfNotSet(ctx context.Context, client *redislib.Client, key string, expiration time.Duration) error {
	ttlResult, err := client.TTL(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get TTL for key %s: %w", key, err)
	}

	if ttlResult > 0 {
		// キーに有効期限が設定済みの場合
		return nil
	}
	if ttlResult == -2 {
		// キーが存在しない場合
		return nil
	}

	// キーは存在するが、有効期限が設定されていない場合
	setCmd := client.Expire(ctx, key, expiration)
	if err := setCmd.Err(); err != nil {
		return fmt.Errorf("failed to set TTL for key %s: %w", key, err)
	}
	return nil
}
