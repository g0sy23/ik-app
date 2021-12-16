package ik_repository

type Repository struct {
	MerchCategory
	MerchItem
}

func NewRepository() *Repository {
	return &Repository{}
}
