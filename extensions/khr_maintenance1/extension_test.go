package khr_maintenance1_test

import (
	"github.com/CannibalVox/VKng/core/common"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_maintenance1"
	mock_maintenance1 "github.com/CannibalVox/VKng/extensions/khr_maintenance1/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestVulkanExtension_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := mocks.EasyMockCommandPool(ctrl, device)

	maintDriver := mock_maintenance1.NewMockDriver(ctrl)
	extension := khr_maintenance1.CreateExtensionFromDriver(maintDriver)

	maintDriver.EXPECT().VkTrimCommandPoolKHR(device.Handle(), commandPool.Handle(), khr_maintenance1.VkCommandPoolTrimFlagsKHR(0))

	extension.TrimCommandPool(commandPool, 0)
}
