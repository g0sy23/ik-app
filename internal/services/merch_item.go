package ik_services

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/internal/repository"
)

type MerchItemService struct {
	repository         ik_repository.MerchItem
	categoryRepository ik_repository.MerchCategory
}

func NewMerchItemService(
	repository ik_repository.MerchItem,
	categoryRepository ik_repository.MerchCategory,
) *MerchItemService {
	return &MerchItemService{
		repository:         repository,
		categoryRepository: categoryRepository,
	}
}

func (c *MerchItemService) Create(item ik_models.MerchItem) (int, error) {
	return c.repository.Create(item)
}

func (c *MerchItemService) GetAll() ([]ik_models.MerchItem, error) {
	return c.repository.GetAll()
}

func (c *MerchItemService) GetById(id int) (ik_models.MerchItem, error) {
	return c.repository.GetById(id)
}

func (c *MerchItemService) GetByCategoryId(category_id int) ([]ik_models.MerchItem, error) {
	return c.repository.GetByCategoryId(category_id)
}

func (c *MerchItemService) Update(id int, itemUpdate ik_models.MerchItemUpdate) error {
	return c.repository.Update(id, itemUpdate)
}

func (c *MerchItemService) Delete(id int) error {
	return c.repository.Delete(id)
}
