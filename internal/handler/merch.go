package ik_handler

import (
	"strconv"

	"github.com/g0sy23/ik-app/internal"
	"github.com/gofiber/fiber/v2"
)

func (h* Handler) createMerchCategory(context *fiber.Ctx) error {
  var category ik_common.MerchCategory

  if err := context.BodyParser(&category); err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to parse body", err)
  }

  if err := ik_common.ValidateStruct(category); err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to validate body", err)
  }

  id, err := h.enterprise.MerchCategory.Create(category)
  if err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to create category", err)
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h* Handler) getMerchCategoryAll(context *fiber.Ctx) error {
  categories, err := h.enterprise.MerchCategory.GetAll()
  if err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to get categories", err)
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": categories})
}

func (h* Handler) getMerchCategoryById(context *fiber.Ctx) error {
  id, err := strconv.Atoi(context.Params("id"))
  if err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to validate id", err)
  }

  category, err := h.enterprise.MerchCategory.GetById(id)
  if err != nil {
    return responseError(context, fiber.StatusBadRequest, "failed to get category", err)
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": category})
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