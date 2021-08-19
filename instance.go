package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "unsafe"

type InstanceBuilder struct {
	applicationName    string
	applicationVersion uint32
	engineName         string
	engineVersion      uint32
}

func (b *InstanceBuilder) ApplicationName(name string) *InstanceBuilder {
	b.applicationName = name
	return b
}

func makeVersion(major, minor, patch uint32) uint32 {
	return (major << 22) | (minor << 12) | patch
}

func (b *InstanceBuilder) ApplicationVersion(major, minor, patch uint32) *InstanceBuilder {
	b.applicationVersion = makeVersion(major, minor, patch)
	return b
}

func (b *InstanceBuilder) EngineName(name string) *InstanceBuilder {
	b.engineName = name
	return b
}

func (b *InstanceBuilder) EngineVersion(major, minor, patch uint32) *InstanceBuilder {
	b.engineVersion = makeVersion(major, minor, patch)
	return b
}

func (b *InstanceBuilder) Build(allocator Allocator) (*Instance, error) {
	cApplication := allocator.CString(b.applicationName)
	cEngine := allocator.CString(b.engineName)
	appInfo := &C.VkApplicationInfo{
		pApplicationName:   (*C.char)(cApplication),
		applicationVersion: C.uint32_t(b.applicationVersion),
		pEngineName:        (*C.char)(cEngine),
		engineVersion:      C.uint32_t(b.engineVersion),
		apiVersion:         C.VK_API_VERSION_1_2,
	}

	createInfo := &C.VkInstanceCreateInfo{
		sType:            C.VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO,
		pApplicationInfo: appInfo,
	}

	var instanceHandle C.VkInstance

	res := C.vkCreateInstance(createInfo, nil, &instanceHandle)

	err := VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Instance{
		handle: instanceHandle,
	}, nil
}

type Instance struct {
	handle C.VkInstance
}

func (i *Instance) Destroy() {
	C.vkDestroyInstance(i.handle, nil)
}

func (i *Instance) PhysicalDevices(allocator Allocator) ([]*PhysicalDevice, error) {
	count := C.uint32_t(0)
	res := C.vkEnumeratePhysicalDevices(i.handle, &count, nil)
	err := VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	allocatedHandles := allocator.Malloc(uint(count)*uint(unsafe.Sizeof([1]C.VkPhysicalDevice{})))
	defer allocator.Free(allocatedHandles)
	deviceHandles := (*[1<<30]C.VkPhysicalDevice)(allocatedHandles)
	res = C.vkEnumeratePhysicalDevices(i.handle, &count, (*C.VkPhysicalDevice)(allocatedHandles))
	err = VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	goCount := uint32(count)
	var devices []*PhysicalDevice
	for i := uint32(0); i < goCount; i++ {
		devices = append(devices, &PhysicalDevice{handle: deviceHandles[i]})
	}

	return devices, nil
}
