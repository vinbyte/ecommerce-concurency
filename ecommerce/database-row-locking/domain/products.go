package domain

import (
	"context"
	"database/sql"
)

// ProductUsecase is usecases for product entity
type ProductUsecase interface {
	GetAllProducts(ctx context.Context, param GetProductListParam) (code int, response interface{})
	ResetStock(ctx context.Context) (code int, response interface{})
}

// ProductRepository is repositories for product entity
type ProductRepository interface {
	GetAllProducts(ctx context.Context, offset int, limit int) ([]Product, error)
	GetProductByCode(ctx context.Context, dbTx *sql.Tx, productCode string, isRowLocking bool) (Product, error)
	UpdateStock(ctx context.Context, dbTx *sql.Tx, productCode string, stock int) error
}

// Product is product data
type Product struct {
	ID    int    `json:"-"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}

// GetProductListParam is parameter for product list
type GetProductListParam struct {
	OffsetStr string `form:"offset"`
	LimitStr  string `form:"limit"`
	Offset    int
	Limit     int
}

type ItemListResponse struct {
	Items interface{} `json:"items"`
}
