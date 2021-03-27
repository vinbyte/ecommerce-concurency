package usecase

import (
	"context"
	"ecommerce-app/domain"
	"os"
	"strconv"
)

// GetAllProducts logic process for getting list of products
func (pu *productUsecase) GetAllProducts(ctx context.Context, param domain.GetProductListParam) (int, interface{}) {
	// prepare the base response for success and error
	var successResponse domain.SuccessBaseResponse
	var errorResponse domain.ErrorBaseResponse
	var response interface{}
	var code int
	var productItems domain.ItemListResponse

	//set the timeout to context
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	//validate the parameter
	if param.LimitStr == "" {
		param.Limit = 10
	} else {
		num, err := strconv.Atoi(param.LimitStr)
		if err != nil {
			param.Limit = 10
		} else {
			param.Limit = num
		}
	}
	if param.OffsetStr != "" {
		num, err := strconv.Atoi(param.OffsetStr)
		if err != nil {
			param.Offset = 0
		} else {
			param.Offset = num
		}
	}

	//start getting from database
	products, err := pu.productRepo.GetAllProducts(ctx, param.Offset, param.Limit)
	if err != nil {
		if os.Getenv("MODE") != "production" {
			errorResponse.Error.Code = 500
			errorResponse.Error.Message = "failed get products"
			errData := domain.ErrorDataResponse{}
			errData.Message = "error GetAllProducts"
			errData.Reason = err.Error()
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errData)
		} else {
			errorResponse.Error.Code = 500
			errorResponse.Error.Message = "failed get products"
			errorResponse.Error.Errors = make([]domain.ErrorDataResponse, 0)
		}
		response = errorResponse
		code = 500
		return code, response
	}
	code = 200
	productItems.Items = products
	successResponse.Data = productItems
	response = successResponse

	return code, response
}
