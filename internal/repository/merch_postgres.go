package ik_repository

import (
  "fmt"

  "github.com/g0sy23/ik-app/internal"
  "github.com/jmoiron/sqlx"
)

type MerchCategoryPostgres struct {
  database *sqlx.DB
}

func NewMerchCategoryPostgres(db *sqlx.DB) *MerchCategoryPostgres {
  return &MerchCategoryPostgres{database: db}
}

func (p *MerchCategoryPostgres) Create(category ik_common.MerchCategory) (int, error) {
  var ret int

  query := fmt.Sprintf(
    "INSERT INTO %s (title, description) values ($1, $2) RETURNING id",
    merchCategoriesTable,
  )

  row := p.database.QueryRow(query, category.Title, category.Description)
  if err := row.Scan(&ret); err != nil {
    return 0, err
  }

  return ret, nil
}

func (p *MerchCategoryPostgres) GetAll() ([]ik_common.MerchCategory, error) {
  var ret []ik_common.MerchCategory

  query := fmt.Sprintf("SELECT * FROM %s", merchCategoriesTable)
	err   := p.database.Select(&ret, query)

	return ret, err
}

func (p *MerchCategoryPostgres) GetById(id int) (ik_common.MerchCategory, error) {
  var ret ik_common.MerchCategory

  query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id = $1",
                       merchCategoriesTable)
	err   := p.database.Get(&ret, query, id)

	return ret, err
}