package ext_debug_utils

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"image/color"
	"unsafe"
)

type Label struct {
	Name  string
	Color color.Color
}

type ObjectNameInfo struct {
	Name   string
	Handle uintptr
	Type   common.ObjectType
}

type CallbackData struct {
	MessageIDName   string
	MessageIDNumber int
	Message         string

	QueueLabels         []*Label
	CommandBufferLabels []*Label
	Objects             []*ObjectNameInfo
}

func createLabel(l C.VkDebugUtilsLabelEXT) *Label {
	var name string

	if l.pLabelName != nil {
		name = C.GoString(l.pLabelName)
	}

	r := uint8(float32(l.color[0])*255.0 + 0.001)
	g := uint8(float32(l.color[1])*255.0 + 0.001)
	b := uint8(float32(l.color[2])*255.0 + 0.001)
	a := uint8(float32(l.color[3])*255.0 + 0.001)

	color := color.RGBA{R: r, G: g, B: b, A: a}

	return &Label{
		Name:  name,
		Color: color,
	}
}

func createObjectNameInfo(o C.VkDebugUtilsObjectNameInfoEXT) *ObjectNameInfo {
	objType := common.ObjectType(o.objectType)
	handle := uintptr(o.objectHandle)
	var objName string

	if o.pObjectName != nil {
		objName = C.GoString(o.pObjectName)
	}

	return &ObjectNameInfo{
		Name:   objName,
		Type:   objType,
		Handle: handle,
	}
}

func createCallbackData(d *C.VkDebugUtilsMessengerCallbackDataEXT) *CallbackData {
	var messageIDName, message string

	if d.pMessageIdName != nil {
		messageIDName = C.GoString(d.pMessageIdName)
	}
	if d.pMessage != nil {
		message = C.GoString(d.pMessage)
	}

	var queueLabels, commandBufferLabels []*Label
	var objects []*ObjectNameInfo

	queueLabelCount := int(d.queueLabelCount)
	cQueueLabels := ([]C.VkDebugUtilsLabelEXT)(unsafe.Slice(d.pQueueLabels, queueLabelCount))
	for i := 0; i < queueLabelCount; i++ {
		queueLabels = append(queueLabels, createLabel(cQueueLabels[i]))
	}

	commandBufferLabelCount := int(d.cmdBufLabelCount)
	cCommandBufferLabels := ([]C.VkDebugUtilsLabelEXT)(unsafe.Slice(d.pCmdBufLabels, commandBufferLabelCount))
	for i := 0; i < commandBufferLabelCount; i++ {
		commandBufferLabels = append(commandBufferLabels, createLabel(cCommandBufferLabels[i]))
	}

	objectCount := int(d.objectCount)
	cObjects := ([]C.VkDebugUtilsObjectNameInfoEXT)(unsafe.Slice(d.pObjects, objectCount))
	for i := 0; i < objectCount; i++ {
		objects = append(objects, createObjectNameInfo(cObjects[i]))
	}

	return &CallbackData{
		MessageIDName:       messageIDName,
		MessageIDNumber:     int(d.messageIdNumber),
		Message:             message,
		QueueLabels:         queueLabels,
		CommandBufferLabels: commandBufferLabels,
		Objects:             objects,
	}
}
