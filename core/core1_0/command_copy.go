package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

func (c *VulkanCommandBuffer) CmdCopyBuffer(srcBuffer Buffer, dstBuffer Buffer, copyRegions []BufferCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionPtr, err := common.AllocSlice[C.VkBufferCopy, BufferCopy](allocator, copyRegions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdCopyBuffer(c.commandBufferHandle, srcBuffer.Handle(), dstBuffer.Handle(), driver.Uint32(len(copyRegions)), (*driver.VkBufferCopy)(unsafe.Pointer(copyRegionPtr)))
	return nil
}

func (c *VulkanCommandBuffer) CmdCopyImage(srcImage Image, srcImageLayout common.ImageLayout, dstImage Image, dstImageLayout common.ImageLayout, regions []ImageCopy) error {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	copyRegionCount := len(regions)
	copyRegionUnsafe, err := common.AllocSlice[C.VkImageCopy, ImageCopy](allocator, regions)
	if err != nil {
		return err
	}

	c.deviceDriver.VkCmdCopyImage(c.commandBufferHandle, srcImage.Handle(), driver.VkImageLayout(srcImageLayout), dstImage.Handle(), driver.VkImageLayout(dstImageLayout), driver.Uint32(copyRegionCount), (*driver.VkImageCopy)(unsafe.Pointer(copyRegionUnsafe)))
	return nil
}
