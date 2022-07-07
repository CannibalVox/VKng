package khr_maintenance3

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_maintenance3

type Extension interface {
	DescriptorSetLayoutSupport(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateInfo, support *DescriptorSetLayoutSupport) error
}
