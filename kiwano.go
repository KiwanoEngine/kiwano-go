package kiwano

import (
	"log"
	"runtime"
	"time"

	"kiwanoengine.com/kiwano/render"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Environment variables
const (
	Version     = "0.1"
	VersionCode = 1
)

// Global variables
var (
	MainWindow   *Window
	CurrentScene Scene
)

func init() {
	runtime.LockOSThread()
}

// Init will initialize kiwano engine
func Init(option *Option) error {
	var err error
	// Create window
	MainWindow, err = NewWindow(option)
	if err != nil {
		return err
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	MainWindow.GLFWWindow.Show()
	return nil
}

// Run starts to play
func Run() {
	now := time.Now()
	last := now

	for !MainWindow.GLFWWindow.ShouldClose() {
		// render
		gl.Clear(gl.COLOR_BUFFER_BIT)

		now = time.Now()
		if CurrentScene != nil {
			CurrentScene.OnUpdate(now.Sub(last))
		}
		last = now

		// swap buffer
		MainWindow.GLFWWindow.SwapBuffers()
		glfw.PollEvents()
	}

	// Clear current scene
	EnterScene(nil)
}

// Destroy clean up engine resources
func Destroy() {
	render.DeleteAllShaders()
	MainWindow.Destroy()
	glfw.Terminate()
}

// EnterScene exits current scene and enters a new scene
func EnterScene(scene Scene) {
	if CurrentScene != nil {
		CurrentScene.OnExit()
	}

	CurrentScene = scene
	if CurrentScene != nil {
		CurrentScene.OnEnter()
	}
}
