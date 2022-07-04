package khr_timeline_semaphore_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_timeline_semaphore"
	khr_timeline_semaphore_driver "github.com/CannibalVox/VKng/extensions/khr_timeline_semaphore/driver"
	mock_timeline_semaphore "github.com/CannibalVox/VKng/extensions/khr_timeline_semaphore/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestVulkanExtension_SemaphoreCounterValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_timeline_semaphore.NewMockDriver(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	semaphore := mocks.EasyMockSemaphore(ctrl)
	semaphore.EXPECT().DeviceHandle().Return(device.Handle()).AnyTimes()

	extDriver.EXPECT().VkGetSemaphoreCounterValueKHR(
		device.Handle(),
		semaphore.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		semaphore driver.VkSemaphore,
		pValue *driver.Uint64) (common.VkResult, error) {

		*pValue = driver.Uint64(37)
		return core1_0.VKSuccess, nil
	})

	value, _, err := extension.SemaphoreCounterValue(
		semaphore,
	)
	require.NoError(t, err)
	require.Equal(t, uint64(37), value)
}

func TestVulkanExtension_SignalSemaphore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_timeline_semaphore.NewMockDriver(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	semaphore := mocks.EasyMockSemaphore(ctrl)

	extDriver.EXPECT().VkSignalSemaphoreKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pSignalInfo *khr_timeline_semaphore_driver.VkSemaphoreSignalInfoKHR) (common.VkResult, error) {

		val := reflect.ValueOf(pSignalInfo).Elem()
		require.Equal(t, uint64(1000207005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, semaphore.Handle(), driver.VkSemaphore(val.FieldByName("semaphore").UnsafePointer()))
		require.Equal(t, uint64(13), val.FieldByName("value").Uint())

		return core1_0.VKSuccess, nil
	})

	_, err := extension.SignalSemaphore(
		device,
		khr_timeline_semaphore.SemaphoreSignalOptions{
			Semaphore: semaphore,
			Value:     uint64(13),
		})
	require.NoError(t, err)
}

