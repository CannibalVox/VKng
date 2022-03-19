package ext_debug_utils

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

VkResult cgoCreateDebugUtilsMessengerEXT(PFN_vkCreateDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerCreateInfoEXT *pCreateInfo, VkAllocationCallbacks *pAllocator, VkDebugUtilsMessengerEXT* pDebugMessenger) {
	return fn(instance, pCreateInfo, pAllocator, pDebugMessenger);
}

void cgoDestroyDebugUtilsMessengerEXT(PFN_vkDestroyDebugUtilsMessengerEXT fn, VkInstance instance, VkDebugUtilsMessengerEXT debugMessenger, VkAllocationCallbacks* pAllocator) {
	fn(instance, debugMessenger, pAllocator);
}

void cgoCmdBeginDebugUtilsLabelEXT(PFN_vkCmdBeginDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(commandBuffer, pLabelInfo);
}

void cgoCmdEndDebugUtilsLabelEXT(PFN_vkCmdEndDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer) {
	fn(commandBuffer);
}

void cgoCmdInsertDebugUtilsLabelEXT(PFN_vkCmdInsertDebugUtilsLabelEXT fn, VkCommandBuffer commandBuffer, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(commandBuffer, pLabelInfo);
}

void cgoQueueBeginDebugUtilsLabelEXT(PFN_vkQueueBeginDebugUtilsLabelEXT fn, VkQueue queue, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(queue, pLabelInfo);
}

void cgoQueueEndDebugUtilsLabelEXT(PFN_vkQueueEndDebugUtilsLabelEXT fn, VkQueue queue) {
	fn(queue);
}

void cgoQueueInsertDebugUtilsLabelEXT(PFN_vkQueueInsertDebugUtilsLabelEXT fn, VkQueue queue, VkDebugUtilsLabelEXT *pLabelInfo) {
	fn(queue, pLabelInfo);
}

VkResult cgoSetDebugUtilsObjectNameEXT(PFN_vkSetDebugUtilsObjectNameEXT fn, VkDevice device, VkDebugUtilsObjectNameInfoEXT *pNameInfo) {
	return fn(device, pNameInfo);
}

VkResult cgoSetDebugUtilsObjectTagEXT(PFN_vkSetDebugUtilsObjectTagEXT fn, VkDevice device, VkDebugUtilsObjectTagInfoEXT *pTagInfo) {
	return fn(device, pTagInfo);
}

