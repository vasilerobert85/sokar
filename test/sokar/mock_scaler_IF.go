// Code generated by MockGen. DO NOT EDIT.
// Source: sokar/iface/scaler_IF.go

// Package mock_sokar is a generated GoMock package.
package mock_sokar

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockScaler is a mock of Scaler interface
type MockScaler struct {
	ctrl     *gomock.Controller
	recorder *MockScalerMockRecorder
}

// MockScalerMockRecorder is the mock recorder for MockScaler
type MockScalerMockRecorder struct {
	mock *MockScaler
}

// NewMockScaler creates a new mock instance
func NewMockScaler(ctrl *gomock.Controller) *MockScaler {
	mock := &MockScaler{ctrl: ctrl}
	mock.recorder = &MockScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScaler) EXPECT() *MockScalerMockRecorder {
	return m.recorder
}

// ScaleTo mocks base method
func (m *MockScaler) ScaleTo(count uint, dryRun bool) error {
	ret := m.ctrl.Call(m, "ScaleTo", count, dryRun)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleTo indicates an expected call of ScaleTo
func (mr *MockScalerMockRecorder) ScaleTo(count, dryRun interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleTo", reflect.TypeOf((*MockScaler)(nil).ScaleTo), count, dryRun)
}

// GetCount mocks base method
func (m *MockScaler) GetCount() (uint, error) {
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount
func (mr *MockScalerMockRecorder) GetCount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockScaler)(nil).GetCount))
}
