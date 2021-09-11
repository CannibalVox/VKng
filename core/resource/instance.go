package resource

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func CreateInstance(allocator cgoalloc.Allocator, load *loader.Loader, options *InstanceOptions) (*Instance, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var instanceHandle loader.VkInstance

	res, err := load.VkCreateInstance((*loader.VkInstanceCreateInfo)(createInfo), nil, &instanceHandle)
	if err != nil {
		return nil, res, err
	}

	instanceLoader, err := load.CreateInstanceLoader((loader.VkInstance)(instanceHandle))
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	return &Instance{
		loader: instanceLoader,
		handle: instanceHandle,
	}, res, nil
}

type InstanceHandle C.VkInstance
type Instance struct {
	loader *loader.Loader
	handle loader.VkInstance
}

func (i *Instance) Loader() *loader.Loader {
	return i.loader
}

func (i *Instance) Handle() loader.VkInstance {
	return i.handle
}

func (i *Instance) Destroy() error {
	err := i.loader.VkDestroyInstance(i.handle, nil)
	if err != nil {
		return err
	}

	i.loader.Destroy()
	return nil
}

func (i *Instance) PhysicalDevices(allocator cgoalloc.Allocator) ([]*PhysicalDevice, loader.VkResult, error) {
	count := (*loader.Uint32)(allocator.Malloc(int(unsafe.Sizeof(loader.Uint32(0)))))
	defer allocator.Free(unsafe.Pointer(count))

	res, err := i.loader.VkEnumeratePhysicalDevices(i.handle, count, nil)
	if err != nil {
		return nil, res, err
	}

	if *count == 0 {
		return nil, res, nil
	}

	allocatedHandles := allocator.Malloc(int(uintptr(*count) * unsafe.Sizeof([1]loader.VkPhysicalDevice{})))
	defer allocator.Free(allocatedHandles)

	deviceHandles := ([]loader.VkPhysicalDevice)(unsafe.Slice((*loader.VkPhysicalDevice)(allocatedHandles), int(*count)))
	res, err = i.loader.VkEnumeratePhysicalDevices(i.handle, count, (*loader.VkPhysicalDevice)(allocatedHandles))
	if err != nil {
		return nil, res, err
	}

	goCount := uint32(*count)
	var devices []*PhysicalDevice
	for ind := uint32(0); ind < goCount; ind++ {
		devices = append(devices, &PhysicalDevice{loader: i.loader, handle: deviceHandles[ind]})
	}

	return devices, res, nil
}