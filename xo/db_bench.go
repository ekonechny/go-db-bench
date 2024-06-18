package xo

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

type XO struct {
	db *sql.DB
}

func (xo *XO) Name() string {
	return "xo/libpq"
}

func New(db *sql.DB) *XO {
	return &XO{db: db}
}

func (xo *XO) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := ProductsByLimit(context.Background(), xo.db, int(limit)); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func (xo *XO) BenchmarkBulkInsert(rows []gofakeit.ProductInfo) func(b *testing.B) {
	// meh https://github.com/xo/xo/issues/322
	return func(b *testing.B) {
		b.Skip("xo not support bulk")
	}
}
