// Code generated by MockGen. DO NOT EDIT.
// Source: ./query_builder.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	goqu "github.com/doug-martin/goqu/v9"
	gomock "github.com/golang/mock/gomock"
)

// MockQueryBuilder is a mock of QueryBuilder interface.
type MockQueryBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockQueryBuilderMockRecorder
}

// MockQueryBuilderMockRecorder is the mock recorder for MockQueryBuilder.
type MockQueryBuilderMockRecorder struct {
	mock *MockQueryBuilder
}

// NewMockQueryBuilder creates a new mock instance.
func NewMockQueryBuilder(ctrl *gomock.Controller) *MockQueryBuilder {
	mock := &MockQueryBuilder{ctrl: ctrl}
	mock.recorder = &MockQueryBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryBuilder) EXPECT() *MockQueryBuilderMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockQueryBuilder) Delete(table interface{}) *goqu.DeleteDataset {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", table)
	ret0, _ := ret[0].(*goqu.DeleteDataset)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockQueryBuilderMockRecorder) Delete(table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQueryBuilder)(nil).Delete), table)
}

// From mocks base method.
func (m *MockQueryBuilder) From(cols ...interface{}) *goqu.SelectDataset {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range cols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "From", varargs...)
	ret0, _ := ret[0].(*goqu.SelectDataset)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockQueryBuilderMockRecorder) From(cols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockQueryBuilder)(nil).From), cols...)
}

// Insert mocks base method.
func (m *MockQueryBuilder) Insert(table interface{}) *goqu.InsertDataset {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", table)
	ret0, _ := ret[0].(*goqu.InsertDataset)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockQueryBuilderMockRecorder) Insert(table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockQueryBuilder)(nil).Insert), table)
}

// Select mocks base method.
func (m *MockQueryBuilder) Select(cols ...interface{}) *goqu.SelectDataset {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range cols {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(*goqu.SelectDataset)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockQueryBuilderMockRecorder) Select(cols ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockQueryBuilder)(nil).Select), cols...)
}

// Truncate mocks base method.
func (m *MockQueryBuilder) Truncate(table ...interface{}) *goqu.TruncateDataset {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range table {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Truncate", varargs...)
	ret0, _ := ret[0].(*goqu.TruncateDataset)
	return ret0
}

// Truncate indicates an expected call of Truncate.
func (mr *MockQueryBuilderMockRecorder) Truncate(table ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockQueryBuilder)(nil).Truncate), table...)
}

// Update mocks base method.
func (m *MockQueryBuilder) Update(table interface{}) *goqu.UpdateDataset {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", table)
	ret0, _ := ret[0].(*goqu.UpdateDataset)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockQueryBuilderMockRecorder) Update(table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockQueryBuilder)(nil).Update), table)
}
