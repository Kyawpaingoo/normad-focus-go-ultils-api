package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-notification/config"
	"go-notification/models"
	"strconv"
)

func GetNotifications(c *fiber.Ctx) error {
	var notifications []models.Notification
	if err := config.DB.Order("notify_at asc").Find(&notifications).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notifications"})
	}
	return c.JSON(notifications)
}

func GetNotificationByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	var notifications []models.Notification
	if err := config.DB.Where("user_id = ? AND sent = ?", id, false).Find(&notifications).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	}

	return c.JSON(notifications)
}
