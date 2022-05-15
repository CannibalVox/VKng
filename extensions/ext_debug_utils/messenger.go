package ext_debug_utils

//go:generate mockgen -source messenger.go -destination ./mocks/messenger.go -package mock_debugutils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

*/
import "C"
import (
	"fmt"
	"github.com/CannibalVox/VKng/core/driver"
	ext_driver "github.com/CannibalVox/VKng/extensions/ext_debug_utils/driver"
	"runtime/cgo"
	"unsafe"
)

type Messenger interface {
	Destroy(callbacks *driver.AllocationCallbacks)
	Handle() ext_driver.VkDebugUtilsMessengerEXT
}

type vulkanMessenger struct {
	instance   driver.VkInstance
	handle     ext_driver.VkDebugUtilsMessengerEXT
	coreDriver driver.Driver
	driver     ext_driver.Driver
}

func (m *vulkanMessenger) Destroy(callbacks *driver.AllocationCallbacks) {
	m.driver.VkDestroyDebugUtilsMessengerEXT(m.instance, m.handle, callbacks.Handle())
	m.coreDriver.ObjectStore().Delete(driver.VulkanHandle(m.handle))
}

func (m *vulkanMessenger) Handle() ext_driver.VkDebugUtilsMessengerEXT {
	return m.handle
}

type CallbackFunction func(msgType MessageTypes, severity MessageSeverities, data *CallbackDataOptions) bool

//export goDebugCallback
func goDebugCallback(messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT, messageType C.VkDebugUtilsMessageTypeFlagsEXT, data *C.VkDebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) C.VkBool32 {
	severity := MessageSeverities(messageSeverity)
	msgType := MessageTypes(messageType)

	callbackData := &CallbackDataOptions{}

	err := callbackData.PopulateFromCPointer(unsafe.Pointer(data))
	if err != nil {
		callbackData = &CallbackDataOptions{
			MessageIDName: "vkng-internal",
			Message:       fmt.Sprintf("error loading debug callback data from C: %v+", err),
		}
	}

	f := cgo.Handle(userData).Value().(CallbackFunction)
	if f(msgType, severity, callbackData) {
		return C.VK_TRUE
	}

	return C.VK_FALSE
}
