package ik_repository

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/jmoiron/sqlx"
)

type MerchCategory interface {
  Create(category ik_common.MerchCategory) (int, error)
  GetAll() ([]ik_common.MerchCategory, error)
  GetById(id int) (ik_common.MerchCategory, error)
}

type MerchItem interface {

}

type Repository struct {
  MerchCategory
  MerchItem
}

func NewRepository(db *sqlx.DB) *Repository {
  return &Repository{
    MerchCategory: NewMerchCategoryPostgres(db),
  }
}