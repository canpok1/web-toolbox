package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	redislib "github.com/redis/go-redis/v9"
)

// Client is a wrapper around the redislib.Client.
type Client struct {
	client *redislib.Client
}

// NewClient creates a new Redis client.
func NewClient(addr, password string, db int) (*Client, error) {
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
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Client{client: rdb}, nil
}

// Close closes the Redis connection.
func (c *Client) Close() error {
	return c.client.Close()
}

// --- Session ---

// Session represents a planning poker session.
type Session struct {
	SessionName    string    `json:"sessionName"`
	HostId         string    `json:"hostId"`
	ScaleType      string    `json:"scaleType"`
	CustomScale    []string  `json:"customScale"`
	CurrentRoundId string    `json:"currentRoundId"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// CreateSession creates a new session in Redis.
func (c *Client) CreateSession(ctx context.Context, sessionId string, session Session) error {
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// GetSession retrieves a session from Redis.
func (c *Client) GetSession(ctx context.Context, sessionId string) (*Session, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil { // redislib.Nil を使用
		return nil, fmt.Errorf("session %s not found", sessionId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get session %s: %w", sessionId, err)
	}
	var session Session
	err = json.Unmarshal([]byte(data), &session)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}
	return &session, nil
}

// UpdateSession updates a session in Redis.
func (c *Client) UpdateSession(ctx context.Context, sessionId string, session Session) error {
	session.UpdatedAt = time.Now()
	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s", sessionId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// --- Round ---

// Round represents a round in a planning poker session.
type Round struct {
	SessionId string    `json:"sessionId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateRound creates a new round in Redis.
func (c *Client) CreateRound(ctx context.Context, roundId string, round Round) error {
	round.CreatedAt = time.Now()
	round.UpdatedAt = time.Now()
	data, err := json.Marshal(round)
	if err != nil {
		return fmt.Errorf("failed to marshal round: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// GetRound retrieves a round from Redis.
func (c *Client) GetRound(ctx context.Context, roundId string) (*Round, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil { // redislib.Nil を使用
		return nil, fmt.Errorf("round %s not found", roundId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get round %s: %w", roundId, err)
	}
	var round Round
	err = json.Unmarshal([]byte(data), &round)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal round: %w", err)
	}
	return &round, nil
}

// UpdateRound updates a round in Redis.
func (c *Client) UpdateRound(ctx context.Context, roundId string, round Round) error {
	round.UpdatedAt = time.Now()
	data, err := json.Marshal(round)
	if err != nil {
		return fmt.Errorf("failed to marshal round: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s", roundId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// --- Participant ---

// Participant represents a participant in a planning poker session.
type Participant struct {
	SessionId string    `json:"sessionId"`
	Name      string    `json:"name"`
	IsHost    bool      `json:"isHost"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateParticipant creates a new participant in Redis.
func (c *Client) CreateParticipant(ctx context.Context, participantId string, participant Participant) error {
	participant.CreatedAt = time.Now()
	participant.UpdatedAt = time.Now()
	data, err := json.Marshal(participant)
	if err != nil {
		return fmt.Errorf("failed to marshal participant: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// GetParticipant retrieves a participant from Redis.
func (c *Client) GetParticipant(ctx context.Context, participantId string) (*Participant, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil {
		return nil, fmt.Errorf("participant %s not found", participantId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get participant %s: %w", participantId, err)
	}
	var participant Participant
	err = json.Unmarshal([]byte(data), &participant)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal participant: %w", err)
	}
	return &participant, nil
}

// UpdateParticipant updates a participant in Redis.
func (c *Client) UpdateParticipant(ctx context.Context, participantId string, participant Participant) error {
	participant.UpdatedAt = time.Now()
	data, err := json.Marshal(participant)
	if err != nil {
		return fmt.Errorf("failed to marshal participant: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:participant:%s", participantId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// --- Vote ---

// Vote represents a vote in a planning poker round.
type Vote struct {
	RoundId       string    `json:"roundId"`
	ParticipantId string    `json:"participantId"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// CreateVote creates a new vote in Redis.
func (c *Client) CreateVote(ctx context.Context, voteId string, vote Vote) error {
	vote.CreatedAt = time.Now()
	vote.UpdatedAt = time.Now()
	data, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// GetVote retrieves a vote from Redis.
func (c *Client) GetVote(ctx context.Context, voteId string) (*Vote, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	data, err := c.client.Get(ctx, key).Result()
	if err == redislib.Nil {
		return nil, fmt.Errorf("vote %s not found", voteId)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get vote %s: %w", voteId, err)
	}
	var vote Vote
	err = json.Unmarshal([]byte(data), &vote)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal vote: %w", err)
	}
	return &vote, nil
}

// UpdateVote updates a vote in Redis.
func (c *Client) UpdateVote(ctx context.Context, voteId string, vote Vote) error {
	vote.UpdatedAt = time.Now()
	data, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote: %w", err)
	}
	key := fmt.Sprintf("web-toolbox:planning-poker:vote:%s", voteId)
	return c.client.Set(ctx, key, data, 0).Err()
}

// --- Session Participants ---

// AddParticipantToSession adds a participant to a session's participant list.
func (c *Client) AddParticipantToSession(ctx context.Context, sessionId, participantId string) error {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s:participants", sessionId)
	return c.client.SAdd(ctx, key, participantId).Err()
}

// GetParticipantsInSession retrieves all participants in a session.
func (c *Client) GetParticipantsInSession(ctx context.Context, sessionId string) ([]string, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:session:%s:participants", sessionId)
	return c.client.SMembers(ctx, key).Result()
}

// --- Round Votes ---

// AddVoteToRound adds a vote to a round's vote list.
func (c *Client) AddVoteToRound(ctx context.Context, roundId, voteId string) error {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s:votes", roundId)
	return c.client.SAdd(ctx, key, voteId).Err()
}

// GetVotesInRound retrieves all votes in a round.
func (c *Client) GetVotesInRound(ctx context.Context, roundId string) ([]string, error) {
	key := fmt.Sprintf("web-toolbox:planning-poker:round:%s:votes", roundId)
	return c.client.SMembers(ctx, key).Result()
}
