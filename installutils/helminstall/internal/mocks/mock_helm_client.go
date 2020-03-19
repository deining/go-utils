// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock_internal is a generated GoMock package.
package mock_internal

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	afero "github.com/spf13/afero"
)

// MockFsHelper is a mock of FsHelper interface
type MockFsHelper struct {
	ctrl     *gomock.Controller
	recorder *MockFsHelperMockRecorder
}

// MockFsHelperMockRecorder is the mock recorder for MockFsHelper
type MockFsHelperMockRecorder struct {
	mock *MockFsHelper
}

// NewMockFsHelper creates a new mock instance
func NewMockFsHelper(ctrl *gomock.Controller) *MockFsHelper {
	mock := &MockFsHelper{ctrl: ctrl}
	mock.recorder = &MockFsHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFsHelper) EXPECT() *MockFsHelperMockRecorder {
	return m.recorder
}

// NewTempFile mocks base method
func (m *MockFsHelper) NewTempFile(dir, prefix string) (afero.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTempFile", dir, prefix)
	ret0, _ := ret[0].(afero.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTempFile indicates an expected call of NewTempFile
func (mr *MockFsHelperMockRecorder) NewTempFile(dir, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTempFile", reflect.TypeOf((*MockFsHelper)(nil).NewTempFile), dir, prefix)
}

// WriteFile mocks base method
func (m *MockFsHelper) WriteFile(filename string, data []byte, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockFsHelperMockRecorder) WriteFile(filename, data, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFsHelper)(nil).WriteFile), filename, data, perm)
}

// RemoveAll mocks base method
func (m *MockFsHelper) RemoveAll(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockFsHelperMockRecorder) RemoveAll(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFsHelper)(nil).RemoveAll), path)
}