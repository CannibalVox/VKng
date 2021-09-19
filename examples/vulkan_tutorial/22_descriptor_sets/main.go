package main

import (
	"embed"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/ext_debug_utils"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/CannibalVox/VKng/extensions/khr_surface_sdl2"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/loov/hrtime"
	"github.com/palantir/stacktrace"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"math"
	"unsafe"
)

//go:embed shaders
var shaders embed.FS

const MaxFramesInFlight = 2

var validationLayers = []string{"VK_LAYER_KHRONOS_validation"}
var deviceExtensions = []string{khr_swapchain.ExtensionName}

const enableValidationLayers = true

type QueueFamilyIndices struct {
	GraphicsFamily *int
	PresentFamily  *int
}

func (i *QueueFamilyIndices) IsComplete() bool {
	return i.GraphicsFamily != nil && i.PresentFamily != nil
}

type SwapChainSupportDetails struct {
	Capabilities *khr_surface.Capabilities
	Formats      []khr_surface.Format
	PresentModes []khr_surface.PresentMode
}

type Vertex struct {
	Position mgl32.Vec2
	Color    mgl32.Vec3
}

type UniformBufferObject struct {
	Model mgl32.Mat4
	View  mgl32.Mat4
	Proj  mgl32.Mat4
}

func getVertexBindingDescription() []core.VertexBindingDescription {
	v := Vertex{}
	return []core.VertexBindingDescription{
		{
			Binding:   0,
			Stride:    unsafe.Sizeof(v),
			InputRate: core.RateVertex,
		},
	}
}

func getVertexAttributeDescriptions() []core.VertexAttributeDescription {
	v := Vertex{}
	return []core.VertexAttributeDescription{
		{
			Binding:  0,
			Location: 0,
			Format:   common.FormatR32G32SignedFloat,
			Offset:   unsafe.Offsetof(v.Position),
		},
		{
			Binding:  0,
			Location: 1,
			Format:   common.FormatR32G32B32SignedFloat,
			Offset:   unsafe.Offsetof(v.Color),
		},
	}
}

var vertices = []Vertex{
	{Position: mgl32.Vec2{-0.5, -0.5}, Color: mgl32.Vec3{1, 0, 0}},
	{Position: mgl32.Vec2{0.5, -0.5}, Color: mgl32.Vec3{0, 1, 0}},
	{Position: mgl32.Vec2{0.5, 0.5}, Color: mgl32.Vec3{0, 0, 1}},
	{Position: mgl32.Vec2{-0.5, 0.5}, Color: mgl32.Vec3{1, 1, 1}},
}

var indices = []uint16{0, 1, 2, 2, 3, 0}

type HelloTriangleApplication struct {
	window *sdl.Window
	loader core.Loader1_0

	instance       core.Instance
	debugMessenger ext_debug_utils.Messenger
	surface        khr_surface.Surface

	physicalDevice core.PhysicalDevice
	device         core.Device

	graphicsQueue core.Queue
	presentQueue  core.Queue

	swapchainLoader       khr_swapchain.Loader
	swapchain             khr_swapchain.Swapchain
	swapchainImages       []core.Image
	swapchainImageFormat  common.DataFormat
	swapchainExtent       common.Extent2D
	swapchainImageViews   []core.ImageView
	swapchainFramebuffers []core.Framebuffer

	renderPass          core.RenderPass
	descriptorPool      core.DescriptorPool
	descriptorSets      []core.DescriptorSet
	descriptorSetLayout core.DescriptorSetLayout
	pipelineLayout      core.PipelineLayout
	graphicsPipeline    core.Pipeline

	commandPool    core.CommandPool
	commandBuffers []core.CommandBuffer

	imageAvailableSemaphore []core.Semaphore
	renderFinishedSemaphore []core.Semaphore
	inFlightFence           []core.Fence
	imagesInFlight          []core.Fence
	currentFrame            int
	frameStart              float64

	vertexBuffer       core.Buffer
	vertexBufferMemory core.DeviceMemory
	indexBuffer        core.Buffer
	indexBufferMemory  core.DeviceMemory

	uniformBuffers       []core.Buffer
	uniformBuffersMemory []core.DeviceMemory
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

	window, err := sdl.CreateWindow("Vulkan", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_VULKAN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		return err
	}
	app.window = window

	app.loader, err = core.CreateLoaderFromProcAddr(sdl.VulkanGetVkGetInstanceProcAddr())
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

	err = app.setupDebugMessenger()
	if err != nil {
		return err
	}

	err = app.createSurface()
	if err != nil {
		return err
	}

	err = app.pickPhysicalDevice()
	if err != nil {
		return err
	}

	err = app.createLogicalDevice()
	if err != nil {
		return err
	}

	err = app.createSwapchain()
	if err != nil {
		return err
	}

	err = app.createImageViews()
	if err != nil {
		return err
	}

	err = app.createRenderPass()
	if err != nil {
		return err
	}

	err = app.createDescriptorSetLayout()
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

	err = app.createCommandPool()
	if err != nil {
		return err
	}

	err = app.createVertexBuffer()
	if err != nil {
		return err
	}

	err = app.createIndexBuffer()
	if err != nil {
		return err
	}

	err = app.createUniformBuffers()
	if err != nil {
		return err
	}

	err = app.createDescriptorPool()
	if err != nil {
		return err
	}

	err = app.createDescriptorSets()
	if err != nil {
		return err
	}

	err = app.createCommandBuffers()
	if err != nil {
		return err
	}

	return app.createSyncObjects()
}

