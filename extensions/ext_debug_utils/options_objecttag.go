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

type DebugUtilsObjectTagInfo struct {
	ObjectType   core1_0.ObjectType
	ObjectHandle uintptr

	TagName uint64
	Tag     []byte

	common.NextOptions
}

func (t DebugUtilsObjectTagInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectTagInfoEXT)
	}

	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(preallocatedPointer)
	tagInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
	tagInfo.pNext = next
	tagInfo.objectType = C.VkObjectType(t.ObjectType)
	tagInfo.objectHandle = C.uint64_t(t.ObjectHandle)
	tagInfo.tagName = C.uint64_t(t.TagName)
	tagInfo.tagSize = C.size_t(len(t.Tag))
	tagInfo.pTag = allocator.CBytes(t.Tag)

	return preallocatedPointer, nil
}

func (t *DebugUtilsObjectTagInfo) PopulateFromCPointer(cPointer unsafe.Pointer) {
	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(cPointer)

	t.ObjectType = core1_0.ObjectType(tagInfo.objectType)
	t.ObjectHandle = uintptr(tagInfo.objectHandle)
	t.TagName = uint64(tagInfo.tagName)
	t.Tag = C.GoBytes(tagInfo.pTag, C.int(tagInfo.tagSize))
}
