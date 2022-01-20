package ik_services

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/internal/repository"
)

type MerchCategoryService struct {
	repository ik_repository.MerchCategory
}

func NewMerchCategoryService(repository ik_repository.MerchCategory) *MerchCategoryService {
	return &MerchCategoryService{repository: repository}
}

func (c *MerchCategoryService) Create(category ik_models.MerchCategory) (int, error) {
	return c.repository.Create(category)
}

func (c *MerchCategoryService) GetAll() ([]ik_models.MerchCategory, error) {
	return c.repository.GetAll()
}

func (c *MerchCategoryService) GetById(id int) (ik_models.MerchCategory, error) {
	return c.repository.GetById(id)
}

func (c *MerchCategoryService) Update(id int, categoryUpdate ik_models.MerchCategoryUpdate) error {
	return c.repository.Update(id, categoryUpdate)
}

func (c *MerchCategoryService) Delete(id int) error {
	return c.repository.Delete(id)
}
