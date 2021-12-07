package main

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/examples/lunarg_samples/utils"
	"github.com/CannibalVox/VKng/extensions/ext_debug_utils"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"runtime/debug"
	"time"
	"unsafe"
)

func logDebug(msgType ext_debug_utils.MessageType, severity ext_debug_utils.MessageSeverity, data *ext_debug_utils.CallbackData) bool {
	log.Printf("[%s %s] - %s", severity, msgType, data.Message)
	debug.PrintStack()
	log.Println()
	return false
}

func main() {
	info := &utils.SampleInfo{}
	err := info.ProcessCommandLineArgs()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitWindowSize(640, 640)
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitWindow()
	if err != nil {
		log.Fatalln(err)
	}

	info.Loader, err = core.CreateLoaderFromProcAddr(sdl.VulkanGetVkGetInstanceProcAddr())
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitGlobalLayerProperties()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitInstanceExtensionNames()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitDeviceExtensionNames()
	if err != nil {
		log.Fatalln(err)
	}

	info.InstanceExtensionNames = append(info.InstanceExtensionNames, ext_debug_utils.ExtensionName)
	info.InstanceLayerNames = append(info.InstanceLayerNames, "VK_LAYER_KHRONOS_validation")
	debugOptions := &ext_debug_utils.CreationOptions{
		CaptureSeverities: ext_debug_utils.SeverityWarning | ext_debug_utils.SeverityError,
		CaptureTypes:      ext_debug_utils.TypeAll,
		Callback:          logDebug,
	}

	err = info.InitInstance("Copy/Blit Image", debugOptions)
	if err != nil {
		log.Fatalln(err)
	}

	debugLoader := ext_debug_utils.CreateLoaderFromInstance(info.Instance)
	debugMessenger, _, err := debugLoader.CreateMessenger(info.Instance, debugOptions)
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitEnumerateDevice()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitSwapchainExtension()
	if err != nil {
		log.Fatalln(err)
	}

	surfCapabilities, _, err := info.Surface.Capabilities(info.Gpus[0])
	if err != nil {
		log.Fatalln(err)
	}

	if (surfCapabilities.SupportedImageUsage & common.ImageUsageTransferDst) == 0 {
		log.Fatalln("Surface cannot be destination of blit - abort")
	}

	err = info.InitDevice()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitCommandPool()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitCommandBuffer()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.ExecuteBeginCommandBuffer()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitDeviceQueue()
	if err != nil {
		log.Fatalln(err)
	}

	err = info.InitSwapchain(common.ImageUsageColorAttachment | common.ImageUsageTransferDst)
	if err != nil {
		log.Fatalln(err)
	}

	/* VULKAN_KEY_START */
	formatProps := info.Gpus[0].FormatProperties(info.Format)
	if (formatProps.LinearTilingFeatures & common.FormatFeatureBlitSource) == 0 {
		log.Fatalln("FOrmat cannot be used as transfer source")
	}

	imageAcquiredSemaphore, _, err := info.Loader.CreateSemaphore(info.Device, &core.SemaphoreOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	// Get the index of the next available swapchain image:
	info.CurrentBuffer, _, err = info.Swapchain.AcquireNextImage(common.NoTimeout, imageAcquiredSemaphore, nil)
	// TODO: Deal with the VK_SUBOPTIMAL_KHR and VK_ERROR_OUT_OF_DATE_KHR
	// return codes
	if err != nil {
		log.Fatalln(err)
	}

	// We'll be blitting into the presentable image, set the layout accordingly
	err = info.SetImageLayout(info.Buffers[info.CurrentBuffer].Image, common.AspectColor, common.LayoutUndefined, common.LayoutTransferDstOptimal, common.PipelineStageTopOfPipe, common.PipelineStageTransfer)
	if err != nil {
		log.Fatalln(err)
	}

	// Create an image, map it, and write some values to the image
	bltSrcImage, _, err := info.Loader.CreateImage(info.Device, &core.ImageOptions{
		ImageType:     common.ImageType2D,
		Format:        info.Format,
		Extent:        common.Extent3D{Width: info.Width, Height: info.Height, Depth: 1},
		MipLevels:     1,
		ArrayLayers:   1,
		Samples:       utils.NumSamples,
		SharingMode:   common.SharingExclusive,
		Usage:         common.ImageUsageTransferSrc,
		Tiling:        common.ImageTilingLinear,
		InitialLayout: common.LayoutUndefined,
	})
	if err != nil {
		log.Fatalln(err)
	}

	memReq := bltSrcImage.MemoryRequirements()
	memoryIndex, err := info.MemoryTypeFromProperties(memReq.MemoryType, core.MemoryHostVisible)
	if err != nil {
		log.Fatalln(err)
	}

	dmem, _, err := info.Device.AllocateMemory(&core.DeviceMemoryOptions{
		AllocationSize:  memReq.Size,
		MemoryTypeIndex: memoryIndex,
	})
	if err != nil {
		log.Fatalln(err)
	}
	_, err = bltSrcImage.BindImageMemory(dmem, 0)
	if err != nil {
		log.Fatalln(err)
	}

	err = info.SetImageLayout(bltSrcImage, common.AspectColor, common.LayoutUndefined, common.LayoutGeneral, common.PipelineStageTopOfPipe, common.PipelineStageHost)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = info.Cmd.End()
	if err != nil {
		log.Fatalln(err)
	}

	cmdFence, err := info.InitFence()
	if err != nil {
		log.Fatalln(err)
	}

	/* Queue the command buffer for execution */
	_, err = info.GraphicsQueue.SubmitToQueue(cmdFence, []*core.SubmitOptions{
		{
			WaitDstStages:  []common.PipelineStages{common.PipelineStageColorAttachmentOutput},
			WaitSemaphores: []core.Semaphore{imageAcquiredSemaphore},
			CommandBuffers: []core.CommandBuffer{info.Cmd},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	/* Make sure command buffer is finished before mapping */
	for {
		res, err := cmdFence.Wait(common.NoTimeout)
		if err != nil {
			log.Fatalln(err)
		}

		if res != core.VKTimeout {
			break
		}
	}
	cmdFence.Destroy()

	pImgMem, _, err := dmem.MapMemory(0, memReq.Size, 0)
	if err != nil {
		log.Fatalln(err)
	}

	imgBytes := ([]byte)(unsafe.Slice((*byte)(pImgMem), info.Height*info.Width*4))
	imgByteIndex := 0

	for row := 0; row < info.Height; row++ {
		for col := 0; col < info.Width; col++ {
			setVal := ((row & 0x8) ^ (col & 0x8)) >> 3
			rgb := byte(setVal * 255)
			imgBytes[imgByteIndex] = rgb
			imgBytes[imgByteIndex+1] = rgb
			imgBytes[imgByteIndex+2] = rgb
			imgBytes[imgByteIndex+3] = 255
			imgByteIndex += 4
		}
	}

	// Flush the mapped memory and then unmap it  Assume it isn't coherent since
	// we didn't really confirm
	_, err = info.Device.FlushMappedMemoryRanges([]*core.MappedMemoryRange{
		{
			Memory: dmem,
			Offset: 0,
			Size:   memReq.Size,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	dmem.UnmapMemory()

	_, err = info.Cmd.Reset(0)
	if err != nil {
		log.Fatalln(err)
	}
	err = info.ExecuteBeginCommandBuffer()
	if err != nil {
		log.Fatalln(err)
	}
	// Intend to blit from this image, set the layout accordingly
	err = info.SetImageLayout(bltSrcImage, common.AspectColor, common.LayoutGeneral, common.LayoutTransferSrcOptimal, common.PipelineStageHost, common.PipelineStageTransfer)
	if err != nil {
		log.Fatalln(err)
	}

	bltDstImage := info.Buffers[info.CurrentBuffer].Image

	// Do a 32x32 blit to all of the dst image - should get big squares
	err = info.Cmd.CmdBlitImage(bltSrcImage, common.LayoutTransferSrcOptimal, bltDstImage, common.LayoutTransferDstOptimal, []*core.ImageBlit{
		{
			SourceSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			SourceOffsets: [2]common.Offset3D{
				{X: 0, Y: 0, Z: 0},
				{X: 32, Y: 32, Z: 1},
			},
			DestinationSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			DestinationOffsets: [2]common.Offset3D{
				{X: 0, Y: 0, Z: 0},
				{X: info.Width, Y: info.Height, Z: 1},
			},
		},
	}, common.FilterLinear)

	if err != nil {
		log.Fatalln(err)
	}

	// Use a barrier to make sure the blit is finished before the copy starts
	err = info.Cmd.CmdPipelineBarrier(common.PipelineStageTransfer, common.PipelineStageTransfer, 0, nil, nil, []*core.ImageMemoryBarrierOptions{
		{
			SrcAccessMask:       common.AccessTransferWrite,
			DstAccessMask:       common.AccessMemoryRead,
			OldLayout:           common.LayoutTransferDstOptimal,
			NewLayout:           common.LayoutTransferDstOptimal,
			SrcQueueFamilyIndex: -1,
			DstQueueFamilyIndex: -1,
			SubresourceRange: common.ImageSubresourceRange{
				AspectMask:     common.AspectColor,
				BaseMipLevel:   0,
				LevelCount:     1,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			Image: bltDstImage,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Do a image copy to part of the dst image - checks should stay small
	err = info.Cmd.CmdCopyImage(bltSrcImage, common.LayoutTransferSrcOptimal, bltDstImage, common.LayoutTransferDstOptimal, []core.ImageCopy{
		{
			SrcSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			SrcOffset: common.Offset3D{X: 0, Y: 0, Z: 0},
			DstSubresource: common.ImageSubresourceLayers{
				AspectMask:     common.AspectColor,
				MipLevel:       0,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			DstOffset: common.Offset3D{X: 256, Y: 256, Z: 0},
			Extent:    common.Extent3D{Width: 128, Height: 128, Depth: 1},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = info.Cmd.CmdPipelineBarrier(common.PipelineStageTransfer, common.PipelineStageBottomOfPipe, 0, nil, nil, []*core.ImageMemoryBarrierOptions{
		{
			SrcAccessMask:       common.AccessTransferWrite,
			DstAccessMask:       common.AccessMemoryRead,
			OldLayout:           common.LayoutTransferDstOptimal,
			NewLayout:           common.LayoutPresentSrcKHR,
			SrcQueueFamilyIndex: -1,
			DstQueueFamilyIndex: -1,
			SubresourceRange: common.ImageSubresourceRange{
				AspectMask:     common.AspectColor,
				BaseMipLevel:   0,
				LevelCount:     1,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
			Image: info.Buffers[info.CurrentBuffer].Image,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = info.Cmd.End()
	if err != nil {
		log.Fatalln(err)
	}

	drawFence, _, err := info.Loader.CreateFence(info.Device, &core.FenceOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	/* Queue the command buffer for execution */
	_, err = info.GraphicsQueue.SubmitToQueue(drawFence, []*core.SubmitOptions{
		{
			CommandBuffers: []core.CommandBuffer{info.Cmd},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = info.GraphicsQueue.WaitForIdle()
	if err != nil {
		log.Fatalln(err)
	}

	/* Now present the image in the window */

	/* Make sure command buffer is finished before presenting */
	for {
		res, err := drawFence.Wait(utils.FenceTimeout)
		if err != nil {
			log.Fatalln(err)
		}

		if res != core.VKTimeout {
			break
		}
	}
	_, _, err = info.Swapchain.PresentToQueue(info.PresentQueue, &khr_swapchain.PresentOptions{
		Swapchains:   []khr_swapchain.Swapchain{info.Swapchain},
		ImageIndices: []int{info.CurrentBuffer},
	})
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	/* VULKAN_KEY_END */
	if info.SaveImages {
		err = info.WritePNG("copy_blit_image")
		if err != nil {
			log.Fatalln(err)
		}
	}

	imageAcquiredSemaphore.Destroy()
	drawFence.Destroy()
	bltSrcImage.Destroy()
	info.Device.FreeMemory(dmem)
	info.DestroySwapchain()
	info.DestroyCommandBuffer()
	info.DestroyCommandPool()
	err = info.DestroyDevice()
	if err != nil {
		log.Fatalln(err)
	}
	info.Surface.Destroy()
	debugMessenger.Destroy()
	info.DestroyInstance()
	err = info.Window.Destroy()
	if err != nil {
		log.Fatalln(err)
	}
}
