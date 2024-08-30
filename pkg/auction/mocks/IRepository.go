// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/kornharem08/auction_example/models"
	mock "github.com/stretchr/testify/mock"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *IRepository) Create(ctx context.Context, data models.Auction) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Auction) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetList provides a mock function with given fields: ctx
func (_m *IRepository) GetList(ctx context.Context) ([]models.Auction, error) {
	ret := _m.Called(ctx)

	var r0 []models.Auction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Auction, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Auction); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Auction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
