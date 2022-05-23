package khr_device_group_creation_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_device_group_creation"
	khr_device_group_creation_driver "github.com/CannibalVox/VKng/extensions/khr_device_group_creation/driver"
	mock_device_group_creation "github.com/CannibalVox/VKng/extensions/khr_device_group_creation/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_EnumeratePhysicalDeviceGroups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group_creation.NewMockDriver(ctrl)
	extension := khr_device_group_creation.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)

	physicalDevice1 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice2 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice4 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice5 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice6 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	groups, _, err := extension.PhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}

func TestVulkanExtension_EnumeratePhysicalDeviceGroups_Incomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_device_group_creation.NewMockDriver(ctrl)
	extension := khr_device_group_creation.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)

	physicalDevice1 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice2 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice4 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice5 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice6 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = driver.Uint32(2)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(2), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		return core1_0.VKIncomplete, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		*pCount = driver.Uint32(3)

		return core1_0.VKSuccess, nil
	})

	extDriver.EXPECT().VkEnumeratePhysicalDeviceGroupsKHR(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(instance driver.VkInstance, pCount *driver.Uint32, pProperties *khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
		require.Equal(t, driver.Uint32(3), *pCount)

		propertySlice := ([]khr_device_group_creation_driver.VkPhysicalDeviceGroupPropertiesKHR)(unsafe.Slice(pProperties, 3))
		props := reflect.ValueOf(propertySlice)
		prop := props.Index(0)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(1)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices := prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice1.Handle()

		prop = props.Index(1)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(2)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(0)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice2.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice3.Handle()

		prop = props.Index(2)
		require.Equal(t, uint64(1000070000), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_GROUP_PROPERTIES_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(prop.FieldByName("physicalDeviceCount").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.VkBool32)(unsafe.Pointer(prop.FieldByName("subsetAllocation").UnsafeAddr())) = driver.VkBool32(1)

		propDevices = prop.FieldByName("physicalDevices")
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(0).UnsafeAddr())) = physicalDevice4.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(1).UnsafeAddr())) = physicalDevice5.Handle()
		*(*driver.VkPhysicalDevice)(unsafe.Pointer(propDevices.Index(2).UnsafeAddr())) = physicalDevice6.Handle()

		return core1_0.VKSuccess, nil
	})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice1.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice2.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice3.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		}).Times(2)

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice4.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice5.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	coreDriver.EXPECT().VkGetPhysicalDeviceProperties(physicalDevice6.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(device driver.VkPhysicalDevice, pProperties *driver.VkPhysicalDeviceProperties) {
			value := reflect.ValueOf(pProperties).Elem()
			*(*driver.Uint32)(unsafe.Pointer(value.FieldByName("apiVersion").UnsafeAddr())) = driver.Uint32(common.Vulkan1_0)
		})

	groups, _, err := extension.PhysicalDeviceGroups(instance, nil)
	require.NoError(t, err)
	require.Len(t, groups, 3)
	require.True(t, groups[0].SubsetAllocation)
	require.Len(t, groups[0].PhysicalDevices, 1)
	require.Equal(t, physicalDevice1.Handle(), groups[0].PhysicalDevices[0].Handle())

	require.False(t, groups[1].SubsetAllocation)
	require.Len(t, groups[1].PhysicalDevices, 2)
	require.Equal(t, physicalDevice2.Handle(), groups[1].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice3.Handle(), groups[1].PhysicalDevices[1].Handle())

	require.True(t, groups[2].SubsetAllocation)
	require.Len(t, groups[2].PhysicalDevices, 3)
	require.Equal(t, physicalDevice4.Handle(), groups[2].PhysicalDevices[0].Handle())
	require.Equal(t, physicalDevice5.Handle(), groups[2].PhysicalDevices[1].Handle())
	require.Equal(t, physicalDevice6.Handle(), groups[2].PhysicalDevices[2].Handle())
}

func TestDeviceGroupOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice1 := core.CreatePhysicalDevice(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)

	physicalDevice2 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	physicalDevice3 := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	handle := mocks.NewFakeDeviceHandle()

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice1.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
		*pDevice = handle

		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		optionsPtr := (*khr_device_group_creation_driver.VkDeviceGroupDeviceCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		options := reflect.ValueOf(optionsPtr).Elem()

		require.Equal(t, uint64(1000070001), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_GROUP_DEVICE_CREATE_INFO_KHR
		require.True(t, options.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), options.FieldByName("physicalDeviceCount").Uint())

		devicePtr := (*driver.VkPhysicalDevice)(options.FieldByName("pPhysicalDevices").UnsafePointer())
		deviceSlice := ([]driver.VkPhysicalDevice)(unsafe.Slice(devicePtr, 3))
		require.Equal(t, physicalDevice1.Handle(), deviceSlice[0])
		require.Equal(t, physicalDevice2.Handle(), deviceSlice[1])
		require.Equal(t, physicalDevice3.Handle(), deviceSlice[2])

		return core1_0.VKSuccess, nil
	})

	device, _, err := physicalDevice1.CreateDevice(nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueCreateOptions{
			{
				CreatedQueuePriorities: []float32{0},
			},
		},
		HaveNext: common.HaveNext{Next: khr_device_group_creation.DeviceGroupDeviceOptions{
			PhysicalDevices: []core1_0.PhysicalDevice{physicalDevice1, physicalDevice2, physicalDevice3},
		}},
	})
	require.NoError(t, err)
	require.Equal(t, handle, device.Handle())
}
