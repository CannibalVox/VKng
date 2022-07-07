package khr_device_group_creation

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_device_group_creation

type Extension interface {
	PhysicalDeviceGroups(instance core1_0.Instance, outDataFactory func() *PhysicalDeviceGroupProperties) ([]*PhysicalDeviceGroupProperties, common.VkResult, error)
}
