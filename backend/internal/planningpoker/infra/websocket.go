package infra

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// WebSocketMessage represents the structure of a WebSocket message.
type WebSocketMessage struct {
	Event   string      `json:"event"`
	Payload interface{} `json:"payload"`
}

// ParticipantJoinedPayload represents the payload for the participantJoined event.
type ParticipantJoinedPayload struct {
	ParticipantId string `json:"participantId"`
	Name          string `json:"name"`
}

// RoundStartedPayload represents the payload for the roundStarted event.
type RoundStartedPayload struct {
	RoundId string `json:"roundId"`
}

// VoteSubmittedPayload represents the payload for the voteSubmitted event.
type VoteSubmittedPayload struct {
	ParticipantId string `json:"participantId"`
}

// Vote represents a single vote.
type Vote struct {
	ParticipantId string `json:"participantId"`
	Value         string `json:"value"`
}

// VotesRevealedPayload represents the payload for the votesRevealed event.
type VotesRevealedPayload struct {
	RoundId string `json:"roundId"`
}

// SessionEndedPayload represents the payload for the sessionEnded event.
type SessionEndedPayload struct{}

// WebSocketHubInterface defines the interface for managing WebSocket connections.
type WebSocketHub interface {
	Run()
	HandleWebSocket(c echo.Context, sessionID string) error
	BroadcastParticipantJoined(sessionID, participantId, name string)
	BroadcastRoundStarted(sessionID, roundId string)
	BroadcastVoteSubmitted(sessionID, participantId string)
	BroadcastVotesRevealed(sessionID, roundId string)
	BroadcastSessionEnded(sessionID string)
}

// WebSocketHub manages WebSocket connections and message broadcasting.
type webSocketHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan WebSocketMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mutex      sync.RWMutex
	sessions   map[string]map[*websocket.Conn]bool
}

// NewWebSocketHub creates a new WebSocketHub.
func NewWebSocketHub() WebSocketHub {
	return &webSocketHub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan WebSocketMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		sessions:   make(map[string]map[*websocket.Conn]bool),
	}
}

// Run starts the WebSocketHub's main loop.
func (hub *webSocketHub) Run() {
	for {
		select {
		case connection := <-hub.register:
			hub.mutex.Lock()
			hub.clients[connection] = true
			hub.mutex.Unlock()
			log.Println("New WebSocket connection registered")
		case connection := <-hub.unregister:
			hub.mutex.Lock()
			if _, ok := hub.clients[connection]; ok {
				delete(hub.clients, connection)
				connection.Close()
			}
			hub.mutex.Unlock()
			log.Println("WebSocket connection unregistered")
		case message := <-hub.broadcast:
			hub.mutex.RLock()
			for connection := range hub.clients {
				if err := connection.WriteJSON(message); err != nil {
					log.Println("Error broadcasting message:", err)
					hub.unregister <- connection
				}
			}
			hub.mutex.RUnlock()
			log.Println("Message broadcasted:", message)
		}
	}
}

// HandleWebSocket handles WebSocket connections.
func (hub *webSocketHub) HandleWebSocket(c echo.Context, sessionID string) error {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for simplicity
		},
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return err
	}

	if sessionID == "" {
		log.Println("Session ID is required")
		conn.Close()
		return echo.NewHTTPError(http.StatusBadRequest, "Session ID is required")
	}

	hub.mutex.Lock()
	if hub.sessions[sessionID] == nil {
		hub.sessions[sessionID] = make(map[*websocket.Conn]bool)
	}
	hub.sessions[sessionID][conn] = true
	hub.mutex.Unlock()

	// Handle incoming messages (if needed)
	go func() {
		defer func() {
			hub.mutex.Lock()
			delete(hub.sessions[sessionID], conn)
			if len(hub.sessions[sessionID]) == 0 {
				delete(hub.sessions, sessionID)
			}
			hub.mutex.Unlock()
			conn.Close()
			log.Println("WebSocket connection unregistered")
		}()
		for {
			var msg WebSocketMessage
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}
			log.Println("Received message:", msg)
			// Process incoming messages here if needed
		}
	}()
	return nil
}

// BroadcastParticipantJoined broadcasts the participantJoined event.
func (hub *webSocketHub) BroadcastParticipantJoined(sessionID, participantId, name string) {
	payload := ParticipantJoinedPayload{
		ParticipantId: participantId,
		Name:          name,
	}
	message := WebSocketMessage{
		Event:   "participantJoined",
		Payload: payload,
	}
	hub.broadcastToSession(sessionID, message)
}

// BroadcastRoundStarted broadcasts the roundStarted event.
func (hub *webSocketHub) BroadcastRoundStarted(sessionID, roundId string) {
	payload := RoundStartedPayload{
		RoundId: roundId,
	}
	message := WebSocketMessage{
		Event:   "roundStarted",
		Payload: payload,
	}
	hub.broadcastToSession(sessionID, message)
}

// BroadcastVoteSubmitted broadcasts the voteSubmitted event.
func (hub *webSocketHub) BroadcastVoteSubmitted(sessionID, participantId string) {
	payload := VoteSubmittedPayload{
		ParticipantId: participantId,
	}
	message := WebSocketMessage{
		Event:   "voteSubmitted",
		Payload: payload,
	}
	hub.broadcastToSession(sessionID, message)
}

// BroadcastVotesRevealed broadcasts the votesRevealed event.
func (hub *webSocketHub) BroadcastVotesRevealed(sessionID, roundId string) {
	payload := VotesRevealedPayload{
		RoundId: roundId,
	}
	message := WebSocketMessage{
		Event:   "votesRevealed",
		Payload: payload,
	}
	hub.broadcastToSession(sessionID, message)
}

// BroadcastSessionEnded broadcasts the sessionEnded event.
func (hub *webSocketHub) BroadcastSessionEnded(sessionID string) {
	payload := SessionEndedPayload{}
	message := WebSocketMessage{
		Event:   "sessionEnded",
		Payload: payload,
	}
	hub.broadcastToSession(sessionID, message)
}

func (hub *webSocketHub) broadcastToSession(sessionID string, message WebSocketMessage) {
	hub.mutex.RLock()
	defer hub.mutex.RUnlock()

	if clients, ok := hub.sessions[sessionID]; ok {
		for connection := range clients {
			if err := connection.WriteJSON(message); err != nil {
				log.Println("Error broadcasting message:", err)
				hub.unregister <- connection
			}
		}
		log.Println("Message broadcasted to session", sessionID, ":", message)
	} else {
		log.Println("Session not found:", sessionID)
	}
}
