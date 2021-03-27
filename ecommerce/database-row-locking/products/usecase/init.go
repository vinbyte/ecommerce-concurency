package usecase

import (
	"ecommerce-app/app/helpers"
	"ecommerce-app/domain"
	"time"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
	helper         *helpers.Helpers
	contextTimeout time.Duration
}

// NewProductUsecase will create an object that represent domain.ProductUsecase
func NewProductUsecase(timeout time.Duration, pr domain.ProductRepository, h *helpers.Helpers) domain.ProductUsecase {
	return &productUsecase{
		productRepo:    pr,
		contextTimeout: timeout,
		helper:         h,
	}
}
