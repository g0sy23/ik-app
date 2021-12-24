package ik_handler

import (
	"strconv"

	"github.com/g0sy23/ik-app/internal"
	"github.com/gofiber/fiber/v2"
)

func (h* Handler) createMerchCategory(context *fiber.Ctx) error {
  var category ik_common.MerchCategory

  if err := context.BodyParser(&category); err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to parse body", err.Error())
  }

  if err := category.Validate(); err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to validate body", err.Error())
  }

  id, err := h.enterprise.MerchCategory.Create(category)
  if err != nil {
    return responseError(context, fiber.StatusInternalServerError,
                         "failed to create category", err.Error())
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (h* Handler) getMerchCategoryAll(context *fiber.Ctx) error {
  categories, err := h.enterprise.MerchCategory.GetAll()
  if err != nil {
    return responseError(context, fiber.StatusInternalServerError,
                         "failed to get categories", err.Error())
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": categories})
}

func (h* Handler) getMerchCategoryById(context *fiber.Ctx) error {
  id, err := strconv.Atoi(context.Params("id"))
  if err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to get id", err.Error())
  }

  category, err := h.enterprise.MerchCategory.GetById(id)
  if err != nil {
    return responseError(context, fiber.StatusInternalServerError,
                         "failed to get category", err.Error())
  }

  return context.Status(fiber.StatusOK).JSON(fiber.Map{"data": category})
}

func (h* Handler) updateMerchCategory(context *fiber.Ctx) error {
  id, err := strconv.Atoi(context.Params("id"))
  if err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to get id", err.Error())
  }

  var categoryUpdate ik_common.MerchCategoryUpdate

  if err := context.BodyParser(&categoryUpdate); err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to parse body", err.Error())
  }

  if err := categoryUpdate.Validate(); err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to validate body", err.Error())
  }

  if err := h.enterprise.MerchCategory.Update(id, categoryUpdate); err != nil {
    return responseError(context, fiber.StatusInternalServerError,
                         "failed to update category", err.Error())
  }

  return context.SendStatus(fiber.StatusOK)
}

func (h* Handler) deleteMerchCategory(context *fiber.Ctx) error {
  id, err := strconv.Atoi(context.Params("id"))
  if err != nil {
    return responseError(context, fiber.StatusBadRequest,
                         "failed to get id", err.Error())
  }

  if err := h.enterprise.MerchCategory.Delete(id); err != nil {
    return responseError(context, fiber.StatusInternalServerError,
                         "failed to delete category", err.Error())
  }

  return context.SendStatus(fiber.StatusOK)
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