func (app *HelloTriangleApplication) mainLoop() error {
	rendering := true

appLoop:
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				break appLoop
			case *sdl.WindowEvent:
				switch e.Event {
				case sdl.WINDOWEVENT_MINIMIZED:
					rendering = false
				case sdl.WINDOWEVENT_RESTORED:
					rendering = true
				case sdl.WINDOWEVENT_RESIZED:
					w, h := app.window.GetSize()
					if w > 0 && h > 0 {
						rendering = true
						app.recreateSwapChain()
					} else {
						rendering = false
					}
				}
			}
		}
		if rendering {
			err := app.drawFrame()
			if err != nil {
				return err
			}
		}
	}

	_, err := app.device.WaitForIdle()
	return err
}

func (app *HelloTriangleApplication) cleanupSwapChain() {
	for _, framebuffer := range app.swapchainFramebuffers {
		framebuffer.Destroy()
	}
	app.swapchainFramebuffers = []core.Framebuffer{}

	if len(app.commandBuffers) > 0 {
		app.commandPool.FreeCommandBuffers(app.commandBuffers)
		app.commandBuffers = []core.CommandBuffer{}
	}

	if app.graphicsPipeline != nil {
		app.graphicsPipeline.Destroy()
		app.graphicsPipeline = nil
	}

	if app.pipelineLayout != nil {
		app.pipelineLayout.Destroy()
		app.pipelineLayout = nil
	}

	if app.renderPass != nil {
		app.renderPass.Destroy()
		app.renderPass = nil
	}

	for _, imageView := range app.swapchainImageViews {
		imageView.Destroy()
	}
	app.swapchainImageViews = []core.ImageView{}

	if app.swapchain != nil {
		app.swapchain.Destroy()
		app.swapchain = nil
	}

	for i := 0; i < len(app.uniformBuffers); i++ {
		app.uniformBuffers[i].Destroy()
	}
	app.uniformBuffers = app.uniformBuffers[:0]

	for i := 0; i < len(app.uniformBuffersMemory); i++ {
		app.uniformBuffersMemory[i].Free()
	}
	app.uniformBuffersMemory = app.uniformBuffersMemory[:0]

	app.descriptorPool.Destroy()
}

func (app *HelloTriangleApplication) cleanup() {
	app.cleanupSwapChain()

	if app.descriptorSetLayout != nil {
		app.descriptorSetLayout.Destroy()
	}

	if app.indexBuffer != nil {
		app.indexBuffer.Destroy()
	}

	if app.indexBufferMemory != nil {
		app.indexBufferMemory.Free()
	}

	if app.vertexBuffer != nil {
		app.vertexBuffer.Destroy()
	}

	if app.vertexBufferMemory != nil {
		app.vertexBufferMemory.Free()
	}

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

	if app.device != nil {
		app.device.Destroy()
	}

	if app.debugMessenger != nil {
		app.debugMessenger.Destroy()
	}

	if app.surface != nil {
		app.surface.Destroy()
	}

	if app.instance != nil {
		app.instance.Destroy()
	}

	if app.window != nil {
		app.window.Destroy()
	}
	sdl.Quit()
}

func (app *HelloTriangleApplication) recreateSwapChain() error {
	w, h := app.window.VulkanGetDrawableSize()
	if w == 0 || h == 0 {
		return nil
	}
	if (app.window.GetFlags() & sdl.WINDOW_MINIMIZED) != 0 {
		return nil
	}

	_, err := app.device.WaitForIdle()
	if err != nil {
		return err
	}

	app.cleanupSwapChain()

	err = app.createSwapchain()
	if err != nil {
		return err
	}

	err = app.createImageViews()
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

	err = app.createUniformBuffers()
	if err != nil {
		return err
	}

	err = app.createDescriptorPool()
	if err != nil {
		return err
	}

	err = app.createDescriptorSets()
	if err != nil {
		return err
	}

	err = app.createCommandBuffers()
	if err != nil {
		return err
	}

	app.imagesInFlight = []core.Fence{}
	for i := 0; i < len(app.swapchainImages); i++ {
		app.imagesInFlight = append(app.imagesInFlight, nil)
	}

	return nil
}

