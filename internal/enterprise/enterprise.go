package ik_enterprise

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
	GetById(id int) (ik_models.MerchItem, error)
	GetByCategoryId(category_id int) ([]ik_models.MerchItem, error)
	Update(id int, category ik_models.MerchItemUpdate) error
	Delete(id int) error
}

type Enterprise struct {
	MerchCategory
	MerchItem
}

func NewEnterprise(repository *ik_repository.Repository) *Enterprise {
	return &Enterprise{
		MerchCategory: NewMerchCategoryEnterprise(repository.MerchCategory),
		MerchItem:     NewMerchItemEnterprise(repository.MerchItem, repository.MerchCategory),
	}
}
