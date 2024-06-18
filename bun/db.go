package bun

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Product struct {
	ID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name        string
	Description string
	Categories  []string `bun:",array"`
	Price       float64
	Features    []string `bun:",array"`
	Color       string
	Material    string
	UPC         string
}

type DB struct {
	bun *bun.DB
}

func (d *DB) Products(ctx context.Context, limit int) ([]Product, error) {
	var products []Product
	if err := d.bun.NewSelect().Model(&products).Limit(limit).Scan(ctx); err != nil {
		return nil, err
	}
	return products, nil
}

func (d *DB) Name() string {
	return "bun"
}

func (d *DB) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := d.Products(context.Background(), int(limit)); err != nil {
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
			// clear data
			b.StopTimer()
			if _, err := d.bun.NewTruncateTable().Model(&Product{}).Exec(context.Background()); err != nil {
				b.Fatal(err)
			}
			b.StartTimer()
			if _, err := d.bun.NewInsert().Model(&products).ExcludeColumn("id").Exec(context.Background()); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func New(db *sql.DB) *DB {
	return &DB{
		bun: bun.NewDB(
			db,
			pgdialect.New(),
		),
	}
}
