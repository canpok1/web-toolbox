package api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	mock_redis "github.com/canpok1/web-toolbox/backend/internal/redis/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestValidatePostApiPlanningPokerSessions(t *testing.T) {
	tests := []struct {
		name          string
		req           *api.CreateSessionRequest
		expectedError string
	}{
		{
			name: "success",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   api.Custom,
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "",
		},
		{
			name: "failure - invalid scale type",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   "invalid",
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "invalid scaleType: invalid",
		},
		{
			name: "failure - missing session name",
			req: &api.CreateSessionRequest{
				HostName:    "Test Host",
				ScaleType:   api.Fibonacci,
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "sessionName is required",
		},
		{
			name: "failure - missing host name",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				ScaleType:   api.Fibonacci,
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "hostName is required",
		},
		{
			name: "failure - missing scale type",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "scaleType is required",
		},
		{
			name: "failure - missing custom scale",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   api.Custom,
				CustomScale: &[]string{},
			},
			expectedError: "customScale is required when scaleType is custom",
		},
		{
			name:          "failure - nil request body",
			req:           nil,
			expectedError: "request body is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			err := server.ValidatePostApiPlanningPokerSessions(tt.req)

			if tt.expectedError != "" {
				assert.Error(t, err, "Expected an error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}

func TestHandlePostApiPlanningPokerSessions(t *testing.T) {
	tests := []struct {
		name          string
		req           *api.CreateSessionRequest
		mockSetup     func(mockRedis *mock_redis.MockClient)
		expectedError string
	}{
		{
			name: "success",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   api.Custom,
				CustomScale: &[]string{"1", "2", "3"},
			},
			mockSetup: func(mockRedis *mock_redis.MockClient) {
				mockRedis.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "failure - redis error",
			req: &api.CreateSessionRequest{
				SessionName: "Test Session",
				HostName:    "Test Host",
				ScaleType:   api.Custom,
				CustomScale: &[]string{"1", "2", "3"},
			},
			mockSetup: func(mockRedis *mock_redis.MockClient) {
				mockRedis.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to save session to redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
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
		})
	}
}

func TestValidatePostApiPlanningPokerSessionsSessionIdParticipants(t *testing.T) {
	// Test cases
	testCases := []struct {
		name        string
		sessionID   uuid.UUID
		body        *api.JoinSessionRequest
		expectedErr string
	}{
		{
			name:        "Valid request",
			sessionID:   uuid.New(),
			body:        &api.JoinSessionRequest{Name: "Test User"},
			expectedErr: "",
		},
		{
			name:        "Nil request body",
			sessionID:   uuid.New(),
			body:        nil,
			expectedErr: "request body is required",
		},
		{
			name:        "Empty name",
			sessionID:   uuid.New(),
			body:        &api.JoinSessionRequest{Name: ""},
			expectedErr: "name is required",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tc.sessionID != uuid.Nil {
				sessionID = tc.sessionID
			}

			// Call the function
			s := &api.Server{} // Assuming you have a Server struct
			err := s.ValidatePostApiPlanningPokerSessionsSessionIdParticipants(sessionID, tc.body)

			// Check for errors
			if tc.expectedErr == "" && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
			if tc.expectedErr != "" && err == nil {
				t.Errorf("Expected error: %v, but got nil", tc.expectedErr)
			}
			if tc.expectedErr != "" && err != nil {
				assert.Contains(t, err.Error(), tc.expectedErr)
			}
		})
	}
}

func TestHandlePostApiPlanningPokerSessionsSessionIdParticipants(t *testing.T) {
	tests := []struct {
		name          string
		sessionID     uuid.UUID
		req           *api.JoinSessionRequest
		mockSetup     func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID)
		expectedError string
	}{
		{
			name:      "success",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().CreateParticipant(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().AddParticipantToSession(gomock.Any(), sessionID.String(), gomock.Any()).Return(nil)
			},
		},
		{
			name:          "failure - nil request body",
			sessionID:     uuid.New(),
			req:           nil,
			mockSetup:     func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {},
			expectedError: "request body is required (sessionID:",
		},
		{
			name:      "failure - create participant error",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().CreateParticipant(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("create participant error"))
			},
			expectedError: "failed to create participant",
		},
		{
			name:      "failure - add participant to session error",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().CreateParticipant(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().AddParticipantToSession(gomock.Any(), sessionID.String(), gomock.Any()).Return(errors.New("add participant to session error"))
			},
			expectedError: "failed to add participant to session",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tt.sessionID != uuid.Nil {
				sessionID = tt.sessionID
			}

			tt.mockSetup(mockRedis, sessionID)

			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdParticipants(sessionID, tt.req)

			if tt.expectedError == "" {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.ParticipantId, "Expected ParticipantId to be non-empty")
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			}
		})
	}
}

func TestHandleGetApiPlanningPokerSessionsSessionId(t *testing.T) {
	tests := []struct {
		name          string
		sessionID     uuid.UUID
		mockSetup     func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID)
		expectedError string
	}{
		{
			name:      "success",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					SessionName:    "Test Session",
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
		},
		{
			name:      "success with current round",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					SessionName:    "Test Session",
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: uuid.New().String(),
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
			},
		},
		{
			name:      "failure - session not found",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, nil)
			},
			expectedError: "session not found (sessionID:",
		},
		{
			name:      "failure - redis error",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get session from redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tt.sessionID != uuid.Nil {
				sessionID = tt.sessionID
			}

			tt.mockSetup(mockRedis, sessionID)

			res, err := server.HandleGetApiPlanningPokerSessionsSessionId(sessionID)

			if tt.expectedError == "" {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.Equal(t, sessionID, res.SessionId)
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain")
			}
		})
	}
}

func TestHandlePostApiPlanningPokerSessionsSessionIdEnd(t *testing.T) {
	tests := []struct {
		name          string
		sessionID     uuid.UUID
		mockSetup     func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID)
		expectedError string
	}{
		{
			name:      "success",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					SessionName:    "Test Session",
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateSession(gomock.Any(), sessionID.String(), gomock.Any()).Return(nil)
			},
		},
		{
			name:      "failure - session not found",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, nil)
			},
			expectedError: "session not found (sessionID:",
		},
		{
			name:      "failure - redis get session error",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get session from redis",
		},
		{
			name:      "failure - redis update session error",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					SessionName:    "Test Session",
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateSession(gomock.Any(), sessionID.String(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to update session in redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			server := api.NewServer(mockRedis)

			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tt.sessionID != uuid.Nil {
				sessionID = tt.sessionID
			}

			tt.mockSetup(mockRedis, sessionID)

			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdEnd(sessionID)

			if tt.expectedError == "" {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain")
			}
		})
	}
}
