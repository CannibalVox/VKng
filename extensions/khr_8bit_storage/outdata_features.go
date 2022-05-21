package khr_8bit_storage

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

type PhysicalDevice8BitStorageFeaturesOutData struct {
	StorageBuffer8BitAccess           bool
	UniformAndStorageBuffer8BitAccess bool
	StoragePushConstant8              bool

	common.HaveNext
}

func (o *PhysicalDevice8BitStorageFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeaturesKHR{})))
	}

	outData := (*C.VkPhysicalDevice8BitStorageFeaturesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevice8BitStorageFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevice8BitStorageFeaturesKHR)(cDataPointer)
	o.StoragePushConstant8 = outData.storagePushConstant8 != C.VkBool32(0)
	o.UniformAndStorageBuffer8BitAccess = outData.uniformAndStorageBuffer8BitAccess != C.VkBool32(0)
	o.StorageBuffer8BitAccess = outData.storageBuffer8BitAccess != C.VkBool32(0)

	return outData.pNext, nil
}
