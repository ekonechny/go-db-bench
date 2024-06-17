package sql

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
)

type DB struct {
	db *sql.DB
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int64
	Description string
	Weight      sql.NullInt32
}

func (d *DB) ListAuthors(limit int32) ([]Product, error) {
	var authors []Product
	rows, err := d.db.Query("SELECT * FROM products limit $1;", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Product
		if err := rows.Scan(&a.ID, &a.Name, &a.Price, &a.Description, &a.Weight); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, err
}

func (d *DB) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := d.ListAuthors(limit); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func (d *DB) Name() string {
	return "database/sql/libpq"
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, err
}
