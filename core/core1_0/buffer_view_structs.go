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

type BufferViewCreateOptions struct {
	Buffer Buffer
	Format DataFormat
	Offset int
	Range  int

	common.NextOptions
}

func (o BufferViewCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferViewCreateInfo)
	}
	createInfo := (*C.VkBufferViewCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_VIEW_CREATE_INFO
	createInfo.pNext = next
	createInfo.flags = 0
	createInfo.buffer = C.VkBuffer(unsafe.Pointer(o.Buffer.Handle()))
	createInfo.format = C.VkFormat(o.Format)
	createInfo.offset = C.VkDeviceSize(o.Offset)
	createInfo._range = C.VkDeviceSize(o.Range)

	return unsafe.Pointer(createInfo), nil
}
