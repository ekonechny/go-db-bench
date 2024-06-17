package pgx

import (
	"context"
	"testing"
)

func (d *Queries) Name() string {
	return "sqlc/pgx"
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
