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
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const ExtensionName string = C.VK_EXT_DEBUG_UTILS_EXTENSION_NAME

type CDriver struct {
	createDebugUtils  C.PFN_vkCreateDebugUtilsMessengerEXT
	destroyDebugUtils C.PFN_vkDestroyDebugUtilsMessengerEXT
}

type VkDebugUtilsMessengerCreateInfoEXT C.VkDebugUtilsMessengerCreateInfoEXT
type VkDebugUtilsMessengerEXT C.VkDebugUtilsMessengerEXT
type Driver interface {
	VkCreateDebugUtilsMessengerEXT(instance driver.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error)
	VkDestroyDebugUtilsMessengerEXT(instance driver.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *driver.VkAllocationCallbacks)
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		createDebugUtils:  (C.PFN_vkCreateDebugUtilsMessengerEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateDebugUtilsMessengerEXT")))),
		destroyDebugUtils: (C.PFN_vkDestroyDebugUtilsMessengerEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroyDebugUtilsMessengerEXT")))),
	}
}

func (d *CDriver) VkCreateDebugUtilsMessengerEXT(instance driver.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error) {
	if d.createDebugUtils == nil {
		panic("attempt to call extension method vkCreateDebugUtilsMessengerEXT when extension not present")
	}

	res := common.VkResult(C.cgoCreateDebugUtilsMessengerEXT(d.createDebugUtils,
		C.VkInstance(unsafe.Pointer(instance)),
		(*C.VkDebugUtilsMessengerCreateInfoEXT)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDebugUtilsMessengerEXT)(pDebugMessenger)))

	return res, res.ToError()
}

func (d *CDriver) VkDestroyDebugUtilsMessengerEXT(instance driver.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyDebugUtils == nil {
		panic("attempt to call extension method vkDestroyDebugUtilsMessengerEXT when extension not present")
	}

	C.cgoDestroyDebugUtilsMessengerEXT(d.destroyDebugUtils,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessengerEXT(debugMessenger),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

type VulkanExtension struct {
	driver Driver
}

type Extension interface {
	CreateMessenger(instance iface.Instance, allocation *driver.AllocationCallbacks, o *CreationOptions) (*Messenger, common.VkResult, error)
}

func CreateExtensionFromInstance(instance iface.Instance) *VulkanExtension {
	driver := CreateDriverFromCore(instance.Driver())

	return &VulkanExtension{
		driver: driver,
	}
}

func CreateExtensionFromDriver(driver Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateMessenger(instance iface.Instance, allocation *driver.AllocationCallbacks, o *CreationOptions) (*Messenger, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, common.VKErrorUnknown, err
	}

	var messenger VkDebugUtilsMessengerEXT
	res, err := l.driver.VkCreateDebugUtilsMessengerEXT(instance.Handle(), (*VkDebugUtilsMessengerCreateInfoEXT)(createInfo), allocation.Handle(), &messenger)

	if err != nil {
		return nil, res, err
	}

	return &Messenger{
		handle:   messenger,
		instance: instance.Handle(),
		driver:   l.driver,
	}, res, nil
}
