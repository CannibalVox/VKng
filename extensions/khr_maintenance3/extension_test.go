package khr_maintenance3_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
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

func TestVulkanExtension_DescriptorSetLayoutSupport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_maintenance3.NewMockDriver(ctrl)
	extension := khr_maintenance3.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetDescriptorSetLayoutSupportKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil())).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkDescriptorSetLayoutCreateInfo, pSupport *khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR) {
			optionVal := reflect.ValueOf(pCreateInfo).Elem()

			require.Equal(t, uint64(32), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), optionVal.FieldByName("bindingCount").Uint())

			bindingPtr := (*driver.VkDescriptorSetLayoutBinding)(optionVal.FieldByName("pBindings").UnsafePointer())
			binding := reflect.ValueOf(bindingPtr).Elem()
			require.Equal(t, uint64(1), binding.FieldByName("binding").Uint())
			require.Equal(t, uint64(3), binding.FieldByName("descriptorCount").Uint())

			outDataVal := reflect.ValueOf(pSupport).Elem()

			require.Equal(t, uint64(1000168001), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT_KHR
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			*(*driver.VkBool32)(unsafe.Pointer(outDataVal.FieldByName("supported").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := &khr_maintenance3.DescriptorSetLayoutSupportOutData{}
	err := extension.DescriptorSetLayoutSupport(device, core1_0.DescriptorSetLayoutOptions{
		Bindings: []core1_0.DescriptorLayoutBinding{
			{
				Binding:         1,
				DescriptorCount: 3,
			},
		},
	}, outData)
	require.NoError(t, err)
	require.True(t, outData.Supported)
}

func TestMaintenance3OutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			props := val.FieldByName("properties")
			*(*driver.Uint32)(unsafe.Pointer(props.FieldByName("vendorID").UnsafeAddr())) = driver.Uint32(3)

			maintPtr := (*khr_maintenance3_driver.VkPhysicalDeviceMaintenance3PropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			maint := reflect.ValueOf(maintPtr).Elem()

			require.Equal(t, uint64(1000168000), maint.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR
			require.True(t, maint.FieldByName("pNext").IsNil())

			*(*driver.Uint32)(unsafe.Pointer(maint.FieldByName("maxPerSetDescriptors").UnsafeAddr())) = driver.Uint32(5)
			*(*driver.Uint64)(unsafe.Pointer(maint.FieldByName("maxMemoryAllocationSize").UnsafeAddr())) = driver.Uint64(7)
		})

	maintOutData := &khr_maintenance3.Maintenance3OutData{}
	outData := &khr_get_physical_device_properties2.DevicePropertiesOutData{
		HaveNext: common.HaveNext{Next: maintOutData},
	}
	err := extension.PhysicalDeviceProperties(physicalDevice, outData)
	require.NoError(t, err)

	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, 5, maintOutData.MaxPerSetDescriptors)
	require.Equal(t, 7, maintOutData.MaxMemoryAllocationSize)
}
