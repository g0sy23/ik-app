package ik_handler

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/gofiber/fiber/v2"
  "github.com/sirupsen/logrus"
)

func responseStringError(context *fiber.Ctx, status int, msg string) error {
  logrus.Error(msg)
  return context.Status(status).JSON(fiber.Map{"message": msg})
}

func responseValidateError(context *fiber.Ctx, status int, err []*ik_common.ErrorValidate) error {
  msg := "Failed to validate fields"
  logrus.WithField("data", err).Error(msg)
  return context.Status(status).JSON(
    fiber.Map{
      "message": msg,
      "data": err,
    },
  )
}

func responseError(context *fiber.Ctx, status int, data interface{}) error {
  switch variable := data.(type) {
  case string:
    return responseStringError(context, status, variable)
  case []*ik_common.ErrorValidate:
    return responseValidateError(context, status, variable)
  default:
    return nil
  }
}