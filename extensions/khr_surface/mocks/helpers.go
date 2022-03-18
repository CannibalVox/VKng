package mock_surface

import (
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeSurface() khr_surface.VkSurfaceKHR {
	return khr_surface.VkSurfaceKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSurface(ctrl *gomock.Controller) *MockSurface {
	surface := NewMockSurface(ctrl)
	surface.EXPECT().Handle().Return(NewFakeSurface()).AnyTimes()

	return surface
}
