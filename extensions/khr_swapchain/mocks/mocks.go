// Code generated by MockGen. DO NOT EDIT.
// Source: extension.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	reflect "reflect"

	core "github.com/CannibalVox/VKng/core"
	common "github.com/CannibalVox/VKng/core/common"
	driver "github.com/CannibalVox/VKng/core/driver"
	khr_swapchain "github.com/CannibalVox/VKng/extensions/khr_swapchain"
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

// VkAcquireNextImageKHR mocks base method.
func (m *MockDriver) VkAcquireNextImageKHR(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkAcquireNextImageKHR", device, swapchain, timeout, semaphore, fence, pImageIndex)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkAcquireNextImageKHR indicates an expected call of VkAcquireNextImageKHR.
func (mr *MockDriverMockRecorder) VkAcquireNextImageKHR(device, swapchain, timeout, semaphore, fence, pImageIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkAcquireNextImageKHR", reflect.TypeOf((*MockDriver)(nil).VkAcquireNextImageKHR), device, swapchain, timeout, semaphore, fence, pImageIndex)
}

// VkCreateSwapchainKHR mocks base method.
func (m *MockDriver) VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *khr_swapchain.VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *khr_swapchain.VkSwapchainKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkCreateSwapchainKHR", device, pCreateInfo, pAllocator, pSwapchain)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkCreateSwapchainKHR indicates an expected call of VkCreateSwapchainKHR.
func (mr *MockDriverMockRecorder) VkCreateSwapchainKHR(device, pCreateInfo, pAllocator, pSwapchain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkCreateSwapchainKHR", reflect.TypeOf((*MockDriver)(nil).VkCreateSwapchainKHR), device, pCreateInfo, pAllocator, pSwapchain)
}

// VkDestroySwapchainKHR mocks base method.
func (m *MockDriver) VkDestroySwapchainKHR(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VkDestroySwapchainKHR", device, swapchain, pAllocator)
}

// VkDestroySwapchainKHR indicates an expected call of VkDestroySwapchainKHR.
func (mr *MockDriverMockRecorder) VkDestroySwapchainKHR(device, swapchain, pAllocator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkDestroySwapchainKHR", reflect.TypeOf((*MockDriver)(nil).VkDestroySwapchainKHR), device, swapchain, pAllocator)
}

// VkGetSwapchainImagesKHR mocks base method.
func (m *MockDriver) VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain khr_swapchain.VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkGetSwapchainImagesKHR", device, swapchain, pSwapchainImageCount, pSwapchainImages)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkGetSwapchainImagesKHR indicates an expected call of VkGetSwapchainImagesKHR.
func (mr *MockDriverMockRecorder) VkGetSwapchainImagesKHR(device, swapchain, pSwapchainImageCount, pSwapchainImages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkGetSwapchainImagesKHR", reflect.TypeOf((*MockDriver)(nil).VkGetSwapchainImagesKHR), device, swapchain, pSwapchainImageCount, pSwapchainImages)
}

// VkQueuePresentKHR mocks base method.
func (m *MockDriver) VkQueuePresentKHR(queue driver.VkQueue, pPresentInfo *khr_swapchain.VkPresentInfoKHR) (common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VkQueuePresentKHR", queue, pPresentInfo)
	ret0, _ := ret[0].(common.VkResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VkQueuePresentKHR indicates an expected call of VkQueuePresentKHR.
func (mr *MockDriverMockRecorder) VkQueuePresentKHR(queue, pPresentInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VkQueuePresentKHR", reflect.TypeOf((*MockDriver)(nil).VkQueuePresentKHR), queue, pPresentInfo)
}

// coreDriver mocks base method.
func (m *MockDriver) coreDriver() driver.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "coreDriver")
	ret0, _ := ret[0].(driver.Driver)
	return ret0
}

// coreDriver indicates an expected call of coreDriver.
func (mr *MockDriverMockRecorder) coreDriver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "coreDriver", reflect.TypeOf((*MockDriver)(nil).coreDriver))
}

// MockLoader is a mock of Loader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader.
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// CreateSwapchain mocks base method.
func (m *MockLoader) CreateSwapchain(device core.Device, allocation *core.AllocationCallbacks, options *khr_swapchain.CreationOptions) (khr_swapchain.Swapchain, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSwapchain", device, allocation, options)
	ret0, _ := ret[0].(khr_swapchain.Swapchain)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSwapchain indicates an expected call of CreateSwapchain.
func (mr *MockLoaderMockRecorder) CreateSwapchain(device, allocation, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSwapchain", reflect.TypeOf((*MockLoader)(nil).CreateSwapchain), device, allocation, options)
}
