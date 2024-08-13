// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	models "github.com/AlejandroAldana99/yalo-challenge/models"
	mock "github.com/stretchr/testify/mock"
)

// IInteractionsRepository is an autogenerated mock type for the IInteractionsRepository type
type IInteractionsRepository struct {
	mock.Mock
}

// CollectUserInteraction provides a mock function with given fields: user
func (_m *IInteractionsRepository) CollectUserInteraction(user []models.UserInteraction) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CollectUserInteraction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]models.UserInteraction) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInteractionsByUserID provides a mock function with given fields: userID
func (_m *IInteractionsRepository) GetInteractionsByUserID(userID string) ([]models.UserInteraction, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetInteractionsByUserID")
	}

	var r0 []models.UserInteraction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.UserInteraction, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) []models.UserInteraction); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.UserInteraction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIInteractionsRepository creates a new instance of IInteractionsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIInteractionsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IInteractionsRepository {
	mock := &IInteractionsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
