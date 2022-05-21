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

type PhysicalDevice8BitStorageFeaturesOptions struct {
	StorageBuffer8BitAccess           bool
	UniformAndStorageBuffer8BitAccess bool
	StoragePushConstant8              bool

	common.HaveNext
}

func (o PhysicalDevice8BitStorageFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevice8BitStorageFeatures{})))
	}

	info := (*C.VkPhysicalDevice8BitStorageFeatures)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES
	info.pNext = next
	info.storageBuffer8BitAccess = C.VkBool32(0)
	info.uniformAndStorageBuffer8BitAccess = C.VkBool32(0)
	info.storagePushConstant8 = C.VkBool32(0)

	if o.StorageBuffer8BitAccess {
		info.storageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.UniformAndStorageBuffer8BitAccess {
		info.uniformAndStorageBuffer8BitAccess = C.VkBool32(1)
	}

	if o.StoragePushConstant8 {
		info.storagePushConstant8 = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDevice8BitStorageFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDevice8BitStorageFeatures)(cDataPointer)
	return info.pNext, nil
}
