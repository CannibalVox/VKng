package khr_descriptor_update_template_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateDescriptorUpdateTemplateKHR(PFN_vkCreateDescriptorUpdateTemplateKHR fn, VkDevice device, VkDescriptorUpdateTemplateCreateInfoKHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkDescriptorUpdateTemplateKHR *pDescriptorUpdateTemplate) {
	return fn(device, pCreateInfo, pAllocator, pDescriptorUpdateTemplate);
}

void cgoDestroyDescriptorUpdateTemplateKHR(PFN_vkDestroyDescriptorUpdateTemplateKHR fn, VkDevice device, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, VkAllocationCallbacks *pAllocator) {
	fn(device, descriptorUpdateTemplate, pAllocator);
}

void cgoUpdateDescriptorSetWithTemplateKHR(PFN_vkUpdateDescriptorSetWithTemplateKHR fn, VkDevice device, VkDescriptorSet descriptorSet, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, void *pData) {
	fn(device, descriptorSet, descriptorUpdateTemplate, pData);
}

void cgoCmdPushDescriptorSetWithTemplateKHR(PFN_vkCmdPushDescriptorSetWithTemplateKHR fn, VkCommandBuffer commandBuffer, VkDescriptorUpdateTemplateKHR descriptorUpdateTemplate, VkPipelineLayout layout, uint32_t set, void *pData) {
	fn(commandBuffer, descriptorUpdateTemplate, layout, set, pData);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_descriptor_update_template

type Driver interface {
	VkCreateDescriptorUpdateTemplateKHR(device driver.VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplateKHR) (common.VkResult, error)
	VkDestroyDescriptorUpdateTemplateKHR(device driver.VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pAllocator *driver.VkAllocationCallbacks)
	VkUpdateDescriptorSetWithTemplateKHR(device driver.VkDevice, descriptorSet driver.VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pData unsafe.Pointer)
	VkCmdPushDescriptorSetWithTemplateKHR(commandBuffer driver.VkCommandBuffer, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, layout driver.VkPipelineLayout, set driver.Uint32, pData unsafe.Pointer)
}

type VkDescriptorUpdateTemplateKHR driver.VulkanHandle
type VkDescriptorUpdateTemplateCreateInfoKHR C.VkDescriptorUpdateTemplateCreateInfoKHR
type VkDescriptorUpdateTemplateEntryKHR C.VkDescriptorUpdateTemplateEntryKHR

type CDriver struct {
	coreDriver driver.Driver

	createDescriptorUpdateTemplate   C.PFN_vkCreateDescriptorUpdateTemplateKHR
	destroyDescriptorUpdateTemplate  C.PFN_vkDestroyDescriptorUpdateTemplateKHR
	updateDescriptorSetWithTemplate  C.PFN_vkUpdateDescriptorSetWithTemplateKHR
	cmdPushDescriptorSetWithTemplate C.PFN_vkCmdPushDescriptorSetWithTemplateKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		createDescriptorUpdateTemplate:   (C.PFN_vkCreateDescriptorUpdateTemplateKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateDescriptorUpdateTemplateKHR")))),
		destroyDescriptorUpdateTemplate:  (C.PFN_vkDestroyDescriptorUpdateTemplateKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroyDescriptorUpdateTemplateKHR")))),
		updateDescriptorSetWithTemplate:  (C.PFN_vkUpdateDescriptorSetWithTemplateKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkUpdateDescriptorSetWithTemplateKHR")))),
		cmdPushDescriptorSetWithTemplate: (C.PFN_vkCmdPushDescriptorSetWithTemplateKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdPushDescriptorSetWithTemplateKHR")))),
	}
}

func (d *CDriver) VkCreateDescriptorUpdateTemplateKHR(device driver.VkDevice, pCreateInfo *VkDescriptorUpdateTemplateCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pDescriptorUpdateTemplate *VkDescriptorUpdateTemplateKHR) (common.VkResult, error) {
	if d.createDescriptorUpdateTemplate == nil {
		panic("attempt to call extension method vkCreateDescriptorUpdateTemplateKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateDescriptorUpdateTemplateKHR(d.createDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDescriptorUpdateTemplateCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDescriptorUpdateTemplateKHR)(unsafe.Pointer(pDescriptorUpdateTemplate))))

	return res, res.ToError()
}

func (d *CDriver) VkDestroyDescriptorUpdateTemplateKHR(device driver.VkDevice, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyDescriptorUpdateTemplate == nil {
		panic("attempt to call extension method vkDestroyDescriptorUpdateTemplateKHR when extension not present")
	}

	C.cgoDestroyDescriptorUpdateTemplateKHR(d.destroyDescriptorUpdateTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CDriver) VkUpdateDescriptorSetWithTemplateKHR(device driver.VkDevice, descriptorSet driver.VkDescriptorSet, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, pData unsafe.Pointer) {
	if d.updateDescriptorSetWithTemplate == nil {
		panic("attempt to call extension method vkUpdateDescriptorSetWithTemplateKHR when extension not present")
	}

	C.cgoUpdateDescriptorSetWithTemplateKHR(d.updateDescriptorSetWithTemplate,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkDescriptorSet(unsafe.Pointer(descriptorSet)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		pData)
}

func (d *CDriver) VkCmdPushDescriptorSetWithTemplateKHR(commandBuffer driver.VkCommandBuffer, descriptorUpdateTemplate VkDescriptorUpdateTemplateKHR, layout driver.VkPipelineLayout, set driver.Uint32, pData unsafe.Pointer) {
	if d.cmdPushDescriptorSetWithTemplate == nil {
		panic("attempt to call extension method vkUpdateDescriptorSetWithTemplateKHR when prerequisite not present")
	}

	C.cgoCmdPushDescriptorSetWithTemplateKHR(d.cmdPushDescriptorSetWithTemplate,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(descriptorUpdateTemplate)),
		C.VkPipelineLayout(unsafe.Pointer(layout)),
		C.uint32_t(set),
		pData)
}
