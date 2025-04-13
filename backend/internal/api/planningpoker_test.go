package api_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	mock_planningpoker "github.com/canpok1/web-toolbox/backend/internal/api/planningpoker/mock"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	mock_redis "github.com/canpok1/web-toolbox/backend/internal/redis/mock"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
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
				HostName:    "Test Host",
				ScaleType:   api.Custom,
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "",
		},
		{
			name: "failure - invalid scale type",
			req: &api.CreateSessionRequest{
				HostName:    "Test Host",
				ScaleType:   "invalid",
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "invalid scaleType: invalid",
		},
		{
			name: "failure - missing host name",
			req: &api.CreateSessionRequest{
				ScaleType:   api.Fibonacci,
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "hostName is required",
		},
		{
			name: "failure - missing scale type",
			req: &api.CreateSessionRequest{
				HostName:    "Test Host",
				CustomScale: &[]string{"1", "2", "3"},
			},
			expectedError: "scaleType is required",
		},
		{
			name: "failure - missing custom scale",
			req: &api.CreateSessionRequest{
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
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

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
	t.Run("異常系", func(t *testing.T) {
		type MockSetting struct {
			createSessionResult           error
			createParticipantResult       error
			addParticipantToSessionResult error
		}

		tests := []struct {
			name          string
			req           *api.CreateSessionRequest
			mockSetting   MockSetting
			expectedError string
		}{
			{
				name: "セッション作成失敗",
				req: &api.CreateSessionRequest{
					HostName:    "Test Host",
					ScaleType:   api.Custom,
					CustomScale: &[]string{"1", "2", "3"},
				},
				mockSetting: MockSetting{
					createSessionResult: errors.New("redis error"),
				},
				expectedError: "failed to save session to redis",
			},
			{
				name: "参加者作成失敗",
				req: &api.CreateSessionRequest{
					HostName:    "Test Host",
					ScaleType:   api.Custom,
					CustomScale: &[]string{"1", "2", "3"},
				},
				mockSetting: MockSetting{
					createParticipantResult: errors.New("redis error"),
				},
				expectedError: "failed to save participant to redis",
			},
			{
				name: "参加者リストへの追加失敗",
				req: &api.CreateSessionRequest{
					HostName:    "Test Host",
					ScaleType:   api.Custom,
					CustomScale: &[]string{"1", "2", "3"},
				},
				mockSetting: MockSetting{
					addParticipantToSessionResult: errors.New("redis error"),
				},
				expectedError: "failed to add participant list to redis",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				mockRedis.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockSetting.createSessionResult).AnyTimes()
				mockRedis.EXPECT().CreateParticipant(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockSetting.createParticipantResult).AnyTimes()
				mockRedis.EXPECT().AddParticipantToSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockSetting.addParticipantToSessionResult).AnyTimes()

				res, err := server.HandlePostApiPlanningPokerSessions(tt.req)

				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain: "+tt.expectedError)
			})
		}
	})

	t.Run("正常系", func(t *testing.T) {
		tests := []struct {
			name string
			req  *api.CreateSessionRequest
		}{
			{
				name: "success",
				req: &api.CreateSessionRequest{
					HostName:    "Test Host",
					ScaleType:   api.Custom,
					CustomScale: &[]string{"1", "2", "3"},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				mockRedis.EXPECT().CreateSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().CreateParticipant(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().AddParticipantToSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				res, err := server.HandlePostApiPlanningPokerSessions(tt.req)

				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.SessionId, "Expected SessionId to be non-empty")
				assert.NotEmpty(t, res.HostId, "Expected HostId to be non-empty")
			})
		}
	})
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
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{}, nil)
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
			name:      "failure - get session error",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, errors.New("get session error"))
			},
			expectedError: "failed to get session",
		},
		{
			name:      "failure - session not found error",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, nil)
			},
			expectedError: "session is not found",
		},
		{
			name:      "failure - create participant error",
			sessionID: uuid.New(),
			req: &api.JoinSessionRequest{
				Name: "Test User",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{}, nil)
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
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{}, nil)
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
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

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

	t.Run("異常系", func(t *testing.T) {
		tests := []struct {
			name                  string
			sessionID             uuid.UUID
			getSessionReturnValue *redis.Session
			getSessionReturnError error
			expectedError         string
		}{
			{
				name:                  "セッションがnil",
				sessionID:             uuid.New(),
				getSessionReturnValue: nil,
				getSessionReturnError: nil,
				expectedError:         "session not found",
			},
			{
				name:                  "セッション取得失敗",
				sessionID:             uuid.New(),
				getSessionReturnValue: nil,
				getSessionReturnError: errors.New("redis error"),
				expectedError:         "failed to get session from redis",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				mockRedis.EXPECT().GetSession(gomock.Any(), tt.sessionID.String()).Return(tt.getSessionReturnValue, tt.getSessionReturnError)

				res, err := server.HandleGetApiPlanningPokerSessionsSessionId(tt.sessionID)

				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain")
			})
		}
	})

	t.Run("正常系", func(t *testing.T) {
		dummyUUID := uuid.New()
		tests := []struct {
			name           string
			sessionID      uuid.UUID
			hostID         string
			participantIDs []string
		}{
			{
				name:           "参加者はホストのみ",
				sessionID:      dummyUUID,
				hostID:         "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				participantIDs: []string{"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"},
			},
			{
				name:      "参加者はホスト以外にも複数",
				sessionID: dummyUUID,
				hostID:    "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
				participantIDs: []string{
					"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
					"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
					"cccccccc-cccc-cccc-cccc-cccccccccccc"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				mockRedis.EXPECT().GetSession(gomock.Any(), tt.sessionID.String()).Return(&redis.Session{
					HostId:         tt.hostID,
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipantsInSession(gomock.Any(), tt.sessionID.String()).Return(tt.participantIDs, nil)
				for _, participantID := range tt.participantIDs {
					mockRedis.EXPECT().GetParticipant(gomock.Any(), participantID).Return(&redis.Participant{}, nil)
				}

				res, err := server.HandleGetApiPlanningPokerSessionsSessionId(tt.sessionID)

				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.Equal(t, tt.sessionID, res.Session.SessionId)
				assert.Equal(t, len(tt.participantIDs), len(res.Session.Participants))
			})
		}
	})

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
			expectedError: "session not found",
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
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tt.sessionID != uuid.Nil {
				sessionID = tt.sessionID
			}

			tt.mockSetup(mockRedis, sessionID)

			ctx := context.Background()
			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdEnd(ctx, sessionID)

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

func TestHandlePostApiPlanningPokerSessionsSessionIdRounds(t *testing.T) {
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
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().CreateRound(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().UpdateSession(gomock.Any(), sessionID.String(), gomock.Any()).Return(nil)
			},
		},
		{
			name:      "failure - session not found",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(nil, nil)
			},
			expectedError: "session not found",
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
			name:      "failure - redis create round error",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().CreateRound(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to create round in redis",
		},
		{
			name:      "failure - redis update session error",
			sessionID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, sessionID uuid.UUID) {
				mockRedis.EXPECT().GetSession(gomock.Any(), sessionID.String()).Return(&redis.Session{
					HostId:         uuid.New().String(),
					ScaleType:      "fibonacci",
					CustomScale:    []string{},
					CurrentRoundId: "",
					Status:         "waiting",
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				}, nil)
				mockRedis.EXPECT().CreateRound(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
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
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

			// Create a new sessionID for each test case to avoid conflicts
			sessionID := uuid.New()
			if tt.sessionID != uuid.Nil {
				sessionID = tt.sessionID
			}

			tt.mockSetup(mockRedis, sessionID)

			ctx := context.Background()
			res, err := server.HandlePostApiPlanningPokerSessionsSessionIdRounds(ctx, sessionID)

			if tt.expectedError == "" {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.RoundId, "Expected RoundId to be non-empty")
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain")
			}
		})
	}
}

