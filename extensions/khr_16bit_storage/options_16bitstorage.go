package khr_16bit_storage

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

type Device16BitStorageOptions struct {
	StorageBuffer16BitAccess           bool
	UniformAndStorageBuffer16BitAccess bool
	StoragePushConstant16              bool
	StorageInputOutput16               bool

	common.HaveNext
}

func (o Device16BitStorageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice16BitStorageFeaturesKHR{})))
	}

	data := (*C.VkPhysicalDevice16BitStorageFeaturesKHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES_KHR
	data.pNext = next

	data.storageBuffer16BitAccess = C.VkBool32(0)
	data.uniformAndStorageBuffer16BitAccess = C.VkBool32(0)
	data.storagePushConstant16 = C.VkBool32(0)
	data.storageInputOutput16 = C.VkBool32(0)

	if o.StorageBuffer16BitAccess {
		data.storageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer16BitAccess {
		data.uniformAndStorageBuffer16BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant16 {
		data.storagePushConstant16 = C.VkBool32(1)
	}

	if o.StorageInputOutput16 {
		data.storageInputOutput16 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o Device16BitStorageOptions) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDevice16BitStorageFeaturesKHR)(cDataPointer)

	return data.pNext, nil
}
