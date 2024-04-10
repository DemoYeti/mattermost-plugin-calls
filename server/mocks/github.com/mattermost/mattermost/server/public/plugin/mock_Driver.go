// Code generated by mockery v2.40.1. DO NOT EDIT.

package plugin

import (
	driver "database/sql/driver"

	mock "github.com/stretchr/testify/mock"

	plugin "github.com/mattermost/mattermost/server/public/plugin"
)

// MockDriver is an autogenerated mock type for the Driver type
type MockDriver struct {
	mock.Mock
}

type MockDriver_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDriver) EXPECT() *MockDriver_Expecter {
	return &MockDriver_Expecter{mock: &_m.Mock}
}

// Conn provides a mock function with given fields: isMaster
func (_m *MockDriver) Conn(isMaster bool) (string, error) {
	ret := _m.Called(isMaster)

	if len(ret) == 0 {
		panic("no return value specified for Conn")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(bool) (string, error)); ok {
		return rf(isMaster)
	}
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(isMaster)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(isMaster)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_Conn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Conn'
type MockDriver_Conn_Call struct {
	*mock.Call
}

// Conn is a helper method to define mock.On call
//   - isMaster bool
func (_e *MockDriver_Expecter) Conn(isMaster interface{}) *MockDriver_Conn_Call {
	return &MockDriver_Conn_Call{Call: _e.mock.On("Conn", isMaster)}
}

func (_c *MockDriver_Conn_Call) Run(run func(isMaster bool)) *MockDriver_Conn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockDriver_Conn_Call) Return(_a0 string, _a1 error) *MockDriver_Conn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_Conn_Call) RunAndReturn(run func(bool) (string, error)) *MockDriver_Conn_Call {
	_c.Call.Return(run)
	return _c
}

// ConnClose provides a mock function with given fields: connID
func (_m *MockDriver) ConnClose(connID string) error {
	ret := _m.Called(connID)

	if len(ret) == 0 {
		panic("no return value specified for ConnClose")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(connID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_ConnClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnClose'
type MockDriver_ConnClose_Call struct {
	*mock.Call
}

// ConnClose is a helper method to define mock.On call
//   - connID string
func (_e *MockDriver_Expecter) ConnClose(connID interface{}) *MockDriver_ConnClose_Call {
	return &MockDriver_ConnClose_Call{Call: _e.mock.On("ConnClose", connID)}
}

func (_c *MockDriver_ConnClose_Call) Run(run func(connID string)) *MockDriver_ConnClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_ConnClose_Call) Return(_a0 error) *MockDriver_ConnClose_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_ConnClose_Call) RunAndReturn(run func(string) error) *MockDriver_ConnClose_Call {
	_c.Call.Return(run)
	return _c
}

// ConnExec provides a mock function with given fields: connID, q, args
func (_m *MockDriver) ConnExec(connID string, q string, args []driver.NamedValue) (plugin.ResultContainer, error) {
	ret := _m.Called(connID, q, args)

	if len(ret) == 0 {
		panic("no return value specified for ConnExec")
	}

	var r0 plugin.ResultContainer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, []driver.NamedValue) (plugin.ResultContainer, error)); ok {
		return rf(connID, q, args)
	}
	if rf, ok := ret.Get(0).(func(string, string, []driver.NamedValue) plugin.ResultContainer); ok {
		r0 = rf(connID, q, args)
	} else {
		r0 = ret.Get(0).(plugin.ResultContainer)
	}

	if rf, ok := ret.Get(1).(func(string, string, []driver.NamedValue) error); ok {
		r1 = rf(connID, q, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_ConnExec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnExec'
type MockDriver_ConnExec_Call struct {
	*mock.Call
}

// ConnExec is a helper method to define mock.On call
//   - connID string
//   - q string
//   - args []driver.NamedValue
func (_e *MockDriver_Expecter) ConnExec(connID interface{}, q interface{}, args interface{}) *MockDriver_ConnExec_Call {
	return &MockDriver_ConnExec_Call{Call: _e.mock.On("ConnExec", connID, q, args)}
}

func (_c *MockDriver_ConnExec_Call) Run(run func(connID string, q string, args []driver.NamedValue)) *MockDriver_ConnExec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]driver.NamedValue))
	})
	return _c
}

