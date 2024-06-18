package sqlc

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func (d *Queries) Name() string {
	return "sqlc/libpq"
}

func (db *Queries) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := db.Products(context.Background(), limit); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func (d *Queries) BenchmarkBulkInsert(rows []gofakeit.ProductInfo) func(b *testing.B) {
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
			Upc:         rows[i].UPC,
		})
	}
	return func(b *testing.B) {
		ctx := context.Background()
		for n := 0; n < b.N; n++ {
			b.StopTimer()
			// clear data
			if err := d.ClearProducts(ctx); err != nil {
				b.Fatal(err)
			}
			b.StartTimer()
			if err := d.Insert(ctx, products); err != nil {
				b.Fatal(err)
			}

		}
	}
}
