package usecase

import (
	"ecommerce-app/app/helpers"
	"ecommerce-app/domain"
	"time"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
	orderRepo      domain.OrderRepository
	helper         *helpers.Helpers
	contextTimeout time.Duration
}

// NewProductUsecase will create an object that represent domain.ProductUsecase
func NewProductUsecase(timeout time.Duration, pr domain.ProductRepository, or domain.OrderRepository, h *helpers.Helpers) domain.ProductUsecase {
	return &productUsecase{
		productRepo:    pr,
		orderRepo:      or,
		contextTimeout: timeout,
		helper:         h,
	}
}
