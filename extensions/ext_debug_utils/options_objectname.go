package ext_debug_utils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type ObjectNameOptions struct {
	Name   string
	Handle uintptr
	Type   core1_0.ObjectType

	common.NextOptions
}

func (i ObjectNameOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectNameInfoEXT)
	}

	nameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(preallocatedPointer)
	nameInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
	nameInfo.pNext = next
	nameInfo.objectType = C.VkObjectType(i.Type)
	nameInfo.objectHandle = C.uint64_t(i.Handle)
	nameInfo.pObjectName = (*C.char)(allocator.CString(i.Name))

	return preallocatedPointer, nil
}

func (i *ObjectNameOptions) PopulateFromCPointer(cDataPointer unsafe.Pointer) {
	objectNameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(cDataPointer)
	i.Type = core1_0.ObjectType(objectNameInfo.objectType)
	i.Handle = uintptr(objectNameInfo.objectHandle)
	i.Name = ""

	if objectNameInfo.pObjectName != nil {
		i.Name = C.GoString(objectNameInfo.pObjectName)
	}
}
