// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "homework/internal/storage/models"

	mock "github.com/stretchr/testify/mock"
)

// CacheInterface is an autogenerated mock type for the CacheInterface type
type CacheInterface struct {
	mock.Mock
}

type CacheInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheInterface) EXPECT() *CacheInterface_Expecter {
	return &CacheInterface_Expecter{mock: &_m.Mock}
}

// AddOrUpdate provides a mock function with given fields: ctx, id, m
func (_m *CacheInterface) AddOrUpdate(ctx context.Context, id uint64, m *models.Movie) {
	_m.Called(ctx, id, m)
}

// CacheInterface_AddOrUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddOrUpdate'
type CacheInterface_AddOrUpdate_Call struct {
	*mock.Call
}

// AddOrUpdate is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
//  - m *models.Movie
func (_e *CacheInterface_Expecter) AddOrUpdate(ctx interface{}, id interface{}, m interface{}) *CacheInterface_AddOrUpdate_Call {
	return &CacheInterface_AddOrUpdate_Call{Call: _e.mock.On("AddOrUpdate", ctx, id, m)}
}

func (_c *CacheInterface_AddOrUpdate_Call) Run(run func(ctx context.Context, id uint64, m *models.Movie)) *CacheInterface_AddOrUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*models.Movie))
	})
	return _c
}

func (_c *CacheInterface_AddOrUpdate_Call) Return() *CacheInterface_AddOrUpdate_Call {
	_c.Call.Return()
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CacheInterface) Delete(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CacheInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
func (_e *CacheInterface_Expecter) Delete(ctx interface{}, id interface{}) *CacheInterface_Delete_Call {
	return &CacheInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *CacheInterface_Delete_Call) Run(run func(ctx context.Context, id uint64)) *CacheInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *CacheInterface_Delete_Call) Return(_a0 error) *CacheInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *CacheInterface) GetById(ctx context.Context, id uint64) (*models.Movie, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Movie); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type CacheInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
func (_e *CacheInterface_Expecter) GetById(ctx interface{}, id interface{}) *CacheInterface_GetById_Call {
	return &CacheInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *CacheInterface_GetById_Call) Run(run func(ctx context.Context, id uint64)) *CacheInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *CacheInterface_GetById_Call) Return(_a0 *models.Movie, _a1 error) *CacheInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewCacheInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheInterface creates a new instance of CacheInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheInterface(t mockConstructorTestingTNewCacheInterface) *CacheInterface {
	mock := &CacheInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
