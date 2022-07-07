// Code generated by MockGen. DO NOT EDIT.
// Source: swapchain.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	reflect "reflect"
	time "time"

	common "github.com/CannibalVox/VKng/core/common"
	core1_0 "github.com/CannibalVox/VKng/core/core1_0"
	driver "github.com/CannibalVox/VKng/core/driver"
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	gomock "github.com/golang/mock/gomock"
)

// MockSwapchain is a mock of Swapchain interface.
type MockSwapchain struct {
	ctrl     *gomock.Controller
	recorder *MockSwapchainMockRecorder
}

// MockSwapchainMockRecorder is the mock recorder for MockSwapchain.
type MockSwapchainMockRecorder struct {
	mock *MockSwapchain
}

// NewMockSwapchain creates a new mock instance.
func NewMockSwapchain(ctrl *gomock.Controller) *MockSwapchain {
	mock := &MockSwapchain{ctrl: ctrl}
	mock.recorder = &MockSwapchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwapchain) EXPECT() *MockSwapchainMockRecorder {
	return m.recorder
}

// AcquireNextImage mocks base method.
func (m *MockSwapchain) AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireNextImage", timeout, semaphore, fence)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AcquireNextImage indicates an expected call of AcquireNextImage.
func (mr *MockSwapchainMockRecorder) AcquireNextImage(timeout, semaphore, fence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireNextImage", reflect.TypeOf((*MockSwapchain)(nil).AcquireNextImage), timeout, semaphore, fence)
}

// Destroy mocks base method.
func (m *MockSwapchain) Destroy(callbacks *driver.AllocationCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy", callbacks)
}

// Destroy indicates an expected call of Destroy.
func (mr *MockSwapchainMockRecorder) Destroy(callbacks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockSwapchain)(nil).Destroy), callbacks)
}

// Handle mocks base method.
func (m *MockSwapchain) Handle() khr_swapchain_driver.VkSwapchainKHR {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(khr_swapchain_driver.VkSwapchainKHR)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockSwapchainMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockSwapchain)(nil).Handle))
}

// Images mocks base method.
func (m *MockSwapchain) SwapchainImages() ([]core1_0.Image, common.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwapchainImages")
	ret0, _ := ret[0].([]core1_0.Image)
	ret1, _ := ret[1].(common.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Images indicates an expected call of Images.
func (mr *MockSwapchainMockRecorder) Images() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwapchainImages", reflect.TypeOf((*MockSwapchain)(nil).SwapchainImages))
}
