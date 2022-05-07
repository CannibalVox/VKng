package khr_multiview_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_multiview"
	khr_multiview_driver "github.com/CannibalVox/VKng/extensions/khr_multiview/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestMultiviewFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR,
	) {
		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*khr_multiview_driver.VkPhysicalDeviceMultiviewFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiview").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewGeometryShader").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multiviewTessellationShader").UnsafeAddr())) = driver.VkBool32(0)
	})

	var outData khr_multiview.MultiviewFeaturesOutData
	features := khr_get_physical_device_properties2.DeviceFeaturesOutData{
		HaveNext: common.HaveNext{&outData},
	}

	err := extension.PhysicalDeviceFeatures(physicalDevice, &features)
	require.NoError(t, err)
	require.Equal(t, khr_multiview.MultiviewFeaturesOutData{
		Multiview:                   true,
		MultiviewTessellationShader: false,
		MultiviewGeometryShader:     true,
	}, outData)
}

func TestMultiviewFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice,
	) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("queueCreateInfoCount").Uint())

		queueCreate := (*driver.VkDeviceQueueCreateInfo)(val.FieldByName("pQueueCreateInfos").UnsafePointer())

		queueFamilyVal := reflect.ValueOf(queueCreate).Elem()
		require.Equal(t, uint64(2), queueFamilyVal.FieldByName("sType").Uint()) //VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
		require.True(t, queueFamilyVal.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), queueFamilyVal.FieldByName("queueCount").Uint())

		next := (*khr_multiview_driver.VkPhysicalDeviceMultiviewFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("multiview").Uint())
		require.Equal(t, uint64(0), val.FieldByName("multiviewGeometryShader").Uint())
		require.Equal(t, uint64(1), val.FieldByName("multiviewTessellationShader").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := loader.CreateDevice(physicalDevice, nil, core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueOptions{
			{
				CreatedQueuePriorities: []float32{3, 2, 1},
			},
		},
		HaveNext: common.HaveNext{
			khr_multiview.MultiviewFeaturesOptions{
				Multiview:                   true,
				MultiviewTessellationShader: true,
				MultiviewGeometryShader:     false,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestMultiviewPropertiesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR,
	) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_multiview_driver.VkPhysicalDeviceMultiviewPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTIVIEW_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewViewCount").UnsafeAddr())) = uint32(5)
		*(*uint32)(unsafe.Pointer(val.FieldByName("maxMultiviewInstanceIndex").UnsafeAddr())) = uint32(3)
	})

	var outData khr_multiview.MultiviewPropertiesOutData
	properties := khr_get_physical_device_properties2.DevicePropertiesOutData{
		HaveNext: common.HaveNext{&outData},
	}

	err := extension.PhysicalDeviceProperties(physicalDevice, &properties)
	require.NoError(t, err)
	require.Equal(t, khr_multiview.MultiviewPropertiesOutData{
		MaxMultiviewInstanceIndex: 3,
		MaxMultiviewViewCount:     5,
	}, outData)
}

func TestRenderPassMultiviewOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)
	mockRenderPass := mocks.EasyMockRenderPass(ctrl)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateRenderPass(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkRenderPassCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pRenderPass *driver.VkRenderPass) (common.VkResult, error) {

		*pRenderPass = mockRenderPass.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(38), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO

		next := (*khr_multiview_driver.VkRenderPassMultiviewCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000053000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_RENDER_PASS_MULTIVIEW_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(3), val.FieldByName("subpassCount").Uint())
		require.Equal(t, uint64(2), val.FieldByName("dependencyCount").Uint())
		require.Equal(t, uint64(1), val.FieldByName("correlationMaskCount").Uint())

		masks := (*driver.Uint32)(val.FieldByName("pViewMasks").UnsafePointer())
		maskSlice := ([]driver.Uint32)(unsafe.Slice(masks, 3))
		require.Equal(t, []driver.Uint32{1, 2, 7}, maskSlice)

		offsets := (*driver.Int32)(val.FieldByName("pViewOffsets").UnsafePointer())
		offsetSlice := ([]driver.Int32)(unsafe.Slice(offsets, 2))
		require.Equal(t, []driver.Int32{11, 13}, offsetSlice)

		correlationMasks := (*driver.Uint32)(val.FieldByName("pCorrelationMasks").UnsafePointer())
		correlationSlice := ([]driver.Uint32)(unsafe.Slice(correlationMasks, 1))
		require.Equal(t, []driver.Uint32{17}, correlationSlice)

		return core1_0.VKSuccess, nil
	})

	renderPass, _, err := loader.CreateRenderPass(device, nil, core1_0.RenderPassCreateOptions{
		HaveNext: common.HaveNext{
			khr_multiview.RenderPassMultiviewOptions{
				SubpassViewMasks:      []uint32{1, 2, 7},
				DependencyViewOffsets: []int{11, 13},
				CorrelationMasks:      []uint32{17},
			},
		},
	})

	require.NoError(t, err)
	require.Equal(t, mockRenderPass.Handle(), renderPass.Handle())

}
