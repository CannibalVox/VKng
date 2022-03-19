package ext_debug_utils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type CallbackDataOptions struct {
	Flags CallbackDataFlags

	MessageIDName   string
	MessageIDNumber int
	Message         string

	QueueLabels         []LabelOptions
	CommandBufferLabels []LabelOptions
	Objects             []ObjectNameOptions

	common.HaveNext
}

func (c CallbackDataOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkDebugUtilsMessengerCallbackDataEXT)
	}

	callbackData := (*C.VkDebugUtilsMessengerCallbackDataEXT)(preallocatedPointer)
	callbackData.sType = C.VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT
	callbackData.pNext = next
	callbackData.flags = (C.VkDebugUtilsMessengerCallbackDataFlagsEXT)(c.Flags)
	callbackData.pMessageIdName = (*C.char)(allocator.CString(c.MessageIDName))
	callbackData.messageIdNumber = C.int32_t(c.MessageIDNumber)
	callbackData.pMessage = (*C.char)(allocator.CString(c.Message))

	queueLabelCount := len(c.QueueLabels)
	queueLabelsPtr, err := common.AllocOptionSlice[C.VkDebugUtilsLabelEXT, LabelOptions](allocator, c.QueueLabels)
	if err != nil {
		return nil, err
	}

	commandBufferLabelCount := len(c.CommandBufferLabels)
	commandBufferLabelPtr, err := common.AllocOptionSlice[C.VkDebugUtilsLabelEXT, LabelOptions](allocator, c.CommandBufferLabels)
	if err != nil {
		return nil, err
	}

	objectCount := len(c.Objects)
	objectPtr, err := common.AllocOptionSlice[C.VkDebugUtilsObjectNameInfoEXT, ObjectNameOptions](allocator, c.Objects)
	if err != nil {
		return nil, err
	}

	callbackData.queueLabelCount = C.uint32_t(queueLabelCount)
	callbackData.pQueueLabels = (*C.VkDebugUtilsLabelEXT)(queueLabelsPtr)
	callbackData.cmdBufLabelCount = C.uint32_t(commandBufferLabelCount)
	callbackData.pCmdBufLabels = (*C.VkDebugUtilsLabelEXT)(commandBufferLabelPtr)
	callbackData.objectCount = C.uint32_t(objectCount)
	callbackData.pObjects = (*C.VkDebugUtilsObjectNameInfoEXT)(objectPtr)

	return preallocatedPointer, nil
}

func (c CallbackDataOptions) PopulateOutData(cPointer unsafe.Pointer) (unsafe.Pointer, error) {
	callbackData := (*C.VkDebugUtilsMessengerCallbackDataEXT)(cPointer)

	c.MessageIDName = ""
	c.Message = ""

	if callbackData.pMessageIdName != nil {
		c.MessageIDName = C.GoString(callbackData.pMessageIdName)
	}
	if callbackData.pMessage != nil {
		c.Message = C.GoString(callbackData.pMessage)
	}

	c.MessageIDNumber = int(callbackData.messageIdNumber)

	queueLabelCount := int(callbackData.queueLabelCount)
	c.QueueLabels = make([]LabelOptions, queueLabelCount)
	err := common.PopulateOutDataSlice[C.VkDebugUtilsLabelEXT, LabelOptions](c.QueueLabels, unsafe.Pointer(callbackData.pQueueLabels))
	if err != nil {
		return nil, err
	}

	commandBufferLabelCount := int(callbackData.cmdBufLabelCount)
	c.CommandBufferLabels = make([]LabelOptions, commandBufferLabelCount)
	err = common.PopulateOutDataSlice[C.VkDebugUtilsLabelEXT, LabelOptions](c.CommandBufferLabels, unsafe.Pointer(callbackData.pCmdBufLabels))
	if err != nil {
		return nil, err
	}

	objectCount := int(callbackData.objectCount)
	c.Objects = make([]ObjectNameOptions, objectCount)
	err = common.PopulateOutDataSlice[C.VkDebugUtilsObjectNameInfoEXT, ObjectNameOptions](c.Objects, unsafe.Pointer(callbackData.pObjects))
	if err != nil {
		return nil, err
	}

	return callbackData.pNext, nil
}
