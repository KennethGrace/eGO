package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	title        string
	width        int
	height       int
	vsync        bool
	resized      bool
	windowHandle *glfw.Window
}

func (w *Window) init() error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.DefaultWindowHints()
	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	w.windowHandle, err = glfw.CreateWindow(w.width, w.height, w.title, nil, nil)
	if err != nil {
		return err
	}
	w.windowHandle.SetFramebufferSizeCallback(func(_ *glfw.Window, width int, height int) {
		w.width = width
		w.height = height
		w.resized = true
	})
	w.windowHandle.SetKeyCallback(
		func(window *glfw.Window, key glfw.Key, _ int, action glfw.Action, _ glfw.ModifierKey) {
			if key == glfw.KeyEscape && action == glfw.Release {
				window.SetShouldClose(true)
			}
		})

	video := glfw.GetPrimaryMonitor().GetVideoMode()
	w.windowHandle.SetPos((video.Width-w.width)/2, (video.Height-w.height)/2)
	w.windowHandle.MakeContextCurrent()
	if w.vsync {
		glfw.SwapInterval(1)
	}
	w.windowHandle.Show()
	if err := gl.Init(); err != nil {
		return err
	}
	gl.ClearColor(0, 0, 0, 0)
	gl.Enable(gl.DEPTH_TEST)
	return nil
}

func (w *Window) update() {
	w.windowHandle.SwapBuffers()
}

func (w *Window) shouldClose() bool {
	return w.windowHandle.ShouldClose()
}
