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

type DebugUtilsMessengerCallbackData struct {
	Flags CallbackDataFlags

	MessageIDName   string
	MessageIDNumber int
	Message         string

	QueueLabels  []DebugUtilsLabel
	CmdBufLabels []DebugUtilsLabel
	Objects      []DebugUtilsObjectNameInfo

	common.NextOptions
}

func (c DebugUtilsMessengerCallbackData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
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
	queueLabelsPtr, err := common.AllocOptionSlice[C.VkDebugUtilsLabelEXT, DebugUtilsLabel](allocator, c.QueueLabels)
	if err != nil {
		return nil, err
	}

	commandBufferLabelCount := len(c.CmdBufLabels)
	commandBufferLabelPtr, err := common.AllocOptionSlice[C.VkDebugUtilsLabelEXT, DebugUtilsLabel](allocator, c.CmdBufLabels)
	if err != nil {
		return nil, err
	}

	objectCount := len(c.Objects)
	objectPtr, err := common.AllocOptionSlice[C.VkDebugUtilsObjectNameInfoEXT, DebugUtilsObjectNameInfo](allocator, c.Objects)
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

func (c *DebugUtilsMessengerCallbackData) PopulateFromCPointer(cPointer unsafe.Pointer) error {
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
	c.QueueLabels = make([]DebugUtilsLabel, queueLabelCount)
	queueCPointer := unsafe.Pointer(callbackData.pQueueLabels)
	labelSize := uintptr(C.sizeof_struct_VkDebugUtilsLabelEXT)
	for i := 0; i < queueLabelCount; i++ {
		c.QueueLabels[i].PopulateFromCPointer(queueCPointer)

		queueCPointer = unsafe.Add(queueCPointer, labelSize)
	}

	commandBufferLabelCount := int(callbackData.cmdBufLabelCount)
	c.CmdBufLabels = make([]DebugUtilsLabel, commandBufferLabelCount)
	cmdBufCPointer := unsafe.Pointer(callbackData.pCmdBufLabels)
	for i := 0; i < commandBufferLabelCount; i++ {
		c.CmdBufLabels[i].PopulateFromCPointer(cmdBufCPointer)

		cmdBufCPointer = unsafe.Add(cmdBufCPointer, labelSize)
	}

	objectCount := int(callbackData.objectCount)
	c.Objects = make([]DebugUtilsObjectNameInfo, objectCount)
	objectsPointer := unsafe.Pointer(callbackData.pObjects)
	objectNameSize := uintptr(C.sizeof_struct_VkDebugUtilsObjectNameInfoEXT)
	for i := 0; i < objectCount; i++ {
		c.Objects[i].PopulateFromCPointer(objectsPointer)

		objectsPointer = unsafe.Add(objectsPointer, objectNameSize)
	}

	return nil
}
