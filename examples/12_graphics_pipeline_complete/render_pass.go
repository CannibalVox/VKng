package main

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/render_pass"
)

func (app *HelloTriangleApplication) createRenderPass() error {
	renderPass, err := render_pass.CreateRenderPass(app.allocator, app.logicalDevice, &render_pass.RenderPassOptions{
		Attachments: []render_pass.AttachmentDescription{
			{
				Format:         app.swapchainFormat.Format,
				Samples:        core.Samples1,
				LoadOp:         core.LoadOpClear,
				StoreOp:        core.StoreOpStore,
				StencilLoadOp:  core.LoadOpDontCare,
				StencilStoreOp: core.StoreOpDontCare,
				InitialLayout:  core.LayoutUndefined,
				FinalLayout:    core.LayoutPresentSrc,
			},
		},
		SubPasses: []render_pass.SubPass{
			{
				BindPoint: core.BindGraphics,
				ColorAttachments: []core.AttachmentReference{
					{
						AttachmentIndex: 0,
						Layout:          core.LayoutColorAttachmentOptimal,
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	app.renderPass = renderPass

	return nil
}
