// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_external_semaphore_capabilities is a generated GoMock package.
package mock_external_semaphore_capabilities

import (
	reflect "reflect"

	driver "github.com/CannibalVox/VKng/core/driver"
	khr_external_semaphore_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities/driver"
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

// VkGetPhysicalDeviceExternalSemaphorePropertiesKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalSemaphoreInfo *khr_external_semaphore_capabilities_driver.VkPhysicalDeviceExternalSemaphoreInfoKHR, pExternalSemaphoreProperties *khr_external_semaphore_capabilities_driver.VkExternalSemaphorePropertiesKHR) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkGetPhysicalDeviceExternalSemaphorePropertiesKHR", physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties)
}

// VkGetPhysicalDeviceExternalSemaphorePropertiesKHR indicates an expected call of VkGetPhysicalDeviceExternalSemaphorePropertiesKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceExternalSemaphorePropertiesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceExternalSemaphorePropertiesKHR), physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties)
}
