package ik_enterprise

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/g0sy23/ik-app/internal/repository"
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