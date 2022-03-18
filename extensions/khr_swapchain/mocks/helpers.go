package mock_swapchain

import (
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeSwapchain() khr_swapchain.VkSwapchainKHR {
	return khr_swapchain.VkSwapchainKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSwapchain(ctrl *gomock.Controller) *MockSwapchain {
	swapchain := NewMockSwapchain(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeSwapchain()).AnyTimes()

	return swapchain
}
