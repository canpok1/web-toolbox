// Code generated by MockGen. DO NOT EDIT.
// Source: ../websocket.go
//
// Generated by this command:
//
//	mockgen -source=../websocket.go -destination=./mock.go
//

// Package mock_planningpoker is a generated GoMock package.
package mock_planningpoker

import (
	reflect "reflect"

	planningpoker "github.com/canpok1/web-toolbox/backend/internal/api/planningpoker"
	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockWebSocketHub is a mock of WebSocketHub interface.
type MockWebSocketHub struct {
	ctrl     *gomock.Controller
	recorder *MockWebSocketHubMockRecorder
	isgomock struct{}
}

// MockWebSocketHubMockRecorder is the mock recorder for MockWebSocketHub.
type MockWebSocketHubMockRecorder struct {
	mock *MockWebSocketHub
}

// NewMockWebSocketHub creates a new mock instance.
func NewMockWebSocketHub(ctrl *gomock.Controller) *MockWebSocketHub {
	mock := &MockWebSocketHub{ctrl: ctrl}
	mock.recorder = &MockWebSocketHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebSocketHub) EXPECT() *MockWebSocketHubMockRecorder {
	return m.recorder
}

// BroadcastParticipantJoined mocks base method.
func (m *MockWebSocketHub) BroadcastParticipantJoined(participantId, name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastParticipantJoined", participantId, name)
}

// BroadcastParticipantJoined indicates an expected call of BroadcastParticipantJoined.
func (mr *MockWebSocketHubMockRecorder) BroadcastParticipantJoined(participantId, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastParticipantJoined", reflect.TypeOf((*MockWebSocketHub)(nil).BroadcastParticipantJoined), participantId, name)
}

// BroadcastRoundStarted mocks base method.
func (m *MockWebSocketHub) BroadcastRoundStarted(roundId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastRoundStarted", roundId)
}

// BroadcastRoundStarted indicates an expected call of BroadcastRoundStarted.
func (mr *MockWebSocketHubMockRecorder) BroadcastRoundStarted(roundId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastRoundStarted", reflect.TypeOf((*MockWebSocketHub)(nil).BroadcastRoundStarted), roundId)
}

// BroadcastSessionEnded mocks base method.
func (m *MockWebSocketHub) BroadcastSessionEnded() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastSessionEnded")
}

// BroadcastSessionEnded indicates an expected call of BroadcastSessionEnded.
func (mr *MockWebSocketHubMockRecorder) BroadcastSessionEnded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastSessionEnded", reflect.TypeOf((*MockWebSocketHub)(nil).BroadcastSessionEnded))
}

// BroadcastVoteSubmitted mocks base method.
func (m *MockWebSocketHub) BroadcastVoteSubmitted(participantId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastVoteSubmitted", participantId)
}

// BroadcastVoteSubmitted indicates an expected call of BroadcastVoteSubmitted.
func (mr *MockWebSocketHubMockRecorder) BroadcastVoteSubmitted(participantId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastVoteSubmitted", reflect.TypeOf((*MockWebSocketHub)(nil).BroadcastVoteSubmitted), participantId)
}

// BroadcastVotesRevealed mocks base method.
func (m *MockWebSocketHub) BroadcastVotesRevealed(votes []planningpoker.Vote, average, median float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastVotesRevealed", votes, average, median)
}

// BroadcastVotesRevealed indicates an expected call of BroadcastVotesRevealed.
func (mr *MockWebSocketHubMockRecorder) BroadcastVotesRevealed(votes, average, median any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastVotesRevealed", reflect.TypeOf((*MockWebSocketHub)(nil).BroadcastVotesRevealed), votes, average, median)
}

// HandleWebSocket mocks base method.
func (m *MockWebSocketHub) HandleWebSocket(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleWebSocket", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleWebSocket indicates an expected call of HandleWebSocket.
func (mr *MockWebSocketHubMockRecorder) HandleWebSocket(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleWebSocket", reflect.TypeOf((*MockWebSocketHub)(nil).HandleWebSocket), c)
}

// Run mocks base method.
func (m *MockWebSocketHub) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockWebSocketHubMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWebSocketHub)(nil).Run))
}
