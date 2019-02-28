// Code generated by MockGen. DO NOT EDIT.
// Source: sokar/iface/scaler_IF.go

// Package mock_sokar is a generated GoMock package.
package mock_sokar

import (
	gomock "github.com/golang/mock/gomock"
	iface "github.com/thomasobenaus/sokar/sokar/iface"
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

// ScaleBy mocks base method
func (m *MockScaler) ScaleBy(amount int) iface.ScaleResult {
	ret := m.ctrl.Call(m, "ScaleBy", amount)
	ret0, _ := ret[0].(iface.ScaleResult)
	return ret0
}

// ScaleBy indicates an expected call of ScaleBy
func (mr *MockScalerMockRecorder) ScaleBy(amount interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleBy", reflect.TypeOf((*MockScaler)(nil).ScaleBy), amount)
}

// ScaleTo mocks base method
func (m *MockScaler) ScaleTo(count uint) iface.ScaleResult {
	ret := m.ctrl.Call(m, "ScaleTo", count)
	ret0, _ := ret[0].(iface.ScaleResult)
	return ret0
}

// ScaleTo indicates an expected call of ScaleTo
func (mr *MockScalerMockRecorder) ScaleTo(count interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleTo", reflect.TypeOf((*MockScaler)(nil).ScaleTo), count)
}