func (_c *MockDriver_ConnExec_Call) Return(_a0 plugin.ResultContainer, _a1 error) *MockDriver_ConnExec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_ConnExec_Call) RunAndReturn(run func(string, string, []driver.NamedValue) (plugin.ResultContainer, error)) *MockDriver_ConnExec_Call {
	_c.Call.Return(run)
	return _c
}

// ConnPing provides a mock function with given fields: connID
func (_m *MockDriver) ConnPing(connID string) error {
	ret := _m.Called(connID)

	if len(ret) == 0 {
		panic("no return value specified for ConnPing")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(connID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_ConnPing_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnPing'
type MockDriver_ConnPing_Call struct {
	*mock.Call
}

// ConnPing is a helper method to define mock.On call
//   - connID string
func (_e *MockDriver_Expecter) ConnPing(connID interface{}) *MockDriver_ConnPing_Call {
	return &MockDriver_ConnPing_Call{Call: _e.mock.On("ConnPing", connID)}
}

func (_c *MockDriver_ConnPing_Call) Run(run func(connID string)) *MockDriver_ConnPing_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_ConnPing_Call) Return(_a0 error) *MockDriver_ConnPing_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_ConnPing_Call) RunAndReturn(run func(string) error) *MockDriver_ConnPing_Call {
	_c.Call.Return(run)
	return _c
}

// ConnQuery provides a mock function with given fields: connID, q, args
func (_m *MockDriver) ConnQuery(connID string, q string, args []driver.NamedValue) (string, error) {
	ret := _m.Called(connID, q, args)

	if len(ret) == 0 {
		panic("no return value specified for ConnQuery")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, []driver.NamedValue) (string, error)); ok {
		return rf(connID, q, args)
	}
	if rf, ok := ret.Get(0).(func(string, string, []driver.NamedValue) string); ok {
		r0 = rf(connID, q, args)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, []driver.NamedValue) error); ok {
		r1 = rf(connID, q, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_ConnQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnQuery'
type MockDriver_ConnQuery_Call struct {
	*mock.Call
}

// ConnQuery is a helper method to define mock.On call
//   - connID string
//   - q string
//   - args []driver.NamedValue
func (_e *MockDriver_Expecter) ConnQuery(connID interface{}, q interface{}, args interface{}) *MockDriver_ConnQuery_Call {
	return &MockDriver_ConnQuery_Call{Call: _e.mock.On("ConnQuery", connID, q, args)}
}

func (_c *MockDriver_ConnQuery_Call) Run(run func(connID string, q string, args []driver.NamedValue)) *MockDriver_ConnQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]driver.NamedValue))
	})
	return _c
}

func (_c *MockDriver_ConnQuery_Call) Return(_a0 string, _a1 error) *MockDriver_ConnQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_ConnQuery_Call) RunAndReturn(run func(string, string, []driver.NamedValue) (string, error)) *MockDriver_ConnQuery_Call {
	_c.Call.Return(run)
	return _c
}

