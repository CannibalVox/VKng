// Code generated by MockGen. DO NOT EDIT.
// Source: driver.go

// Package mock_surface is a generated GoMock package.
package mock_surface

import (
	reflect "reflect"

	common "github.com/CannibalVox/VKng/core/common"
	driver "github.com/CannibalVox/VKng/core/driver"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
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

// VkDestroySurfaceKHR mocks base method.
func (m *MockDriver) VkDestroySurfaceKHR(instance driver.VkInstance, surface khr_surface_driver.VkSurfaceKHR, pAllocator *driver.VkAllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkDestroySurfaceKHR", instance, surface, pAllocator)
}

// VkDestroySurfaceKHR indicates an expected call of VkDestroySurfaceKHR.
func (mr *MockDriverMockRecorder) VkDestroySurfaceKHR(instance, surface, pAllocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkDestroySurfaceKHR", reflect.TypeOf((*MockDriver)(nil).VkDestroySurfaceKHR), instance, surface, pAllocator)
}

// VkGetPhysicalDeviceSurfaceCapabilitiesKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pSurfaceCapabilities *khr_surface_driver.VkSurfaceCapabilitiesKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDeviceSurfaceCapabilitiesKHR", physicalDevice, surface, pSurfaceCapabilities)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDeviceSurfaceCapabilitiesKHR indicates an expected call of VkGetPhysicalDeviceSurfaceCapabilitiesKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice, surface, pSurfaceCapabilities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceSurfaceCapabilitiesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceSurfaceCapabilitiesKHR), physicalDevice, surface, pSurfaceCapabilities)
}

// VkGetPhysicalDeviceSurfaceFormatsKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pSurfaceFormatCount *driver.Uint32, pSurfaceFormats *khr_surface_driver.VkSurfaceFormatKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDeviceSurfaceFormatsKHR", physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDeviceSurfaceFormatsKHR indicates an expected call of VkGetPhysicalDeviceSurfaceFormatsKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceSurfaceFormatsKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceSurfaceFormatsKHR), physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats)
}

// VkGetPhysicalDeviceSurfacePresentModesKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *khr_surface_driver.VkPresentModeKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDeviceSurfacePresentModesKHR", physicalDevice, surface, pPresentModeCount, pPresentModes)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDeviceSurfacePresentModesKHR indicates an expected call of VkGetPhysicalDeviceSurfacePresentModesKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice, surface, pPresentModeCount, pPresentModes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceSurfacePresentModesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceSurfacePresentModesKHR), physicalDevice, surface, pPresentModeCount, pPresentModes)
}

// VkGetPhysicalDeviceSurfaceSupportKHR mocks base method.
func (m *MockDriver) VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice driver.VkPhysicalDevice, queueFamilyIndex driver.Uint32, surface khr_surface_driver.VkSurfaceKHR, pSupported *driver.VkBool32) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetPhysicalDeviceSurfaceSupportKHR", physicalDevice, queueFamilyIndex, surface, pSupported)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetPhysicalDeviceSurfaceSupportKHR indicates an expected call of VkGetPhysicalDeviceSurfaceSupportKHR.
func (mr *MockDriverMockRecorder) VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice, queueFamilyIndex, surface, pSupported interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetPhysicalDeviceSurfaceSupportKHR", reflect.TypeOf((*MockDriver)(nil).VkGetPhysicalDeviceSurfaceSupportKHR), physicalDevice, queueFamilyIndex, surface, pSupported)
}
