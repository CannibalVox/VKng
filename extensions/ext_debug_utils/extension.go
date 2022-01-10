package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

VkResult cgoCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerCreateInfoEXT *pCreateInfo, VkAllocationCallbacks *pAllocator, VkDebugUtilsMessengerEXT* pDebugMessenger) {
	return fn(instance, pCreateInfo, pAllocator, pDebugMessenger);
}

void cgoDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerEXT debugMessenger, VkAllocationCallbacks* pAllocator) {
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

const ExtensionName string = C.VK_EXT_DEBUG_UTILS_EXTENSION_NAME

type extDebugUtilsDriver struct {
	createDebugUtils  C.PFN_vkCreateDebugUtilsMessengerEXT
	destroyDebugUtils C.PFN_vkDestroyDebugUtilsMessengerEXT
}

type VkDebugUtilsMessengerCreateInfoEXT C.VkDebugUtilsMessengerCreateInfoEXT
type VkDebugUtilsMessengerEXT C.VkDebugUtilsMessengerEXT
type Driver interface {
	VkCreateDebugUtilsMessengerEXT(instance core.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *core.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (core.VkResult, error)
	VkDestroyDebugUtilsMessengerEXT(instance core.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *core.VkAllocationCallbacks)
}

func CreateDriverFromCore(driver core.Driver) Driver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &extDebugUtilsDriver{
		createDebugUtils:  (C.PFN_vkCreateDebugUtilsMessengerEXT)(driver.LoadProcAddr((*core.Char)(arena.CString("vkCreateDebugUtilsMessengerEXT")))),
		destroyDebugUtils: (C.PFN_vkDestroyDebugUtilsMessengerEXT)(driver.LoadProcAddr((*core.Char)(arena.CString("vkDestroyDebugUtilsMessengerEXT")))),
	}
}

func (d *extDebugUtilsDriver) VkCreateDebugUtilsMessengerEXT(instance core.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *core.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (core.VkResult, error) {
	if d.createDebugUtils == nil {
		panic("attempt to call extension method vkCreateDebugUtilsMessengerEXT when extension not present")
	}

	res := core.VkResult(C.cgoCreateDebugUtilsMessengerEXT(d.createDebugUtils,
		C.VkInstance(unsafe.Pointer(instance)),
		(*C.VkDebugUtilsMessengerCreateInfoEXT)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDebugUtilsMessengerEXT)(pDebugMessenger)))

	return res, res.ToError()
}

func (d *extDebugUtilsDriver) VkDestroyDebugUtilsMessengerEXT(instance core.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *core.VkAllocationCallbacks) {
	if d.destroyDebugUtils == nil {
		panic("attempt to call extension method vkDestroyDebugUtilsMessengerEXT when extension not present")
	}

	C.cgoDestroyDebugUtilsMessengerEXT(d.destroyDebugUtils,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessengerEXT(debugMessenger),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

type extDebugUtilsLoader struct {
	driver Driver
}

type Loader interface {
	CreateMessenger(instance core.Instance, allocation *core.AllocationCallbacks, o *CreationOptions) (Messenger, core.VkResult, error)
}

func CreateLoaderFromInstance(instance core.Instance) Loader {
	driver := CreateDriverFromCore(instance.Driver())

	return &extDebugUtilsLoader{
		driver: driver,
	}
}

func CreateLoaderFromDriver(driver Driver) Loader {
	return &extDebugUtilsLoader{
		driver: driver,
	}
}

func (l *extDebugUtilsLoader) CreateMessenger(instance core.Instance, allocation *core.AllocationCallbacks, o *CreationOptions) (Messenger, core.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var messenger VkDebugUtilsMessengerEXT
	res, err := l.driver.VkCreateDebugUtilsMessengerEXT(instance.Handle(), (*VkDebugUtilsMessengerCreateInfoEXT)(createInfo), nil, &messenger)

	if err != nil {
		return nil, res, err
	}

	return &vulkanMessenger{
		handle:   messenger,
		instance: instance.Handle(),
		driver:   l.driver,
	}, res, nil
}
