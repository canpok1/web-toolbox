package api_test

import (
	"errors"
	"testing"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	mock_redis "github.com/canpok1/web-toolbox/backend/internal/redis/mock"
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
