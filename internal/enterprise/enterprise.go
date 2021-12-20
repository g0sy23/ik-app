package ik_enterprise

import (
	"github.com/g0sy23/ik-app/internal"
	"github.com/g0sy23/ik-app/internal/repository"
)

type MerchCategory interface {
	CreateMerchCategory(category ik_common.MerchCategory) (int, error)
}

type MerchItem interface {

}

type Enterprise struct {
	MerchCategory
	MerchItem
}

func NewEnterprise(repository *ik_repository.Repository) *Enterprise {
	return &Enterprise{
		MerchCategory: NewMerchCategoryEnterprise(repository.MerchCategory),
	}
}