func TestHandleGetApiPlanningPokerRoundsRoundId(t *testing.T) {
	ctx := context.Background()
	now := time.Now()

	t.Run("正常系", func(t *testing.T) {
		validSessionID := uuid.New()
		validRoundID := uuid.New()
		validParticipantID1 := uuid.New()
		validParticipantID2 := uuid.New()
		voteID1 := uuid.New().String()
		voteID2 := uuid.New().String()

		tests := []struct {
			name             string
			roundID          openapi_types.UUID
			participantID    *openapi_types.UUID
			mockSetup        func(mockRedis *mock_redis.MockClient)
			expectedResponse *api.GetRoundResponse
			expectedError    string
		}{
			{
				name:    "成功 - ステータス voting",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Voting), // "voting"
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return([]string{}, nil)
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Voting,
						CreatedAt: now,
						UpdatedAt: now,
						Votes:     []api.Vote{},
					},
				},
			},
			{
				name:          "成功 - ステータス voting, 投票あり, 参加者指定あり",
				roundID:       openapi_types.UUID(validRoundID),
				participantID: &validParticipantID1,
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Voting), // "voting"
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return([]string{voteID1, voteID2}, nil)
					mockRedis.EXPECT().GetVote(ctx, voteID1).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID1.String(),
						Value:         "5",
						CreatedAt:     now,
						UpdatedAt:     now,
					}, nil)
					mockRedis.EXPECT().GetVote(ctx, voteID2).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID2.String(),
						Value:         "8",
						CreatedAt:     now,
						UpdatedAt:     now,
					}, nil)
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Voting,
						CreatedAt: now,
						UpdatedAt: now,
						Votes: []api.Vote{
							{ParticipantId: openapi_types.UUID(validParticipantID1), Value: PtrString("5")},
							{ParticipantId: openapi_types.UUID(validParticipantID2)},
						},
					},
				},
			},
			{
				name:    "成功 - ステータス revealed, 投票あり",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Revealed), // "revealed"
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return([]string{voteID1, voteID2}, nil)
					mockRedis.EXPECT().GetVote(ctx, voteID1).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID1.String(),
						Value:         "5",
						CreatedAt:     now,
						UpdatedAt:     now,
					}, nil)
					mockRedis.EXPECT().GetVote(ctx, voteID2).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID2.String(),
						Value:         "8",
						CreatedAt:     now,
						UpdatedAt:     now,
					}, nil)
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Revealed,
						CreatedAt: now,
						UpdatedAt: now,
						Votes: []api.Vote{
							{ParticipantId: openapi_types.UUID(validParticipantID1), Value: PtrString("5")},
							{ParticipantId: openapi_types.UUID(validParticipantID2), Value: PtrString("8")},
						},
					},
				},
			},
			{
				name:    "成功 - ステータス revealed, 投票なし",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Revealed),
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return([]string{}, nil) // 空のリストを返す
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Revealed,
						CreatedAt: now,
						UpdatedAt: now,
						Votes:     []api.Vote{}, // 空のスライスが返る
					},
				},
			},
			{
				name:    "成功 - ステータス revealed, GetVotesInRound エラー (ログ出力のみ)",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Revealed),
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return(nil, errors.New("redis error")) // エラーを返す
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Revealed,
						CreatedAt: now,
						UpdatedAt: now,
						Votes:     []api.Vote{},
					},
				},
			},
			{
				name:    "成功 - ステータス revealed, 一部の GetVote エラー/nil/パースエラー (スキップされる)",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					invalidParticipantID := "invalid-uuid"
					voteID3 := uuid.New().String()
					voteID4 := uuid.New().String()
					voteID5 := uuid.New().String()

					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: validSessionID.String(),
						Status:    string(api.Revealed),
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
					mockRedis.EXPECT().GetVotesInRound(ctx, validRoundID.String()).Return([]string{voteID1, voteID2, voteID3, voteID4, voteID5}, nil)

					// 正常な投票
					mockRedis.EXPECT().GetVote(ctx, voteID1).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID1.String(),
						Value:         "5",
					}, nil)
					// GetVote エラー
					mockRedis.EXPECT().GetVote(ctx, voteID2).Return(nil, errors.New("get vote error"))
					// GetVote nil
					mockRedis.EXPECT().GetVote(ctx, voteID3).Return(nil, nil)
					// ParticipantId パースエラー
					mockRedis.EXPECT().GetVote(ctx, voteID4).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: invalidParticipantID, // 不正なUUID
						Value:         "13",
					}, nil)
					// 正常な投票 (エラーの後でも処理される)
					mockRedis.EXPECT().GetVote(ctx, voteID5).Return(&redis.Vote{
						RoundId:       validRoundID.String(),
						ParticipantId: validParticipantID2.String(),
						Value:         "21",
					}, nil)
				},
				expectedResponse: &api.GetRoundResponse{
					Round: api.Round{
						RoundId:   openapi_types.UUID(validRoundID),
						SessionId: openapi_types.UUID(validSessionID),
						Status:    api.Revealed,
						CreatedAt: now,
						UpdatedAt: now,
						Votes: []api.Vote{ // エラー/nil/パースエラーの投票は含まれない
							{ParticipantId: openapi_types.UUID(validParticipantID1), Value: PtrString("5")},
							{ParticipantId: openapi_types.UUID(validParticipantID2), Value: PtrString("21")},
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				tt.mockSetup(mockRedis)

				res, err := server.HandleGetApiPlanningPokerRoundsRoundId(ctx, tt.roundID, tt.participantID)

				if tt.expectedError == "" {
					assert.NoError(t, err)
					assert.Equal(t, tt.expectedResponse, res)
				} else {
					assert.Error(t, err)
					assert.Nil(t, res)
					assert.Contains(t, err.Error(), tt.expectedError)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		invalidRoundID := uuid.New()
		validRoundID := uuid.New()

		tests := []struct {
			name          string
			roundID       openapi_types.UUID
			participantID *openapi_types.UUID
			mockSetup     func(mockRedis *mock_redis.MockClient)
			expectedError string
		}{
			{
				name:    "失敗 - GetRound エラー",
				roundID: openapi_types.UUID(invalidRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, invalidRoundID.String()).Return(nil, errors.New("redis connection error"))
				},
				expectedError: "failed to get round from redis", // エラーメッセージは抽象化される想定
			},
			{
				name:    "失敗 - Round Not Found",
				roundID: openapi_types.UUID(invalidRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, invalidRoundID.String()).Return(nil, nil) // nil, nil を返す
				},
				expectedError: "round not found",
			},
			{
				name:    "失敗 - SessionId パースエラー",
				roundID: openapi_types.UUID(validRoundID),
				mockSetup: func(mockRedis *mock_redis.MockClient) {
					mockRedis.EXPECT().GetRound(ctx, validRoundID.String()).Return(&redis.Round{
						SessionId: "invalid-session-id-format", // 不正な形式
						Status:    string(api.Voting),
						CreatedAt: now,
						UpdatedAt: now,
					}, nil)
				},
				expectedError: "invalid session ID format",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()
				mockRedis := mock_redis.NewMockClient(ctrl)
				mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
				server := api.NewServer(mockRedis, mockWebSocketHub)

				tt.mockSetup(mockRedis)

				res, err := server.HandleGetApiPlanningPokerRoundsRoundId(ctx, tt.roundID, tt.participantID)

				assert.Error(t, err)
				assert.Nil(t, res)
				assert.Contains(t, err.Error(), tt.expectedError)
			})
		}
	})
}

