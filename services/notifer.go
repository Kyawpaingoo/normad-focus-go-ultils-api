package services

import (
	"fmt"
	"go-notification/models"
)

func SendNotification(n models.Notification) {
	fmt.Printf("🔔 Notify user %d: %s - %s at %t\n", n.User_Id, n.Source_Type, n.Message, n.Sent_At)
}
