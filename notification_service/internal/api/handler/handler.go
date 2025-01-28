package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

type WebSocket struct {
	Map   map[string]*websocket.Conn
	Mutex *sync.Mutex
	Ctx   context.Context
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *WebSocket) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket is working")
	username := r.Header.Get("username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Current connections map:", u.Map)

	kafkaReader, err := kgo.NewClient(
		kgo.SeedBrokers("broker:29092"),
		kgo.ConsumeTopics("notification"),
	)
	if err != nil {
		log.Println("Kafka client creation error:", err)
		return
	}
	defer kafkaReader.Close()
	for {
		fetches := kafkaReader.PollFetches(r.Context())
		if fetches.IsClientClosed() {
			break
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			if string(record.Key) == username {
				message := record.Value
				if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("Error writing message to WebSocket:", err)
					return
				}
			}
		}
	}
}
