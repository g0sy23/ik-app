package ik_postgres

import (
	"github.com/g0sy23/ik-app/internal/models"
	"github.com/g0sy23/ik-app/pkg/database/postgres"
)

const merchItemsTable = "merch_items"

type MerchItemPostgres struct {
	database *postgresdb.PostgresDB
}

func NewMerchItemPostgres(database *postgresdb.PostgresDB) *MerchItemPostgres {
	return &MerchItemPostgres{database: database}
}

func (p *MerchItemPostgres) Create(item ik_models.MerchItem) (int, error) {
	fields := postgresdb.Fields{
		"title": *item.Title,
		"price": *item.Price,
		"quantity": *item.Quantity,
	}

	if item.CategoryId != nil {
		fields["category_id"] = *item.CategoryId
	}

	if item.Description != nil {
		fields["description"] = *item.Description
	}

	if item.Size != nil {
		fields["size"] = *item.Size
	}

	if item.Color != nil {
		fields["color"] = *item.Color
	}

	var ret int
	if err := p.database.Insert(&ret, merchItemsTable, fields); err != nil {
		return 0, err
	}

	return ret, nil
}

func (p *MerchItemPostgres) GetAll() ([]ik_models.MerchItem, error) {
	var ret []ik_models.MerchItem
	err := p.database.GetAll(&ret, merchItemsTable)
	return ret, err
}

func (p *MerchItemPostgres) GetAllPaged(pageNumber int) ([]ik_models.MerchItem, error) {
	var ret []ik_models.MerchItem
	err := p.database.GetAllPaged(&ret, merchItemsTable, pageNumber)
	return ret, err
}

func (p *MerchItemPostgres) GetById(id int) (ik_models.MerchItem, error) {
	var ret ik_models.MerchItem
	err := p.database.GetById(&ret, merchItemsTable, id)
	return ret, err
}

func (p *MerchItemPostgres) GetByCategoryId(categoryId/*, pageNumber*/ int) ([]ik_models.MerchItem, error) {
	fields := postgresdb.Fields{"category_id": categoryId}
	var ret []ik_models.MerchItem
	err := p.database.Get(&ret, merchItemsTable/*, pageNumber*/, fields)
	return ret, err
}

func (p *MerchItemPostgres) Update(id int, itemUpdate ik_models.MerchItemUpdate) error {
	fields := make(postgresdb.Fields)

	if itemUpdate.CategoryId != nil {
		fields["category_id"] = *itemUpdate.CategoryId
	}

	if itemUpdate.Title != nil {
		fields["title"] = *itemUpdate.Title
	}

	if itemUpdate.Description != nil {
		fields["description"] = *itemUpdate.Description
	}

	if itemUpdate.Price != nil {
		fields["price"] = *itemUpdate.Price
	}

	if itemUpdate.Quantity != nil {
		fields["quantity"] = *itemUpdate.Quantity
	}

	if itemUpdate.Size != nil {
		fields["size"] = *itemUpdate.Size
	}

	if itemUpdate.Color != nil {
		fields["color"] = *itemUpdate.Color
	}

	return p.database.Update(merchItemsTable, id, fields)
}

func (p *MerchItemPostgres) Delete(id int) error {
	return p.database.Delete(merchItemsTable, id)
}
