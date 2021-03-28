package http

import (
	"ecommerce-app/domain"

	"github.com/gin-gonic/gin"
)

// ProductHandler represent the httphandler for product
type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

// NewProductHandler initiate new http delivery layer
func NewProductHandler(r *gin.Engine, pu domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: pu,
	}
	v1 := r.Group("v1")
	v1.GET("/products", handler.GetProductList)
	v1.GET("/reset", handler.Reset)
}

// GetProductList is handler for products list endpoint
func (ph *ProductHandler) GetProductList(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.GetProductListParam
	c.Bind(&param)
	code, res := ph.ProductUsecase.GetAllProducts(ctx, param)
	c.JSON(code, res)
}

// Reset is handler for reset endpoint
func (ph *ProductHandler) Reset(c *gin.Context) {
	ctx := c.Request.Context()
	code, res := ph.ProductUsecase.ResetStock(ctx)
	c.JSON(code, res)
}
