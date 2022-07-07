package khr_timeline_semaphore

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_timeline_semaphore_driver "github.com/CannibalVox/VKng/extensions/khr_timeline_semaphore/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
)

type VulkanExtension struct {
	driver khr_timeline_semaphore_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_timeline_semaphore_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_timeline_semaphore_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) SemaphoreCounterValue(semaphore core1_0.Semaphore) (uint64, common.VkResult, error) {

	var value driver.Uint64
	res, err := e.driver.VkGetSemaphoreCounterValueKHR(
		semaphore.DeviceHandle(),
		semaphore.Handle(),
		&value,
	)
	if err != nil {
		return 0, res, err
	}

	return uint64(value), res, nil
}

func (e *VulkanExtension) SignalSemaphore(device core1_0.Device, o SemaphoreSignalInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	signalPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkSignalSemaphoreKHR(
		device.Handle(),
		(*khr_timeline_semaphore_driver.VkSemaphoreSignalInfoKHR)(signalPtr),
	)
}

func (e *VulkanExtension) WaitSemaphores(device core1_0.Device, timeout time.Duration, o SemaphoreWaitInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	waitPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkWaitSemaphoresKHR(
		device.Handle(),
		(*khr_timeline_semaphore_driver.VkSemaphoreWaitInfoKHR)(waitPtr),
		driver.Uint64(common.TimeoutNanoseconds(timeout)),
	)
}