// RowsClose provides a mock function with given fields: rowsID
func (_m *MockDriver) RowsClose(rowsID string) error {
	ret := _m.Called(rowsID)

	if len(ret) == 0 {
		panic("no return value specified for RowsClose")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(rowsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_RowsClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsClose'
type MockDriver_RowsClose_Call struct {
	*mock.Call
}

// RowsClose is a helper method to define mock.On call
//   - rowsID string
func (_e *MockDriver_Expecter) RowsClose(rowsID interface{}) *MockDriver_RowsClose_Call {
	return &MockDriver_RowsClose_Call{Call: _e.mock.On("RowsClose", rowsID)}
}

func (_c *MockDriver_RowsClose_Call) Run(run func(rowsID string)) *MockDriver_RowsClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_RowsClose_Call) Return(_a0 error) *MockDriver_RowsClose_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsClose_Call) RunAndReturn(run func(string) error) *MockDriver_RowsClose_Call {
	_c.Call.Return(run)
	return _c
}

// RowsColumnTypeDatabaseTypeName provides a mock function with given fields: rowsID, index
func (_m *MockDriver) RowsColumnTypeDatabaseTypeName(rowsID string, index int) string {
	ret := _m.Called(rowsID, index)

	if len(ret) == 0 {
		panic("no return value specified for RowsColumnTypeDatabaseTypeName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, int) string); ok {
		r0 = rf(rowsID, index)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockDriver_RowsColumnTypeDatabaseTypeName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsColumnTypeDatabaseTypeName'
type MockDriver_RowsColumnTypeDatabaseTypeName_Call struct {
	*mock.Call
}

// RowsColumnTypeDatabaseTypeName is a helper method to define mock.On call
//   - rowsID string
//   - index int
func (_e *MockDriver_Expecter) RowsColumnTypeDatabaseTypeName(rowsID interface{}, index interface{}) *MockDriver_RowsColumnTypeDatabaseTypeName_Call {
	return &MockDriver_RowsColumnTypeDatabaseTypeName_Call{Call: _e.mock.On("RowsColumnTypeDatabaseTypeName", rowsID, index)}
}

func (_c *MockDriver_RowsColumnTypeDatabaseTypeName_Call) Run(run func(rowsID string, index int)) *MockDriver_RowsColumnTypeDatabaseTypeName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockDriver_RowsColumnTypeDatabaseTypeName_Call) Return(_a0 string) *MockDriver_RowsColumnTypeDatabaseTypeName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsColumnTypeDatabaseTypeName_Call) RunAndReturn(run func(string, int) string) *MockDriver_RowsColumnTypeDatabaseTypeName_Call {
	_c.Call.Return(run)
	return _c
}

// RowsColumnTypePrecisionScale provides a mock function with given fields: rowsID, index
func (_m *MockDriver) RowsColumnTypePrecisionScale(rowsID string, index int) (int64, int64, bool) {
	ret := _m.Called(rowsID, index)

	if len(ret) == 0 {
		panic("no return value specified for RowsColumnTypePrecisionScale")
	}

	var r0 int64
	var r1 int64
	var r2 bool
	if rf, ok := ret.Get(0).(func(string, int) (int64, int64, bool)); ok {
		return rf(rowsID, index)
	}
	if rf, ok := ret.Get(0).(func(string, int) int64); ok {
		r0 = rf(rowsID, index)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, int) int64); ok {
		r1 = rf(rowsID, index)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, int) bool); ok {
		r2 = rf(rowsID, index)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// MockDriver_RowsColumnTypePrecisionScale_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsColumnTypePrecisionScale'
type MockDriver_RowsColumnTypePrecisionScale_Call struct {
	*mock.Call
}

// RowsColumnTypePrecisionScale is a helper method to define mock.On call
//   - rowsID string
//   - index int
func (_e *MockDriver_Expecter) RowsColumnTypePrecisionScale(rowsID interface{}, index interface{}) *MockDriver_RowsColumnTypePrecisionScale_Call {
	return &MockDriver_RowsColumnTypePrecisionScale_Call{Call: _e.mock.On("RowsColumnTypePrecisionScale", rowsID, index)}
}

func (_c *MockDriver_RowsColumnTypePrecisionScale_Call) Run(run func(rowsID string, index int)) *MockDriver_RowsColumnTypePrecisionScale_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockDriver_RowsColumnTypePrecisionScale_Call) Return(_a0 int64, _a1 int64, _a2 bool) *MockDriver_RowsColumnTypePrecisionScale_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockDriver_RowsColumnTypePrecisionScale_Call) RunAndReturn(run func(string, int) (int64, int64, bool)) *MockDriver_RowsColumnTypePrecisionScale_Call {
	_c.Call.Return(run)
	return _c
}