func TestVulkanExtension_WaitSemaphores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_timeline_semaphore.NewMockDriver(ctrl)
	extension := khr_timeline_semaphore.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	semaphore1 := mocks.EasyMockSemaphore(ctrl)
	semaphore2 := mocks.EasyMockSemaphore(ctrl)

	extDriver.EXPECT().VkWaitSemaphoresKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		driver.Uint64(60000000000),
	).DoAndReturn(func(device driver.VkDevice,
		pWaitInfo *khr_timeline_semaphore_driver.VkSemaphoreWaitInfoKHR,
		timeout driver.Uint64) (common.VkResult, error) {

		val := reflect.ValueOf(pWaitInfo).Elem()
		require.Equal(t, uint64(1000207004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_SEMAPHORE_WAIT_ANY_BIT_KHR
		require.Equal(t, uint64(2), val.FieldByName("semaphoreCount").Uint())

		semaphorePtr := (*driver.VkSemaphore)(val.FieldByName("pSemaphores").UnsafePointer())
		semaphoreSlice := unsafe.Slice(semaphorePtr, 2)
		require.Equal(t, []driver.VkSemaphore{semaphore1.Handle(), semaphore2.Handle()}, semaphoreSlice)

		valuesPtr := (*driver.Uint64)(val.FieldByName("pValues").UnsafePointer())
		valuesSlice := unsafe.Slice(valuesPtr, 2)
		require.Equal(t, []driver.Uint64{13, 19}, valuesSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := extension.WaitSemaphores(
		device,
		time.Minute,
		khr_timeline_semaphore.SemaphoreWaitOptions{
			Flags: khr_timeline_semaphore.SemaphoreWaitAny,
			Semaphores: []core1_0.Semaphore{
				semaphore1,
				semaphore2,
			},
			Values: []uint64{
				13,
				19,
			},
		})
	require.NoError(t, err)
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOptions(t *testing.T) {
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
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice) (common.VkResult, error) {
		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("timelineSemaphore").Uint())

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
			NextOptions: common.NextOptions{
				khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures{
					TimelineSemaphore: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestPhysicalDeviceTimelineSemaphoreFeaturesOutData(t *testing.T) {
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

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("timelineSemaphore").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures
	err := extension.PhysicalDeviceFeatures2(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeatures{
			NextOutData: common.NextOutData{&outData},
		},
	)
	require.NoError(t, err)
	require.Equal(t, khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreFeatures{
		TimelineSemaphore: true,
	}, outData)
}

func TestSemaphoreTypeCreateOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockSemaphore := mocks.EasyMockSemaphore(ctrl)

	coreDriver.EXPECT().VkCreateSemaphore(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkSemaphoreCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pSemaphore *driver.VkSemaphore) (common.VkResult, error) {

		*pSemaphore = mockSemaphore.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(9), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO

		next := (*khr_timeline_semaphore_driver.VkSemaphoreTypeCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(1), val.FieldByName("semaphoreType").Uint()) // VK_SEMAPHORE_TYPE_TIMELINE_KHR
		require.Equal(t, uint64(13), val.FieldByName("initialValue").Uint())

		return core1_0.VKSuccess, nil
	})

	semaphore, _, err := device.CreateSemaphore(
		nil,
		core1_0.SemaphoreCreateOptions{
			NextOptions: common.NextOptions{khr_timeline_semaphore.SemaphoreTypeCreateOptions{
				SemaphoreType: khr_timeline_semaphore.SemaphoreTypeTimeline,
				InitialValue:  uint64(13),
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockSemaphore.Handle(), semaphore.Handle())
}

func TestTimelineSemaphoreSubmitOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	queue := extensions.CreateQueueObject(coreDriver, device.Handle(), mocks.NewFakeQueue(), common.Vulkan1_0)
	fence := mocks.EasyMockFence(ctrl)

	coreDriver.EXPECT().VkQueueSubmit(
		queue.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
		fence.Handle(),
	).DoAndReturn(func(queue driver.VkQueue,
		submitCount driver.Uint32,
		pSubmits *driver.VkSubmitInfo,
		fence driver.VkFence) (common.VkResult, error) {

		val := reflect.ValueOf(pSubmits).Elem()
		require.Equal(t, uint64(4), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SUBMIT_INFO

		next := (*khr_timeline_semaphore_driver.VkTimelineSemaphoreSubmitInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000207003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_TIMELINE_SEMAPHORE_SUBMIT_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("waitSemaphoreValueCount").Uint())
		require.Equal(t, uint64(3), val.FieldByName("signalSemaphoreValueCount").Uint())

		waitPtr := (*driver.Uint64)(val.FieldByName("pWaitSemaphoreValues").UnsafePointer())
		waitSlice := unsafe.Slice(waitPtr, 2)
		require.Equal(t, []driver.Uint64{3, 5}, waitSlice)

		signalPtr := (*driver.Uint64)(val.FieldByName("pSignalSemaphoreValues").UnsafePointer())
		signalSlice := unsafe.Slice(signalPtr, 3)
		require.Equal(t, []driver.Uint64{7, 11, 13}, signalSlice)

		return core1_0.VKSuccess, nil
	})

	_, err := queue.SubmitToQueue(
		fence,
		[]core1_0.SubmitOptions{
			{
				NextOptions: common.NextOptions{
					khr_timeline_semaphore.TimelineSemaphoreSubmitOptions{
						WaitSemaphoreValues:   []uint64{3, 5},
						SignalSemaphoreValues: []uint64{7, 11, 13},
					},
				},
			},
		})
	require.NoError(t, err)
}

func TestPhysicalDeviceTimelineSemaphoreOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_timeline_semaphore_driver.VkPhysicalDeviceTimelineSemaphorePropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000207001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		*(*driver.Uint64)(unsafe.Pointer(val.FieldByName("maxTimelineSemaphoreValueDifference").UnsafeAddr())) = driver.Uint64(3)
	})

	var outData khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreOutData
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.DevicePropertiesOutData{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_timeline_semaphore.PhysicalDeviceTimelineSemaphoreOutData{
		MaxTimelineSemaphoreValueDifference: 3,
	}, outData)
}
