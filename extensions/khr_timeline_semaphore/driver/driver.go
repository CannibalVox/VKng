package khr_timeline_semaphore_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoGetSemaphoreCounterValueKHR(PFN_vkGetSemaphoreCounterValueKHR fn, VkDevice device, VkSemaphore semaphore, uint64_t *pValue) {
	return fn(device, semaphore, pValue);
}

VkResult cgoSignalSemaphoreKHR(PFN_vkSignalSemaphoreKHR fn, VkDevice device, VkSemaphoreSignalInfo *pSignalInfo) {
	return fn(device, pSignalInfo);
}

VkResult cgoWaitSemaphoresKHR(PFN_vkWaitSemaphoresKHR fn, VkDevice device, VkSemaphoreWaitInfo *pWaitInfo, uint64_t timeout) {
	return fn(device, pWaitInfo, timeout);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_timeline_semaphore

type Driver interface {
	VkGetSemaphoreCounterValueKHR(device driver.VkDevice, semaphore driver.VkSemaphore, pValue *driver.Uint64) (common.VkResult, error)
	VkSignalSemaphoreKHR(device driver.VkDevice, pSignalInfo *VkSemaphoreSignalInfoKHR) (common.VkResult, error)
	VkWaitSemaphoresKHR(device driver.VkDevice, pWaitInfo *VkSemaphoreWaitInfoKHR, timeout driver.Uint64) (common.VkResult, error)
}

type VkSemaphoreSignalInfoKHR C.VkSemaphoreSignalInfoKHR
type VkSemaphoreWaitInfoKHR C.VkSemaphoreWaitInfoKHR
type VkPhysicalDeviceTimelineSemaphoreFeaturesKHR C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR
type VkPhysicalDeviceTimelineSemaphorePropertiesKHR C.VkPhysicalDeviceTimelineSemaphorePropertiesKHR
type VkSemaphoreTypeCreateInfoKHR C.VkSemaphoreTypeCreateInfoKHR
type VkTimelineSemaphoreSubmitInfoKHR C.VkTimelineSemaphoreSubmitInfoKHR

type CDriver struct {
	coreDriver driver.Driver

	getSemaphoreCounterValue C.PFN_vkGetSemaphoreCounterValueKHR
	signalSemaphore          C.PFN_vkSignalSemaphoreKHR
	waitSemaphores           C.PFN_vkWaitSemaphoresKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getSemaphoreCounterValue: (C.PFN_vkGetSemaphoreCounterValueKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetSemaphoreCounterValueKHR")))),
		signalSemaphore:          (C.PFN_vkSignalSemaphoreKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkSignalSemaphoreKHR")))),
		waitSemaphores:           (C.PFN_vkWaitSemaphoresKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkWaitSemaphoresKHR")))),
	}
}

func (d *CDriver) VkGetSemaphoreCounterValueKHR(device driver.VkDevice, semaphore driver.VkSemaphore, pValue *driver.Uint64) (common.VkResult, error) {
	if d.getSemaphoreCounterValue == nil {
		panic("attempt to call extension method vkGetSemaphoreCounterValueKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetSemaphoreCounterValueKHR(
		d.getSemaphoreCounterValue,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		(*C.uint64_t)(pValue),
	))
	return res, res.ToError()
}

func (d *CDriver) VkSignalSemaphoreKHR(device driver.VkDevice, pSignalInfo *VkSemaphoreSignalInfoKHR) (common.VkResult, error) {
	if d.signalSemaphore == nil {
		panic("attempt to call extension method vkSignalSemaphoreKHR when extension not present")
	}

	res := common.VkResult(C.cgoSignalSemaphoreKHR(
		d.signalSemaphore,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreSignalInfoKHR)(pSignalInfo),
	))

	return res, res.ToError()
}

func (d *CDriver) VkWaitSemaphoresKHR(device driver.VkDevice, pWaitInfo *VkSemaphoreWaitInfoKHR, timeout driver.Uint64) (common.VkResult, error) {
	if d.waitSemaphores == nil {
		panic("attempt to call extension method vkWaitSemaphoresKHR when extension not present")
	}

	res := common.VkResult(C.cgoWaitSemaphoresKHR(
		d.waitSemaphores,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSemaphoreWaitInfoKHR)(pWaitInfo),
		C.uint64_t(timeout),
	))

	return res, res.ToError()
}
