package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type TessellationDomainOriginOptions struct {
	DomainOrigin TessellationDomainOrigin
	common.HaveNext
}

func (o TessellationDomainOriginOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPipelineTessellationDomainOriginStateCreateInfoKHR{})))
	}

	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.domainOrigin = (C.VkTessellationDomainOriginKHR)(o.DomainOrigin)

	return preallocatedPointer, nil
}

func (o TessellationDomainOriginOptions) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPipelineTessellationDomainOriginStateCreateInfoKHR)(cPointer)
	return createInfo.pNext, nil
}

var _ common.Options = TessellationDomainOriginOptions{}
