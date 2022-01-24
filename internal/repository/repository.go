package ik_repository

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/internal/repository/postgres"
	"github.com/g0sy23/ik-app/pkg/database/postgres"
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

type Repository struct {
	MerchCategory
	MerchItem
}

func New(database *postgresdb.PostgresDB) *Repository {
	return &Repository{
		MerchCategory: ik_postgres.NewMerchCategoryPostgres(database),
		MerchItem:     ik_postgres.NewMerchItemPostgres(database),
	}
}
