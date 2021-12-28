package ik_repository

import (
	"github.com/g0sy23/ik-app/internal"
	"github.com/jmoiron/sqlx"
)

type MerchCategory interface {
	Create(category ik_common.MerchCategory) (int, error)
	GetAll() ([]ik_common.MerchCategory, error)
	GetById(id int) (ik_common.MerchCategory, error)
	Update(id int, category ik_common.MerchCategoryUpdate) error
	Delete(id int) error
}

type MerchItem interface {
	Create(category ik_common.MerchItem) (int, error)
	GetAll() ([]ik_common.MerchItem, error)
	GetById(id int) (ik_common.MerchItem, error)
	GetByCategoryId(category_id int) ([]ik_common.MerchItem, error)
	Update(id int, category ik_common.MerchItemUpdate) error
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
