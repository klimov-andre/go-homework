// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "homework/internal/storage/models"
)

// DatabaseInterface is an autogenerated mock type for the DatabaseInterface type
type DatabaseInterface struct {
	mock.Mock
}

type DatabaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DatabaseInterface) EXPECT() *DatabaseInterface_Expecter {
	return &DatabaseInterface_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, m
func (_m *DatabaseInterface) Add(ctx context.Context, m *models.Movie) (uint64, error) {
	ret := _m.Called(ctx, m)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, *models.Movie) uint64); ok {
		r0 = rf(ctx, m)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Movie) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DatabaseInterface_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type DatabaseInterface_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - ctx context.Context
//  - m *models.Movie
func (_e *DatabaseInterface_Expecter) Add(ctx interface{}, m interface{}) *DatabaseInterface_Add_Call {
	return &DatabaseInterface_Add_Call{Call: _e.mock.On("Add", ctx, m)}
}

func (_c *DatabaseInterface_Add_Call) Run(run func(ctx context.Context, m *models.Movie)) *DatabaseInterface_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Movie))
	})
	return _c
}

func (_c *DatabaseInterface_Add_Call) Return(_a0 uint64, _a1 error) *DatabaseInterface_Add_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DatabaseInterface) Delete(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type DatabaseInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
func (_e *DatabaseInterface_Expecter) Delete(ctx interface{}, id interface{}) *DatabaseInterface_Delete_Call {
	return &DatabaseInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *DatabaseInterface_Delete_Call) Run(run func(ctx context.Context, id uint64)) *DatabaseInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *DatabaseInterface_Delete_Call) Return(_a0 error) *DatabaseInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetOneMovie provides a mock function with given fields: ctx, id
func (_m *DatabaseInterface) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
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

// DatabaseInterface_GetOneMovie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOneMovie'
type DatabaseInterface_GetOneMovie_Call struct {
	*mock.Call
}

// GetOneMovie is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
func (_e *DatabaseInterface_Expecter) GetOneMovie(ctx interface{}, id interface{}) *DatabaseInterface_GetOneMovie_Call {
	return &DatabaseInterface_GetOneMovie_Call{Call: _e.mock.On("GetOneMovie", ctx, id)}
}

func (_c *DatabaseInterface_GetOneMovie_Call) Run(run func(ctx context.Context, id uint64)) *DatabaseInterface_GetOneMovie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *DatabaseInterface_GetOneMovie_Call) Return(_a0 *models.Movie, _a1 error) *DatabaseInterface_GetOneMovie_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, limit, offset, order
func (_m *DatabaseInterface) List(ctx context.Context, limit int, offset int, order string) ([]*models.Movie, error) {
	ret := _m.Called(ctx, limit, offset, order)

	var r0 []*models.Movie
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) []*models.Movie); ok {
		r0 = rf(ctx, limit, offset, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, limit, offset, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DatabaseInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type DatabaseInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - ctx context.Context
//  - limit int
//  - offset int
//  - order string
func (_e *DatabaseInterface_Expecter) List(ctx interface{}, limit interface{}, offset interface{}, order interface{}) *DatabaseInterface_List_Call {
	return &DatabaseInterface_List_Call{Call: _e.mock.On("List", ctx, limit, offset, order)}
}

func (_c *DatabaseInterface_List_Call) Run(run func(ctx context.Context, limit int, offset int, order string)) *DatabaseInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int), args[3].(string))
	})
	return _c
}

func (_c *DatabaseInterface_List_Call) Return(_a0 []*models.Movie, _a1 error) *DatabaseInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: ctx, id, newMovie
func (_m *DatabaseInterface) Update(ctx context.Context, id uint64, newMovie *models.Movie) error {
	ret := _m.Called(ctx, id, newMovie)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *models.Movie) error); ok {
		r0 = rf(ctx, id, newMovie)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type DatabaseInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint64
//  - newMovie *models.Movie
func (_e *DatabaseInterface_Expecter) Update(ctx interface{}, id interface{}, newMovie interface{}) *DatabaseInterface_Update_Call {
	return &DatabaseInterface_Update_Call{Call: _e.mock.On("Update", ctx, id, newMovie)}
}

func (_c *DatabaseInterface_Update_Call) Run(run func(ctx context.Context, id uint64, newMovie *models.Movie)) *DatabaseInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*models.Movie))
	})
	return _c
}

func (_c *DatabaseInterface_Update_Call) Return(_a0 error) *DatabaseInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewDatabaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabaseInterface creates a new instance of DatabaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabaseInterface(t mockConstructorTestingTNewDatabaseInterface) *DatabaseInterface {
	mock := &DatabaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
