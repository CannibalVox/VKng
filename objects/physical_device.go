package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/creation"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type PhysicalDeviceHandle C.VkPhysicalDevice
type PhysicalDevice struct {
	handle C.VkPhysicalDevice
}

func (d *PhysicalDevice) Handle() PhysicalDeviceHandle {
	return PhysicalDeviceHandle(d.handle)
}

func (d *PhysicalDevice) QueueFamilyProperties(allocator cgoalloc.Allocator) ([]*VKng.QueueFamily, error) {
	count := (*C.uint32_t)(allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer allocator.Free(unsafe.Pointer(count))

	C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, nil)

	if *count == 0 {
		return nil, nil
	}

	goCount := int(*count)

	allocatedHandles := allocator.Malloc(goCount * int(unsafe.Sizeof(C.VkQueueFamilyProperties{})))
	defer allocator.Free(allocatedHandles)
	familyProperties := ([]C.VkQueueFamilyProperties)(unsafe.Slice((*C.VkQueueFamilyProperties)(allocatedHandles), int(*count)))

	C.vkGetPhysicalDeviceQueueFamilyProperties(d.handle, count, (*C.VkQueueFamilyProperties)(allocatedHandles))

	var queueFamilies []*VKng.QueueFamily
	for i := 0; i < goCount; i++ {
		queueFamilies = append(queueFamilies, &VKng.QueueFamily{
			Flags:              VKng.QueueFlags(familyProperties[i].queueFlags),
			QueueCount:         uint32(familyProperties[i].queueCount),
			TimestampValidBits: uint32(familyProperties[i].timestampValidBits),
			MinImageTransferGranularity: VKng.Extent3D{
				Width: uint32(familyProperties[i].minImageTransferGranularity.width),
				Height: uint32(familyProperties[i].minImageTransferGranularity.height),
				Depth: uint32(familyProperties[i].minImageTransferGranularity.depth),
			},
		})
	}

	return queueFamilies, nil
}

func (d *PhysicalDevice) CreateDevice(allocator cgoalloc.Allocator, options *creation.DeviceOptions) (*Device, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var deviceHandle C.VkDevice
	res := C.vkCreateDevice(d.handle, (*C.VkDeviceCreateInfo)(createInfo), nil, &deviceHandle)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Device{handle: deviceHandle}, nil
}

func (d *PhysicalDevice) Properties(allocator cgoalloc.Allocator) (*PhysicalDeviceProperties, error) {
	properties := (*C.VkPhysicalDeviceProperties)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceProperties{}))))
	defer allocator.Free(unsafe.Pointer(properties))

	C.vkGetPhysicalDeviceProperties(d.handle, properties)

	return createPhysicalDeviceProperties(properties)
}

func (d *PhysicalDevice) Features(allocator cgoalloc.Allocator) (*VKng.PhysicalDeviceFeatures, error) {
	features := (*C.VkPhysicalDeviceFeatures)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkPhysicalDeviceFeatures{}))))
	defer allocator.Free(unsafe.Pointer(features))

	C.vkGetPhysicalDeviceFeatures(d.handle, features)

	return createPhysicalDeviceFeatures(features), nil
}

func (d *PhysicalDevice) AvailableExtensions(allocator cgoalloc.Allocator) (map[string]*VKng.ExtensionProperties, error) {
	extensionCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(extensionCountPtr)

	extensionCount := (*C.uint32_t)(extensionCountPtr)

	C.vkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, nil)

	if *extensionCount == 0 {
		return nil, nil
	}

	extensionTotal := int(*extensionCount)
	extensionsPtr := allocator.Malloc(extensionTotal * int(unsafe.Sizeof([1]C.VkExtensionProperties{})))
	defer allocator.Free(extensionsPtr)

	typedExtensionsPtr := (*C.VkExtensionProperties)(extensionsPtr)
	C.vkEnumerateDeviceExtensionProperties(d.handle, nil, extensionCount, typedExtensionsPtr)

	retVal := make(map[string]*VKng.ExtensionProperties)
	extensionSlice := ([]C.VkExtensionProperties)(unsafe.Slice(typedExtensionsPtr, extensionTotal))

	for i := 0; i < extensionTotal; i++ {
		extension := extensionSlice[i]

		outExtension := &VKng.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   VKng.Version(extension.specVersion),
		}

		existingExtension, ok := retVal[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		retVal[outExtension.ExtensionName] = outExtension
	}

	return retVal, nil
}
