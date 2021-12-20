package ik_enterprise

import (
  "github.com/g0sy23/ik-app/internal"
  "github.com/g0sy23/ik-app/internal/repository"
)

type MerchCategoryEnterprise struct {
  repository ik_repository.MerchCategory
}

func NewMerchCategoryEnterprise(repository ik_repository.MerchCategory) *MerchCategoryEnterprise {
  return &MerchCategoryEnterprise{repository: repository}
}

func (c *MerchCategoryEnterprise) CreateMerchCategory(category ik_common.MerchCategory) (int, error) {
  return c.repository.CreateMerchCategory(category);
}

type MerchItemEnterprise struct {

}