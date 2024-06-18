package go_db_bench_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	"go-db-bench/bun"
	"go-db-bench/gorm"
	sqlDB "go-db-bench/sql"
	sqlcLibPQ "go-db-bench/sqlc/libpq"
	sqlcPGX "go-db-bench/sqlc/pgx"
	"go-db-bench/sqlx"
	"go-db-bench/xo"
)

type Bencher interface {
	Name() string
	BenchmarkSelect(limit int32) func(b *testing.B)
	BenchmarkBulkInsert(rows []gofakeit.ProductInfo) func(b *testing.B)
}

var (
	limits    = []int32{100000}
	batchSize = []int32{1, 10, 100, 1000}
)

func Benchmark(b *testing.B) {
	dsn := os.Getenv("DATABASE_URL")
	var benchmarks = []Bencher{
		Must(sqlDB.New(dsn)),
		sqlcPGX.New(Must(pgx.Connect(context.Background(), dsn))),
		sqlcLibPQ.New(Must(sql.Open("postgres", dsn))),
		xo.New(Must(sql.Open("postgres", dsn))),
		Must(sqlx.New(dsn)),
		Must(gorm.New(dsn)),
		bun.New(Must(sql.Open("postgres", dsn))),
	}

	for _, limit := range limits {
		fmt.Printf("\n======= limit:%d =========\n", limit)
		for _, bench := range benchmarks {
			b.Run(fmt.Sprintf("%s limit:%d", bench.Name(), limit), bench.BenchmarkSelect(limit))
		}
	}
	var products = make([]gofakeit.ProductInfo, 100000)
	gofakeit.Slice(&products)
	for _, limit := range batchSize {
		fmt.Printf("\n======= insert:%d =========\n", limit)
		for _, bench := range benchmarks {
			b.Run(fmt.Sprintf("%s batchSize:%d", bench.Name(), limit), bench.BenchmarkBulkInsert(products[0:limit]))
		}
	}
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
