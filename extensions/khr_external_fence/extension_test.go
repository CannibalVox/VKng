package khr_external_fence_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_external_fence"
	khr_external_fence_driver "github.com/CannibalVox/VKng/extensions/khr_external_fence/driver"
	"github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestExportFenceOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := core.CreateDevice(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	mockFence := mocks.EasyMockFence(ctrl)

	coreDriver.EXPECT().VkCreateFence(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice, pCreateInfo *driver.VkFenceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pFence *driver.VkFence) (common.VkResult, error) {
		*pFence = mockFence.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(8), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FENCE_CREATE_INFO
		require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_FENCE_CREATE_SIGNALED_BIT

		next := (*khr_external_fence_driver.VkExportFenceCreateInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000113000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(2), val.FieldByName("handleTypes").Uint()) // VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	fence, _, err := device.CreateFence(
		nil,
		core1_0.FenceCreateOptions{
			Flags: core1_0.FenceCreateSignaled,

			HaveNext: common.HaveNext{
				khr_external_fence.ExportFenceOptions{
					HandleTypes: khr_external_fence_capabilities.ExternalFenceHandleTypeOpaqueWin32,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockFence.Handle(), fence.Handle())
}
