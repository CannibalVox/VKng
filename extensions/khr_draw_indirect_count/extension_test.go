package khr_draw_indirect_count

import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/core/mocks"
	mock_draw_indirect_count "github.com/CannibalVox/VKng/extensions/khr_draw_indirect_count/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestVulkanExtension_CmdDrawIndexedIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_draw_indirect_count.NewMockDriver(ctrl)
	extension := CreateExtensionFromDriver(extDriver)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)
	buffer := mocks.EasyMockBuffer(ctrl)
	countBuffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkCmdDrawIndexedIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(1),
		countBuffer.Handle(),
		driver.VkDeviceSize(3),
		driver.Uint32(5),
		driver.Uint32(7),
	)

	extension.CmdDrawIndexedIndirectCount(
		commandBuffer,
		buffer,
		uint64(1),
		countBuffer,
		uint64(3),
		5,
		7,
	)
}

func TestVulkanExtension_CmdDrawIndirectCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_draw_indirect_count.NewMockDriver(ctrl)
	extension := CreateExtensionFromDriver(extDriver)

	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)
	buffer := mocks.EasyMockBuffer(ctrl)
	countBuffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkCmdDrawIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(11),
		countBuffer.Handle(),
		driver.VkDeviceSize(13),
		driver.Uint32(17),
		driver.Uint32(19),
	)

	extension.CmdDrawIndirectCount(
		commandBuffer,
		buffer,
		uint64(11),
		countBuffer,
		uint64(13),
		17,
		19,
	)
}
