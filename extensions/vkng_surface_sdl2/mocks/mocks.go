// Code generated by MockGen. DO NOT EDIT.
// Source: extension.go

// Package mock_surface_sdl2 is a generated GoMock package.
package mock_surface_sdl2

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_surface "github.com/CannibalVox/VKng/extensions/khr_surface"
	gomock "github.com/golang/mock/gomock"
	sdl "github.com/veandco/go-sdl2/sdl"
)

// MockExtension is a mock of Extension interface.
type MockExtension struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionMockRecorder
}

// MockExtensionMockRecorder is the mock recorder for MockExtension.
type MockExtensionMockRecorder struct {
	mock *MockExtension
}

// NewMockExtension creates a new mock instance.
func NewMockExtension(ctrl *gomock.Controller) *MockExtension {
	mock := &MockExtension{ctrl: ctrl}
	mock.recorder = &MockExtensionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension) EXPECT() *MockExtensionMockRecorder {
	return m.recorder
}

// CreateSurface mocks base method.
func (m *MockExtension) CreateSurface(instance core1_0.Instance, window *sdl.Window) (khr_surface.Surface, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSurface", instance, window)
	ret0, _ := ret[0].(khr_surface.Surface)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSurface indicates an expected call of CreateSurface.
func (mr *MockExtensionMockRecorder) CreateSurface(instance, window interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSurface", reflect.TypeOf((*MockExtension)(nil).CreateSurface), instance, window)
}
