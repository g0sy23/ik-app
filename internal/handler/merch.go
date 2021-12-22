package ik_handler

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/gofiber/fiber/v2"
)

func (h* Handler) createMerchCategory(context *fiber.Ctx) error {
  var category ik_common.MerchCategory

  if err := context.BodyParser(&category); err != nil {
    return responseError(context, fiber.StatusBadGateway, err.Error())
  }

  if err := ik_common.ValidateStruct(category); err != nil {
    return responseError(context, fiber.StatusBadGateway, err)
  }

  id, err := h.enterprise.MerchCategory.Create(category)
  if err != nil {
    return responseError(context, fiber.StatusInternalServerError, err.Error())
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h* Handler) getMerchCategoryAll(context *fiber.Ctx) error {
  categories, err := h.enterprise.MerchCategory.GetAll()
  if err != nil {
    return responseError(context, fiber.StatusInternalServerError, err.Error())
  }
  return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": categories})
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