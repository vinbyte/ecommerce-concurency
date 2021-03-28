package usecase

import (
	"context"
	"ecommerce-app/domain"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// AddCart is logic prosess for adding item to cart
func (ou *orderUsecase) AddCart(ctx context.Context, param domain.AddToCartParam) (int, interface{}) {
	// prepare the base response for success and error
	var successResponse domain.SuccessBaseResponse
	var errorResponse domain.ErrorBaseResponse
	var errorItem domain.ErrorDataResponse
	var response interface{}
	var code int

	//set the timeout to context
	ctx, cancel := context.WithTimeout(ctx, ou.contextTimeout)
	defer cancel()

	//validate param
	if param.UserIDStr == "" {
		errorItem.Message = "user_id is required"
		errorItem.Reason = "ErrValidate"
		errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
	} else {
		num, err := strconv.Atoi(param.UserIDStr)
		if err != nil {
			errorItem.Message = "user_id is invalid"
			errorItem.Reason = "ErrValidate"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		} else {
			_, err := ou.orderRepo.CheckUserByID(ctx, num)
			if err != nil {
				log.Error(err)
				errorItem.Message = "user not found"
				errorItem.Reason = "ErrValidate"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			} else {
				param.UserID = num
			}
		}
	}
	if len(param.ProductCodes) == 0 {
		errorItem.Message = "product_code is required"
		errorItem.Reason = "ErrValidate"
		errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
	}
	if len(param.QuantitiesStr) == 0 {
		errorItem.Message = "qty is required"
		errorItem.Reason = "ErrValidate"
		errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
	}
	if len(param.ProductCodes) != len(param.QuantitiesStr) {
		errorItem.Message = "product_code length must same with qty length"
		errorItem.Reason = "ErrValidate"
		errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
	} else {
		param.Quantities = make([]int, len(param.QuantitiesStr))
		for i, q := range param.QuantitiesStr {
			num, err := strconv.Atoi(q)
			if err != nil {
				errorItem.Message = "qty is invalid"
				errorItem.Reason = "ErrValidate"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
				break
			} else {
				param.Quantities[i] = num
			}
		}
		for _, pc := range param.ProductCodes {
			_, err := ou.productRepo.GetProductByCode(ctx, pc, false)
			if err != nil {
				log.Error(err)
				errorItem.Message = "product not found"
				errorItem.Reason = "ErrValidate"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
				break
			}
		}
	}
	if len(errorResponse.Error.Errors) > 0 {
		code = 400
		errorResponse.Error.Code = 400
		errorResponse.Error.Message = "failed add to cart"
		response = errorResponse
		return code, response
	}

	//begin the logic process
	//start the db trans
	err := ou.helper.BeginTrx(ctx, nil)
	if err != nil {
		log.Error(err)
		if os.Getenv("MODE") != "production" {
			errorItem.Message = err.Error()
			errorItem.Reason = "ErrTransaction"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		}
		errorResponse.Error.Code = 500
		errorResponse.Error.Message = "failed add to cart"
		response = errorResponse
		return code, response
	}
	//stock checking
	for i, pc := range param.ProductCodes {
		product, err := ou.productRepo.GetProductByCode(ctx, pc, true)
		if err != nil {
			_ = ou.helper.RollbackTrx()
			log.Error(err)
			if os.Getenv("MODE") != "production" {
				errorItem.Message = err.Error()
				errorItem.Reason = "ErrQuery"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			}
			errorResponse.Error.Code = 500
			code = 500
			errorResponse.Error.Message = "failed add to cart"
			response = errorResponse
			return code, response
		}
		//return error if stock not enough
		if product.Stock < param.Quantities[i] {
			_ = ou.helper.RollbackTrx()
			errorItem.Message = "stock of " + product.Name + " not enough for your order"
			errorItem.Reason = "ErrStock"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			errorResponse.Error.Code = 400
			code = 400
			errorResponse.Error.Message = "failed add to cart"
			response = errorResponse
			return code, response
		}
	}
	// insert into tabel cart, rollback if any error
	cartID, err := ou.orderRepo.CreateCart(ctx, param.UserID)
	if err != nil {
		_ = ou.helper.RollbackTrx()
		log.Error(err)
		if os.Getenv("MODE") != "production" {
			errorItem.Message = err.Error()
			errorItem.Reason = "ErrQuery"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		}
		errorResponse.Error.Code = 500
		code = 500
		errorResponse.Error.Message = "failed add to cart"
		response = errorResponse
		return code, response
	}
	// insert into tabel cart_items, rollback if any error
	for i, pc := range param.ProductCodes {
		qty := param.Quantities[i]
		err := ou.orderRepo.InsertCartItems(ctx, cartID, pc, qty)
		if err != nil {
			_ = ou.helper.RollbackTrx()
			log.Error(err)
			if os.Getenv("MODE") != "production" {
				errorItem.Message = err.Error()
				errorItem.Reason = "ErrQuery"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			}
			errorResponse.Error.Code = 500
			code = 500
			errorResponse.Error.Message = "failed add to cart"
			response = errorResponse
			return code, response
		}
	}
	_ = ou.helper.CommitTrx()

	//set the response
	var dataResponse domain.AddToCartResponse
	dataResponse.CartID = cartID
	successResponse.Data = dataResponse
	code = 200
	response = successResponse
	return code, response
}
