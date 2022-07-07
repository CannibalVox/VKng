package khr_external_semaphore_capabilities

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_external_semaphore_capabilities

type Extension interface {
	PhysicalDeviceExternalSemaphoreProperties(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error
}
