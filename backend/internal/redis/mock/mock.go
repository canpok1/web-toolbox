// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/redis/client.go
//
// Generated by this command:
//
//	mockgen -source=../internal/redis/client.go -destination=../internal/redis/mock/mock.go
//

// Package mock_redis is a generated GoMock package.
package mock_redis

import (
	context "context"
	reflect "reflect"

	redis "github.com/canpok1/web-toolbox/backend/internal/redis"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddParticipantToSession mocks base method.
func (m *MockClient) AddParticipantToSession(ctx context.Context, sessionId, participantId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddParticipantToSession", ctx, sessionId, participantId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddParticipantToSession indicates an expected call of AddParticipantToSession.
func (mr *MockClientMockRecorder) AddParticipantToSession(ctx, sessionId, participantId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddParticipantToSession", reflect.TypeOf((*MockClient)(nil).AddParticipantToSession), ctx, sessionId, participantId)
}

// AddVoteToRound mocks base method.
func (m *MockClient) AddVoteToRound(ctx context.Context, roundId, voteId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVoteToRound", ctx, roundId, voteId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVoteToRound indicates an expected call of AddVoteToRound.
func (mr *MockClientMockRecorder) AddVoteToRound(ctx, roundId, voteId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVoteToRound", reflect.TypeOf((*MockClient)(nil).AddVoteToRound), ctx, roundId, voteId)
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// CreateParticipant mocks base method.
func (m *MockClient) CreateParticipant(ctx context.Context, participantId string, participant redis.Participant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParticipant", ctx, participantId, participant)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateParticipant indicates an expected call of CreateParticipant.
func (mr *MockClientMockRecorder) CreateParticipant(ctx, participantId, participant any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParticipant", reflect.TypeOf((*MockClient)(nil).CreateParticipant), ctx, participantId, participant)
}

// CreateRound mocks base method.
func (m *MockClient) CreateRound(ctx context.Context, roundId string, round redis.Round) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRound", ctx, roundId, round)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRound indicates an expected call of CreateRound.
func (mr *MockClientMockRecorder) CreateRound(ctx, roundId, round any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRound", reflect.TypeOf((*MockClient)(nil).CreateRound), ctx, roundId, round)
}

// CreateSession mocks base method.
func (m *MockClient) CreateSession(ctx context.Context, sessionId string, session redis.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, sessionId, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockClientMockRecorder) CreateSession(ctx, sessionId, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockClient)(nil).CreateSession), ctx, sessionId, session)
}

// CreateVote mocks base method.
func (m *MockClient) CreateVote(ctx context.Context, voteId string, vote redis.Vote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVote", ctx, voteId, vote)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVote indicates an expected call of CreateVote.
func (mr *MockClientMockRecorder) CreateVote(ctx, voteId, vote any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVote", reflect.TypeOf((*MockClient)(nil).CreateVote), ctx, voteId, vote)
}

// GetParticipant mocks base method.
func (m *MockClient) GetParticipant(ctx context.Context, participantId string) (*redis.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipant", ctx, participantId)
	ret0, _ := ret[0].(*redis.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipant indicates an expected call of GetParticipant.
func (mr *MockClientMockRecorder) GetParticipant(ctx, participantId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipant", reflect.TypeOf((*MockClient)(nil).GetParticipant), ctx, participantId)
}

// GetParticipantsInSession mocks base method.
func (m *MockClient) GetParticipantsInSession(ctx context.Context, sessionId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantsInSession", ctx, sessionId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantsInSession indicates an expected call of GetParticipantsInSession.
func (mr *MockClientMockRecorder) GetParticipantsInSession(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantsInSession", reflect.TypeOf((*MockClient)(nil).GetParticipantsInSession), ctx, sessionId)
}

// GetRound mocks base method.
func (m *MockClient) GetRound(ctx context.Context, roundId string) (*redis.Round, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRound", ctx, roundId)
	ret0, _ := ret[0].(*redis.Round)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRound indicates an expected call of GetRound.
func (mr *MockClientMockRecorder) GetRound(ctx, roundId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRound", reflect.TypeOf((*MockClient)(nil).GetRound), ctx, roundId)
}

// GetSession mocks base method.
func (m *MockClient) GetSession(ctx context.Context, sessionId string) (*redis.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", ctx, sessionId)
	ret0, _ := ret[0].(*redis.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockClientMockRecorder) GetSession(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockClient)(nil).GetSession), ctx, sessionId)
}

// GetVote mocks base method.
func (m *MockClient) GetVote(ctx context.Context, voteId string) (*redis.Vote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVote", ctx, voteId)
	ret0, _ := ret[0].(*redis.Vote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVote indicates an expected call of GetVote.
func (mr *MockClientMockRecorder) GetVote(ctx, voteId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVote", reflect.TypeOf((*MockClient)(nil).GetVote), ctx, voteId)
}

// GetVoteIdByRoundIdAndParticipantId mocks base method.
func (m *MockClient) GetVoteIdByRoundIdAndParticipantId(ctx context.Context, roundId, participantId string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVoteIdByRoundIdAndParticipantId", ctx, roundId, participantId)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoteIdByRoundIdAndParticipantId indicates an expected call of GetVoteIdByRoundIdAndParticipantId.
func (mr *MockClientMockRecorder) GetVoteIdByRoundIdAndParticipantId(ctx, roundId, participantId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoteIdByRoundIdAndParticipantId", reflect.TypeOf((*MockClient)(nil).GetVoteIdByRoundIdAndParticipantId), ctx, roundId, participantId)
}

// GetVotesInRound mocks base method.
func (m *MockClient) GetVotesInRound(ctx context.Context, roundId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotesInRound", ctx, roundId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVotesInRound indicates an expected call of GetVotesInRound.
func (mr *MockClientMockRecorder) GetVotesInRound(ctx, roundId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotesInRound", reflect.TypeOf((*MockClient)(nil).GetVotesInRound), ctx, roundId)
}

// UpdateParticipant mocks base method.
func (m *MockClient) UpdateParticipant(ctx context.Context, participantId string, participant redis.Participant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParticipant", ctx, participantId, participant)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateParticipant indicates an expected call of UpdateParticipant.
func (mr *MockClientMockRecorder) UpdateParticipant(ctx, participantId, participant any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParticipant", reflect.TypeOf((*MockClient)(nil).UpdateParticipant), ctx, participantId, participant)
}

// UpdateRound mocks base method.
func (m *MockClient) UpdateRound(ctx context.Context, roundId string, round redis.Round) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRound", ctx, roundId, round)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRound indicates an expected call of UpdateRound.
func (mr *MockClientMockRecorder) UpdateRound(ctx, roundId, round any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRound", reflect.TypeOf((*MockClient)(nil).UpdateRound), ctx, roundId, round)
}

// UpdateSession mocks base method.
func (m *MockClient) UpdateSession(ctx context.Context, sessionId string, session redis.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", ctx, sessionId, session)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession.
func (mr *MockClientMockRecorder) UpdateSession(ctx, sessionId, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockClient)(nil).UpdateSession), ctx, sessionId, session)
}

// UpdateVote mocks base method.
func (m *MockClient) UpdateVote(ctx context.Context, voteId string, vote redis.Vote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVote", ctx, voteId, vote)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVote indicates an expected call of UpdateVote.
func (mr *MockClientMockRecorder) UpdateVote(ctx, voteId, vote any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVote", reflect.TypeOf((*MockClient)(nil).UpdateVote), ctx, voteId, vote)
}
