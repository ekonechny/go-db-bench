package gorm

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	g *gorm.DB
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	// meh https://github.com/go-gorm/gorm/issues/4498
	Categories pq.StringArray `gorm:"type:text[]"`
	Price      float64
	Features   pq.StringArray `gorm:"type:text[]"`
	Color      string
	Material   string
	UPC        string
}

func (d *DB) Products(limit int) ([]Product, error) {
	var products []Product
	if err := d.g.Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
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
			// meh https://gorm.io/docs/delete.html#Block-Global-Delete
			if err := d.g.Where("1 = 1").Delete(&Product{}).Error; err != nil {
				b.Fatal(err)
			}
			b.StartTimer()
			if err := d.g.Omit("id").CreateInBatches(&products, len(products)).Error; err != nil {
				b.Fatal(err)
			}
		}
	}
}

func New(dsn string) (*DB, error) {
	g, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return &DB{
		g: g,
	}, nil
}
