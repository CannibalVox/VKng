package khr_external_fence_capabilities

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_fence_capabilities

type Extension interface {
	ExternalFenceProperties(physicalDevice core1_0.PhysicalDevice, o ExternalFencePropertiesOptions, outData *ExternalFencePropertiesOutData) error
}