func (app *HelloTriangleApplication) createInstance() error {
	instanceOptions := &core.InstanceOptions{
		ApplicationName:    "Hello Triangle",
		ApplicationVersion: common.CreateVersion(1, 0, 0),
		EngineName:         "No Engine",
		EngineVersion:      common.CreateVersion(1, 0, 0),
		VulkanVersion:      common.Vulkan1_2,
	}

	// Add extensions
	sdlExtensions := app.window.VulkanGetInstanceExtensions()
	extensions, _, err := app.loader.AvailableExtensions()
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

	if enableValidationLayers {
		instanceOptions.ExtensionNames = append(instanceOptions.ExtensionNames, ext_debug_utils.ExtensionName)
	}

	// Add layers
	layers, _, err := app.loader.AvailableLayers()
	if err != nil {
		return err
	}

	if enableValidationLayers {
		for _, layer := range validationLayers {
			_, hasValidation := layers[layer]
			if !hasValidation {
				return stacktrace.NewError("createInstance: cannot add validation- layer %s not available- install LunarG Vulkan SDK", layer)
			}
			instanceOptions.LayerNames = append(instanceOptions.LayerNames, layer)
		}

		// Add debug messenger
		instanceOptions.Next = app.debugMessengerOptions()
	}

	app.instance, _, err = app.loader.CreateInstance(instanceOptions)
	if err != nil {
		return err
	}

	return nil
}

func (app *HelloTriangleApplication) debugMessengerOptions() *ext_debug_utils.Options {
	return &ext_debug_utils.Options{
		CaptureSeverities: ext_debug_utils.SeverityError | ext_debug_utils.SeverityWarning,
		CaptureTypes:      ext_debug_utils.TypeAll,
		Callback:          app.logDebug,
	}
}

func (app *HelloTriangleApplication) setupDebugMessenger() error {
	if !enableValidationLayers {
		return nil
	}

	var err error
	app.debugMessenger, _, err = ext_debug_utils.CreateMessenger(app.instance, app.debugMessengerOptions())
	if err != nil {
		return err
	}

	return nil
}

func (app *HelloTriangleApplication) createSurface() error {
	surfaceLoader := khr_surface_sdl2.CreateLoaderFromInstance(app.instance)
	surface, _, err := surfaceLoader.CreateSurface(app.window)
	if err != nil {
		return err
	}

	app.surface = surface
	return nil
}

func (app *HelloTriangleApplication) pickPhysicalDevice() error {
	physicalDevices, _, err := app.instance.PhysicalDevices()
	if err != nil {
		return err
	}

	for _, device := range physicalDevices {
		if app.isDeviceSuitable(device) {
			app.physicalDevice = device
			break
		}
	}

	if app.physicalDevice == nil {
		return stacktrace.NewError("failed to find a suitable GPU!")
	}

	return nil
}

func (app *HelloTriangleApplication) createLogicalDevice() error {
	indices, err := app.findQueueFamilies(app.physicalDevice)
	if err != nil {
		return err
	}

	uniqueQueueFamilies := []int{*indices.GraphicsFamily}
	if uniqueQueueFamilies[0] != *indices.PresentFamily {
		uniqueQueueFamilies = append(uniqueQueueFamilies, *indices.PresentFamily)
	}

	var queueFamilyOptions []*core.QueueFamilyOptions
	queuePriority := float32(1.0)
	for _, queueFamily := range uniqueQueueFamilies {
		queueFamilyOptions = append(queueFamilyOptions, &core.QueueFamilyOptions{
			QueueFamilyIndex: queueFamily,
			QueuePriorities:  []float32{queuePriority},
		})
	}

	var extensionNames []string
	extensionNames = append(extensionNames, deviceExtensions...)

	// Makes this example compatible with vulkan portability, necessary to run on mobile & mac
	extensions, _, err := app.physicalDevice.AvailableExtensions()
	if err != nil {
		return err
	}

	_, supported := extensions["VK_KHR_portability_subset"]
	if supported {
		extensionNames = append(extensionNames, "VK_KHR_portability_subset")
	}

	var layerNames []string
	if enableValidationLayers {
		layerNames = append(layerNames, validationLayers...)
	}

	app.device, _, err = app.loader.CreateDevice(app.physicalDevice, &core.DeviceOptions{
		QueueFamilies:   queueFamilyOptions,
		EnabledFeatures: &common.PhysicalDeviceFeatures{},
		ExtensionNames:  extensionNames,
		LayerNames:      layerNames,
	})
	if err != nil {
		return err
	}

	app.graphicsQueue, err = app.device.GetQueue(*indices.GraphicsFamily, 0)
	if err != nil {
		return err
	}

	app.presentQueue, err = app.device.GetQueue(*indices.PresentFamily, 0)
	return err
}

func (app *HelloTriangleApplication) createSwapchain() error {
	app.swapchainLoader = khr_swapchain.CreateLoaderFromDevice(app.device)

	swapchainSupport, err := app.querySwapChainSupport(app.physicalDevice)
	if err != nil {
		return err
	}

	surfaceFormat := app.chooseSwapSurfaceFormat(swapchainSupport.Formats)
	presentMode := app.chooseSwapPresentMode(swapchainSupport.PresentModes)
	extent := app.chooseSwapExtent(swapchainSupport.Capabilities)

	imageCount := swapchainSupport.Capabilities.MinImageCount + 1
	if swapchainSupport.Capabilities.MaxImageCount > 0 && swapchainSupport.Capabilities.MaxImageCount < imageCount {
		imageCount = swapchainSupport.Capabilities.MaxImageCount
	}

	sharingMode := common.SharingExclusive
	var queueFamilyIndices []int

	indices, err := app.findQueueFamilies(app.physicalDevice)
	if err != nil {
		return err
	}

	if *indices.GraphicsFamily != *indices.PresentFamily {
		sharingMode = common.SharingConcurrent
		queueFamilyIndices = append(queueFamilyIndices, *indices.GraphicsFamily, *indices.PresentFamily)
	}

	swapchain, _, err := app.swapchainLoader.CreateSwapchain(app.device, &khr_swapchain.CreationOptions{
		Surface: app.surface,

		MinImageCount:    imageCount,
		ImageFormat:      surfaceFormat.Format,
		ImageColorSpace:  surfaceFormat.ColorSpace,
		ImageExtent:      extent,
		ImageArrayLayers: 1,
		ImageUsage:       common.ImageColorAttachment,

		SharingMode:        sharingMode,
		QueueFamilyIndices: queueFamilyIndices,

		PreTransform:   swapchainSupport.Capabilities.CurrentTransform,
		CompositeAlpha: khr_surface.Opaque,
		PresentMode:    presentMode,
		Clipped:        true,
	})
	if err != nil {
		return err
	}
	app.swapchainExtent = extent
	app.swapchain = swapchain
	app.swapchainImageFormat = surfaceFormat.Format

	return nil
}

