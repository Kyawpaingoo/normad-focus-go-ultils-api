package ws

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
	"go-notification/models"
	"log"
	"strconv"
	"sync"
)

type Client struct {
	Conn   *websocket.Conn
	UserID uint
}

var (
	clients = make(map[*websocket.Conn]Client)
	mu      sync.Mutex
)

func HandleWebSocket(c *websocket.Conn) {
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		log.Printf("❌ WebSocket connection rejected: missing user_id parameter")
		c.Close()
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Printf("Missing or invalid user_id")
		err := c.Close()
		if err != nil {
			return
		}
		return
	}

	client := Client{Conn: c, UserID: uint(userId)}

	mu.Lock()
	clients[c] = client
	mu.Unlock()

	log.Printf("✅ WebSocket client connected: user_id=%d", client.UserID)

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
	//go func() {
	//	for notification := range broadcast {
	//		data, err := json.Marshal(&notification)
	//		if err != nil {
	//			log.Println("Marshal error:", err)
	//			continue
	//		}
	//		mu.Lock()
	//		for client := range clients {
	//			if err := client.WriteMessage((websocket.TextMessage), data); err != nil {
	//				log.Println("Send error:", err)
	//				client.Close()
	//				delete(clients, client)
	//			}
	//		}
	//		mu.Unlock()
	//	}
	//}()
}

func PushNotificationClients(userID uint, notification models.Notification) {
	data, err := json.Marshal(notification)
	if err != nil {
		log.Println("Marshal error:", err)
		return
	}

	mu.Lock()
	for _, client := range clients {
		if client.UserID == userID {
			if err := client.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println("Write error:", err)
				client.Conn.Close()
				delete(clients, client.Conn)
			}
		}
	}
	mu.Unlock()
}
