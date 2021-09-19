package ext_debug_utils

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
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
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

func CreateMessenger(instance core.Instance, options *Options) (Messenger, core.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	instanceHandle := C.VkInstance(unsafe.Pointer(instance.Handle()))
	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	createFunc := (C.PFN_vkCreateDebugUtilsMessengerEXT)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkCreateDebugUtilsMessengerEXT"))))

	var messenger C.VkDebugUtilsMessengerEXT
	res := core.VkResult(C.cgoCreateDebugUtilsMessengerEXT(createFunc, instanceHandle, (*C.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), nil, &messenger))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	destroyFunc := (C.PFN_vkDestroyDebugUtilsMessengerEXT)(instance.Driver().LoadProcAddr((*core.Char)(arena.CString("vkDestroyDebugUtilsMessengerEXT"))))

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
