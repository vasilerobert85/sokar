// Code generated by MockGen. DO NOT EDIT.
// Source: sokar/iface/scaleEventEmitter_IF.go

// Package mock_sokar is a generated GoMock package.
package mock_sokar

import (
	gomock "github.com/golang/mock/gomock"
	iface "github.com/thomasobenaus/sokar/sokar/iface"
	reflect "reflect"
)

// MockScaleEventEmitter is a mock of ScaleEventEmitter interface
type MockScaleEventEmitter struct {
	ctrl     *gomock.Controller
	recorder *MockScaleEventEmitterMockRecorder
}

// MockScaleEventEmitterMockRecorder is the mock recorder for MockScaleEventEmitter
type MockScaleEventEmitterMockRecorder struct {
	mock *MockScaleEventEmitter
}

// NewMockScaleEventEmitter creates a new mock instance
func NewMockScaleEventEmitter(ctrl *gomock.Controller) *MockScaleEventEmitter {
	mock := &MockScaleEventEmitter{ctrl: ctrl}
	mock.recorder = &MockScaleEventEmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScaleEventEmitter) EXPECT() *MockScaleEventEmitterMockRecorder {
	return m.recorder
}

// Subscribe mocks base method
func (m *MockScaleEventEmitter) Subscribe(eventChannel chan iface.ScaleEvent) {
	m.ctrl.Call(m, "Subscribe", eventChannel)
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockScaleEventEmitterMockRecorder) Subscribe(eventChannel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockScaleEventEmitter)(nil).Subscribe), eventChannel)
}
