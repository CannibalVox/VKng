// Code generated by MockGen. DO NOT EDIT.
// Source: extension.go

// Package mock_swapchain is a generated GoMock package.
package mock_swapchain

import (
	"github.com/CannibalVox/VKng/core"
	reflect "reflect"
	time "time"

	ext_swapchain "github.com/CannibalVox/VKng/extensions/swapchain"
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
func (m *MockSwapchain) AcquireNextImage(timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, core.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireNextImage", timeout, semaphore, fence)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(core.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AcquireNextImage indicates an expected call of AcquireNextImage.
func (mr *MockSwapchainMockRecorder) AcquireNextImage(timeout, semaphore, fence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireNextImage", reflect.TypeOf((*MockSwapchain)(nil).AcquireNextImage), timeout, semaphore, fence)
}

// Destroy mocks base method.
func (m *MockSwapchain) Destroy() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy.
func (mr *MockSwapchainMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockSwapchain)(nil).Destroy))
}

// Handle mocks base method.
func (m *MockSwapchain) Handle() ext_swapchain.SwapchainHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle")
	ret0, _ := ret[0].(ext_swapchain.SwapchainHandle)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockSwapchainMockRecorder) Handle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockSwapchain)(nil).Handle))
}

// Images mocks base method.
func (m *MockSwapchain) Images() ([]core.Image, core.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Images")
	ret0, _ := ret[0].([]core.Image)
	ret1, _ := ret[1].(core.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Images indicates an expected call of Images.
func (mr *MockSwapchainMockRecorder) Images() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Images", reflect.TypeOf((*MockSwapchain)(nil).Images))
}

// PresentToQueue mocks base method.
func (m *MockSwapchain) PresentToQueue(queue core.Queue, o *ext_swapchain.PresentOptions) ([]core.VkResult, core.VkResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresentToQueue", queue, o)
	ret0, _ := ret[0].([]core.VkResult)
	ret1, _ := ret[1].(core.VkResult)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PresentToQueue indicates an expected call of PresentToQueue.
func (mr *MockSwapchainMockRecorder) PresentToQueue(queue, o interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentToQueue", reflect.TypeOf((*MockSwapchain)(nil).PresentToQueue), queue, o)
}