func (app *HelloTriangleApplication) createImageViews() error {
	images, _, err := app.swapchain.Images()
	if err != nil {
		return err
	}
	app.swapchainImages = images

	var imageViews []core.ImageView
	for _, image := range images {
		view, _, err := app.loader.CreateImageView(app.device, &core.ImageViewOptions{
			ViewType: common.View2D,
			Image:    image,
			Format:   app.swapchainImageFormat,
			Components: common.ComponentMapping{
				R: common.SwizzleIdentity,
				G: common.SwizzleIdentity,
				B: common.SwizzleIdentity,
				A: common.SwizzleIdentity,
			},
			SubresourceRange: common.ImageSubresourceRange{
				AspectMask:     common.AspectColor,
				BaseMipLevel:   0,
				LevelCount:     1,
				BaseArrayLayer: 0,
				LayerCount:     1,
			},
		})
		if err != nil {
			return err
		}

		imageViews = append(imageViews, view)
	}
	app.swapchainImageViews = imageViews

	return nil
}

func (app *HelloTriangleApplication) createRenderPass() error {
	renderPass, _, err := app.loader.CreateRenderPass(app.device, &core.RenderPassOptions{
		Attachments: []core.AttachmentDescription{
			{
				Format:         app.swapchainImageFormat,
				Samples:        common.Samples1,
				LoadOp:         common.LoadOpClear,
				StoreOp:        common.StoreOpStore,
				StencilLoadOp:  common.LoadOpDontCare,
				StencilStoreOp: common.StoreOpDontCare,
				InitialLayout:  common.LayoutUndefined,
				FinalLayout:    common.LayoutPresentSrc,
			},
		},
		SubPasses: []core.SubPass{
			{
				BindPoint: common.BindGraphics,
				ColorAttachments: []common.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          common.LayoutColorAttachmentOptimal,
					},
				},
			},
		},
		SubPassDependencies: []core.SubPassDependency{
			{
				SrcSubPassIndex: core.SubpassExternal,
				DstSubPassIndex: 0,

				SrcStageMask: common.PipelineStageColorAttachmentOutput,
				SrcAccess:    0,

				DstStageMask: common.PipelineStageColorAttachmentOutput,
				DstAccess:    common.AccessColorAttachmentWrite,
			},
		},
	})
	if err != nil {
		return err
	}

	app.renderPass = renderPass

	return nil
}

