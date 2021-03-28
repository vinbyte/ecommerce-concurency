package http

import (
	"ecommerce-app/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase domain.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, ou domain.OrderUsecase) {
	handler := &OrderHandler{
		OrderUsecase: ou,
	}
	v1 := r.Group("v1")
	v1.POST("/cart/add", handler.AddToCart)
	v1.POST("/checkout", handler.Checkout)
}

func (oh *OrderHandler) AddToCart(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.AddToCartParam
	c.Bind(&param)
	code, res := oh.OrderUsecase.AddCart(ctx, param)
	c.JSON(code, res)
}

func (oh *OrderHandler) Checkout(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.CheckoutParam
	c.Bind(&param)
	code, res := oh.OrderUsecase.Checkout(ctx, param)
	c.JSON(code, res)
}
