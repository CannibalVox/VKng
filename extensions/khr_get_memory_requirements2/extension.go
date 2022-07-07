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
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_get_memory_requirements2_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_get_memory_requirements2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BufferMemoryRequirements2(device core1_0.Device, o BufferMemoryRequirementsInfo2, out *MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetBufferMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) ImageMemoryRequirements2(device core1_0.Device, o ImageMemoryRequirementsInfo2, out *MemoryRequirements2) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, out)
	if err != nil {
		return err
	}

	e.driver.VkGetImageMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR)(optionPtr),
		(*khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR)(outDataPtr),
	)

	return common.PopulateOutData(out, outDataPtr)
}

func (e *VulkanExtension) ImageSparseMemoryRequirements2(device core1_0.Device, o ImageSparseMemoryRequirementsInfo2, outDataFactory func() *SparseImageMemoryRequirements2) ([]*SparseImageMemoryRequirements2, error) {
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

	outDataSlice := make([]*SparseImageMemoryRequirements2, count)
	for i := 0; i < count; i++ {
		if outDataFactory != nil {
			outDataSlice[i] = outDataFactory()
		} else {
			outDataSlice[i] = &SparseImageMemoryRequirements2{}
		}
	}

	outDataPtr, err := common.AllocOutDataHeaderSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageMemoryRequirements2](arena, outDataSlice)
	if err != nil {
		return nil, err
	}

	castOutDataPtr := (*C.VkSparseImageMemoryRequirements2KHR)(outDataPtr)

	e.driver.VkGetImageSparseMemoryRequirements2KHR(device.Handle(),
		(*khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR)(optionPtr),
		requirementCountPtr,
		(*khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR)(unsafe.Pointer(castOutDataPtr)),
	)

	err = common.PopulateOutDataSlice[C.VkSparseImageMemoryRequirements2KHR, *SparseImageMemoryRequirements2](outDataSlice, unsafe.Pointer(outDataPtr))
	if err != nil {
		return nil, err
	}

	return outDataSlice, nil
}
