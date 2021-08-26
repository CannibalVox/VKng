package ext_debugutils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"

VkResult vkCreateDebugUtilsMessengerEXT(VkInstance instance, const VkDebugUtilsMessengerCreateInfoEXT *pCreateInfo,const VkAllocationCallbacks *pAllocator, VkDebugUtilsMessengerEXT* pDebugMessenger) {
	PFN_vkCreateDebugUtilsMessengerEXT func = (PFN_vkCreateDebugUtilsMessengerEXT) vkGetInstanceProcAddr(instance, "vkCreateDebugUtilsMessengerEXT");
    if (func != NULL) {
        return func(instance, pCreateInfo, pAllocator, pDebugMessenger);
    } else {
        return VK_ERROR_EXTENSION_NOT_PRESENT;
    }
}

void vkDestroyDebugUtilsMessengerEXT(VkInstance instance, VkDebugUtilsMessengerEXT debugMessenger, const VkAllocationCallbacks* pAllocator) {
    PFN_vkDestroyDebugUtilsMessengerEXT func = (PFN_vkDestroyDebugUtilsMessengerEXT) vkGetInstanceProcAddr(instance, "vkDestroyDebugUtilsMessengerEXT");
    if (func != NULL) {
        func(instance, debugMessenger, pAllocator);
    }
}
*/
import "C"

import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/objects"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type MessengerHandle C.VkDebugUtilsMessengerEXT
type Messenger struct {
	instance C.VkInstance
	handle C.VkDebugUtilsMessengerEXT
}

func CreateMessenger(allocator cgoalloc.Allocator, instance *objects.Instance, options *Options) (*Messenger, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	instanceHandle := C.VkInstance(unsafe.Pointer(instance.Handle()))
	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var messenger C.VkDebugUtilsMessengerEXT
	res := C.vkCreateDebugUtilsMessengerEXT(instanceHandle, (*C.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), nil, &messenger)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Messenger{handle: messenger, instance: instanceHandle}, nil
}

func (m *Messenger) Destroy() {
	C.vkDestroyDebugUtilsMessengerEXT(m.instance, m.handle, nil)
}

func (m *Messenger) Handle() MessengerHandle {
	return MessengerHandle(m.handle)
}
