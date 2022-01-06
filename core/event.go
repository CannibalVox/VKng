package core

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"strings"
	"unsafe"
)

type vulkanEvent struct {
	handle VkEvent
	device VkDevice
	driver Driver
}

func (e *vulkanEvent) Handle() VkEvent {
	return e.handle
}

func (e *vulkanEvent) Destroy() {
	e.driver.VkDestroyEvent(e.device, e.handle, nil)
}

func (e *vulkanEvent) Set() (VkResult, error) {
	return e.driver.VkSetEvent(e.device, e.handle)
}

func (e *vulkanEvent) Reset() (VkResult, error) {
	return e.driver.VkResetEvent(e.device, e.handle)
}

func (e *vulkanEvent) Status() (VkResult, error) {
	return e.driver.VkGetEventStatus(e.device, e.handle)
}

type EventFlags int32

const (
	EventDeviceOnlyKHR EventFlags = C.VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR
)

var eventFlagsToString = map[EventFlags]string{
	EventDeviceOnlyKHR: "Device Only (Khronos Extension)",
}

func (f EventFlags) String() string {
	if f == 0 {
		return "None"
	}

	var hasOne bool
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		checkBit := EventFlags(1 << i)
		if (f & checkBit) != 0 {
			str, hasStr := eventFlagsToString[checkBit]
			if hasStr {
				if hasOne {
					sb.WriteRune('|')
				}
				sb.WriteString(str)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type EventOptions struct {
	Flags EventFlags

	common.HaveNext
}

func (o *EventOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkEventCreateInfo)(allocator.Malloc(C.sizeof_struct_VkEventCreateInfo))
	createInfo.sType = C.VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
	createInfo.flags = C.VkEventCreateFlags(o.Flags)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}