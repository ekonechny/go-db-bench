package gorm

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	g *gorm.DB
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Price       int64
	Description string
	Weight      sql.NullInt32
}

func (d *DB) Products(limit int) ([]Product, error) {
	var authors []Product
	if err := d.g.Limit(limit).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (d *DB) Name() string {
	return "gorm/pgx"
}

func (d *DB) BenchmarkSelect(limit int32) func(b *testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			if _, err := d.Products(int(limit)); err != nil {
				b.Fatal(err)
			}
		}
	}
}

func New(dsn string) (*DB, error) {
	g, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &DB{
		g: g,
	}, nil
}
