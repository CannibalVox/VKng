package ext_descriptor_indexing

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	ext_descriptor_indexing_driver "github.com/CannibalVox/VKng/extensions/ext_descriptor_indexing/driver"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_maintenance3"
	khr_maintenance3_driver "github.com/CannibalVox/VKng/extensions/khr_maintenance3/driver"
	mock_maintenance3 "github.com/CannibalVox/VKng/extensions/khr_maintenance3/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestDescriptorSetLayoutBindingFlagsCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockDescriptorSetLayout := mocks.EasyMockDescriptorSetLayout(ctrl)

	coreDriver.EXPECT().VkCreateDescriptorSetLayout(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pAllocator *driver.VkAllocationCallbacks, pSetLayout *driver.VkDescriptorSetLayout) (common.VkResult, error) {
		*pSetLayout = mockDescriptorSetLayout.Handle()
		val := reflect.ValueOf(pCreateInfo).Elem()

		require.Equal(t, uint64(32), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO

		next := (*ext_descriptor_indexing_driver.VkDescriptorSetLayoutBindingFlagsCreateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("bindingCount").Uint())
		flagsPtr := (*ext_descriptor_indexing_driver.VkDescriptorBindingFlagsEXT)(val.FieldByName("pBindingFlags").UnsafePointer())
		flagSlice := unsafe.Slice(flagsPtr, 2)

		require.Equal(t, []ext_descriptor_indexing_driver.VkDescriptorBindingFlagsEXT{8, 1}, flagSlice)

		return core1_0.VKSuccess, nil
	})

	descriptorSetLayout, _, err := device.CreateDescriptorSetLayout(
		nil,
		core1_0.DescriptorSetLayoutCreateOptions{
			NextOptions: common.NextOptions{
				DescriptorSetLayoutBindingFlagsCreateOptions{
					BindingFlags: []DescriptorBindingFlags{
						DescriptorBindingVariableDescriptorCount,
						DescriptorBindingUpdateAfterBind,
					},
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDescriptorSetLayout.Handle(), descriptorSetLayout.Handle())
}

func TestDescriptorSetVariableDescriptorCountAllocateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	descriptorPool := mocks.EasyMockDescriptorPool(ctrl, device)
	descriptorLayout1 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout2 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout3 := mocks.EasyMockDescriptorSetLayout(ctrl)
	descriptorLayout4 := mocks.EasyMockDescriptorSetLayout(ctrl)

	mockDescriptorSet1 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet2 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet3 := mocks.EasyMockDescriptorSet(ctrl)
	mockDescriptorSet4 := mocks.EasyMockDescriptorSet(ctrl)

	coreDriver.EXPECT().VkAllocateDescriptorSets(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pAllocateInfo *driver.VkDescriptorSetAllocateInfo,
		pDescriptorSets *driver.VkDescriptorSet) (common.VkResult, error) {

		sets := unsafe.Slice(pDescriptorSets, 4)
		sets[0] = mockDescriptorSet1.Handle()
		sets[1] = mockDescriptorSet2.Handle()
		sets[2] = mockDescriptorSet3.Handle()
		sets[3] = mockDescriptorSet4.Handle()

		val := reflect.ValueOf(pAllocateInfo).Elem()
		require.Equal(t, uint64(34), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_ALLOCATE_INFO

		next := (*ext_descriptor_indexing_driver.VkDescriptorSetVariableDescriptorCountAllocateInfoEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_ALLOCATE_INFO_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(4), val.FieldByName("descriptorSetCount").Uint())

		countsPtr := (*driver.Uint32)(val.FieldByName("pDescriptorCounts").UnsafePointer())
		countSlice := unsafe.Slice(countsPtr, 4)

		require.Equal(t, []driver.Uint32{1, 3, 5, 7}, countSlice)

		return core1_0.VKSuccess, nil
	})

	sets, _, err := device.AllocateDescriptorSets(core1_0.DescriptorSetAllocateOptions{
		DescriptorPool: descriptorPool,
		AllocationLayouts: []core1_0.DescriptorSetLayout{
			descriptorLayout1,
			descriptorLayout2,
			descriptorLayout3,
			descriptorLayout4,
		},
		NextOptions: common.NextOptions{
			DescriptorSetVariableDescriptorCountAllocateOptions{
				DescriptorCounts: []int{1, 3, 5, 7},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, sets, 4)
	require.Equal(t, []driver.VkDescriptorSet{
		mockDescriptorSet1.Handle(),
		mockDescriptorSet2.Handle(),
		mockDescriptorSet3.Handle(),
		mockDescriptorSet4.Handle(),
	}, []driver.VkDescriptorSet{
		sets[0].Handle(),
		sets[1].Handle(),
		sets[2].Handle(),
		sets[3].Handle(),
	})
}

func TestDescriptorSetVariableDescriptorCountLayoutSupportOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance3.NewMockDriver(ctrl)
	extension := khr_maintenance3.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)

	extDriver.EXPECT().VkGetDescriptorSetLayoutSupportKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo,
		pSupport *khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR) {
		val := reflect.ValueOf(pSupport).Elem()

		require.Equal(t, uint64(1000168001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT
		next := (*ext_descriptor_indexing_driver.VkDescriptorSetVariableDescriptorCountLayoutSupportEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000161004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxVariableDescriptorCount").UnsafeAddr())) = driver.Uint32(7)
	})

	var outData DescriptorSetVariableDescriptorCountLayoutSupportOutData
	err := extension.DescriptorSetLayoutSupport(
		device,
		core1_0.DescriptorSetLayoutCreateOptions{},
		&khr_maintenance3.DescriptorSetLayoutSupportOutData{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, DescriptorSetVariableDescriptorCountLayoutSupportOutData{
		MaxVariableDescriptorCount: 7,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)

	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := extensions.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pCreateInfo *driver.VkDeviceCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pDevice *driver.VkDevice) (common.VkResult, error) {

			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderSampledImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderStorageImageArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(0), val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingUpdateUnusedWhilePending").Uint())
			require.Equal(t, uint64(0), val.FieldByName("descriptorBindingPartiallyBound").Uint())
			require.Equal(t, uint64(1), val.FieldByName("descriptorBindingVariableDescriptorCount").Uint())
			require.Equal(t, uint64(0), val.FieldByName("runtimeDescriptorArray").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(
		nil,
		core1_0.DeviceCreateOptions{
			QueueFamilies: []core1_0.DeviceQueueCreateOptions{
				{
					CreatedQueuePriorities: []float32{0},
				},
			},
			NextOptions: common.NextOptions{PhysicalDeviceDescriptorIndexingFeatures{
				ShaderInputAttachmentArrayDynamicIndexing:          true,
				ShaderUniformTexelBufferArrayDynamicIndexing:       false,
				ShaderStorageTexelBufferArrayDynamicIndexing:       true,
				ShaderUniformBufferArrayNonUniformIndexing:         false,
				ShaderSampledImageArrayNonUniformIndexing:          true,
				ShaderStorageBufferArrayNonUniformIndexing:         false,
				ShaderStorageImageArrayNonUniformIndexing:          true,
				ShaderInputAttachmentArrayNonUniformIndexing:       false,
				ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
				ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
				DescriptorBindingUniformBufferUpdateAfterBind:      true,
				DescriptorBindingSampledImageUpdateAfterBind:       false,
				DescriptorBindingStorageImageUpdateAfterBind:       true,
				DescriptorBindingStorageBufferUpdateAfterBind:      false,
				DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
				DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
				DescriptorBindingUpdateUnusedWhilePending:          true,
				DescriptorBindingPartiallyBound:                    false,
				DescriptorBindingVariableDescriptorCount:           true,
				RuntimeDescriptorArray:                             false,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceDescriptorIndexingFeaturesOutData(t *testing.T) {
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
		pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

		val := reflect.ValueOf(pFeatures).Elem()
		require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

		next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageTexelBufferArrayNonUniformIndexing").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformBufferUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingSampledImageUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageImageUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageBufferUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUniformTexelBufferUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingStorageTexelBufferUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingUpdateUnusedWhilePending").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingPartiallyBound").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("descriptorBindingVariableDescriptorCount").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("runtimeDescriptorArray").UnsafeAddr())) = driver.VkBool32(0)
	})

	var outData PhysicalDeviceDescriptorIndexingFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeatures{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceDescriptorIndexingFeatures{
		ShaderInputAttachmentArrayDynamicIndexing:          true,
		ShaderUniformTexelBufferArrayDynamicIndexing:       false,
		ShaderStorageTexelBufferArrayDynamicIndexing:       true,
		ShaderUniformBufferArrayNonUniformIndexing:         false,
		ShaderSampledImageArrayNonUniformIndexing:          true,
		ShaderStorageBufferArrayNonUniformIndexing:         false,
		ShaderStorageImageArrayNonUniformIndexing:          true,
		ShaderInputAttachmentArrayNonUniformIndexing:       false,
		ShaderUniformTexelBufferArrayNonUniformIndexing:    true,
		ShaderStorageTexelBufferArrayNonUniformIndexing:    false,
		DescriptorBindingUniformBufferUpdateAfterBind:      true,
		DescriptorBindingSampledImageUpdateAfterBind:       false,
		DescriptorBindingStorageImageUpdateAfterBind:       true,
		DescriptorBindingStorageBufferUpdateAfterBind:      false,
		DescriptorBindingUniformTexelBufferUpdateAfterBind: true,
		DescriptorBindingStorageTexelBufferUpdateAfterBind: false,
		DescriptorBindingUpdateUnusedWhilePending:          true,
		DescriptorBindingPartiallyBound:                    false,
		DescriptorBindingVariableDescriptorCount:           true,
		RuntimeDescriptorArray:                             false,
	}, outData)
}

func TestPhysicalDeviceDescriptorIndexingOutData(t *testing.T) {
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
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR
		next := (*ext_descriptor_indexing_driver.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(val.FieldByName("pNext").UnsafePointer())

		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000161002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxUpdateAfterBindDescriptorsInAllPools").UnsafeAddr())) = driver.Uint32(1)

		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderUniformBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampledImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageBufferArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderStorageImageArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderInputAttachmentArrayNonUniformIndexingNative").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("robustBufferAccessUpdateAfterBind").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("quadDivergentImplicitLod").UnsafeAddr())) = driver.VkBool32(1)

		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(5)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(7)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(11)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(13)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageDescriptorUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(17)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxPerStageUpdateAfterBindResources").UnsafeAddr())) = driver.Uint32(19)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSamplers").UnsafeAddr())) = driver.Uint32(23)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffers").UnsafeAddr())) = driver.Uint32(29)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindUniformBuffersDynamic").UnsafeAddr())) = driver.Uint32(31)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffers").UnsafeAddr())) = driver.Uint32(37)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageBuffersDynamic").UnsafeAddr())) = driver.Uint32(41)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindSampledImages").UnsafeAddr())) = driver.Uint32(43)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindStorageImages").UnsafeAddr())) = driver.Uint32(47)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("maxDescriptorSetUpdateAfterBindInputAttachments").UnsafeAddr())) = driver.Uint32(51)
	})

	var outData PhysicalDeviceDescriptorIndexingOutData
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.DevicePropertiesOutData{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t,
		PhysicalDeviceDescriptorIndexingOutData{
			MaxUpdateAfterBindDescriptorsInAllPools: 1,

			ShaderUniformBufferArrayNonUniformIndexingNative:   true,
			ShaderSampledImageArrayNonUniformIndexingNative:    false,
			ShaderStorageBufferArrayNonUniformIndexingNative:   true,
			ShaderStorageImageArrayNonUniformIndexingNative:    false,
			ShaderInputAttachmentArrayNonUniformIndexingNative: true,
			RobustBufferAccessUpdateAfterBind:                  false,
			QuadDivergentImplicitLod:                           true,

			MaxPerStageDescriptorUpdateAfterBindSamplers:         3,
			MaxPerStageDescriptorUpdateAfterBindUniformBuffers:   5,
			MaxPerStageDescriptorUpdateAfterBindStorageBuffers:   7,
			MaxPerStageDescriptorUpdateAfterBindSampledImages:    11,
			MaxPerStageDescriptorUpdateAfterBindStorageImages:    13,
			MaxPerStageDescriptorUpdateAfterBindInputAttachments: 17,
			MaxPerStageUpdateAfterBindResources:                  19,
			MaxDescriptorSetUpdateAfterBindSamplers:              23,
			MaxDescriptorSetUpdateAfterBindUniformBuffers:        29,
			MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic: 31,
			MaxDescriptorSetUpdateAfterBindStorageBuffers:        37,
			MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic: 41,
			MaxDescriptorSetUpdateAfterBindSampledImages:         43,
			MaxDescriptorSetUpdateAfterBindStorageImages:         47,
			MaxDescriptorSetUpdateAfterBindInputAttachments:      51,
		},
		outData)
}
