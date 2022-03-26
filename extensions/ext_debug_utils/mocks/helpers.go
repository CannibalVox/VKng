package mock_debugutils

import (
	ext_driver "github.com/CannibalVox/VKng/extensions/ext_debug_utils/driver"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeMessenger() ext_driver.VkDebugUtilsMessengerEXT {
	return ext_driver.VkDebugUtilsMessengerEXT(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	messenger := NewMockMessenger(ctrl)
	messenger.EXPECT().Handle().Return(NewFakeMessenger()).AnyTimes()

	return messenger
}
