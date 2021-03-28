package domain

import (
	"context"
	"time"
)

// OrderUsecase is usecases for order entity
type OrderUsecase interface {
	AddCart(ctx context.Context, param AddToCartParam) (code int, response interface{})
	Checkout(ctx context.Context, param CheckoutParam) (code int, response interface{})
}

// OrderRepository is repositories for order entity
type OrderRepository interface {
	CreateCart(ctx context.Context, userID int) (cartID int, err error)
	InsertCartItems(ctx context.Context, cartID int, productCode string, Qty int) error
	CheckUserByID(ctx context.Context, userID int) (int, error)
	GetCartData(ctx context.Context, cartID int, isRowLocking bool) (CartData, error)
	CreateOrders(ctx context.Context, userID int) (int, error)
	InsertOrderItems(ctx context.Context, orderID int, productCode string, qty int) error
	DeleteCart(ctx context.Context, cartID int) error
}

// AddToCartParam parameter for "/cart/add" endpoint
type AddToCartParam struct {
	UserIDStr     string `form:"user_id"`
	UserID        int
	ProductCodes  []string `form:"product_code"`
	QuantitiesStr []string `form:"qty"`
	Quantities    []int
}

// AddToCartResponse is response for "/cart/add" endpoint
type AddToCartResponse struct {
	CartID int `json:"cart_id"`
}

// CheckoutParam is param for checkout endpoint
type CheckoutParam struct {
	CartIDStr string `form:"cart_id"`
	CartID    int
}

// CartItem is data for cart item
type CartItem struct {
	ProductCode string
	Qty         int
	Date        time.Time
}

// CartData containing cart data including its details
type CartData struct {
	CartID int
	UserID int
	Date   time.Time
	Items  []CartItem
}

// CheckoutResponse is response for checkout endpoint
type CheckoutResponse struct {
	OrderID int `json:"order_id"`
}
