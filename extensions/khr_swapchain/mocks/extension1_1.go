// Code generated by MockGen. DO NOT EDIT.
// Source: extiface.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	khr_surface "github.com/CannibalVox/VKng/extensions/khr_surface"
	khr_swapchain1_1 "github.com/CannibalVox/VKng/extensions/khr_swapchain/khr_swapchain1_1"
	gomock "github.com/golang/mock/gomock"
)

// MockExtension1_1 is a mock of Extension interface.
type MockExtension1_1 struct {
	ctrl     *gomock.Controller
	recorder *MockExtension1_1MockRecorder
}

// MockExtension1_1MockRecorder is the mock recorder for MockExtension1_1.
type MockExtension1_1MockRecorder struct {
	mock *MockExtension1_1
}

// NewMockExtension1_1 creates a new mock instance.
func NewMockExtension1_1(ctrl *gomock.Controller) *MockExtension1_1 {
	mock := &MockExtension1_1{ctrl: ctrl}
	mock.recorder = &MockExtension1_1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtension1_1) EXPECT() *MockExtension1_1MockRecorder {
	return m.recorder
}

// AcquireNextImage mocks base method.
func (m *MockExtension1_1) AcquireNextImage(device core1_0.Device, o khr_swapchain1_1.AcquireNextImageOptions) (int, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireNextImage", device, o)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AcquireNextImage indicates an expected call of AcquireNextImage.
func (mr *MockExtension1_1MockRecorder) AcquireNextImage(device, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireNextImage", reflect.TypeOf((*MockExtension1_1)(nil).AcquireNextImage), device, o)
}

// DeviceGroupPresentCapabilities mocks base method.
func (m *MockExtension1_1) DeviceGroupPresentCapabilities(device core1_0.Device, outData *khr_swapchain1_1.DeviceGroupPresentCapabilitiesOutData) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeviceGroupPresentCapabilities", device, outData)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeviceGroupPresentCapabilities indicates an expected call of DeviceGroupPresentCapabilities.
func (mr *MockExtension1_1MockRecorder) DeviceGroupPresentCapabilities(device, outData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceGroupPresentCapabilities", reflect.TypeOf((*MockExtension1_1)(nil).DeviceGroupPresentCapabilities), device, outData)
}

// DeviceGroupSurfacePresentModes mocks base method.
func (m *MockExtension1_1) DeviceGroupSurfacePresentModes(device core1_0.Device, surface khr_surface.Surface) (khr_swapchain1_1.DeviceGroupPresentModeFlags, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeviceGroupSurfacePresentModes", device, surface)
	ret0, _ := ret[0].(khr_swapchain1_1.DeviceGroupPresentModeFlags)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeviceGroupSurfacePresentModes indicates an expected call of DeviceGroupSurfacePresentModes.
func (mr *MockExtension1_1MockRecorder) DeviceGroupSurfacePresentModes(device, surface interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceGroupSurfacePresentModes", reflect.TypeOf((*MockExtension1_1)(nil).DeviceGroupSurfacePresentModes), device, surface)
}

// PhysicalDevicePresentRectangles mocks base method.
func (m *MockExtension1_1) PhysicalDevicePresentRectangles(physicalDevice core1_0.PhysicalDevice, surface khr_surface.Surface) ([]core1_0.Rect2D, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhysicalDevicePresentRectangles", physicalDevice, surface)
	ret0, _ := ret[0].([]core1_0.Rect2D)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PhysicalDevicePresentRectangles indicates an expected call of PhysicalDevicePresentRectangles.
func (mr *MockExtension1_1MockRecorder) PhysicalDevicePresentRectangles(physicalDevice, surface interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhysicalDevicePresentRectangles", reflect.TypeOf((*MockExtension1_1)(nil).PhysicalDevicePresentRectangles), physicalDevice, surface)
}
