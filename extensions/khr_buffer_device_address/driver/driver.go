package khr_buffer_device_address_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkDeviceAddress cgoGetBufferDeviceAddressKHR(PFN_vkGetBufferDeviceAddressKHR fn, VkDevice device, VkBufferDeviceAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}

uint64_t cgoGetBufferOpaqueCaptureAddressKHR(PFN_vkGetBufferOpaqueCaptureAddressKHR fn, VkDevice device, VkBufferDeviceAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}

uint64_t cgoGetDeviceMemoryOpaqueCaptureAddressKHR(PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR fn, VkDevice device, VkDeviceMemoryOpaqueCaptureAddressInfoKHR *pInfo) {
	return fn(device, pInfo);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_buffer_device_address

type Driver interface {
	VkGetBufferDeviceAddressKHR(device driver.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) driver.VkDeviceAddress
	VkGetBufferOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) driver.Uint64
	VkGetDeviceMemoryOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfoKHR) driver.Uint64
}

type VkBufferDeviceAddressInfoKHR C.VkBufferDeviceAddressInfoKHR
type VkDeviceMemoryOpaqueCaptureAddressInfoKHR C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR
type VkBufferOpaqueCaptureAddressCreateInfoKHR C.VkBufferOpaqueCaptureAddressCreateInfoKHR
type VkMemoryOpaqueCaptureAddressAllocateInfoKHR C.VkMemoryOpaqueCaptureAddressAllocateInfoKHR
type VkPhysicalDeviceBufferDeviceAddressFeaturesKHR C.VkPhysicalDeviceBufferDeviceAddressFeaturesKHR

type CDriver struct {
	coreDriver driver.Driver

	getBufferDeviceAddress              C.PFN_vkGetBufferDeviceAddressKHR
	getBufferOpaqueCaptureAddress       C.PFN_vkGetBufferOpaqueCaptureAddressKHR
	getDeviceMemoryOpaqueCaptureAddress C.PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getBufferDeviceAddress:              (C.PFN_vkGetBufferDeviceAddressKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetBufferDeviceAddressKHR")))),
		getBufferOpaqueCaptureAddress:       (C.PFN_vkGetBufferOpaqueCaptureAddressKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetBufferOpaqueCaptureAddressKHR")))),
		getDeviceMemoryOpaqueCaptureAddress: (C.PFN_vkGetDeviceMemoryOpaqueCaptureAddressKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceMemoryOpaqueCaptureAddressKHR")))),
	}
}

func (d *CDriver) VkGetBufferDeviceAddressKHR(device driver.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) driver.VkDeviceAddress {
	if d.getBufferDeviceAddress == nil {
		panic("attempt to call extension method vkGetBufferDeviceAddressKHR when extension not present")
	}

	return driver.VkDeviceAddress(C.cgoGetBufferDeviceAddressKHR(
		d.getBufferDeviceAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfoKHR)(pInfo),
	))
}

func (d *CDriver) VkGetBufferOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *VkBufferDeviceAddressInfoKHR) driver.Uint64 {
	if d.getBufferOpaqueCaptureAddress == nil {
		panic("attempt to call extension method vkGetBufferOpaqueCaptureAddressKHR when extension not present")
	}

	return driver.Uint64(C.cgoGetBufferOpaqueCaptureAddressKHR(
		d.getBufferOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkBufferDeviceAddressInfoKHR)(pInfo),
	))
}

func (d *CDriver) VkGetDeviceMemoryOpaqueCaptureAddressKHR(device driver.VkDevice, pInfo *VkDeviceMemoryOpaqueCaptureAddressInfoKHR) driver.Uint64 {
	if d.getDeviceMemoryOpaqueCaptureAddress == nil {
		panic("attempt to call extension method vkGetDeviceMemoryOpaqueCaptureAddressKHR when extension not present")
	}

	return driver.Uint64(C.cgoGetDeviceMemoryOpaqueCaptureAddressKHR(
		d.getDeviceMemoryOpaqueCaptureAddress,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceMemoryOpaqueCaptureAddressInfoKHR)(pInfo),
	))
}
