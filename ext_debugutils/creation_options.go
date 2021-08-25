package ext_debugutils

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
	"github.com/CannibalVox/VKng/creation"
	"github.com/CannibalVox/cgoalloc"
	"runtime/cgo"
	"unsafe"
)

type Options struct {
	CaptureSeverities MessageSeverity
	CaptureTypes MessageType
	Callback CallbackFunction

	Next creation.Options
}

func (o *Options) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkDebugUtilsMessengerCreateInfoEXT)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkDebugUtilsMessengerCreateInfoEXT{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
	createInfo.flags = 0
	createInfo.messageSeverity = C.VkDebugUtilsMessageSeverityFlagsEXT(o.CaptureSeverities)
	createInfo.messageType = C.VkDebugUtilsMessageTypeFlagsEXT(o.CaptureTypes)
	createInfo.pfnUserCallback = (C.PFN_vkDebugUtilsMessengerCallbackEXT)(C.debugCallback)
	createInfo.pUserData = unsafe.Pointer(cgo.NewHandle(o.Callback))

	var next unsafe.Pointer
	var err error

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}

	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}