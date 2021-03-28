package usecase

import (
	"context"
	"database/sql"
	"ecommerce-app/domain"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Checkout is logic prosess for checkout order and decrease the stock
func (ou *orderUsecase) Checkout(ctx context.Context, param domain.CheckoutParam) (int, interface{}) {
	// prepare the base response for success and error
	var successResponse domain.SuccessBaseResponse
	var errorResponse domain.ErrorBaseResponse
	var errorItem domain.ErrorDataResponse
	var response interface{}
	var code int

	//validate param
	if param.CartIDStr == "" {
		errorItem.Message = "cart_id is required"
		errorItem.Reason = "ErrValidate"
		errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
	} else {
		num, err := strconv.Atoi(param.CartIDStr)
		if err != nil {
			errorItem.Message = "cart_id is invalid"
			errorItem.Reason = "ErrValidate"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		} else {
			param.CartID = num
		}
	}
	if len(errorResponse.Error.Errors) > 0 {
		code = 400
		errorResponse.Error.Code = 400
		errorResponse.Error.Message = "failed to checkout"
		response = errorResponse
		return code, response
	}

	// start getting the cart data
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
		errorResponse.Error.Message = "failed to checkout"
		response = errorResponse
		return code, response
	}
	cartData, err := ou.orderRepo.GetCartData(ctx, param.CartID, true)
	if err != nil {
		//rollback if any error
		_ = ou.helper.RollbackTrx()
		log.Error("GetCartData ", err)
		errorResponse.Error.Code = 500
		code = 500
		if os.Getenv("MODE") != "production" {
			errorItem.Message = err.Error()
			if err == sql.ErrNoRows {
				errorResponse.Error.Code = 404
				code = 404
				errorItem.Message = "data not found"
			}
			errorItem.Reason = "ErrQuery"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		}
		errorResponse.Error.Message = "failed to checkout"
		response = errorResponse
		return code, response
	}
	stockList := make(map[string]int)
	for _, item := range cartData.Items {
		product, err := ou.productRepo.GetProductByCode(ctx, item.ProductCode, true)
		if err != nil {
			//rollback if any error
			_ = ou.helper.RollbackTrx()
			log.Error("GetProductByCode ", err)
			if os.Getenv("MODE") != "production" {
				errorItem.Message = err.Error()
				errorItem.Reason = "ErrQuery"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			}
			errorResponse.Error.Code = 500
			code = 500
			errorResponse.Error.Message = "failed to checkout"
			response = errorResponse
			return code, response
		}
		if product.Stock < item.Qty {
			_ = ou.helper.RollbackTrx()
			errorItem.Message = "stock of " + product.Name + " not enough for your order"
			errorItem.Reason = "ErrStock"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			errorResponse.Error.Code = 400
			code = 400
			errorResponse.Error.Message = "failed to checkout"
			response = errorResponse
			return code, response
		}
		stockList[product.Code] = product.Stock
	}
	orderID, err := ou.orderRepo.CreateOrders(ctx, cartData.UserID)
	if err != nil {
		_ = ou.helper.RollbackTrx()
		log.Error("CreateOrders ", err)
		if os.Getenv("MODE") != "production" {
			errorItem.Message = err.Error()
			errorItem.Reason = "ErrQuery"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		}
		errorResponse.Error.Code = 500
		code = 500
		errorResponse.Error.Message = "failed to checkout"
		response = errorResponse
		return code, response
	}
	for _, item := range cartData.Items {
		err := ou.orderRepo.InsertOrderItems(ctx, orderID, item.ProductCode, item.Qty)
		if err != nil {
			_ = ou.helper.RollbackTrx()
			log.Error("InsertOrderItems ", err)
			if os.Getenv("MODE") != "production" {
				errorItem.Message = err.Error()
				errorItem.Reason = "ErrQuery"
				errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
			}
			errorResponse.Error.Code = 500
			code = 500
			errorResponse.Error.Message = "failed to checkout"
			response = errorResponse
			return code, response
		} else {
			//decrease stock
			oldStock := stockList[item.ProductCode]
			newStock := oldStock - item.Qty
			err = ou.productRepo.UpdateStock(ctx, item.ProductCode, newStock)
			if err != nil {
				_ = ou.helper.RollbackTrx()
				log.Error("UpdateStock ", err)
				if os.Getenv("MODE") != "production" {
					errorItem.Message = err.Error()
					errorItem.Reason = "ErrQuery"
					errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
				}
				errorResponse.Error.Code = 500
				code = 500
				errorResponse.Error.Message = "failed to checkout"
				response = errorResponse
				return code, response
			}
		}
	}
	err = ou.orderRepo.DeleteCart(ctx, cartData.CartID)
	if err != nil {
		_ = ou.helper.RollbackTrx()
		log.Error("DeleteCart ", err)
		if os.Getenv("MODE") != "production" {
			errorItem.Message = err.Error()
			errorItem.Reason = "ErrQuery"
			errorResponse.Error.Errors = append(errorResponse.Error.Errors, errorItem)
		}
		errorResponse.Error.Code = 500
		code = 500
		errorResponse.Error.Message = "failed to checkout"
		response = errorResponse
		return code, response
	}

	_ = ou.helper.CommitTrx()
	var responseData domain.CheckoutResponse
	responseData.OrderID = orderID
	successResponse.Data = responseData
	response = successResponse
	code = 200

	return code, response
}
