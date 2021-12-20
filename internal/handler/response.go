package ik_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type responseError struct {
	Message string `json:"message"`
}

func errorResponse(context *fiber.Ctx, status int, message string) error {
	logrus.Error(message)
	return context.Status(status).JSON(responseError{message})
}
