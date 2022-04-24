package khr_external_memory_capabilities

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_memory_capabilities

type Extension interface {
	ExternalBufferProperties(physicalDevice core1_0.PhysicalDevice, o ExternalBufferOptions, outData *ExternalBufferOutData) error
}
