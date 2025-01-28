package connections

import (
	"context"
	"notification/internal/api/handler"
	"notification/internal/services"
	"sync"

	"github.com/gorilla/websocket"
)

func NewWebSocket() *handler.WebSocket {
	ctx := context.Background()
	return &handler.WebSocket{
		Map:   make(map[string]*websocket.Conn),
		Mutex: &sync.Mutex{},
		Ctx:   ctx,
	}
}

func NewService() *services.Service {
	a := NewWebSocket()
	return &services.Service{W: a}
}
