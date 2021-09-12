// Code generated by MockGen. DO NOT EDIT.
// Source: iface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	loader "github.com/CannibalVox/VKng/core/loader"
	gomock "github.com/golang/mock/gomock"
)

// MockFramebuffer is a mock of Framebuffer interface.
type MockFramebuffer struct {
	ctrl     *gomock.Controller
	recorder *MockFramebufferMockRecorder
}

// MockFramebufferMockRecorder is the mock recorder for MockFramebuffer.
type MockFramebufferMockRecorder struct {
	mock *MockFramebuffer
}

// NewMockFramebuffer creates a new mock instance.
func NewMockFramebuffer(ctrl *gomock.Controller) *MockFramebuffer {
	mock := &MockFramebuffer{ctrl: ctrl}
	mock.recorder = &MockFramebufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFramebuffer) EXPECT() *MockFramebufferMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockFramebuffer) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockFramebufferMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockFramebuffer)(nil).Destroy))
}

// Handle mocks base method.
func (m *MockFramebuffer) Handle() loader.VkFramebuffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(loader.VkFramebuffer)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockFramebufferMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockFramebuffer)(nil).Handle))
}

// MockRenderPass is a mock of RenderPass interface.
type MockRenderPass struct {
	ctrl     *gomock.Controller
	recorder *MockRenderPassMockRecorder
}

// MockRenderPassMockRecorder is the mock recorder for MockRenderPass.
type MockRenderPassMockRecorder struct {
	mock *MockRenderPass
}

// NewMockRenderPass creates a new mock instance.
func NewMockRenderPass(ctrl *gomock.Controller) *MockRenderPass {
	mock := &MockRenderPass{ctrl: ctrl}
	mock.recorder = &MockRenderPassMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderPass) EXPECT() *MockRenderPassMockRecorder {
	return m.recorder
}

// Destroy mocks base method.
func (m *MockRenderPass) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockRenderPassMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockRenderPass)(nil).Destroy))
}

// Handle mocks base method.
func (m *MockRenderPass) Handle() loader.VkRenderPass {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(loader.VkRenderPass)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockRenderPassMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockRenderPass)(nil).Handle))
}