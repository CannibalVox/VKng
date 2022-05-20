// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_device_group_creation is a generated GoMock package.
package mock_device_group_creation

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	driver "github.com/CannibalVox/VKng/core/driver"
	khr_device_group_creation_driver "github.com/CannibalVox/VKng/extensions/khr_device_group_creation/driver"
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

// VkEnumeratePhysicalDeviceGroupsKHR mocks base method.
func (m *MockDriver) VkEnumeratePhysicalDeviceGroupsKHR(instance driver.VkInstance, pPhysicalDeviceGroupCount *driver.Uint32, pPhysicalDeviceGroupProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkEnumeratePhysicalDeviceGroupsKHR", instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkEnumeratePhysicalDeviceGroupsKHR indicates an expected call of VkEnumeratePhysicalDeviceGroupsKHR.
func (mr *MockDriverMockRecorder) VkEnumeratePhysicalDeviceGroupsKHR(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkEnumeratePhysicalDeviceGroupsKHR", reflect.TypeOf((*MockDriver)(nil).VkEnumeratePhysicalDeviceGroupsKHR), instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties)
}
