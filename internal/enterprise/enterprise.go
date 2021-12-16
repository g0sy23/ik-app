package ik_enterprise

import "github.com/g0sy23/ik-app/internal/repository"

type Enterprise struct {
	MerchCategory
	MerchItem
}

func NewEnterprise(repository *ik_repository.Repository) *Enterprise {
	return &Enterprise{}
}
