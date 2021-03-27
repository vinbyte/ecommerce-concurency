package http

import (
	"ecommerce-app/domain"

	"github.com/gin-gonic/gin"
)

// ProductHandler represent the httphandler for product
type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

func NewProductHandler(r *gin.Engine, pu domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: pu,
	}
	v1 := r.Group("v1")
	v1.GET("/products", handler.GetProductList)
}

func (ph *ProductHandler) GetProductList(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.GetProductListParam
	c.Bind(&param)
	code, res := ph.ProductUsecase.GetAllProducts(ctx, param)
	c.JSON(code, res)
}
