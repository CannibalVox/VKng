package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type InstanceCreateFlags int32

var instanceCreateFlagsMapping = common.NewFlagStringMapping[InstanceCreateFlags]()

func (f InstanceCreateFlags) Register(str string) {
	instanceCreateFlagsMapping.Register(f, str)
}

func (f InstanceCreateFlags) String() string {
	return instanceCreateFlagsMapping.FlagsToString(f)
}

////

type InstanceCreateInfo struct {
	ApplicationName    string
	ApplicationVersion common.Version
	EngineName         string
	EngineVersion      common.Version
	APIVersion         common.APIVersion

	Flags InstanceCreateFlags

	EnabledExtensionNames []string
	EnabledLayerNames     []string

	common.NextOptions
}

func (o InstanceCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkInstanceCreateInfo{})))
	}
	cApplication := allocator.CString(o.ApplicationName)
	cEngine := allocator.CString(o.EngineName)

	appInfo := (*C.VkApplicationInfo)(allocator.Malloc(int(unsafe.Sizeof(C.VkApplicationInfo{}))))

	appInfo.sType = C.VK_STRUCTURE_TYPE_APPLICATION_INFO
	appInfo.pNext = nil
	appInfo.pApplicationName = (*C.char)(cApplication)
	appInfo.pEngineName = (*C.char)(cEngine)
	appInfo.applicationVersion = C.uint32_t(o.ApplicationVersion)
	appInfo.engineVersion = C.uint32_t(o.EngineVersion)
	appInfo.apiVersion = C.uint32_t(o.APIVersion)

	createInfo := (*C.VkInstanceCreateInfo)(preallocatedPointer)

	// Alloc array of extension names
	numExtensions := len(o.EnabledExtensionNames)
	extNamePtr := allocator.Malloc(numExtensions * int(unsafe.Sizeof(uintptr(0))))
	extNames := ([]*C.char)(unsafe.Slice((**C.char)(extNamePtr), numExtensions))
	for i := 0; i < numExtensions; i++ {
		extNames[i] = (*C.char)(allocator.CString(o.EnabledExtensionNames[i]))
	}

	// Alloc array of layer names
	numLayers := len(o.EnabledLayerNames)
	layerNamePtr := allocator.Malloc(numLayers * int(unsafe.Sizeof(uintptr(0))))
	layerNames := ([]*C.char)(unsafe.Slice((**C.char)(layerNamePtr), numLayers))
	for i := 0; i < numLayers; i++ {
		layerNames[i] = (*C.char)(allocator.CString(o.EnabledLayerNames[i]))
	}

	createInfo.sType = C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO
	createInfo.flags = C.VkInstanceCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.pApplicationInfo = appInfo
	createInfo.enabledExtensionCount = C.uint32_t(numExtensions)
	createInfo.ppEnabledExtensionNames = (**C.char)(extNamePtr)
	createInfo.enabledLayerCount = C.uint32_t(numLayers)
	createInfo.ppEnabledLayerNames = (**C.char)(layerNamePtr)

	return preallocatedPointer, nil
}
