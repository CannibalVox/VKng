// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_create_renderpass2 is a generated GoMock package.
package mock_create_renderpass2

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	driver "github.com/CannibalVox/VKng/core/driver"
	khr_create_renderpass2 "github.com/CannibalVox/VKng/extensions/khr_create_renderpass2"
	gomock "github.com/golang/mock/gomock"
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

// CmdBeginRenderPass2 mocks base method.
func (m *MockExtension) CmdBeginRenderPass2(commandBuffer core1_0.CommandBuffer, renderPassBegin core1_0.RenderPassBeginOptions, subpassBegin khr_create_renderpass2.SubpassBeginOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdBeginRenderPass2", commandBuffer, renderPassBegin, subpassBegin)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdBeginRenderPass2 indicates an expected call of CmdBeginRenderPass2.
func (mr *MockExtensionMockRecorder) CmdBeginRenderPass2(commandBuffer, renderPassBegin, subpassBegin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdBeginRenderPass2", reflect.TypeOf((*MockExtension)(nil).CmdBeginRenderPass2), commandBuffer, renderPassBegin, subpassBegin)
}

// CmdEndRenderPass2 mocks base method.
func (m *MockExtension) CmdEndRenderPass2(commandBuffer core1_0.CommandBuffer, subpassEnd khr_create_renderpass2.SubpassEndOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdEndRenderPass2", commandBuffer, subpassEnd)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdEndRenderPass2 indicates an expected call of CmdEndRenderPass2.
func (mr *MockExtensionMockRecorder) CmdEndRenderPass2(commandBuffer, subpassEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdEndRenderPass2", reflect.TypeOf((*MockExtension)(nil).CmdEndRenderPass2), commandBuffer, subpassEnd)
}

// CmdNextSubpass2 mocks base method.
func (m *MockExtension) CmdNextSubpass2(commandBuffer core1_0.CommandBuffer, subpassBegin khr_create_renderpass2.SubpassBeginOptions, subpassEnd khr_create_renderpass2.SubpassEndOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CmdNextSubpass2", commandBuffer, subpassBegin, subpassEnd)
	ret0, _ := ret[0].(error)
	return ret0
}

// CmdNextSubpass2 indicates an expected call of CmdNextSubpass2.
func (mr *MockExtensionMockRecorder) CmdNextSubpass2(commandBuffer, subpassBegin, subpassEnd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CmdNextSubpass2", reflect.TypeOf((*MockExtension)(nil).CmdNextSubpass2), commandBuffer, subpassBegin, subpassEnd)
}

// CreateRenderPass2 mocks base method.
func (m *MockExtension) CreateRenderPass2(device core1_0.Device, allocator *driver.AllocationCallbacks, options khr_create_renderpass2.RenderPassCreateOptions) (core1_0.RenderPass, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRenderPass2", device, allocator, options)
	ret0, _ := ret[0].(core1_0.RenderPass)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRenderPass2 indicates an expected call of CreateRenderPass2.
func (mr *MockExtensionMockRecorder) CreateRenderPass2(device, allocator, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRenderPass2", reflect.TypeOf((*MockExtension)(nil).CreateRenderPass2), device, allocator, options)
}
