package main

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
#include <windows.h>
#include "vulkan/vulkan_win32.h"
*/
import "C"
import (
	"errors"
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/commands"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/ext_debugutils"
	"github.com/CannibalVox/VKng/ext_surface"
	"github.com/CannibalVox/VKng/ext_surface_sdl2"
	"github.com/CannibalVox/VKng/ext_swapchain"
	"github.com/CannibalVox/VKng/pipeline"
	"github.com/CannibalVox/VKng/render_pass"
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type HelloTriangleApplication struct {
	allocator cgoalloc.Allocator
	window    *sdl.Window

	instance       *VKng.Instance
	debugMessenger *ext_debugutils.Messenger

	physicalDevice *VKng.PhysicalDevice
	logicalDevice  *VKng.Device
	graphicsQueue  *VKng.Queue
	presentQueue   *VKng.Queue

	surface             *ext_surface.Surface
	swapchainExtent     core.Extent2D
	swapchain           *ext_swapchain.Swapchain
	swapchainFormat     *ext_surface.SurfaceFormat
	swapchainImages     []*VKng.Image
	swapchainImageViews []*VKng.ImageView
	framebuffers        []*render_pass.Framebuffer
	imagesInFlight      []*VKng.Fence

	pipelineLayout *pipeline.PipelineLayout
	renderPass     *render_pass.RenderPass
	pipeline       *pipeline.Pipeline

	commandPool    *commands.CommandPool
	commandBuffers []*commands.CommandBuffer

	currentFrame            int
	imageAvailableSemaphore []*VKng.Semaphore
	renderFinishedSemaphore []*VKng.Semaphore
	inFlightFence           []*VKng.Fence
}

func (app *HelloTriangleApplication) Run() error {
	err := app.initWindow()
	if err != nil {
		return err
	}

	err = app.initVulkan()
	if err != nil {
		return err
	}
	defer app.cleanup()

	return app.mainLoop()
}

func (app *HelloTriangleApplication) initWindow() error {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return err
	}

	window, err := sdl.CreateWindow("Vulkan", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_VULKAN)
	if err != nil {
		return err
	}
	app.window = window

	return nil
}

func (app *HelloTriangleApplication) createInstance() error {
	instanceOptions := &VKng.InstanceOptions{
		ApplicationName:    "Hello Triangle",
		ApplicationVersion: core.CreateVersion(1, 0, 0),
		EngineName:         "No Engine",
		EngineVersion:      core.CreateVersion(1, 0, 0),
		VulkanVersion:      core.Vulkan1_2,
	}

	// Add extensions
	sdlExtensions := app.window.VulkanGetInstanceExtensions()
	extensions, err := VKng.AvailableExtensions(app.allocator)
	if err != nil {
		return err
	}

	for _, ext := range sdlExtensions {
		_, hasExt := extensions[ext]
		if !hasExt {
			return stacktrace.NewError("createinstance: cannot initialize sdl: missing extension %s", ext)
		}
		instanceOptions.ExtensionNames = append(instanceOptions.ExtensionNames, ext)
	}

	instanceOptions.ExtensionNames = append(instanceOptions.ExtensionNames, ext_debugutils.ExtensionName)

	// Add layers
	layers, err := VKng.AvailableLayers(app.allocator)
	if err != nil {
		return err
	}

	_, hasValidation := layers["VK_LAYER_KHRONOS_validation"]
	if !hasValidation {
		return errors.New("createInstance: cannot add khronos validation layer- install LunarG Vulkan SDK")
	}
	instanceOptions.LayerNames = append(instanceOptions.LayerNames, "VK_LAYER_KHRONOS_validation")

	// Add debug messenger
	debugMessengerOptions := &ext_debugutils.Options{
		CaptureSeverities: ext_debugutils.SeverityError | ext_debugutils.SeverityWarning,
		CaptureTypes:      ext_debugutils.TypeAll,
		Callback:          app.logDebug,
	}
	instanceOptions.Next = debugMessengerOptions

	app.instance, err = VKng.CreateInstance(app.allocator, instanceOptions)
	if err != nil {
		return err
	}

	app.debugMessenger, err = ext_debugutils.CreateMessenger(app.allocator, app.instance, debugMessengerOptions)
	if err != nil {
		return err
	}

	return nil
}

func (app *HelloTriangleApplication) initVulkan() error {
	err := app.createInstance()
	if err != nil {
		return err
	}

	surface, err := ext_surface_sdl2.CreateSurface(app.allocator, app.instance, &ext_surface_sdl2.CreationOptions{
		Window: app.window,
	})
	if err != nil {
		return err
	}
	app.surface = surface

	caps, err := app.pickPhysicalDevice()
	if err != nil {
		return err
	}

	err = app.createLogicalDevice(caps)
	if err != nil {
		return err
	}

	err = app.createSwapchain(caps)
	if err != nil {
		return err
	}

	err = app.createRenderPass()
	if err != nil {
		return err
	}

	err = app.createGraphicsPipeline()
	if err != nil {
		return err
	}

	err = app.createFramebuffers()
	if err != nil {
		return err
	}

	err = app.createCommandBuffers(caps)
	if err != nil {
		return err
	}

	return app.createSemaphores()
}

func (app *HelloTriangleApplication) cleanup() {
	for _, fence := range app.inFlightFence {
		fence.Destroy()
	}

	for _, semaphore := range app.renderFinishedSemaphore {
		semaphore.Destroy()
	}

	for _, semaphore := range app.imageAvailableSemaphore {
		semaphore.Destroy()
	}

	if app.commandPool != nil {
		app.commandPool.Destroy()
	}

	for _, framebuffer := range app.framebuffers {
		framebuffer.Destroy()
	}

	if app.pipeline != nil {
		app.pipeline.Destroy()
	}

	if app.pipelineLayout != nil {
		app.pipelineLayout.Destroy()
	}

	if app.renderPass != nil {
		app.renderPass.Destroy()
	}

	for _, imageView := range app.swapchainImageViews {
		imageView.Destroy()
	}

	if app.swapchain != nil {
		app.swapchain.Destroy()
	}

	if app.logicalDevice != nil {
		app.logicalDevice.Destroy()
	}

	if app.surface != nil {
		app.surface.Destroy()
	}

	if app.debugMessenger != nil {
		app.debugMessenger.Destroy()
	}

	if app.instance != nil {
		app.instance.Destroy()
	}

	if app.window != nil {
		app.window.Destroy()
	}
	sdl.Quit()

	app.allocator.Destroy()
}

func (app *HelloTriangleApplication) logDebug(msgType ext_debugutils.MessageType, severity ext_debugutils.MessageSeverity, data *ext_debugutils.CallbackData) bool {
	log.Printf("[%s %s] - %s", severity, msgType, data.Message)
	return false
}

func (app *HelloTriangleApplication) mainLoop() error {
	defer app.logicalDevice.WaitForIdle()

	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}
		}
		err := app.drawFrame()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}

func main() {
	defAlloc := &cgoalloc.DefaultAllocator{}
	lowTier, err := cgoalloc.CreateFixedBlockAllocator(defAlloc, 64*1024, 64, 8)
	if err != nil {
		log.Fatalln(err)
	}

	highTier, err := cgoalloc.CreateFixedBlockAllocator(defAlloc, 4096*1024, 4096, 8)
	if err != nil {
		log.Fatalln(err)
	}

	alloc := cgoalloc.CreateFallbackAllocator(highTier, defAlloc)
	alloc = cgoalloc.CreateFallbackAllocator(lowTier, alloc)

	app := &HelloTriangleApplication{
		allocator: alloc,
	}

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
