package mock_swapchain

import (
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeSwapchain() khr_swapchain_driver.VkSwapchainKHR {
	return khr_swapchain_driver.VkSwapchainKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSwapchain(ctrl *gomock.Controller) *MockSwapchain {
	swapchain := NewMockSwapchain(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeSwapchain()).AnyTimes()

	return swapchain
}
