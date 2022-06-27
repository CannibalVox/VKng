package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanCommandPool struct {
	core1_1.CommandPool
}

func PromoteCommandPool(commandPool core1_0.CommandPool) CommandPool {
	if !commandPool.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedCommandPool := core1_1.PromoteCommandPool(commandPool)

	return commandPool.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(commandPool.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanCommandPool{
				CommandPool: promotedCommandPool,
			}
		}).(CommandPool)
}
