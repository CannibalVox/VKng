package khr_get_memory_requirements2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2"
	khr_get_memory_requirements2_driver "github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2/driver"
	"github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_BufferMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	buffer := mocks.EasyMockBuffer(ctrl)

	extDriver.EXPECT().VkGetBufferMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *khr_get_memory_requirements2_driver.VkBufferMemoryRequirementsInfo2KHR,
		pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BUFFER_MEMORY_REQUIREMENTS_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, buffer.Handle(), driver.VkBuffer(val.FieldByName("buffer").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)
	})

	var outData khr_get_memory_requirements2.MemoryRequirementsOutData
	err := extension.BufferMemoryRequirements(device,
		khr_get_memory_requirements2.BufferRequirementsOptions{
			Buffer: buffer,
		}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryType)
}

func TestVulkanExtension_ImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)

	extDriver.EXPECT().VkGetImageMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR,
		pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		val = val.FieldByName("memoryRequirements")
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("size").UnsafeAddr())) = driver.VkDeviceSize(1)
		*(*driver.VkDeviceSize)(unsafe.Pointer(val.FieldByName("alignment").UnsafeAddr())) = driver.VkDeviceSize(3)
		*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("memoryTypeBits").UnsafeAddr())) = driver.Uint32(5)
	})

	var outData khr_get_memory_requirements2.MemoryRequirementsOutData
	err := extension.ImageMemoryRequirements(device,
		khr_get_memory_requirements2.ImageMemoryRequirementsOptions{
			Image: image,
		}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.MemoryRequirements.Size)
	require.Equal(t, 3, outData.MemoryRequirements.Alignment)
	require.Equal(t, uint32(5), outData.MemoryRequirements.MemoryType)
}

func TestVulkanExtension_SparseImageMemoryRequirements(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	image := mocks.EasyMockImage(ctrl)

	extDriver.EXPECT().VkGetImageSparseMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR) {

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (driver.VkImage)(options.FieldByName("image").UnsafePointer()))

			*pSparseMemoryRequirementCount = driver.Uint32(2)
		})

	extDriver.EXPECT().VkGetImageSparseMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pInfo *khr_get_memory_requirements2_driver.VkImageSparseMemoryRequirementsInfo2KHR,
			pSparseMemoryRequirementCount *driver.Uint32,
			pSparseMemoryRequirements *khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR) {

			require.Equal(t, driver.Uint32(2), *pSparseMemoryRequirementCount)

			options := reflect.ValueOf(pInfo).Elem()
			require.Equal(t, uint64(1000146002), options.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_SPARSE_MEMORY_REQUIREMENTS_INFO_2_KHR
			require.True(t, options.FieldByName("pNext").IsNil())
			require.Equal(t, image.Handle(), (driver.VkImage)(options.FieldByName("image").UnsafePointer()))

			requirementSlice := ([]khr_get_memory_requirements2_driver.VkSparseImageMemoryRequirements2KHR)(unsafe.Slice(pSparseMemoryRequirements, 2))
			outData := reflect.ValueOf(requirementSlice)
			element := outData.Index(0)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs := element.FieldByName("memoryRequirements")
			imageAspectFlags := (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000008) // VK_IMAGE_ASPECT_METADATA_BIT
			width := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(1)
			height := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(3)
			depth := (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(5)
			flags := (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0x00000004) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(7)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(17)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(11)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(13)

			element = outData.Index(1)
			require.Equal(t, uint64(1000146004), element.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_MEMORY_REQUIREMENTS_2_KHR
			require.True(t, element.FieldByName("pNext").IsNil())

			memReqs = element.FieldByName("memoryRequirements")
			imageAspectFlags = (*driver.VkImageAspectFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("aspectMask").UnsafeAddr()))
			*imageAspectFlags = driver.VkImageAspectFlags(0x00000004) // VK_IMAGE_ASPECT_STENCIL_BIT
			width = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("width").UnsafeAddr()))
			*width = driver.Uint32(19)
			height = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("height").UnsafeAddr()))
			*height = driver.Uint32(23)
			depth = (*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr()))
			*depth = driver.Uint32(29)
			flags = (*driver.VkSparseImageFormatFlags)(unsafe.Pointer(memReqs.FieldByName("formatProperties").FieldByName("flags").UnsafeAddr()))
			*flags = driver.VkSparseImageFormatFlags(0)
			*(*driver.Uint32)(unsafe.Pointer(memReqs.FieldByName("imageMipTailFirstLod").UnsafeAddr())) = driver.Uint32(43)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailSize").UnsafeAddr())) = driver.VkDeviceSize(31)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailOffset").UnsafeAddr())) = driver.VkDeviceSize(41)
			*(*driver.VkDeviceSize)(unsafe.Pointer(memReqs.FieldByName("imageMipTailStride").UnsafeAddr())) = driver.VkDeviceSize(37)
		})

	outData, err := extension.SparseImageMemoryRequirements(device,
		khr_get_memory_requirements2.SparseImageRequirementsOptions{
			Image: image,
		}, nil)
	require.NoError(t, err)
	require.Equal(t, []*khr_get_memory_requirements2.SparseImageRequirementsOutData{
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.AspectMetadata,
					ImageGranularity: common.Extent3D{
						Width:  1,
						Height: 3,
						Depth:  5,
					},
					Flags: core1_0.SparseImageFormatNonstandardBlockSize,
				},
				ImageMipTailFirstLod: 7,
				ImageMipTailOffset:   11,
				ImageMipTailStride:   13,
				ImageMipTailSize:     17,
			},
		},
		{
			MemoryRequirements: core1_0.SparseImageMemoryRequirements{
				FormatProperties: core1_0.SparseImageFormatProperties{
					AspectMask: core1_0.AspectStencil,
					ImageGranularity: common.Extent3D{
						Width:  19,
						Height: 23,
						Depth:  29,
					},
					Flags: 0,
				},
				ImageMipTailSize:     31,
				ImageMipTailStride:   37,
				ImageMipTailOffset:   41,
				ImageMipTailFirstLod: 43,
			},
		},
	}, outData)
}
