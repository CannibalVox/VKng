package mock_debugutils

import (
	"github.com/CannibalVox/VKng/extensions/ext_debug_utils"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeMessenger() ext_debug_utils.VkDebugUtilsMessengerEXT {
	return ext_debug_utils.VkDebugUtilsMessengerEXT(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	messenger := NewMockMessenger(ctrl)
	messenger.EXPECT().Handle().Return(NewFakeMessenger()).AnyTimes()

	return messenger
}
