package sql

import (
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Categories  pq.StringArray
	Price       float64
	Features    pq.StringArray
	Color       string
	Material    string
	UPC         string
}

func (d *DB) Products(limit int32) ([]Product, error) {
	var products []Product
	rows, err := d.db.Query("SELECT * FROM products limit $1;", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a Product
		err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.Description,
			&a.Categories,
			&a.Price,
			&a.Features,
			&a.Color,
			&a.Material,
			&a.UPC,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, err
}

func (d *DB) Insert(products []Product) error {
	// from https://pkg.go.dev/github.com/lib/pq#hdr-Bulk_imports
	txn, err := d.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := txn.Prepare(pq.CopyIn(
		"products",
		"name",
		"description",
		"categories",
		"price",
		"features",
		"color",
		"material",
		"upc"))
	if err != nil {
		return err
	}

	for _, p := range products {
		_, err = stmt.Exec(p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.UPC)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = stmt.Close()
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}
	return nil
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

func (d *DB) BenchmarkBulkInsert(rows []gofakeit.ProductInfo) func(b *testing.B) {
	var products = make([]Product, 0, len(rows))
	for i := range rows {
		products = append(products, Product{
			Name:        rows[i].Name,
			Description: rows[i].Description,
			Categories:  rows[i].Categories,
			Price:       rows[i].Price,
			Features:    rows[i].Features,
			Color:       rows[i].Color,
			Material:    rows[i].Material,
			UPC:         rows[i].UPC,
		})
	}
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			b.StopTimer()
			if _, err := d.db.Exec("TRUNCATE products;"); err != nil {
				b.Fatal(err)
			}
			b.StartTimer()
			if err := d.Insert(products); err != nil {
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
