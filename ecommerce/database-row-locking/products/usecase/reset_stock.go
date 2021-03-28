package usecase

import (
	"context"
	"ecommerce-app/domain"
)

// ResetStock is logic process for reset the stock for re-testing
func (pu *productUsecase) ResetStock(ctx context.Context) (int, interface{}) {
	// prepare the base response for success and error
	var successResponse domain.SuccessBaseResponse
	// var errorResponse domain.ErrorBaseResponse
	var response interface{}
	var code int

	//set the timeout to context
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	_ = pu.productRepo.UpdateStock(ctx, "P1", 3)
	_ = pu.productRepo.UpdateStock(ctx, "P2", 2)
	code = 200
	successResponse.Data = new(struct{})
	response = successResponse

	return code, response
}
