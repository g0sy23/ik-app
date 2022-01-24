package ik_handler

import (
	"strconv"

	"github.com/g0sy23/ik-app/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createMerchItem(context *fiber.Ctx) error {
	var item ik_models.MerchItem

	if err := context.BodyParser(&item); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to parse body",
			},
		)
	}

	if err := item.Validate(); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to validate body",
			},
		)
	}

	id, err := h.enterprise.MerchItem.Create(item)
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to create item",
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

// @Summary      Get all merch items
// @Description  Get all merch items
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/item [get]
func (h *Handler) getMerchItemsAll(context *fiber.Ctx) error {
	items, err := h.enterprise.MerchItem.GetAll()
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to get items",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"data": items,
			"status": fiber.StatusOK,
		},
	)
}

// @Summary      Get item by ID
// @Description  Get merch item by given ID
// @Tags         item
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "item ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/item/{id} [get]
func (h *Handler) getMerchItemById(context *fiber.Ctx) error {
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

	item, err := h.enterprise.MerchItem.GetById(id)
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to get item",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"data": item,
			"status": fiber.StatusOK,
		},
	)
}

func (h *Handler) getMerchItemsByCategoryId(context *fiber.Ctx) error {
	category_id, err := strconv.Atoi(context.Params("category_id"))
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to get category_id",
			},
		)
	}

	items, err := h.enterprise.MerchItem.GetByCategoryId(category_id)
	if err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to get item",
			},
		)
	}

	return responseData(
		context,
		map[string]interface{}{
			"data": items,
			"status": fiber.StatusOK,
		},
	)
}

func (h *Handler) updateMerchItem(context *fiber.Ctx) error {
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

	var itemUpdate ik_models.MerchItemUpdate

	if err := context.BodyParser(&itemUpdate); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to parse body",
			},
		)
	}

	if err := itemUpdate.Validate(); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusBadRequest,
				"message": "failed to validate body",
			},
		)
	}

	if err := h.enterprise.MerchItem.Update(id, itemUpdate); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to update item",
			},
		)
	}

	return responseStatusOK(context)
}

func (h *Handler) deleteMerchItem(context *fiber.Ctx) error {
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

	if err := h.enterprise.MerchItem.Delete(id); err != nil {
		return responseData(
			context,
			map[string]interface{}{
				"error": err.Error(),
				"status": fiber.StatusInternalServerError,
				"message": "failed to delete item",
			},
		)
	}

	return responseStatusOK(context)
}
