package main

import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/cgoalloc"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type HelloTriangleApplication struct {
	allocator cgoalloc.Allocator
	window *sdl.Window

	instance *VKng.Instance
}

func (app *HelloTriangleApplication) Run() error {
	err := app.initWindow()
	if err != nil {return err }

	err = app.initVulkan()
	if err != nil { return err }
	defer app.cleanup()

	return app.mainLoop()
}

func (app *HelloTriangleApplication) initWindow() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow("Vulkan", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	app.window = window

	return nil
}

func (app *HelloTriangleApplication) createInstance() error {
	i, err := (&VKng.InstanceBuilder{}).
		ApplicationName("Hello Triangle").
		ApplicationVersion(1, 0, 0).
		EngineName("No Engine").
		EngineVersion(1, 0, 0).
		Build(app.allocator)
	if err != nil {
		return err
	}
	app.instance = i
	return nil
}

func (app *HelloTriangleApplication) initVulkan() error {
	return app.createInstance()
}

func (app *HelloTriangleApplication) cleanup() {
	app.instance.Destroy()

	app.allocator.Destroy()
	app.window.Destroy()
	sdl.Quit()
}

func (app *HelloTriangleApplication) mainLoop() error {
	for true {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return nil
			}
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
