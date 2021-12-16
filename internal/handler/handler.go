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

func (h* Handler) InitRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		category := api.Group("/category")
		{
			category.Post("/", 			h.CreateMerchCategory)
			category.Get("/", 			h.GetMerchCategoryAll)
			category.Get("/:id", 		h.GetMerchCategoryById)
			category.Patch("/:id", 	h.UpdateMerchCategory)
			category.Delete("/:id", h.DeleteMerchCategory)
		}
		item := api.Group("/item")
		{
			item.Post("/", 			h.CreateMerchItem)
			item.Get("/:id", 		h.GetMerchItemById)
			item.Patch("/:id", 	h.UpdateMerchItem)
			item.Delete("/:id",	h.DeleteMerchItem)
		}
	}
}
