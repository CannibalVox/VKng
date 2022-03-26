package mock_surface

import (
	ext_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeSurface() ext_driver.VkSurfaceKHR {
	return ext_driver.VkSurfaceKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSurface(ctrl *gomock.Controller) *MockSurface {
	surface := NewMockSurface(ctrl)
	surface.EXPECT().Handle().Return(NewFakeSurface()).AnyTimes()

	return surface
}
