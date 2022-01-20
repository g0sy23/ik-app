package ik_repository

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	MerchCategory
	MerchItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MerchCategory: NewMerchCategoryPostgres(db),
		MerchItem:     NewMerchItemPostgres(db),
	}
}
