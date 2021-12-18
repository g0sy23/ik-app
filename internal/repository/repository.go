package ik_repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	MerchCategory
	MerchItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