// RowsColumns provides a mock function with given fields: rowsID
func (_m *MockDriver) RowsColumns(rowsID string) []string {
	ret := _m.Called(rowsID)

	if len(ret) == 0 {
		panic("no return value specified for RowsColumns")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(rowsID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockDriver_RowsColumns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsColumns'
type MockDriver_RowsColumns_Call struct {
	*mock.Call
}

// RowsColumns is a helper method to define mock.On call
//   - rowsID string
func (_e *MockDriver_Expecter) RowsColumns(rowsID interface{}) *MockDriver_RowsColumns_Call {
	return &MockDriver_RowsColumns_Call{Call: _e.mock.On("RowsColumns", rowsID)}
}

func (_c *MockDriver_RowsColumns_Call) Run(run func(rowsID string)) *MockDriver_RowsColumns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_RowsColumns_Call) Return(_a0 []string) *MockDriver_RowsColumns_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsColumns_Call) RunAndReturn(run func(string) []string) *MockDriver_RowsColumns_Call {
	_c.Call.Return(run)
	return _c
}

// RowsHasNextResultSet provides a mock function with given fields: rowsID
func (_m *MockDriver) RowsHasNextResultSet(rowsID string) bool {
	ret := _m.Called(rowsID)

	if len(ret) == 0 {
		panic("no return value specified for RowsHasNextResultSet")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(rowsID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockDriver_RowsHasNextResultSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsHasNextResultSet'
type MockDriver_RowsHasNextResultSet_Call struct {
	*mock.Call
}

// RowsHasNextResultSet is a helper method to define mock.On call
//   - rowsID string
func (_e *MockDriver_Expecter) RowsHasNextResultSet(rowsID interface{}) *MockDriver_RowsHasNextResultSet_Call {
	return &MockDriver_RowsHasNextResultSet_Call{Call: _e.mock.On("RowsHasNextResultSet", rowsID)}
}

func (_c *MockDriver_RowsHasNextResultSet_Call) Run(run func(rowsID string)) *MockDriver_RowsHasNextResultSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_RowsHasNextResultSet_Call) Return(_a0 bool) *MockDriver_RowsHasNextResultSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsHasNextResultSet_Call) RunAndReturn(run func(string) bool) *MockDriver_RowsHasNextResultSet_Call {
	_c.Call.Return(run)
	return _c
}

// RowsNext provides a mock function with given fields: rowsID, dest
func (_m *MockDriver) RowsNext(rowsID string, dest []driver.Value) error {
	ret := _m.Called(rowsID, dest)

	if len(ret) == 0 {
		panic("no return value specified for RowsNext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []driver.Value) error); ok {
		r0 = rf(rowsID, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_RowsNext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsNext'
type MockDriver_RowsNext_Call struct {
	*mock.Call
}

// RowsNext is a helper method to define mock.On call
//   - rowsID string
//   - dest []driver.Value
func (_e *MockDriver_Expecter) RowsNext(rowsID interface{}, dest interface{}) *MockDriver_RowsNext_Call {
	return &MockDriver_RowsNext_Call{Call: _e.mock.On("RowsNext", rowsID, dest)}
}

func (_c *MockDriver_RowsNext_Call) Run(run func(rowsID string, dest []driver.Value)) *MockDriver_RowsNext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]driver.Value))
	})
	return _c
}

func (_c *MockDriver_RowsNext_Call) Return(_a0 error) *MockDriver_RowsNext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsNext_Call) RunAndReturn(run func(string, []driver.Value) error) *MockDriver_RowsNext_Call {
	_c.Call.Return(run)
	return _c
}

