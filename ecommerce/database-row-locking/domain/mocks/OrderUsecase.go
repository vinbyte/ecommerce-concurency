// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "ecommerce-app/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderUsecase is an autogenerated mock type for the OrderUsecase type
type OrderUsecase struct {
	mock.Mock
}

// AddCart provides a mock function with given fields: ctx, param
func (_m *OrderUsecase) AddCart(ctx context.Context, param domain.AddToCartParam) (int, interface{}) {
	ret := _m.Called(ctx, param)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, domain.AddToCartParam) int); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(context.Context, domain.AddToCartParam) interface{}); ok {
		r1 = rf(ctx, param)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}