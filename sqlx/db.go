package sqlx

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int64
	Description string
	Weight      sql.NullInt32
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Name() string {
	return "sqlx/libpq"
}

func (d *DB) Products(limit int32) ([]Product, error) {
	var books []Product
	err := d.db.Select(&books, "SELECT * FROM products LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (d *DB) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := d.Products(limit); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func New(dsn string) (*DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}