// RowsNextResultSet provides a mock function with given fields: rowsID
func (_m *MockDriver) RowsNextResultSet(rowsID string) error {
	ret := _m.Called(rowsID)

	if len(ret) == 0 {
		panic("no return value specified for RowsNextResultSet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(rowsID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_RowsNextResultSet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowsNextResultSet'
type MockDriver_RowsNextResultSet_Call struct {
	*mock.Call
}

// RowsNextResultSet is a helper method to define mock.On call
//   - rowsID string
func (_e *MockDriver_Expecter) RowsNextResultSet(rowsID interface{}) *MockDriver_RowsNextResultSet_Call {
	return &MockDriver_RowsNextResultSet_Call{Call: _e.mock.On("RowsNextResultSet", rowsID)}
}

func (_c *MockDriver_RowsNextResultSet_Call) Run(run func(rowsID string)) *MockDriver_RowsNextResultSet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_RowsNextResultSet_Call) Return(_a0 error) *MockDriver_RowsNextResultSet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_RowsNextResultSet_Call) RunAndReturn(run func(string) error) *MockDriver_RowsNextResultSet_Call {
	_c.Call.Return(run)
	return _c
}

// Stmt provides a mock function with given fields: connID, q
func (_m *MockDriver) Stmt(connID string, q string) (string, error) {
	ret := _m.Called(connID, q)

	if len(ret) == 0 {
		panic("no return value specified for Stmt")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(connID, q)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(connID, q)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(connID, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_Stmt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stmt'
type MockDriver_Stmt_Call struct {
	*mock.Call
}

// Stmt is a helper method to define mock.On call
//   - connID string
//   - q string
func (_e *MockDriver_Expecter) Stmt(connID interface{}, q interface{}) *MockDriver_Stmt_Call {
	return &MockDriver_Stmt_Call{Call: _e.mock.On("Stmt", connID, q)}
}

func (_c *MockDriver_Stmt_Call) Run(run func(connID string, q string)) *MockDriver_Stmt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockDriver_Stmt_Call) Return(_a0 string, _a1 error) *MockDriver_Stmt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_Stmt_Call) RunAndReturn(run func(string, string) (string, error)) *MockDriver_Stmt_Call {
	_c.Call.Return(run)
	return _c
}

// StmtClose provides a mock function with given fields: stID
func (_m *MockDriver) StmtClose(stID string) error {
	ret := _m.Called(stID)

	if len(ret) == 0 {
		panic("no return value specified for StmtClose")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(stID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_StmtClose_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StmtClose'
type MockDriver_StmtClose_Call struct {
	*mock.Call
}

// StmtClose is a helper method to define mock.On call
//   - stID string
func (_e *MockDriver_Expecter) StmtClose(stID interface{}) *MockDriver_StmtClose_Call {
	return &MockDriver_StmtClose_Call{Call: _e.mock.On("StmtClose", stID)}
}

func (_c *MockDriver_StmtClose_Call) Run(run func(stID string)) *MockDriver_StmtClose_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_StmtClose_Call) Return(_a0 error) *MockDriver_StmtClose_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_StmtClose_Call) RunAndReturn(run func(string) error) *MockDriver_StmtClose_Call {
	_c.Call.Return(run)
	return _c
}

// StmtExec provides a mock function with given fields: stID, args
func (_m *MockDriver) StmtExec(stID string, args []driver.NamedValue) (plugin.ResultContainer, error) {
	ret := _m.Called(stID, args)

	if len(ret) == 0 {
		panic("no return value specified for StmtExec")
	}

	var r0 plugin.ResultContainer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []driver.NamedValue) (plugin.ResultContainer, error)); ok {
		return rf(stID, args)
	}
	if rf, ok := ret.Get(0).(func(string, []driver.NamedValue) plugin.ResultContainer); ok {
		r0 = rf(stID, args)
	} else {
		r0 = ret.Get(0).(plugin.ResultContainer)
	}

	if rf, ok := ret.Get(1).(func(string, []driver.NamedValue) error); ok {
		r1 = rf(stID, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_StmtExec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StmtExec'
type MockDriver_StmtExec_Call struct {
	*mock.Call
}

// StmtExec is a helper method to define mock.On call
//   - stID string
//   - args []driver.NamedValue
func (_e *MockDriver_Expecter) StmtExec(stID interface{}, args interface{}) *MockDriver_StmtExec_Call {
	return &MockDriver_StmtExec_Call{Call: _e.mock.On("StmtExec", stID, args)}
}

func (_c *MockDriver_StmtExec_Call) Run(run func(stID string, args []driver.NamedValue)) *MockDriver_StmtExec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]driver.NamedValue))
	})
	return _c
}

func (_c *MockDriver_StmtExec_Call) Return(_a0 plugin.ResultContainer, _a1 error) *MockDriver_StmtExec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_StmtExec_Call) RunAndReturn(run func(string, []driver.NamedValue) (plugin.ResultContainer, error)) *MockDriver_StmtExec_Call {
	_c.Call.Return(run)
	return _c
}

// StmtNumInput provides a mock function with given fields: stID
func (_m *MockDriver) StmtNumInput(stID string) int {
	ret := _m.Called(stID)

	if len(ret) == 0 {
		panic("no return value specified for StmtNumInput")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(stID)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockDriver_StmtNumInput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StmtNumInput'
type MockDriver_StmtNumInput_Call struct {
	*mock.Call
}

// StmtNumInput is a helper method to define mock.On call
//   - stID string
func (_e *MockDriver_Expecter) StmtNumInput(stID interface{}) *MockDriver_StmtNumInput_Call {
	return &MockDriver_StmtNumInput_Call{Call: _e.mock.On("StmtNumInput", stID)}
}

func (_c *MockDriver_StmtNumInput_Call) Run(run func(stID string)) *MockDriver_StmtNumInput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_StmtNumInput_Call) Return(_a0 int) *MockDriver_StmtNumInput_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_StmtNumInput_Call) RunAndReturn(run func(string) int) *MockDriver_StmtNumInput_Call {
	_c.Call.Return(run)
	return _c
}

