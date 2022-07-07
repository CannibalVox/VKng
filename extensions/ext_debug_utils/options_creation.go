package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "debug_callback.h"

VKAPI_ATTR VkBool32 VKAPI_CALL debugCallback(
	VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity,
	VkDebugUtilsMessageTypeFlagsEXT messageType,
	const VkDebugUtilsMessengerCallbackDataEXT *pCallbackData,
	void *pUserData) {

	return goDebugCallback(messageSeverity, messageType, (VkDebugUtilsMessengerCallbackDataEXT*)pCallbackData, pUserData);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"runtime/cgo"
	"unsafe"
)

type DebugUtilsMessengerCreateInfo struct {
	Flags           CallbackDataFlags
	MessageSeverity MessageSeverities
	MessageType     MessageTypes
	UserCallback    CallbackFunction

	common.NextOptions
}

func (o DebugUtilsMessengerCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkDebugUtilsMessengerCreateInfoEXT{})))
	}
	createInfo := (*C.VkDebugUtilsMessengerCreateInfoEXT)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	createInfo.flags = C.VkDebugUtilsMessengerCreateFlagsEXT(o.Flags)
	createInfo.pNext = next

	createInfo.messageSeverity = C.VkDebugUtilsMessageSeverityFlagsEXT(o.MessageSeverity)
	createInfo.messageType = C.VkDebugUtilsMessageTypeFlagsEXT(o.MessageType)
	createInfo.pfnUserCallback = (C.PFN_vkDebugUtilsMessengerCallbackEXT)(C.debugCallback)
	createInfo.pUserData = unsafe.Pointer(cgo.NewHandle(o.UserCallback))

	return preallocatedPointer, nil
}
