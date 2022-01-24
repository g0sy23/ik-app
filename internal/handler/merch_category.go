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

	id, err := h.services.MerchCategory.Create(category)
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

// @Summary      Get all merch categories
// @Description  Get all merch categories
// @Tags         category
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/category [get]
func (h *Handler) getMerchCategoriesAll(context *fiber.Ctx) error {
	categories, err := h.services.MerchCategory.GetAll()
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

// @Summary      Get category by ID
// @Description  Get merch category by given ID
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "category ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/category/{id} [get]
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

	category, err := h.services.MerchCategory.GetById(id)
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

	if err := h.services.MerchCategory.Update(id, categoryUpdate); err != nil {
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

	if err := h.services.MerchCategory.Delete(id); err != nil {
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