func TestHandlePostApiPlanningPokerRoundsRoundIdReveal(t *testing.T) {
	tests := []struct {
		name          string
		roundID       uuid.UUID
		mockSetup     func(mockRedis *mock_redis.MockClient, roundID uuid.UUID)
		expectedError string
	}{
		{
			name:    "success",
			roundID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateRound(gomock.Any(), roundID.String(), gomock.Any()).Return(nil)
			},
		},
		{
			name:    "failure - round not found",
			roundID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(nil, errors.New("round not found"))
			},
			expectedError: "failed to get round from redis",
		},
		{
			name:    "failure - redis get round error",
			roundID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get round from redis",
		},
		{
			name:    "failure - redis update round error",
			roundID: uuid.New(),
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateRound(gomock.Any(), roundID.String(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to update round in redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

			// Create a new roundID for each test case to avoid conflicts
			roundID := uuid.New()
			if tt.roundID != uuid.Nil {
				roundID = tt.roundID
			}

			tt.mockSetup(mockRedis, roundID)

			ctx := context.Background()
			res, err := server.HandlePostApiPlanningPokerRoundsRoundIdReveal(ctx, roundID)

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

func TestHandlePostApiPlanningPokerRoundsRoundIdVotes(t *testing.T) {
	tests := []struct {
		name          string
		roundID       uuid.UUID
		req           *api.SendVoteRequest
		mockSetup     func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest)
		expectedError string
	}{
		{
			name:    "success - first time vote",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(nil, nil)
				mockRedis.EXPECT().CreateVote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().AddVoteToRound(gomock.Any(), roundID.String(), gomock.Any()).Return(nil)
			},
		},
		{
			name:    "success - multiple votes",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "8",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				voteId := uuid.New().String()
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(&voteId, nil)
				mockRedis.EXPECT().GetVote(gomock.Any(), voteId).Return(&redis.Vote{
					RoundId:       roundID.String(),
					ParticipantId: req.ParticipantId.String(),
					Value:         "5",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateVote(gomock.Any(), voteId, gomock.Any()).Return(nil)
			},
		},
		{
			name:    "failure - nil request body",
			roundID: uuid.New(),
			req:     nil,
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
			},
			expectedError: "request body is required",
		},
		{
			name:    "failure - round not found",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(nil, nil)
			},
			expectedError: "round not found",
		},
		{
			name:    "failure - redis get round error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get round from redis",
		},
		{
			name:    "failure - round not in voting state",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "revealed",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			expectedError: "round is not in voting state",
		},
		{
			name:    "failure - participant not found",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(nil, nil)
			},
			expectedError: "participant not found",
		},
		{
			name:    "failure - redis get participant error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get participant from redis",
		},
		{
			name:    "failure - redis get vote id error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get vote id from redis",
		},
		{
			name:    "failure - redis create vote error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(nil, nil)
				mockRedis.EXPECT().CreateVote(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to create vote in redis",
		},
		{
			name:    "failure - redis add vote to round error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "5",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(nil, nil)
				mockRedis.EXPECT().CreateVote(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				mockRedis.EXPECT().AddVoteToRound(gomock.Any(), roundID.String(), gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to add vote to round in redis",
		},
		{
			name:    "failure - redis get vote error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "8",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				voteId := uuid.New().String()
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(&voteId, nil)
				mockRedis.EXPECT().GetVote(gomock.Any(), voteId).Return(nil, errors.New("redis error"))
			},
			expectedError: "failed to get vote from redis",
		},
		{
			name:    "failure - redis update vote error",
			roundID: uuid.New(),
			req: &api.SendVoteRequest{
				ParticipantId: uuid.New(),
				Value:         "8",
			},
			mockSetup: func(mockRedis *mock_redis.MockClient, roundID uuid.UUID, req *api.SendVoteRequest) {
				voteId := uuid.New().String()
				mockRedis.EXPECT().GetRound(gomock.Any(), roundID.String()).Return(&redis.Round{
					SessionId: uuid.New().String(),
					Status:    "voting",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
				mockRedis.EXPECT().GetParticipant(gomock.Any(), req.ParticipantId.String()).Return(&redis.Participant{}, nil)
				mockRedis.EXPECT().GetVoteIdByRoundIdAndParticipantId(gomock.Any(), roundID.String(), req.ParticipantId.String()).Return(&voteId, nil)
				mockRedis.EXPECT().GetVote(gomock.Any(), voteId).Return(&redis.Vote{
					RoundId:       roundID.String(),
					ParticipantId: req.ParticipantId.String(),
					Value:         "5",
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}, nil)
				mockRedis.EXPECT().UpdateVote(gomock.Any(), voteId, gomock.Any()).Return(errors.New("redis error"))
			},
			expectedError: "failed to update vote in redis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRedis := mock_redis.NewMockClient(ctrl)
			mockWebSocketHub := mock_planningpoker.NewMockWebSocketHub(ctrl)
			server := api.NewServer(mockRedis, mockWebSocketHub)

			// Create a new roundID for each test case to avoid conflicts
			roundID := uuid.New()
			if tt.roundID != uuid.Nil {
				roundID = tt.roundID
			}

			tt.mockSetup(mockRedis, roundID, tt.req)

			ctx := context.Background()
			res, err := server.HandlePostApiPlanningPokerRoundsRoundIdVotes(ctx, roundID, tt.req)

			if tt.expectedError == "" {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, res, "Expected non-nil response on success")
				assert.NotEmpty(t, res.VoteId, "Expected VoteId to be non-empty")
			} else {
				assert.Error(t, err, "Expected an error")
				assert.Nil(t, res, "Expected nil response on error")
				assert.Contains(t, err.Error(), tt.expectedError, "Expected error message to contain")
			}
		})
	}
}

// PtrString は string のポインタを返すヘルパー関数です。
func PtrString(s string) *string {
	return &s
}
