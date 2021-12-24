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

func (c *MerchCategoryEnterprise) Create(category ik_common.MerchCategory) (int, error) {
  return c.repository.Create(category);
}

func (c *MerchCategoryEnterprise) GetAll() ([]ik_common.MerchCategory, error) {
  return c.repository.GetAll();
}

func (c *MerchCategoryEnterprise) GetById(id int) (ik_common.MerchCategory, error) {
  return c.repository.GetById(id);
}

func (c *MerchCategoryEnterprise) Update(id int, category ik_common.MerchCategoryUpdate) error {
  return c.repository.Update(id, category);
}

func (c *MerchCategoryEnterprise) Delete(id int) error {
  return c.repository.Delete(id);
}

type MerchItemEnterprise struct {

}