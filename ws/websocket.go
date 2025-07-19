package ws

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"go-notification/models"
	"log"
	"sync"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Notification)
	mu        sync.Mutex
)

func HandleWebSocket(c *websocket.Conn) {
	mu.Lock()
	clients[c] = true
	mu.Unlock()
	log.Println("WebSocket connected")

	defer func() {
		mu.Lock()
		delete(clients, c)
		mu.Unlock()
		c.Close()
		log.Println("WebSocket disconnected")
	}()
	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}

func StartBroadcaster() {
	go func() {
		for notification := range broadcast {
			data, err := json.Marshal(&notification)
			if err != nil {
				log.Println("Marshal error:", err)
				continue
			}
			mu.Lock()
			for client := range clients {
				if err := client.WriteMessage((websocket.TextMessage), data); err != nil {
					log.Println("Send error:", err)
					client.Close()
					delete(clients, client)
				}
			}
			mu.Unlock()
		}
	}()
}

func PushNotificationClients(notification models.Notification) {
	broadcast <- notification
}
