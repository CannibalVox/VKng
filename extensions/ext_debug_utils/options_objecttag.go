package ext_debug_utils

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

type ObjectTagOptions struct {
	Type   common.ObjectType
	Handle uintptr

	TagName uint64
	Tag     []byte

	common.HaveNext
}

func (t ObjectTagOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsObjectTagInfoEXT)
	}

	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(preallocatedPointer)
	tagInfo.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
	tagInfo.pNext = next
	tagInfo.objectType = C.VkObjectType(t.Type)
	tagInfo.objectHandle = C.uint64_t(t.Handle)
	tagInfo.tagName = C.uint64_t(t.TagName)
	tagInfo.tagSize = C.size_t(len(t.Tag))
	tagInfo.pTag = allocator.CBytes(t.Tag)

	return preallocatedPointer, nil
}

func (t ObjectTagOptions) PopulateOutData(cPointer unsafe.Pointer) (unsafe.Pointer, error) {
	tagInfo := (*C.VkDebugUtilsObjectTagInfoEXT)(cPointer)

	t.Type = common.ObjectType(tagInfo.objectType)
	t.Handle = uintptr(tagInfo.objectHandle)
	t.TagName = uint64(tagInfo.tagName)
	t.Tag = C.GoBytes(tagInfo.pTag, C.int(tagInfo.tagSize))

	return tagInfo.pNext, nil
}
