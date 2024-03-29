package core1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

func TestVulkanDescriptorSet_Free(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)

	device := mocks.EasyMockDevice(ctrl, mockDriver)
	pool := mocks.EasyMockDescriptorPool(ctrl, device)

	set := internal_mocks.EasyDummyDescriptorSet(mockDriver, pool)

	mockDriver.EXPECT().VkFreeDescriptorSets(device.Handle(), pool.Handle(), driver.Uint32(1), gomock.Not(unsafe.Pointer(nil))).DoAndReturn(
		func(device driver.VkDevice, descriptorPool driver.VkDescriptorPool, descriptorSetCount driver.Uint32, pDescriptorSets *driver.VkDescriptorSet) (common.VkResult, error) {
			descriptorSetSlice := unsafe.Slice(pDescriptorSets, 1)
			require.Equal(t, set.Handle(), descriptorSetSlice[0])

			return core1_0.VKSuccess, nil
		})

	_, err := set.Free()
	require.NoError(t, err)
}
