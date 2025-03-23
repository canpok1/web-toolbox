package api_test

import (
	"errors"
	"testing"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	mock_redis "github.com/canpok1/web-toolbox/backend/internal/redis/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/mock/gomock"
)

func TestHandlePostApiPlanningPokerSessions(t *testing.T) {
	tests := []struct {
		name          string
		req           *api.CreateSessionRequest
		mockSetup     func(mockRedis *mock_redis.MockClient)
		expectedRes   *api.CreateSessionResponse
		expectedError string
	}{
		{
			name: "success",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   "custom",
				CustomScale: &[]string{"1", "2", "3"},
			},
			mockSetup: func(mockRedis *mock_redis.MockClient) {
				mockRedis.EXPECT().CreateSession(mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedRes: &api.CreateSessionResponse{},
		},
		{
			name: "failure - invalid scale type",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   "invalid",
			},
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "invalid scaleType: invalid",
		},
		{
			name: "failure - missing session name",
			req: &api.CreateSessionRequest{
				HostName:  "Test Host",
				ScaleType: "fibonacci",
			},
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "sessionName is required",
		},
		{
			name: "failure - missing host name",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				ScaleType:   "fibonacci",
			},
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "hostName is required",
		},
		{
			name: "failure - missing scale type",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
			},
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "scaleType is required",
		},
		{
			name: "failure - missing custom scale",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   "custom",
				CustomScale: &[]string{},
			},
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "customScale is required when scaleType is custom",
		},
		{
			name: "failure - redis error",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   "custom",
				CustomScale: &[]string{"1", "2", "3"},
			},
			mockSetup: func(mockRedis *mock_redis.MockClient) {
				mockRedis.EXPECT().CreateSession(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("redis error"))
			},
			expectedError: "failed to save session to redis",
		},
		{
			name:          "failure - nil request body",
			req:           nil,
			mockSetup:     func(mockRedis *mock_redis.MockClient) {},
			expectedError: "request body is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			tt.mockSetup(mockRedis)

			res, err := server.HandlePostApiPlanningPokerSessions(tt.req)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.SessionId, "Expected SessionId to be non-empty")
				assert.NotEmpty(t, res.HostId, "Expected HostId to be non-empty")
			}
			ctrl.Finish()

		})
	}
}

func TestHandlePostApiPlanningPokerSessionsSessionIdParticipants(t *testing.T) {
	tests := []struct {
		name          string
		sessionId     uuid.UUID
		req           *api.JoinSessionRequest
		mockSetup     func(mockRedis *mock_redis.MockClient, sessionId uuid.UUID)
		expectedRes   *api.JoinSessionResponse
		expectedError string
	}{
		{
			name:      "success",
			sessionId: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionId uuid.UUID) {
				mockRedis.EXPECT().GetSession(mock.Anything, sessionId.String()).Return(&redis.Session{}, nil)
				mockRedis.EXPECT().CreateParticipant(mock.Anything, mock.Anything, mock.Anything).Return(nil)
				mockRedis.EXPECT().AddParticipantToSession(mock.Anything, sessionId.String(), mock.Anything).Return(nil)
				mockRedis.EXPECT().UpdateSession(mock.Anything, sessionId.String(), mock.Anything).Return(nil)

			},
			req: &api.JoinSessionRequest{
				Name: "Test Participant",
			},
			expectedRes: &api.JoinSessionResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)
			tt.mockSetup(mockRedis, tt.sessionId)

			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdParticipants(tt.sessionId, tt.req)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.ParticipantId, "Expected ParticipantId to be non-empty")
			}
			ctrl.Finish()
			//mockRedis.AssertExpectations(t)
		})
	}
}

func TestHandlePostApiPlanningPokerRoundsRoundIdReveal(t *testing.T) {
	tests := []struct {
		name          string
		roundId       uuid.UUID
		expectedRes   *api.RevealRoundResponse
		expectedError string
	}{
		{
			name:          "not implemented",
			roundId:       uuid.New(),
			expectedError: "HandlePostApiPlanningPokerRoundsRoundIdReveal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			res, err := server.HandlePostApiPlanningPokerRoundsRoundIdReveal(tt.roundId)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
			}
			ctrl.Finish()
		})
	}
}

func TestHandlePostApiPlanningPokerRoundsRoundIdVotes(t *testing.T) {
	tests := []struct {
		name          string
		roundId       uuid.UUID
		req           *api.SendVoteRequest
		expectedRes   *api.SendVoteResponse
		expectedError string
	}{
		{
			name:    "success",
			roundId: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "1",
			},
			expectedRes: &api.SendVoteResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			res, err := server.HandlePostApiPlanningPokerRoundsRoundIdVotes(tt.roundId, tt.req)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.VoteId, "Expected VoteId to be non-empty")
			}
			ctrl.Finish()
		})
	}
}

func TestHandleGetApiPlanningPokerSessionsSessionId(t *testing.T) {
	tests := []struct {
		name          string
		sessionId     uuid.UUID
		expectedRes   *api.GetSessionResponse
		expectedError string
	}{
		{
			name:          "not implemented",
			sessionId:     uuid.New(),
			expectedError: "GetApiPlanningPokerSessionsSessionId",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			res, err := server.HandleGetApiPlanningPokerSessionsSessionId(tt.sessionId)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
			}
			ctrl.Finish()
		})
	}
}

func TestHandlePostApiPlanningPokerSessionsSessionIdEnd(t *testing.T) {
	tests := []struct {
		name          string
		sessionId     uuid.UUID
		expectedRes   *api.EndSessionResponse
		expectedError string
	}{
		{
			name:          "not implemented",
			sessionId:     uuid.New(),
			expectedError: "PostApiPlanningPokerSessionsSessionIdEnd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdEnd(tt.sessionId)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
			}
			ctrl.Finish()
		})
	}
}

func TestHandlePostApiPlanningPokerSessionsSessionIdRounds(t *testing.T) {
	tests := []struct {
		name          string
		sessionId     uuid.UUID
		expectedRes   *api.StartRoundResponse
		expectedError string
	}{
		{
			name:        "success",
			sessionId:   uuid.New(),
			expectedRes: &api.StartRoundResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdRounds(tt.sessionId)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.RoundId, "Expected RoundId to be non-empty")
			}
			ctrl.Finish()
		})
	}
}
