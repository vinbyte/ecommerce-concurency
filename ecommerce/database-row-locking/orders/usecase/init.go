package usecase

import (
	"ecommerce-app/app/helpers"
	"ecommerce-app/domain"
	"time"
)

type orderUsecase struct {
	orderRepo      domain.OrderRepository
	productRepo    domain.ProductRepository
	helper         *helpers.Helpers
	contextTimeout time.Duration
}

func NewOrderUsecase(timeout time.Duration, or domain.OrderRepository, pr domain.ProductRepository, h *helpers.Helpers) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo:      or,
		productRepo:    pr,
		contextTimeout: timeout,
		helper:         h,
	}
}
