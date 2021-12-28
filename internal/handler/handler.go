package ik_handler

import (
	"github.com/g0sy23/ik-app/internal/enterprise"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	enterprise *ik_enterprise.Enterprise
}

func NewHandler(enterprise *ik_enterprise.Enterprise) *Handler {
	return &Handler{
		enterprise: enterprise,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	api := app.Group("/api") // for adminka
	{
		category := api.Group("/category")
		{
			category.Get("/", h.getMerchCategoriesAll)
			category.Get("/:id", h.getMerchCategoryById)
			category.Post("/", h.createMerchCategory)
			category.Patch("/:id", h.updateMerchCategory)
			category.Delete("/:id", h.deleteMerchCategory)
		}
		item := api.Group("/item")
		{
			item.Get("/", h.getMerchItemsAll)
			item.Get("/:id", h.getMerchItemById)
			item.Post("/", h.createMerchItem)
			item.Patch("/:id", h.updateMerchItem)
			item.Delete("/:id", h.deleteMerchItem)
		}
	}
	merch := app.Group("/merch")
	{
		merch.Get("/", h.getMerchItemsAll)
		merch.Get("/:id", h.getMerchItemsByCategoryId)
		merch.Get("/item/:id", h.getMerchItemById)
	}
}
