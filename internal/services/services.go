package ik_services

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/internal/repository"
)

type MerchCategory interface {
	Create(category ik_models.MerchCategory) (int, error)
	GetAll() ([]ik_models.MerchCategory, error)
	GetById(id int) (ik_models.MerchCategory, error)
	Update(id int, category ik_models.MerchCategoryUpdate) error
	Delete(id int) error
}

type MerchItem interface {
	Create(category ik_models.MerchItem) (int, error)
	GetAll() ([]ik_models.MerchItem, error)
	GetAllPaged(pageNumber int) ([]ik_models.MerchItem, error)
	GetById(id int) (ik_models.MerchItem, error)
	GetByCategoryId(categoryId/*, pageNumber*/ int) ([]ik_models.MerchItem, error)
	Update(id int, category ik_models.MerchItemUpdate) error
	Delete(id int) error
}

type Services struct {
	MerchCategory
	MerchItem
}

func New(repository *ik_repository.Repository) *Services {
	return &Services{
		MerchCategory: NewMerchCategoryService(repository.MerchCategory),
		MerchItem:     NewMerchItemService(repository.MerchItem, repository.MerchCategory),
	}
}
