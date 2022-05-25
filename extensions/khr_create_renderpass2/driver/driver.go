package khr_create_renderpass2_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdBeginRenderPass2KHR(PFN_vkCmdBeginRenderPass2KHR fn, VkCommandBuffer commandBuffer, VkRenderPassBeginInfo *pRenderPassBegin, VkSubpassBeginInfoKHR *pSubpassBegininfo) {
	fn(commandBuffer, pRenderPassBegin, pSubpassBegininfo);
}

void cgoCmdEndRenderPass2KHR(PFN_vkCmdEndRenderPass2KHR fn, VkCommandBuffer commandBuffer, VkSubpassEndInfoKHR *pSubpassEndInfo) {
	fn(commandBuffer, pSubpassEndInfo);
}

void cgoCmdNextSubpass2KHR(PFN_vkCmdNextSubpass2KHR fn, VkCommandBuffer commandBuffer, VkSubpassBeginInfoKHR *pSubpassBeginInfo, VkSubpassEndInfoKHR *pSubpassEndInfo) {
	fn(commandBuffer, pSubpassBeginInfo, pSubpassEndInfo);
}

VkResult cgoCreateRenderPass2KHR(PFN_vkCreateRenderPass2KHR fn, VkDevice device, VkRenderPassCreateInfo2KHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkRenderPass *pRenderPass) {
	return fn(device, pCreateInfo, pAllocator, pRenderPass);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_create_renderpass2

type Driver interface {
	VkCmdBeginRenderPass2KHR(commandBuffer driver.VkCommandBuffer, pRenderPassBegin *driver.VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfoKHR)
	VKCmdEndRenderPass2KHR(commandBuffer driver.VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfoKHR)
	VkCmdNextSubpass2KHR(commandBuffer driver.VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfoKHR, pSubpassEndInfo *VkSubpassEndInfoKHR)
	VkCreateRenderPass2KHR(device driver.VkDevice, pCreateInfo *VkRenderPassCreateInfo2KHR, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error)
}

type VkAttachmentDescription2KHR C.VkAttachmentDescription2KHR
type VkAttachmentReference2KHR C.VkAttachmentReference2KHR
type VkRenderPassCreateInfo2KHR C.VkRenderPassCreateInfo2KHR
type VkSubpassBeginInfoKHR C.VkSubpassBeginInfoKHR
type VkSubpassDependency2KHR C.VkSubpassDependency2KHR
type VkSubpassDescription2KHR C.VkSubpassDescription2KHR
type VkSubpassEndInfoKHR C.VkSubpassEndInfoKHR

type CDriver struct {
	coreDriver driver.Driver

	beginRenderPass  C.PFN_vkCmdBeginRenderPass2KHR
	endRenderPass    C.PFN_vkCmdEndRenderPass2KHR
	nextSubpass      C.PFN_vkCmdNextSubpass2KHR
	createRenderPass C.PFN_vkCreateRenderPass2KHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		beginRenderPass:  (C.PFN_vkCmdBeginRenderPass2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdBeginRenderPass2KHR")))),
		endRenderPass:    (C.PFN_vkCmdEndRenderPass2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdEndRenderPass2KHR")))),
		nextSubpass:      (C.PFN_vkCmdNextSubpass2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdNextSubpass2KHR")))),
		createRenderPass: (C.PFN_vkCreateRenderPass2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateRenderPass2KHR")))),
	}
}

func (d *CDriver) VkCmdBeginRenderPass2KHR(commandBuffer driver.VkCommandBuffer, pRenderPassBegin *driver.VkRenderPassBeginInfo, pSubpassBeginInfo *VkSubpassBeginInfoKHR) {
	if d.beginRenderPass == nil {
		panic("attempt to call extension method vkCmdBeginRenderPass2KHR when extension not present")
	}

	C.cgoCmdBeginRenderPass2KHR(
		d.beginRenderPass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkRenderPassBeginInfo)(unsafe.Pointer(pRenderPassBegin)),
		(*C.VkSubpassBeginInfoKHR)(pSubpassBeginInfo),
	)
}

func (d *CDriver) VKCmdEndRenderPass2KHR(commandBuffer driver.VkCommandBuffer, pSubpassEndInfo *VkSubpassEndInfoKHR) {
	if d.endRenderPass == nil {
		panic("attempt to call extension method vkCmdEndRenderPass2KHR when extension not present")
	}

	C.cgoCmdEndRenderPass2KHR(
		d.endRenderPass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassEndInfo)(pSubpassEndInfo),
	)
}

func (d *CDriver) VkCmdNextSubpass2KHR(commandBuffer driver.VkCommandBuffer, pSubpassBeginInfo *VkSubpassBeginInfoKHR, pSubpassEndInfo *VkSubpassEndInfoKHR) {
	if d.nextSubpass == nil {
		panic("attempt to call extension method vkCmdNextSubpass2KHR when extension not present")
	}

	C.cgoCmdNextSubpass2KHR(
		d.nextSubpass,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkSubpassBeginInfoKHR)(pSubpassBeginInfo),
		(*C.VkSubpassEndInfoKHR)(pSubpassEndInfo),
	)
}

func (d *CDriver) VkCreateRenderPass2KHR(device driver.VkDevice, pCreateInfo *VkRenderPassCreateInfo2KHR, pAllocator *driver.VkAllocationCallbacks, pRenderPass *driver.VkRenderPass) (common.VkResult, error) {
	if d.createRenderPass == nil {
		panic("attempt to call extension method vkCreateRenderPass2KHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateRenderPass2KHR(
		d.createRenderPass,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo2KHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkRenderPass)(unsafe.Pointer(pRenderPass)),
	))

	return res, res.ToError()
}
