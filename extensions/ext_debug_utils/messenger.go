package ext_debug_utils

//go:generate mockgen -source messenger.go -destination ./mocks/mocks.go -package mock_debugutils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"runtime/cgo"
	"unsafe"
)

type Messenger struct {
	instance driver.VkInstance
	handle   VkDebugUtilsMessengerEXT
	driver   Driver
}

func (m *Messenger) Destroy(callbacks *driver.AllocationCallbacks) {
	m.driver.VkDestroyDebugUtilsMessengerEXT(m.instance, m.handle, callbacks.Handle())
}

func (m *Messenger) Handle() VkDebugUtilsMessengerEXT {
	return m.handle
}

type CallbackFunction func(msgType MessageTypes, severity MessageSeverities, data *CallbackData) bool

//export goDebugCallback
func goDebugCallback(messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT, messageType C.VkDebugUtilsMessageTypeFlagsEXT, data *C.VkDebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) C.VkBool32 {
	severity := MessageSeverities(messageSeverity)
	msgType := MessageTypes(messageType)
	callbackData := createCallbackData(data)

	f := cgo.Handle(userData).Value().(CallbackFunction)
	if f(msgType, severity, callbackData) {
		return C.VK_TRUE
	}

	return C.VK_FALSE
}
