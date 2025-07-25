package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"go-notification/config"
	"go-notification/models"
	"go-notification/routes"
	"go-notification/worker"
	"go-notification/ws"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Notification{})

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	routes.RegisterNotificationRoutes(app)
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws", websocket.New(ws.HandleWebSocket))

	worker.StartScheduler()
	ws.StartBroadcaster()

	log.Fatal(app.Listen(":7000"))
}
