package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu          sync.RWMutex
	connections map[string]*websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[string]*websocket.Conn),
	}
}

func (h *Hub) Register(userID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.connections[userID] = conn
}

func (h *Hub) Unregister(userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.connections, userID)
}

func (h *Hub) SendToUser(userID string, message []byte) error {
	h.mu.RLock()
	conn, ok := h.connections[userID]
	h.mu.RUnlock()

	if !ok {
		return nil
	}
	return conn.WriteMessage(websocket.TextMessage, message)
}
