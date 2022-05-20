// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_external_memory_capabilities is a generated GoMock package.
package mock_external_memory_capabilities

import (
	reflect "reflect"

	driver "github.com/CannibalVox/VKng/core/driver"
	khr_external_memory_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_memory_capabilities/driver"
	gomock "github.com/golang/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// VkGetPhysicalDeviceExternalBufferPropertiesKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalBufferInfo *khr_external_memory_capabilities_driver.VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *khr_external_memory_capabilities_driver.VkExternalBufferPropertiesKHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceExternalBufferPropertiesKHR", physicalDevice, pExternalBufferInfo, pExternalBufferProperties)
}

// VkGetPhysicalDeviceExternalBufferPropertiesKHR indicates an expected call of VkGetPhysicalDeviceExternalBufferPropertiesKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice, pExternalBufferInfo, pExternalBufferProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceExternalBufferPropertiesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceExternalBufferPropertiesKHR), physicalDevice, pExternalBufferInfo, pExternalBufferProperties)
}
