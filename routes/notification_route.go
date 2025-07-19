package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-notification/handlers"
)

func RegisterNotificationRoutes(app *fiber.App) {
	api := app.Group("/notification")

	api.Get("/", handlers.GetNotifications)
	api.Get("/:id", handlers.GetNotificationByID)
}
