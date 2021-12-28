package ik_repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/g0sy23/ik-app/internal"
	"github.com/jmoiron/sqlx"
)

type MerchCategoryPostgres struct {
	database *sqlx.DB
}

func NewMerchCategoryPostgres(database *sqlx.DB) *MerchCategoryPostgres {
	return &MerchCategoryPostgres{database: database}
}

func (p *MerchCategoryPostgres) Create(category ik_common.MerchCategory) (int, error) {
	var ret int
	var columns []string
	var values []string

	columns = append(columns, "title")
	values = append(values, fmt.Sprintf("'%s'", *category.Title))

	if category.Description != nil {
		columns = append(columns, "description")
		values = append(values, fmt.Sprintf("'%s'", *category.Description))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		merchCategoriesTable,
		strings.Join(columns, ","),
		strings.Join(values, ","))

	row := p.database.QueryRow(query)
	if err := row.Scan(&ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (p *MerchCategoryPostgres) GetAll() ([]ik_common.MerchCategory, error) {
	var ret []ik_common.MerchCategory

	query := fmt.Sprintf("SELECT * FROM %s", merchCategoriesTable)
	err := p.database.Select(&ret, query)

	return ret, err
}

func (p *MerchCategoryPostgres) GetById(id int) (ik_common.MerchCategory, error) {
	var ret ik_common.MerchCategory

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", merchCategoriesTable)
	err := p.database.Get(&ret, query, id)

	return ret, err
}

func (p *MerchCategoryPostgres) Update(id int, categoryUpdate ik_common.MerchCategoryUpdate) error {
	var values []string

	if categoryUpdate.Title != nil {
		values = append(values, fmt.Sprintf("title='%s'", *categoryUpdate.Title))
	}

	if categoryUpdate.Description != nil {
		values = append(values, fmt.Sprintf("description='%s'", *categoryUpdate.Description))
	}

	setValues := strings.Join(values, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$1", merchCategoriesTable, setValues)

	res, err := p.database.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("row not affected")
	}

	return nil
}

func (p *MerchCategoryPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", merchCategoriesTable)

	res, err := p.database.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("row not affected")
	}

	return nil
}

type MerchItemPostgres struct {
	database *sqlx.DB
}

func NewMerchItemPostgres(database *sqlx.DB) *MerchItemPostgres {
	return &MerchItemPostgres{database: database}
}

func (p *MerchItemPostgres) Create(item ik_common.MerchItem) (int, error) {
	var ret int
	var columns []string
	var values []string

	columns = append(columns, "title")
	values = append(values, fmt.Sprintf("'%s'", *item.Title))

	columns = append(columns, "price")
	values = append(values, fmt.Sprintf("%d", *item.Price))

	columns = append(columns, "quantity")
	values = append(values, fmt.Sprintf("%d", *item.Quantity))

	columns = append(columns, "category_id")
	if item.CategoryId != nil {
		values = append(values, fmt.Sprintf("%d", *item.CategoryId))
	} else {
		values = append(values, "DEFAULT")
	}

	if item.Description != nil {
		columns = append(columns, "description")
		values = append(values, fmt.Sprintf("'%s'", *item.Description))
	}

	if item.Size != nil {
		columns = append(columns, "size")
		values = append(values, fmt.Sprintf("'%s'", *item.Size))
	}

	if item.Color != nil {
		columns = append(columns, "color")
		values = append(values, fmt.Sprintf("'%s'", *item.Color))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		merchItemsTable,
		strings.Join(columns, ","),
		strings.Join(values, ","))

	row := p.database.QueryRow(query)
	if err := row.Scan(&ret); err != nil {
		return 0, err
	}

	return ret, nil
}

func (p *MerchItemPostgres) GetAll() ([]ik_common.MerchItem, error) {
	var ret []ik_common.MerchItem

	query := fmt.Sprintf("SELECT * FROM %s", merchItemsTable)
	err := p.database.Select(&ret, query)

	return ret, err
}

func (p *MerchItemPostgres) GetById(id int) (ik_common.MerchItem, error) {
	var ret ik_common.MerchItem

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", merchItemsTable)
	err := p.database.Get(&ret, query, id)

	return ret, err
}

func (p *MerchItemPostgres) GetByCategoryId(category_id int) ([]ik_common.MerchItem, error) {
	var ret []ik_common.MerchItem

	query := fmt.Sprintf("SELECT * FROM %s WHERE category_id=$1", merchItemsTable)
	err := p.database.Get(&ret, query, category_id)

	return ret, err
}

func (p *MerchItemPostgres) Update(id int, itemUpdate ik_common.MerchItemUpdate) error {
	var values []string

	if itemUpdate.CategoryId != nil {
		values = append(values, fmt.Sprintf("category_id='%d'", *itemUpdate.CategoryId))
	}

	if itemUpdate.Title != nil {
		values = append(values, fmt.Sprintf("title='%s'", *itemUpdate.Title))
	}

	if itemUpdate.Description != nil {
		values = append(values, fmt.Sprintf("description='%s'", *itemUpdate.Description))
	}

	if itemUpdate.Price != nil {
		values = append(values, fmt.Sprintf("price='%d'", *itemUpdate.Price))
	}

	if itemUpdate.Quantity != nil {
		values = append(values, fmt.Sprintf("quantity='%d'", *itemUpdate.Quantity))
	}

	if itemUpdate.Size != nil {
		values = append(values, fmt.Sprintf("size='%s'", *itemUpdate.Size))
	}

	if itemUpdate.Color != nil {
		values = append(values, fmt.Sprintf("color='%s'", *itemUpdate.Color))
	}

	setValues := strings.Join(values, ",")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$1", merchItemsTable, setValues)

	res, err := p.database.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("row not affected")
	}

	return nil
}

func (p *MerchItemPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", merchItemsTable)
	res, err := p.database.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("row not affected")
	}

	return nil
}
