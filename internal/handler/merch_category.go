package ik_handler

import (
	"strconv"

	"github.com/g0sy23/ik-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createMerchCategory(context *fiber.Ctx) error {
	var category ik_models.MerchCategory

	if err := context.BodyParser(&category); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to parse body",
			},
		)
	}

	if err := category.Validate(); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to validate body",
			},
		)
	}

	id, err := h.enterprise.MerchCategory.Create(category)
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to create category",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"id": id,
			"status": fiber.StatusOK,
		},
	)
}

func (h *Handler) getMerchCategoriesAll(context *fiber.Ctx) error {
	categories, err := h.enterprise.MerchCategory.GetAll()
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to get categories",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"data": categories,
			"status": fiber.StatusOK,
		},
	)
}

func (h *Handler) getMerchCategoryById(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to get id",
			},
		)
	}

	category, err := h.enterprise.MerchCategory.GetById(id)
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to get category",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"data": category,
			"status": fiber.StatusOK,
		},
	)
}

func (h *Handler) updateMerchCategory(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to get id",
			},
		)
	}

	var categoryUpdate ik_models.MerchCategoryUpdate

	if err := context.BodyParser(&categoryUpdate); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to parse body",
			},
		)
	}

	if err := categoryUpdate.Validate(); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to validate body",
			},
		)
	}

	if err := h.enterprise.MerchCategory.Update(id, categoryUpdate); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to update category",
			},
		)
	}

	return responseStatusOK(context)
}

func (h *Handler) deleteMerchCategory(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to get id",
			},
		)
	}

	if err := h.enterprise.MerchCategory.Delete(id); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to delete category",
			},
		)
	}

	return responseStatusOK(context)
}
