package worker

import (
	"go-notification/config"
	"go-notification/models"
	"go-notification/services"
	"go-notification/ws"
	"log"
	"time"
)

func StartScheduler() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			now := time.Now()
			var due []models.Notification
			err := config.DB.Where("notify_at <= ? AND sent = ?", now, false).Find(&due).Error
			if err != nil {
				log.Println(err)
				continue
			}
			for _, n := range due {
				services.SendNotification(n)
				ws.PushNotificationClients(n)
				n.Sent_At = now
				n.Sent = true
				config.DB.Save(&n)
			}
		}
	}()
}
