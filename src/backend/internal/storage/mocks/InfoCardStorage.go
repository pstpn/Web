// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "course/internal/service/dto"

	mock "github.com/stretchr/testify/mock"

	model "course/internal/model"
)

// InfoCardStorage is an autogenerated mock type for the InfoCardStorage type
type InfoCardStorage struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) Create(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.InfoCard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateInfoCardRequest) (*model.InfoCard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.CreateInfoCardRequest) *model.InfoCard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.InfoCard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.CreateInfoCardRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.DeleteInfoCardRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmployeeID provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) GetByEmployeeID(ctx context.Context, request *dto.GetInfoCardByEmployeeIDRequest) (*model.InfoCard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmployeeID")
	}

	var r0 *model.InfoCard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetInfoCardByEmployeeIDRequest) (*model.InfoCard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetInfoCardByEmployeeIDRequest) *model.InfoCard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.InfoCard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetInfoCardByEmployeeIDRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) GetByID(ctx context.Context, request *dto.GetInfoCardByIDRequest) (*model.InfoCard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *model.InfoCard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetInfoCardByIDRequest) (*model.InfoCard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.GetInfoCardByIDRequest) *model.InfoCard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.InfoCard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.GetInfoCardByIDRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*model.FullInfoCard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListInfoCardsRequest) []*model.FullInfoCard); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.FullInfoCard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListInfoCardsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, request
func (_m *InfoCardStorage) Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ValidateInfoCardRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewInfoCardStorage creates a new instance of InfoCardStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfoCardStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *InfoCardStorage {
	mock := &InfoCardStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}