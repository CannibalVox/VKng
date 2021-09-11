package ext_debugutils

//go:generate mockgen -source messenger.go -destination ./mocks/mocks.go -package mock_debugutils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

VkResult cgoCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT fn, VkInstance instance, const VkDebugUtilsMessengerCreateInfoEXT *pCreateInfo,const VkAllocationCallbacks *pAllocator, VkDebugUtilsMessengerEXT* pDebugMessenger) {
	return fn(instance, pCreateInfo, pAllocator, pDebugMessenger);
}

void cgoDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerEXT debugMessenger, const VkAllocationCallbacks* pAllocator) {
	fn(instance, debugMessenger, pAllocator);
}
*/
import "C"

import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type MessengerHandle C.VkDebugUtilsMessengerEXT
type vulkanMessenger struct {
	instance C.VkInstance
	handle   C.VkDebugUtilsMessengerEXT

	destroyFunc C.PFN_vkDestroyDebugUtilsMessengerEXT
}

type Messenger interface {
	Handle() MessengerHandle
	Destroy()
}

func CreateMessenger(allocator cgoalloc.Allocator, instance resources.Instance, options *Options) (Messenger, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	instanceHandle := C.VkInstance(unsafe.Pointer(instance.Handle()))
	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	createFunc := (C.PFN_vkCreateDebugUtilsMessengerEXT)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkCreateDebugUtilsMessengerEXT"))))

	var messenger C.VkDebugUtilsMessengerEXT
	res := loader.VkResult(C.cgoCreateDebugUtilsMessengerEXT(createFunc, instanceHandle, (*C.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), nil, &messenger))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	destroyFunc := (C.PFN_vkDestroyDebugUtilsMessengerEXT)(instance.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkDestroyDebugUtilsMessengerEXT"))))

	return &vulkanMessenger{
		handle:   messenger,
		instance: instanceHandle,

		destroyFunc: destroyFunc,
	}, res, nil
}

func (m *vulkanMessenger) Destroy() {
	C.cgoDestroyDebugUtilsMessengerEXT(m.destroyFunc, m.instance, m.handle, nil)
}

func (m *vulkanMessenger) Handle() MessengerHandle {
	return MessengerHandle(m.handle)
}
