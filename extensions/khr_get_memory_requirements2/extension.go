package khr_get_memory_requirements2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_get_memory_requirements2_driver "github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver khr_get_memory_requirements2_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: khr_get_memory_requirements2_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_get_memory_requirements2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BufferMemoryRequirements(device core1_0.Device, o BufferRequirementsOptions, out *MemoryRequirementsOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetBufferMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) ImageMemoryRequirements(device core1_0.Device, o ImageMemoryRequirementsOptions, out *MemoryRequirementsOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetImageMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) SparseImageMemoryRequirements(device core1_0.Device, o SparseImageRequirementsOptions, outDataFactory func() *SparseImageRequirementsOutData) ([]*SparseImageRequirementsOutData, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, err
	}

	requirementCountPtr := (*driver.Uint32)(arena.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))

	e.driver.VkGetImageSparseMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		nil,
	)

	count := int(*requirementCountPtr)
	if count == 0 {
		return nil, nil
	}

	outDataSlice := make([]*SparseImageRequirementsOutData, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &SparseImageRequirementsOutData{}
		}
	}

	outDataPtr, err := common.AllocOptionSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageRequirementsOutData](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2KHR)(outDataPtr)

	e.driver.VkGetImageSparseMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		(*khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageRequirementsOutData](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}
