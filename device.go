package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
 */
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type PhysicalDevice struct {
	handle C.VkPhysicalDevice
}

func (d *PhysicalDevice) QueueFamilyProperties(allocator cgoalloc.Allocator) ([]*QueueFamily, error) {
	count := C.uint32_t(0)
	C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, &count, nil)

	if count == 0 {
		return nil, nil
	}

	allocatedHandles := allocator.Malloc(count*unsafe.Sizeof(C.VkQueueFamilyProperties{}))
	defer allocator.Free(allocatedHandles)
	familyProperties := (*[1<<30]C.VkQueueFamilyProperties)(allocatedHandles)

	 C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, &count, (*C.VkQueueFamilyProperties)(allocatedHandles))

	goCount := uint32(count)
	var queueFamilies []*QueueFamily
	for i := uint32(0); i < goCount; i++ {
		queueFamilies = append(queueFamilies, &QueueFamily{
			Flags: QueueFlags(familyProperties[i].queueFlags),
			QueueCount: uint32(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: Extent3D{
				Width: uint32(familyProperties[i].minImageTransferGranularity.width),
				Height: uint32(familyProperties[i].minImageTransferGranularity.height),
				Depth: uint32(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies, nil
}

func (d *PhysicalDevice) DeviceBuilder() *DeviceBuilder {
	return &DeviceBuilder{
		physicalDeviceHandle: d.handle,
	}
}

type QueueFamilyBuilder struct {
	parent *DeviceBuilder
	queueFamilyIndex uint32
	queuePriorities []float32
}

func (b *QueueFamilyBuilder) AddQueuePriority(priority float32) *QueueFamilyBuilder {
	b.queuePriorities = append(b.queuePriorities, priority)
	return b
}

func (b *QueueFamilyBuilder) Complete() *DeviceBuilder {
	b.parent.queueFamilies = append(b.parent.queueFamilies, b)
	return b.parent
}

type DeviceBuilder struct {
	physicalDeviceHandle C.VkPhysicalDevice
	queueFamilies []*QueueFamilyBuilder
}

func (b *DeviceBuilder) AddQueueFamily(familyIndex uint32) *QueueFamilyBuilder {
	return &QueueFamilyBuilder{
		parent: b,
		queueFamilyIndex: familyIndex,
	}
}

func (b *DeviceBuilder) Build(allocator cgoalloc.Allocator) (*Device, error) {
	if len(b.queueFamilies) == 0 {
		return nil, stacktrace.NewError("building a vulkan device before adding queue families")
	}

	queueFamilyPtr := allocator.Malloc(len(b.queueFamilies)*int(unsafe.Sizeof([1]C.VkDeviceQueueCreateInfo{})))
	defer allocator.Free(queueFamilyPtr)
	queueFamilyArray := (*[1<<30]C.VkDeviceQueueCreateInfo)(queueFamilyPtr)

	for idx, queueFamily := range b.queueFamilies {
		if len(queueFamily.queuePriorities) == 0 {
			return nil, stacktrace.NewError("building vulkan device: queue family %d had no queue priorities", queueFamily.queueFamilyIndex)
		}

		prioritiesPtr := allocator.Malloc(len(queueFamily.queuePriorities)*int(unsafe.Sizeof(C.float(0))))
		defer allocator.Free(prioritiesPtr)
		prioritiesArray := (*[1<<30]C.float)(prioritiesPtr)
		for idx, priority := range queueFamily.queuePriorities {
			prioritiesArray[idx] = C.float(priority)
		}

		queueFamilyArray[idx] = C.VkDeviceQueueCreateInfo {
			sType: C.VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO,
			queueCount: C.uint32_t(len(queueFamily.queuePriorities)),
			queueFamilyIndex: C.uint32_t(queueFamily.queueFamilyIndex),
			pQueuePriorities: (*C.float)(prioritiesPtr),
		}
	}

	deviceCreate := &C.VkDeviceCreateInfo {
		sType: C.VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO,
		queueCreateInfoCount: C.uint32_t(len(b.queueFamilies)),
		pQueueCreateInfos: (*C.VkDeviceQueueCreateInfo)(queueFamilyPtr),
	}

	var deviceHandle C.VkDevice
	res := C.vkCreateDevice(b.physicalDeviceHandle, deviceCreate, nil, &deviceHandle)
	err := VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Device{handle: deviceHandle}, nil
}

type Device struct {
	handle C.VkDevice
}

func (d *Device) Destroy() {
	C.vkDestroyDevice(d.handle, nil)
}

func (d *Device) CommandPoolBuilder() *CommandPoolBuilder {
	return &CommandPoolBuilder{
		deviceHandle: d.handle,
		graphicsQueueFamily: 0xFFFFFFFF,
	}
}
