// Code generated by MockGen. DO NOT EDIT.
// Source: runnable.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRunnable is a mock of Runnable interface
type MockRunnable struct {
	ctrl     *gomock.Controller
	recorder *MockRunnableMockRecorder
}

// MockRunnableMockRecorder is the mock recorder for MockRunnable
type MockRunnableMockRecorder struct {
	mock *MockRunnable
}

// NewMockRunnable creates a new mock instance
func NewMockRunnable(ctrl *gomock.Controller) *MockRunnable {
	mock := &MockRunnable{ctrl: ctrl}
	mock.recorder = &MockRunnableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRunnable) EXPECT() *MockRunnableMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockRunnable) Run() {
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run
func (mr *MockRunnableMockRecorder) Run() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRunnable)(nil).Run))
}

// Join mocks base method
func (m *MockRunnable) Join() {
	m.ctrl.Call(m, "Join")
}

// Join indicates an expected call of Join
func (mr *MockRunnableMockRecorder) Join() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockRunnable)(nil).Join))
}

// Stop mocks base method
func (m *MockRunnable) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockRunnableMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRunnable)(nil).Stop))
}

// GetName mocks base method
func (m *MockRunnable) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockRunnableMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockRunnable)(nil).GetName))
}
