package xo

import (
	"context"
	"database/sql"
	"testing"
)

type XO struct {
	db *sql.DB
}

func (d *XO) Name() string {
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
