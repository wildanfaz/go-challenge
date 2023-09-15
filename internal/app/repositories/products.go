package repositories

import (
	"database/sql"

	"github.com/wildanfaz/go-challenge/internal/app/entities"
)

type Products interface {
	Add(entities.Product) error
}

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) Products {
	return &ProductsRepo{db: db}
}

func (p *ProductsRepo) Add(product entities.Product) error {
	return nil
}
