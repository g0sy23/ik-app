package ik_handler

import (
  "github.com/gofiber/fiber/v2"
  "github.com/sirupsen/logrus"
)

func responseError(context *fiber.Ctx, status int, message string, data interface{}) error {
  logrus.WithField("data", data).Error(message)
  return context.Status(status).JSON(
    fiber.Map{
      "message": message,
      "data": data,
    },
  )
}