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

type DebugUtilsObjectNameInfo struct {
	ObjectName   string
	ObjectHandle uintptr
	ObjectType   core1_0.ObjectType

	common.NextOptions
}

func (i DebugUtilsObjectNameInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectNameInfoEXT)
	}

	nameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(preallocatedPointer)
	nameInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
	nameInfo.pNext = next
	nameInfo.objectType = C.VkObjectType(i.ObjectType)
	nameInfo.objectHandle = C.uint64_t(i.ObjectHandle)
	nameInfo.pObjectName = (*C.char)(allocator.CString(i.ObjectName))

	return preallocatedPointer, nil
}

func (i *DebugUtilsObjectNameInfo) PopulateFromCPointer(cDataPointer unsafe.Pointer) {
	objectNameInfo := (*C.VkDebugUtilsObjectNameInfoEXT)(cDataPointer)
	i.ObjectType = core1_0.ObjectType(objectNameInfo.objectType)
	i.ObjectHandle = uintptr(objectNameInfo.objectHandle)
	i.ObjectName = ""

	if objectNameInfo.pObjectName != nil {
		i.ObjectName = C.GoString(objectNameInfo.pObjectName)
	}
}
