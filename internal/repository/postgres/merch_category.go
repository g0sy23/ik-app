package ik_postgres

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/pkg/database/postgres"
)

const merchCategoriesTable = "merch_categories"

type MerchCategoryPostgres struct {
	database *postgresdb.PostgresDB
}

func NewMerchCategoryPostgres(database *postgresdb.PostgresDB) *MerchCategoryPostgres {
	return &MerchCategoryPostgres{database: database}
}

func (p *MerchCategoryPostgres) Create(category ik_models.MerchCategory) (int, error) {
	fields := postgresdb.Fields{"title": *category.Title}

	if category.Description != nil {
		fields["description"] = *category.Description
	}

	var ret int
	if err := p.database.Insert(&ret, merchCategoriesTable, fields); err != nil {
		return 0, err
	}

	return ret, nil
}

func (p *MerchCategoryPostgres) GetAll() ([]ik_models.MerchCategory, error) {
	var ret []ik_models.MerchCategory
	err := p.database.GetAll(&ret, merchCategoriesTable)
	return ret, err
}

func (p *MerchCategoryPostgres) GetById(id int) (ik_models.MerchCategory, error) {
	var ret ik_models.MerchCategory
	err := p.database.GetById(&ret, merchCategoriesTable, id)
	return ret, err
}

func (p *MerchCategoryPostgres) Update(id int, categoryUpdate ik_models.MerchCategoryUpdate) error {
	fields := make(postgresdb.Fields)

	if categoryUpdate.Title != nil {
		fields["title"] = *categoryUpdate.Title
	}

	if categoryUpdate.Description != nil {
		fields["description"] = *categoryUpdate.Description
	}

	return p.database.Update(merchCategoriesTable, id, fields)
}

func (p *MerchCategoryPostgres) Delete(id int) error {
	return p.database.Delete(merchCategoriesTable, id)
}
