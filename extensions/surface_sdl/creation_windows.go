//go:build windows
// +build windows

package ext_surface_sdl2

/*
#include <stdlib.h>
#include <windows.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_win32.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"github.com/cockroachdb/errors"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

func (o *CreationOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkWin32SurfaceCreateInfoKHR)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkWin32SurfaceCreateInfoKHR{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR
	createInfo.flags = 0

	sysInfo, err := o.Window.GetWMInfo()
	if err != nil {
		return nil, err
	}

	if sysInfo.Subsystem != sdl.SYSWM_WINDOWS {
		return nil, errors.Newf("Unexpected window subsystems in windows OS: %v", sysInfo.Subsystem)
	}

	winInfo := sysInfo.GetWindowsInfo()
	createInfo.hinstance = (C.HINSTANCE)(winInfo.Instance)
	createInfo.hwnd = (C.HWND)(winInfo.Window)

	var next unsafe.Pointer

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), err
}
