package ext_debugutils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

const ExtensionName string = C.VK_EXT_DEBUG_UTILS_EXTENSION_NAME

type CallbackFunction func(msgType MessageType, severity MessageSeverity, data *CallbackData) bool

//export goDebugCallback
func goDebugCallback(messageSeverity C.VkDebugUtilsMessageSeverityFlagBitsEXT, messageType C.VkDebugUtilsMessageTypeFlagsEXT, data *C.VkDebugUtilsMessengerCallbackDataEXT, userData unsafe.Pointer) C.VkBool32 {
	severity := MessageSeverity(messageSeverity)
	msgType := MessageType(messageType)
	callbackData := CreateCallbackData(data)

	f := cgo.Handle(userData).Value().(CallbackFunction)
	if f(msgType, severity, callbackData) {
		return C.VK_TRUE
	}

	return C.VK_FALSE
}