// StmtQuery provides a mock function with given fields: stID, args
func (_m *MockDriver) StmtQuery(stID string, args []driver.NamedValue) (string, error) {
	ret := _m.Called(stID, args)

	if len(ret) == 0 {
		panic("no return value specified for StmtQuery")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []driver.NamedValue) (string, error)); ok {
		return rf(stID, args)
	}
	if rf, ok := ret.Get(0).(func(string, []driver.NamedValue) string); ok {
		r0 = rf(stID, args)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []driver.NamedValue) error); ok {
		r1 = rf(stID, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_StmtQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StmtQuery'
type MockDriver_StmtQuery_Call struct {
	*mock.Call
}

// StmtQuery is a helper method to define mock.On call
//   - stID string
//   - args []driver.NamedValue
func (_e *MockDriver_Expecter) StmtQuery(stID interface{}, args interface{}) *MockDriver_StmtQuery_Call {
	return &MockDriver_StmtQuery_Call{Call: _e.mock.On("StmtQuery", stID, args)}
}

func (_c *MockDriver_StmtQuery_Call) Run(run func(stID string, args []driver.NamedValue)) *MockDriver_StmtQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]driver.NamedValue))
	})
	return _c
}

func (_c *MockDriver_StmtQuery_Call) Return(_a0 string, _a1 error) *MockDriver_StmtQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_StmtQuery_Call) RunAndReturn(run func(string, []driver.NamedValue) (string, error)) *MockDriver_StmtQuery_Call {
	_c.Call.Return(run)
	return _c
}

// Tx provides a mock function with given fields: connID, opts
func (_m *MockDriver) Tx(connID string, opts driver.TxOptions) (string, error) {
	ret := _m.Called(connID, opts)

	if len(ret) == 0 {
		panic("no return value specified for Tx")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, driver.TxOptions) (string, error)); ok {
		return rf(connID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, driver.TxOptions) string); ok {
		r0 = rf(connID, opts)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, driver.TxOptions) error); ok {
		r1 = rf(connID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDriver_Tx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tx'
type MockDriver_Tx_Call struct {
	*mock.Call
}

// Tx is a helper method to define mock.On call
//   - connID string
//   - opts driver.TxOptions
func (_e *MockDriver_Expecter) Tx(connID interface{}, opts interface{}) *MockDriver_Tx_Call {
	return &MockDriver_Tx_Call{Call: _e.mock.On("Tx", connID, opts)}
}

func (_c *MockDriver_Tx_Call) Run(run func(connID string, opts driver.TxOptions)) *MockDriver_Tx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(driver.TxOptions))
	})
	return _c
}

func (_c *MockDriver_Tx_Call) Return(_a0 string, _a1 error) *MockDriver_Tx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDriver_Tx_Call) RunAndReturn(run func(string, driver.TxOptions) (string, error)) *MockDriver_Tx_Call {
	_c.Call.Return(run)
	return _c
}

// TxCommit provides a mock function with given fields: txID
func (_m *MockDriver) TxCommit(txID string) error {
	ret := _m.Called(txID)

	if len(ret) == 0 {
		panic("no return value specified for TxCommit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(txID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_TxCommit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxCommit'
type MockDriver_TxCommit_Call struct {
	*mock.Call
}

// TxCommit is a helper method to define mock.On call
//   - txID string
func (_e *MockDriver_Expecter) TxCommit(txID interface{}) *MockDriver_TxCommit_Call {
	return &MockDriver_TxCommit_Call{Call: _e.mock.On("TxCommit", txID)}
}

func (_c *MockDriver_TxCommit_Call) Run(run func(txID string)) *MockDriver_TxCommit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_TxCommit_Call) Return(_a0 error) *MockDriver_TxCommit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_TxCommit_Call) RunAndReturn(run func(string) error) *MockDriver_TxCommit_Call {
	_c.Call.Return(run)
	return _c
}

// TxRollback provides a mock function with given fields: txID
func (_m *MockDriver) TxRollback(txID string) error {
	ret := _m.Called(txID)

	if len(ret) == 0 {
		panic("no return value specified for TxRollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(txID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDriver_TxRollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxRollback'
type MockDriver_TxRollback_Call struct {
	*mock.Call
}

// TxRollback is a helper method to define mock.On call
//   - txID string
func (_e *MockDriver_Expecter) TxRollback(txID interface{}) *MockDriver_TxRollback_Call {
	return &MockDriver_TxRollback_Call{Call: _e.mock.On("TxRollback", txID)}
}

func (_c *MockDriver_TxRollback_Call) Run(run func(txID string)) *MockDriver_TxRollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDriver_TxRollback_Call) Return(_a0 error) *MockDriver_TxRollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDriver_TxRollback_Call) RunAndReturn(run func(string) error) *MockDriver_TxRollback_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDriver creates a new instance of MockDriver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDriver(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDriver {
	mock := &MockDriver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
