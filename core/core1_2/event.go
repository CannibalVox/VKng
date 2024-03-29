package core1_2

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanEvent struct {
	core1_1.Event
}

func PromoteEvent(event core1_0.Event) Event {
	if !event.APIVersion().IsAtLeast(common.Vulkan1_2) {
		return nil
	}

	promotedEvent := core1_1.PromoteEvent(event)
	return event.Driver().ObjectStore().GetOrCreate(
		driver.VulkanHandle(event.Handle()),
		driver.Core1_2,
		func() any {
			return &VulkanEvent{promotedEvent}
		}).(Event)
}