void cgoSubmitDebugUtilsMessageEXT(PFN_vkSubmitDebugUtilsMessageEXT fn, VkInstance instance, VkDebugUtilsMessageSeverityFlagBitsEXT messageSeverity, VkDebugUtilsMessageTypeFlagsEXT messageTypes, VkDebugUtilsMessengerCallbackDataEXT *pCallbackData) {
	fn(instance, messageSeverity, messageTypes, pCallbackData);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ./mocks/driver.go -package mock_debugutils

type CDriver struct {
	createDebugUtilsMessenger  C.PFN_vkCreateDebugUtilsMessengerEXT
	destroyDebugUtilsMessenger C.PFN_vkDestroyDebugUtilsMessengerEXT
	cmdBeginDebugUtilsLabel    C.PFN_vkCmdBeginDebugUtilsLabelEXT
	cmdEndDebugUtilsLabel      C.PFN_vkCmdEndDebugUtilsLabelEXT
	cmdInsertDebugUtilsLabel   C.PFN_vkCmdInsertDebugUtilsLabelEXT
	queueBeginDebugUtilsLabel  C.PFN_vkQueueBeginDebugUtilsLabelEXT
	queueEndDebugUtilsLabel    C.PFN_vkQueueEndDebugUtilsLabelEXT
	queueInsertDebugUtilsLabel C.PFN_vkQueueInsertDebugUtilsLabelEXT
	setDebugUtilsObjectName    C.PFN_vkSetDebugUtilsObjectNameEXT
	setDebugUtilsObjectTag     C.PFN_vkSetDebugUtilsObjectTagEXT
	submitDebugUtilsMessage    C.PFN_vkSubmitDebugUtilsMessageEXT
}

type VkDebugUtilsMessengerCreateInfoEXT C.VkDebugUtilsMessengerCreateInfoEXT
type VkDebugUtilsMessengerEXT C.VkDebugUtilsMessengerEXT
type VkDebugUtilsLabelEXT C.VkDebugUtilsLabelEXT
type VkDebugUtilsObjectNameInfoEXT C.VkDebugUtilsObjectNameInfoEXT
type VkDebugUtilsObjectTagInfoEXT C.VkDebugUtilsObjectTagInfoEXT
type VkDebugUtilsMessageSeverityFlagBitsEXT C.VkDebugUtilsMessageSeverityFlagBitsEXT
type VkDebugUtilsMessageTypeFlagsEXT C.VkDebugUtilsMessageTypeFlagsEXT
type VkDebugUtilsMessengerCallbackDataEXT C.VkDebugUtilsMessengerCallbackDataEXT
type Driver interface {
	VkCreateDebugUtilsMessengerEXT(instance driver.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error)
	VkDestroyDebugUtilsMessengerEXT(instance driver.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *driver.VkAllocationCallbacks)
	VKCmdBeginDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT)
	VkCmdEndDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer)
	VkCmdInsertDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT)
	VkQueueBeginDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT)
	VkQueueEndDebugUtilsLabelEXT(queue driver.VkQueue)
	VkQueueInsertDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT)
	VkSetDebugUtilsObjectNameEXT(device driver.VkDevice, pNameInfo *VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error)
	VkSetDebugUtilsObjectTagEXT(device driver.VkDevice, pTagInfo *VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error)
	VkSubmitDebugUtilsMessageEXT(instance driver.VkInstance, messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *VkDebugUtilsMessengerCallbackDataEXT)
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		createDebugUtilsMessenger:  (C.PFN_vkCreateDebugUtilsMessengerEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateDebugUtilsMessengerEXT")))),
		destroyDebugUtilsMessenger: (C.PFN_vkDestroyDebugUtilsMessengerEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroyDebugUtilsMessengerEXT")))),
		cmdBeginDebugUtilsLabel:    (C.PFN_vkCmdBeginDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdBeginDebugUtilsLabelEXT")))),
		cmdEndDebugUtilsLabel:      (C.PFN_vkCmdEndDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdEndDebugUtilsLabelEXT")))),
		cmdInsertDebugUtilsLabel:   (C.PFN_vkCmdInsertDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdInsertDebugUtilsLabelEXT")))),
		queueBeginDebugUtilsLabel:  (C.PFN_vkQueueBeginDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkQueueBeginDebugUtilsLabelEXT")))),
		queueEndDebugUtilsLabel:    (C.PFN_vkQueueEndDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkQueueEndDebugUtilsLabelEXT")))),
		queueInsertDebugUtilsLabel: (C.PFN_vkQueueInsertDebugUtilsLabelEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkQueueInsertDebugUtilsLabelEXT")))),
		setDebugUtilsObjectName:    (C.PFN_vkSetDebugUtilsObjectNameEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkSetDebugUtilsObjectNameEXT")))),
		setDebugUtilsObjectTag:     (C.PFN_vkSetDebugUtilsObjectTagEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkSetDebugUtilsObjectTagEXT")))),
		submitDebugUtilsMessage:    (C.PFN_vkSubmitDebugUtilsMessageEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkSubmitDebugUtilsMessageEXT")))),
	}
}

func (d *CDriver) VkCreateDebugUtilsMessengerEXT(instance driver.VkInstance, pCreateInfo *VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *VkDebugUtilsMessengerEXT) (common.VkResult, error) {
	if d.createDebugUtilsMessenger == nil {
		panic("attempt to call extension method vkCreateDebugUtilsMessengerEXT when extension not present")
	}

	res := common.VkResult(C.cgoCreateDebugUtilsMessengerEXT(d.createDebugUtilsMessenger,
		C.VkInstance(unsafe.Pointer(instance)),
		(*C.VkDebugUtilsMessengerCreateInfoEXT)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkDebugUtilsMessengerEXT)(pDebugMessenger)))

	return res, res.ToError()
}

func (d *CDriver) VkDestroyDebugUtilsMessengerEXT(instance driver.VkInstance, debugMessenger VkDebugUtilsMessengerEXT, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyDebugUtilsMessenger == nil {
		panic("attempt to call extension method vkDestroyDebugUtilsMessengerEXT when extension not present")
	}

	C.cgoDestroyDebugUtilsMessengerEXT(d.destroyDebugUtilsMessenger,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessengerEXT(debugMessenger),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CDriver) VKCmdBeginDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.cmdBeginDebugUtilsLabel == nil {
		panic("attempt to call extension method VKCmdBeginDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdBeginDebugUtilsLabelEXT(d.cmdBeginDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CDriver) VkCmdEndDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer) {
	if d.cmdEndDebugUtilsLabel == nil {
		panic("attempt to call extension method VkCmdEndDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdEndDebugUtilsLabelEXT(d.cmdEndDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)))
}

func (d *CDriver) VkCmdInsertDebugUtilsLabelEXT(commandBuffer driver.VkCommandBuffer, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.cmdInsertDebugUtilsLabel == nil {
		panic("attempt to call extension method VkCmdInsertDebugUtilsLabelEXT when extension not present")
	}

	C.cgoCmdInsertDebugUtilsLabelEXT(d.cmdInsertDebugUtilsLabel,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CDriver) VkQueueBeginDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.queueBeginDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueBeginDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueBeginDebugUtilsLabelEXT(d.queueBeginDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CDriver) VkQueueEndDebugUtilsLabelEXT(queue driver.VkQueue) {
	if d.queueEndDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueEndDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueEndDebugUtilsLabelEXT(d.queueEndDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)))
}

func (d *CDriver) VkQueueInsertDebugUtilsLabelEXT(queue driver.VkQueue, pLabelInfo *VkDebugUtilsLabelEXT) {
	if d.queueInsertDebugUtilsLabel == nil {
		panic("attempt to call extension method VkQueueInsertDebugUtilsLabelEXT when extension not present")
	}

	C.cgoQueueInsertDebugUtilsLabelEXT(d.queueInsertDebugUtilsLabel,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkDebugUtilsLabelEXT)(unsafe.Pointer(pLabelInfo)))
}

func (d *CDriver) VkSetDebugUtilsObjectNameEXT(device driver.VkDevice, pNameInfo *VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error) {
	if d.setDebugUtilsObjectName == nil {
		panic("attempt to call extension method VkSetDebugUtilsObjectNameEXT when extension not present")
	}

	res := common.VkResult(C.cgoSetDebugUtilsObjectNameEXT(d.setDebugUtilsObjectName,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDebugUtilsObjectNameInfoEXT)(unsafe.Pointer(pNameInfo))))

	return res, res.ToError()
}

func (d *CDriver) VkSetDebugUtilsObjectTagEXT(device driver.VkDevice, pTagInfo *VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error) {
	if d.setDebugUtilsObjectTag == nil {
		panic("attempt to call extension method VkSetDebugUtilsObjectTagEXT when extension not present")
	}

	res := common.VkResult(C.cgoSetDebugUtilsObjectTagEXT(d.setDebugUtilsObjectTag,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDebugUtilsObjectTagInfoEXT)(unsafe.Pointer(pTagInfo))))

	return res, res.ToError()
}

func (d *CDriver) VkSubmitDebugUtilsMessageEXT(instance driver.VkInstance, messageSeverity VkDebugUtilsMessageSeverityFlagBitsEXT, messageTypes VkDebugUtilsMessageTypeFlagsEXT, pCallbackData *VkDebugUtilsMessengerCallbackDataEXT) {
	if d.submitDebugUtilsMessage == nil {
		panic("attempt to call extension method VkSubmitDebugUtilsMessageEXT when extension not present")
	}

	C.cgoSubmitDebugUtilsMessageEXT(d.submitDebugUtilsMessage,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkDebugUtilsMessageSeverityFlagBitsEXT(messageSeverity),
		C.VkDebugUtilsMessageTypeFlagsEXT(messageTypes),
		(*C.VkDebugUtilsMessengerCallbackDataEXT)(unsafe.Pointer(pCallbackData)))
}
