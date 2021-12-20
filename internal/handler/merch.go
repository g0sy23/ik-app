package ik_handler

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/gofiber/fiber/v2"
)

func (h* Handler) createMerchCategory(context *fiber.Ctx) error {
  var category ik_common.MerchCategory
  if err := context.BodyParser(&category); err != nil {
    return errorResponse(context, fiber.StatusBadRequest, err.Error())
  }

  id, err := h.enterprise.MerchCategory.CreateMerchCategory(category)
  if err != nil {
    return errorResponse(context, fiber.StatusInternalServerError, err.Error())
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h* Handler) getMerchCategoryAll(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) getMerchCategoryById(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) updateMerchCategory(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) deleteMerchCategory(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) createMerchItem(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) getMerchItemById(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) updateMerchItem(c *fiber.Ctx) error {
  return nil
}

func (h* Handler) deleteMerchItem(c *fiber.Ctx) error {
  return nil
}