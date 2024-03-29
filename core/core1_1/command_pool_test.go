package core1_1_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestVulkanCommandPool_TrimCommandPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_1)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandPool := core1_1.PromoteCommandPool(dummies.EasyDummyCommandPool(coreDriver, device))

	coreDriver.EXPECT().VkTrimCommandPool(device.Handle(), commandPool.Handle(), driver.VkCommandPoolTrimFlags(0))

	commandPool.TrimCommandPool(0)
}
