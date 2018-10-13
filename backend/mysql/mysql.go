package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// ErrNotFoundOrNoOp entity hasn't been found or the query did nothing
	ErrNotFoundOrNoOp = errors.New("entity not found or no op")
)

type Query int

const (
	CreateDatabase Query = iota
	CreateGeolocations
)

// Be careful to order the enum in the execution order
var migration = []Query{
	CreateDatabase,
	CreateGeolocations,
}

var setupQueries = map[Query]string{
	CreateDatabase: `CREATE DATABASE IF NOT EXISTS test`,
	CreateGeolocations: `
		CREATE TABLE IF NOT EXISTS test.geolocations(
			ip NVARCHAR(16) NOT NULL PRIMARY KEY,
			city NVARCHAR(128) NOT NULL,
			country NVARCHAR(128) NOT NULL
	)`,
}

// MySQL implement the Repository interface with MySQL
type MySQL struct {
	*sqlx.DB
	Stmts map[Query]*sqlx.NamedStmt
}

func connect(cfg *Config) (*MySQL, error) {
	// Open the connection
	db, err := sqlx.Connect("mysql", cfg.MySQLConnectionString())
	if err != nil {
		return nil, err
	}

	return &MySQL{
		db, make(map[Query]*sqlx.NamedStmt),
	}, nil
}

// New instantiate a new MySQL
func New(cfg *Config) (*MySQL, error) {
	mysql, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	if err := mysql.migrate(); err != nil {
		return nil, err
	}
	err = mysql.prepareStatements(geolocationQueries)
	return mysql, err
}

func (m *MySQL) prepareStatements(resources ...map[Query]string) error {
	for _, queries := range resources {
		for key, query := range queries {
			stmt, err := m.PrepareNamed(query)
			if err != nil {
				return fmt.Errorf("Query: %d - Error: %v", key, err)
			}
			m.Stmts[key] = stmt
		}
	}
	return nil
}

func (m *MySQL) migrate() error {
	// We can't use prepared statement here, because the statements will fails
	// as long as the migration hasn't been completed
	for _, query := range migration {
		if _, err := m.Exec(setupQueries[query]); err != nil {
			return fmt.Errorf("Query: %d - Error: %v", query, err)
		}
	}
	return nil
}

// GeneratedID monad style helper to retrieve an auto incremented ID
func GeneratedID(result sql.Result, err error) (uint, error) {
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint(id), err
}

// RowAffected monad style helper to retrieve if the last query has been a NoOp
func RowAffected(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	// Would happen if the driver doesn't support this feature, I assume.
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFoundOrNoOp
	}
	return nil
}
