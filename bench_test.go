package go_db_bench_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"

	"go-db-bench/gorm"
	sqlDB "go-db-bench/sql"
	libpq "go-db-bench/sqlc/libpq"
	sqlcPgx "go-db-bench/sqlc/pgx"
	"go-db-bench/sqlx"
	"go-db-bench/xo"

	_ "github.com/lib/pq"
)

type Bencher interface {
	Name() string
	BenchmarkSelect(limit int32) func(b *testing.B)
}

var limits = []int32{1, 10, 100, 1000, 10000, 100000}

func Benchmark(b *testing.B) {
	dsn := os.Getenv("DATABASE_URL")
	var benchmarks = []Bencher{
		Must(sqlDB.New(dsn)),
		Must(sqlx.New(dsn)),
		Must(gorm.New(dsn)),
		sqlcPgx.New(Must(pgx.Connect(context.Background(), dsn))),
		libpq.New(Must(sql.Open("postgres", dsn))),
		xo.New(Must(sql.Open("postgres", dsn))),
	}

	for _, limit := range limits {
		fmt.Printf("\n======= limit:%d =========\n", limit)
		for _, bench := range benchmarks {
			b.Run(fmt.Sprintf("%s limit:%d", bench.Name(), limit), bench.BenchmarkSelect(limit))
		}
	}
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
