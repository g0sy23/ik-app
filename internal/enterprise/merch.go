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

func (c *MerchCategoryEnterprise) Update(id int, categoryUpdate ik_common.MerchCategoryUpdate) error {
  return c.repository.Update(id, categoryUpdate);
}

func (c *MerchCategoryEnterprise) Delete(id int) error {
  return c.repository.Delete(id);
}

type MerchItemEnterprise struct {
  repository         ik_repository.MerchItem
  categoryRepository ik_repository.MerchCategory
}

func NewMerchItemEnterprise(
  repository ik_repository.MerchItem,
  categoryRepository ik_repository.MerchCategory,
) *MerchItemEnterprise {
  return &MerchItemEnterprise{
    repository:         repository,
    categoryRepository: categoryRepository,
  }
}

func (c *MerchItemEnterprise) Create(item ik_common.MerchItem) (int, error) {
  return c.repository.Create(item);
}

func (c *MerchItemEnterprise) GetAll() ([]ik_common.MerchItem, error) {
  return c.repository.GetAll();
}

func (c *MerchItemEnterprise) GetById(id int) (ik_common.MerchItem, error) {
  return c.repository.GetById(id);
}

func (c *MerchItemEnterprise) GetByCategoryId(category_id int) ([]ik_common.MerchItem, error) {
  return c.repository.GetByCategoryId(category_id);
}

func (c *MerchItemEnterprise) Update(id int, itemUpdate ik_common.MerchItemUpdate) error {
  return c.repository.Update(id, itemUpdate);
}

func (c *MerchItemEnterprise) Delete(id int) error {
  return c.repository.Delete(id);
}