func (app *HelloTriangleApplication) createDescriptorSetLayout() error {
	var err error
	app.descriptorSetLayout, _, err = app.loader.CreateDescriptorSetLayout(app.device, &core.DescriptorSetLayoutOptions{
		Bindings: []*core.DescriptorLayoutBinding{
			{
				Binding: 0,
				Type:    common.DescriptorUniformBuffer,
				Count:   1,

				ShaderStages: common.StageVertex,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func bytesToBytecode(b []byte) []uint32 {
	byteCode := make([]uint32, len(b)/4)
	for i := 0; i < len(byteCode); i++ {
		byteIndex := i * 4
		byteCode[i] = 0
		byteCode[i] |= uint32(b[byteIndex])
		byteCode[i] |= uint32(b[byteIndex+1]) << 8
		byteCode[i] |= uint32(b[byteIndex+2]) << 16
		byteCode[i] |= uint32(b[byteIndex+3]) << 24
	}

	return byteCode
}

func (app *HelloTriangleApplication) createGraphicsPipeline() error {
	// Load vertex shader
	vertShaderBytes, err := shaders.ReadFile("shaders/vert.spv")
	if err != nil {
		return err
	}

	vertShader, _, err := app.loader.CreateShaderModule(app.device, &core.ShaderModuleOptions{
		SpirVByteCode: bytesToBytecode(vertShaderBytes),
	})
	if err != nil {
		return err
	}
	defer vertShader.Destroy()

	// Load fragment shader
	fragShaderBytes, err := shaders.ReadFile("shaders/frag.spv")
	if err != nil {
		return err
	}

	fragShader, _, err := app.loader.CreateShaderModule(app.device, &core.ShaderModuleOptions{
		SpirVByteCode: bytesToBytecode(fragShaderBytes),
	})
	if err != nil {
		return err
	}
	defer fragShader.Destroy()

	vertexInput := &core.VertexInputOptions{
		VertexBindingDescriptions:   getVertexBindingDescription(),
		VertexAttributeDescriptions: getVertexAttributeDescriptions(),
	}

	inputAssembly := &core.InputAssemblyOptions{
		Topology:               common.TopologyTriangleList,
		EnablePrimitiveRestart: false,
	}

	vertStage := &core.ShaderStage{
		Stage:  common.StageVertex,
		Shader: vertShader,
		Name:   "main",
	}

	fragStage := &core.ShaderStage{
		Stage:  common.StageFragment,
		Shader: fragShader,
		Name:   "main",
	}

	viewport := &core.ViewportOptions{
		Viewports: []common.Viewport{
			{
				X:        0,
				Y:        0,
				Width:    float32(app.swapchainExtent.Width),
				Height:   float32(app.swapchainExtent.Height),
				MinDepth: 0,
				MaxDepth: 1,
			},
		},
		Scissors: []common.Rect2D{
			{
				Offset: common.Offset2D{X: 0, Y: 0},
				Extent: app.swapchainExtent,
			},
		},
	}

	rasterization := &core.RasterizationOptions{
		DepthClamp:        false,
		RasterizerDiscard: false,

		PolygonMode: core.ModeFill,
		CullMode:    common.CullBack,
		FrontFace:   common.CounterClockwise,

		DepthBias: false,

		LineWidth: 1.0,
	}

	multisample := &core.MultisampleOptions{
		SampleShading:        false,
		RasterizationSamples: common.Samples1,
		MinSampleShading:     1.0,
	}

	colorBlend := &core.ColorBlendOptions{
		LogicOpEnabled: false,
		LogicOp:        common.LogicOpCopy,

		BlendConstants: [4]float32{0, 0, 0, 0},
		Attachments: []core.ColorBlendAttachment{
			{
				BlendEnabled: false,
				WriteMask:    common.ComponentRed | common.ComponentGreen | common.ComponentBlue | common.ComponentAlpha,
			},
		},
	}

	app.pipelineLayout, _, err = app.loader.CreatePipelineLayout(app.device, &core.PipelineLayoutOptions{
		SetLayouts: []core.DescriptorSetLayout{
			app.descriptorSetLayout,
		},
	})

	pipelines, _, err := app.loader.CreateGraphicsPipelines(app.device, []*core.Options{
		{
			ShaderStages: []*core.ShaderStage{
				vertStage,
				fragStage,
			},
			VertexInput:       vertexInput,
			InputAssembly:     inputAssembly,
			Viewport:          viewport,
			Rasterization:     rasterization,
			Multisample:       multisample,
			ColorBlend:        colorBlend,
			Layout:            app.pipelineLayout,
			RenderPass:        app.renderPass,
			SubPass:           0,
			BasePipelineIndex: -1,
		},
	})
	if err != nil {
		return err
	}
	app.graphicsPipeline = pipelines[0]

	return nil
}

func (app *HelloTriangleApplication) createFramebuffers() error {
	for _, imageView := range app.swapchainImageViews {
		framebuffer, _, err := app.loader.CreateFrameBuffer(app.device, &core.FramebufferOptions{
			RenderPass: app.renderPass,
			Layers:     1,
			Attachments: []core.ImageView{
				imageView,
			},
			Width:  app.swapchainExtent.Width,
			Height: app.swapchainExtent.Height,
		})
		if err != nil {
			return err
		}

		app.swapchainFramebuffers = append(app.swapchainFramebuffers, framebuffer)
	}

	return nil
}

func (app *HelloTriangleApplication) createCommandPool() error {
	indices, err := app.findQueueFamilies(app.physicalDevice)
	if err != nil {
		return err
	}

	pool, _, err := app.loader.CreateCommandPool(app.device, &core.CommandPoolOptions{
		GraphicsQueueFamily: indices.GraphicsFamily,
	})

	if err != nil {
		return err
	}
	app.commandPool = pool

	return nil
}

func (app *HelloTriangleApplication) createVertexBuffer() error {
	var err error
	bufferSize := binary.Size(vertices)

	stagingBuffer, stagingBufferMemory, err := app.createBuffer(bufferSize, common.UsageTransferSrc, core.MemoryHostVisible|core.MemoryHostCoherent)
	if stagingBuffer != nil {
		defer stagingBuffer.Destroy()
	}
	if stagingBufferMemory != nil {
		defer stagingBufferMemory.Free()
	}

	if err != nil {
		return err
	}

	_, err = stagingBufferMemory.WriteData(0, vertices)
	if err != nil {
		return err
	}

	app.vertexBuffer, app.vertexBufferMemory, err = app.createBuffer(bufferSize, common.UsageTransferDst|common.UsageVertexBuffer, core.MemoryDeviceLocal)
	if err != nil {
		return err
	}

	return app.copyBuffer(stagingBuffer, app.vertexBuffer, bufferSize)
}

func (app *HelloTriangleApplication) createIndexBuffer() error {
	bufferSize := binary.Size(indices)

	stagingBuffer, stagingBufferMemory, err := app.createBuffer(bufferSize, common.UsageTransferSrc, core.MemoryHostVisible|core.MemoryHostCoherent)
	if stagingBuffer != nil {
		defer stagingBuffer.Destroy()
	}
	if stagingBufferMemory != nil {
		defer stagingBufferMemory.Free()
	}

	if err != nil {
		return err
	}

	_, err = stagingBufferMemory.WriteData(0, indices)
	if err != nil {
		return err
	}

	app.indexBuffer, app.indexBufferMemory, err = app.createBuffer(bufferSize, common.UsageTransferDst|common.UsageIndexBuffer, core.MemoryDeviceLocal)
	if err != nil {
		return err
	}

	return app.copyBuffer(stagingBuffer, app.indexBuffer, bufferSize)
}

func (app *HelloTriangleApplication) createUniformBuffers() error {
	bufferSize := int(unsafe.Sizeof(UniformBufferObject{}))

	for i := 0; i < len(app.swapchainImages); i++ {
		buffer, memory, err := app.createBuffer(bufferSize, common.UsageUniformBuffer, core.MemoryHostVisible|core.MemoryHostCoherent)
		if err != nil {
			return err
		}

		app.uniformBuffers = append(app.uniformBuffers, buffer)
		app.uniformBuffersMemory = append(app.uniformBuffersMemory, memory)
	}

	return nil
}

func (app *HelloTriangleApplication) createDescriptorPool() error {
	var err error
	app.descriptorPool, _, err = app.loader.CreateDescriptorPool(app.device, &core.DescriptorPoolOptions{
		MaxSets: len(app.swapchainImages),
		PoolSizes: []core.PoolSize{
			{
				Type:  common.DescriptorUniformBuffer,
				Count: len(app.swapchainImages),
			},
		},
	})
	return err
}

func (app *HelloTriangleApplication) createDescriptorSets() error {
	var allocLayouts []core.DescriptorSetLayout
	for i := 0; i < len(app.swapchainImages); i++ {
		allocLayouts = append(allocLayouts, app.descriptorSetLayout)
	}

	var err error
	app.descriptorSets, _, err = app.descriptorPool.AllocateDescriptorSets(&core.DescriptorSetOptions{
		AllocationLayouts: allocLayouts,
	})
	if err != nil {
		return err
	}

	for i := 0; i < len(app.swapchainImages); i++ {
		err = app.device.UpdateDescriptorSets([]core.WriteDescriptorSetOptions{
			{
				Destination:             app.descriptorSets[i],
				DestinationBinding:      0,
				DestinationArrayElement: 0,

				DescriptorType: common.DescriptorUniformBuffer,

				BufferInfo: []core.DescriptorBufferInfo{
					{
						Buffer: app.uniformBuffers[i],
						Offset: 0,
						Range:  uint64(unsafe.Sizeof(UniformBufferObject{})),
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *HelloTriangleApplication) createBuffer(size int, usage common.BufferUsages, properties core.MemoryPropertyFlags) (core.Buffer, core.DeviceMemory, error) {
	buffer, _, err := app.loader.CreateBuffer(app.device, &core.BufferOptions{
		BufferSize:  size,
		Usages:      usage,
		SharingMode: common.SharingExclusive,
	})
	if err != nil {
		return nil, nil, err
	}

	memRequirements, err := buffer.MemoryRequirements()
	if err != nil {
		return nil, nil, err
	}

	memoryTypeIndex, err := app.findMemoryType(memRequirements.MemoryType, properties)
	if err != nil {
		return buffer, nil, err
	}

	memory, _, err := app.device.AllocateMemory(&core.DeviceMemoryOptions{
		AllocationSize:  memRequirements.Size,
		MemoryTypeIndex: memoryTypeIndex,
	})
	if err != nil {
		return buffer, nil, err
	}

	_, err = buffer.BindBufferMemory(memory, 0)
	return buffer, memory, err
}

func (app *HelloTriangleApplication) copyBuffer(srcBuffer core.Buffer, dstBuffer core.Buffer, size int) error {
	buffers, _, err := app.commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		Level:       common.LevelPrimary,
		BufferCount: 1,
	})
	if err != nil {
		return err
	}

	buffer := buffers[0]
	_, err = buffer.Begin(&core.BeginOptions{
		Flags: core.OneTimeSubmit,
	})
	if err != nil {
		return err
	}
	defer app.commandPool.FreeCommandBuffers(buffers)

	buffer.CmdCopyBuffer(srcBuffer, dstBuffer, []core.BufferCopy{
		{
			SrcOffset: 0,
			DstOffset: 0,
			Size:      size,
		},
	})

	_, err = buffer.End()
	if err != nil {
		return err
	}

	_, err = core.SubmitToQueue(app.graphicsQueue, nil, []*core.SubmitOptions{
		{
			CommandBuffers: []core.CommandBuffer{buffer},
		},
	})
	if err != nil {
		return err
	}

	_, err = app.graphicsQueue.WaitForIdle()
	return err
}

func (app *HelloTriangleApplication) findMemoryType(typeFilter uint32, properties core.MemoryPropertyFlags) (int, error) {
	memProperties := app.physicalDevice.MemoryProperties()
	for i, memoryType := range memProperties.MemoryTypes {
		typeBit := uint32(1 << i)

		if (typeFilter&typeBit) != 0 && (memoryType.Properties&properties) == properties {
			return i, nil
		}
	}

	return 0, stacktrace.NewError("failed to find any suitable memory type!")
}

func (app *HelloTriangleApplication) createCommandBuffers() error {

	buffers, _, err := app.commandPool.AllocateCommandBuffers(&core.CommandBufferOptions{
		Level:       common.LevelPrimary,
		BufferCount: len(app.swapchainImages),
	})
	if err != nil {
		return err
	}
	app.commandBuffers = buffers

	for bufferIdx, buffer := range buffers {
		_, err = buffer.Begin(&core.BeginOptions{})
		if err != nil {
			return err
		}

		err = buffer.CmdBeginRenderPass(core.ContentsInline,
			&core.RenderPassBeginOptions{
				RenderPass:  app.renderPass,
				Framebuffer: app.swapchainFramebuffers[bufferIdx],
				RenderArea: common.Rect2D{
					Offset: common.Offset2D{X: 0, Y: 0},
					Extent: app.swapchainExtent,
				},
				ClearValues: []core.ClearValue{
					core.ClearValueFloat{0, 0, 0, 1},
				},
			})
		if err != nil {
			return err
		}

		buffer.CmdBindPipeline(common.BindGraphics, app.graphicsPipeline)
		buffer.CmdBindVertexBuffers(0, []core.Buffer{app.vertexBuffer}, []int{0})
		buffer.CmdBindIndexBuffer(app.indexBuffer, 0, common.IndexUInt16)
		buffer.CmdBindDescriptorSets(common.BindGraphics, app.pipelineLayout, 0, []core.DescriptorSet{
			app.descriptorSets[bufferIdx],
		}, nil)
		buffer.CmdDrawIndexed(len(indices), 1, 0, 0, 0)
		buffer.CmdEndRenderPass()

		_, err = buffer.End()
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *HelloTriangleApplication) createSyncObjects() error {
	for i := 0; i < MaxFramesInFlight; i++ {
		semaphore, _, err := app.loader.CreateSemaphore(app.device, &core.SemaphoreOptions{})
		if err != nil {
			return err
		}

		app.imageAvailableSemaphore = append(app.imageAvailableSemaphore, semaphore)

		semaphore, _, err = app.loader.CreateSemaphore(app.device, &core.SemaphoreOptions{})
		if err != nil {
			return err
		}

		app.renderFinishedSemaphore = append(app.renderFinishedSemaphore, semaphore)

		fence, _, err := app.loader.CreateFence(app.device, &core.FenceOptions{
			Flags: core.FenceSignaled,
		})
		if err != nil {
			return err
		}

		app.inFlightFence = append(app.inFlightFence, fence)
	}

	for i := 0; i < len(app.swapchainImages); i++ {
		app.imagesInFlight = append(app.imagesInFlight, nil)
	}

	return nil
}

func (app *HelloTriangleApplication) drawFrame() error {
	fences := []core.Fence{app.inFlightFence[app.currentFrame]}

	_, err := app.device.WaitForFences(true, common.NoTimeout, fences)
	if err != nil {
		return err
	}

	imageIndex, res, err := app.swapchain.AcquireNextImage(common.NoTimeout, app.imageAvailableSemaphore[app.currentFrame], nil)
	if res == core.VKErrorOutOfDate {
		return app.recreateSwapChain()
	} else if err != nil {
		return err
	}

	if app.imagesInFlight[imageIndex] != nil {
		_, err := app.device.WaitForFences(true, common.NoTimeout, []core.Fence{app.imagesInFlight[imageIndex]})
		if err != nil {
			return err
		}
	}
	app.imagesInFlight[imageIndex] = app.inFlightFence[app.currentFrame]

	_, err = app.device.ResetFences(fences)
	if err != nil {
		return err
	}

	err = app.updateUniformBuffer(imageIndex)
	if err != nil {
		return err
	}

	_, err = core.SubmitToQueue(app.graphicsQueue, app.inFlightFence[app.currentFrame], []*core.SubmitOptions{
		{
			WaitSemaphores:   []core.Semaphore{app.imageAvailableSemaphore[app.currentFrame]},
			WaitDstStages:    []common.PipelineStages{common.PipelineStageColorAttachmentOutput},
			CommandBuffers:   []core.CommandBuffer{app.commandBuffers[imageIndex]},
			SignalSemaphores: []core.Semaphore{app.renderFinishedSemaphore[app.currentFrame]},
		},
	})
	if err != nil {
		return err
	}

	_, res, err = app.swapchain.PresentToQueue(app.presentQueue, &khr_swapchain.PresentOptions{
		WaitSemaphores: []core.Semaphore{app.renderFinishedSemaphore[app.currentFrame]},
		Swapchains:     []khr_swapchain.Swapchain{app.swapchain},
		ImageIndices:   []int{imageIndex},
	})
	if res == core.VKErrorOutOfDate || res == core.VKSuboptimal {
		return app.recreateSwapChain()
	} else if err != nil {
		return err
	}

	app.currentFrame = (app.currentFrame + 1) % MaxFramesInFlight

	return nil
}

func (app *HelloTriangleApplication) updateUniformBuffer(currentImage int) error {
	currentTime := hrtime.Now().Seconds()
	timePeriod := float32(math.Mod(currentTime, 4.0))

	ubo := UniformBufferObject{}
	ubo.Model = mgl32.HomogRotate3D(timePeriod*mgl32.DegToRad(90.0), mgl32.Vec3{0, 0, 1})
	ubo.View = mgl32.LookAt(2, 2, 2, 0, 0, 0, 0, 0, 1)
	aspectRatio := float32(app.swapchainExtent.Width) / float32(app.swapchainExtent.Height)
	ubo.Proj = mgl32.Perspective(mgl32.DegToRad(45), aspectRatio, 0.1, 10)
	ubo.Proj[5] *= -1

	_, err := app.uniformBuffersMemory[currentImage].WriteData(0, &ubo)
	return err
}

func (app *HelloTriangleApplication) chooseSwapSurfaceFormat(availableFormats []khr_surface.Format) khr_surface.Format {
	for _, format := range availableFormats {
		if format.Format == common.FormatB8G8R8A8SRGB && format.ColorSpace == khr_surface.SRGBNonlinear {
			return format
		}
	}

	return availableFormats[0]
}

func (app *HelloTriangleApplication) chooseSwapPresentMode(availablePresentModes []khr_surface.PresentMode) khr_surface.PresentMode {
	for _, presentMode := range availablePresentModes {
		if presentMode == khr_surface.Mailbox {
			return presentMode
		}
	}

	return khr_surface.FIFO
}

func (app *HelloTriangleApplication) chooseSwapExtent(capabilities *khr_surface.Capabilities) common.Extent2D {
	if capabilities.CurrentExtent.Width != (^uint32(0)) {
		return capabilities.CurrentExtent
	}

	widthInt, heightInt := app.window.VulkanGetDrawableSize()
	width := uint32(widthInt)
	height := uint32(heightInt)

	if width < capabilities.MinImageExtent.Width {
		width = capabilities.MinImageExtent.Width
	}
	if width > capabilities.MaxImageExtent.Width {
		width = capabilities.MaxImageExtent.Width
	}
	if height < capabilities.MinImageExtent.Height {
		height = capabilities.MinImageExtent.Height
	}
	if height > capabilities.MaxImageExtent.Height {
		height = capabilities.MaxImageExtent.Height
	}

	return common.Extent2D{Width: width, Height: height}
}

func (app *HelloTriangleApplication) querySwapChainSupport(device core.PhysicalDevice) (SwapChainSupportDetails, error) {
	var details SwapChainSupportDetails
	var err error

	details.Capabilities, _, err = app.surface.Capabilities(device)
	if err != nil {
		return details, err
	}

	details.Formats, _, err = app.surface.Formats(device)
	if err != nil {
		return details, err
	}

	details.PresentModes, _, err = app.surface.PresentModes(device)
	return details, err
}

func (app *HelloTriangleApplication) isDeviceSuitable(device core.PhysicalDevice) bool {
	indices, err := app.findQueueFamilies(device)
	if err != nil {
		return false
	}

	extensionsSupported := app.checkDeviceExtensionSupport(device)

	var swapChainAdequate bool
	if extensionsSupported {
		swapChainSupport, err := app.querySwapChainSupport(device)
		if err != nil {
			return false
		}

		swapChainAdequate = len(swapChainSupport.Formats) > 0 && len(swapChainSupport.PresentModes) > 0
	}

	return indices.IsComplete() && extensionsSupported && swapChainAdequate
}

func (app *HelloTriangleApplication) checkDeviceExtensionSupport(device core.PhysicalDevice) bool {
	extensions, _, err := device.AvailableExtensions()
	if err != nil {
		return false
	}

	for _, extension := range deviceExtensions {
		_, hasExtension := extensions[extension]
		if !hasExtension {
			return false
		}
	}

	return true
}

func (app *HelloTriangleApplication) findQueueFamilies(device core.PhysicalDevice) (QueueFamilyIndices, error) {
	indices := QueueFamilyIndices{}
	queueFamilies, err := device.QueueFamilyProperties()
	if err != nil {
		return indices, err
	}

	for queueFamilyIdx, queueFamily := range queueFamilies {
		if (queueFamily.Flags & common.Graphics) != 0 {
			indices.GraphicsFamily = new(int)
			*indices.GraphicsFamily = queueFamilyIdx
		}

		supported, _, err := app.surface.SupportsDevice(device, queueFamilyIdx)
		if err != nil {
			return indices, err
		}

		if supported {
			indices.PresentFamily = new(int)
			*indices.PresentFamily = queueFamilyIdx
		}

		if indices.IsComplete() {
			break
		}
	}

	return indices, nil
}

func (app *HelloTriangleApplication) logDebug(msgType ext_debug_utils.MessageType, severity ext_debug_utils.MessageSeverity, data *ext_debug_utils.CallbackData) bool {
	log.Printf("[%s %s] - %s", severity, msgType, data.Message)
	return false
}

func main() {
	app := &HelloTriangleApplication{}

	err := app.Run()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}