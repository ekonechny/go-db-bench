package sqlx

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

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

type DB struct {
	db *sqlx.DB
}

func (d *DB) Name() string {
	return "sqlx/libpq"
}

func (d *DB) Products(limit int32) ([]Product, error) {
	var products []Product
	err := d.db.Select(&products, "SELECT * FROM products LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d *DB) Insert(products []Product) error {
	if _, err := d.db.NamedExec(`INSERT INTO products (name, description, categories, price, features, color, material, upc)
        VALUES (:name, :description, :categories, :price, :features, :color, :material, :upc)`, products); err != nil {
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
			// clear data
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

func New(dsn string) (*DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}
