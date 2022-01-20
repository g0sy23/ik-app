package ik_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func responseStatusOK(context *fiber.Ctx) error {
	return context.SendStatus(fiber.StatusOK)
}

func responseData(context *fiber.Ctx, data map[string]interface{}) error {
	logrus.WithFields(data).Debug()
	return context.Status(data["status"].(int)).JSON(data)
}
