package postgresdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	database *sqlx.DB
}

type Config struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	SSLMode    string
}

func New(config Config) (*PostgresDB, error) {
	name := config.DriverName
	conf := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Username,
		config.DBName,
		config.Password,
		config.SSLMode,
	)

	database, err := sqlx.Open(name, conf)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresDB{database: database}, nil
}

func (d *PostgresDB) Close() error {
	return d.database.Close()
}

type Fields map[string]interface{}

func (d *PostgresDB) Insert(ret interface{}, table string, fields Fields) error {
	columns := make([]string, 0, len(fields))
	columnNames := make([]string, 0, len(fields))
	values := make([]interface{}, 0, len(fields))

	i := 1
	for key, value := range fields {
		columns = append(columns, fmt.Sprintf("$%d", i))
		columnNames = append(columnNames, key)
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		table,
		strings.Join(columnNames, ","),
		strings.Join(columns, ","),
	)

	return d.database.QueryRow(query, values...).Scan(ret)
}

func (d *PostgresDB) GetAll(ret interface{}, table string) error {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	return d.database.Select(ret, query)
}

func (d *PostgresDB) GetById(ret interface{}, table string, id int) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", table)
	return d.database.Get(ret, query, id)
}

func (d *PostgresDB) Get(ret interface{}, table string, fields Fields) error {
	columns := make([]string, 0, len(fields))
	values := make([]interface{}, 0, len(fields))

	i := 1
	for key, value := range fields {
		columns = append(columns, fmt.Sprintf("%s=$%d", key, i))
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE %s",
		table,
		strings.Join(columns, ","),
	)

	return d.database.Get(&ret, query, values...)
}

func (d *PostgresDB) Update(table string, id int, fields Fields) error {
	columns := make([]string, 0, len(fields))
	values := make([]interface{}, 0, len(fields) + 1)

	i := 2
	values = append(values, id)
	for key, value := range fields {
		columns = append(columns, fmt.Sprintf("%s=$%d", key, i))
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id=$1",
		table,
		strings.Join(columns, ","),
	)

	res, err := d.database.Exec(query, values...)
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

func (d *PostgresDB) Delete(table string, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", table)

	res, err := d.database.Exec(query, id)
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
