package kiwano

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Option struct {
	Width, Height int
	Title         string
	ClearColor    Color
	NoTitleBar    bool
	Fullscreen    bool
	Resizable     bool
	Vsync         bool
}

type Window struct {
	Option
	GLFWWindow *glfw.Window
}

func NewWindow(option *Option) (*Window, error) {

	window := &Window{
		Option: *option,
	}

	// Init GLFW
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Visible, glfw.False)

	if option.NoTitleBar {
		glfw.WindowHint(glfw.Decorated, glfw.False)
	} else {
		glfw.WindowHint(glfw.Decorated, glfw.True)
	}

	if option.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	// Get monitor information
	monitor := glfw.GetPrimaryMonitor()

	var monitorMode *glfw.VidMode
	if monitor != nil {
		monitorMode = monitor.GetVideoMode()
	}

	// Fullscreen mode
	if option.Fullscreen && monitorMode != nil {
		option.Width, option.Height = monitorMode.Width, monitorMode.Height
	} else {
		monitor = nil
	}

	// Create window
	w, err := glfw.CreateWindow(option.Width, option.Height, option.Title, monitor, nil)
	if err != nil {
		return nil, err
	}

	if !option.Fullscreen && monitorMode != nil {
		w.SetPos((monitorMode.Width-option.Width)/2, (monitorMode.Height-option.Height)/2)
	}

	w.MakeContextCurrent()
	w.SetFramebufferSizeCallback(window.onFramebufferSizeCallback)

	if option.Vsync {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	// Init OpenGL
	if err := gl.Init(); err != nil {
		return nil, err
	}

	gl.ClearColor(option.ClearColor.ToVec4())

	window.GLFWWindow = w
	return window, nil
}

func (w *Window) Destroy() {
	w.GLFWWindow.Destroy()
}

func (w *Window) onFramebufferSizeCallback(win *glfw.Window, width int, height int) {
	w.Width, w.Height = width, height
	gl.Viewport(0, 0, int32(width), int32(height))
}
