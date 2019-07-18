package kiwano

import (
	"log"
	"runtime"
	"time"

	"kiwanoengine.com/kiwano/external/gl"
	"kiwanoengine.com/kiwano/render"

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

// Setup starts to play
func Setup(option *Option, setup func()) error {
	if err := Init(option); err != nil {
		return err
	}

	// Destroy all resources
	defer Destroy()

	// Perform user setup function
	setup()

	// Enter the main loop
	MainLoop()
	return nil
}

// Init will initialize kiwano engine
func Init(option *Option) error {
	var err error
	// Create window
	MainWindow, err = NewWindow(option)
	if err != nil {
		return err
	}

	log.Println("OpenGL version", gl.GetString(gl.VERSION))

	MainWindow.Show()
	return nil
}

// MainLoop ...
func MainLoop() {
	now := time.Now()
	last := now

	for !MainWindow.ShouldClose() {
		// render
		gl.Clear(gl.COLOR_BUFFER_BIT)

		now = time.Now()
		if CurrentScene != nil {
			CurrentScene.OnUpdate(now.Sub(last))
		}
		last = now

		// swap buffer
		MainWindow.SwapBuffers()
		glfw.PollEvents()
	}

	// Clear current scene
	EnterScene(nil)
}

// Destroy clean up engine resources
func Destroy() {
	render.DestroyAllShaders()
	MainWindow.Destroy()
	glfw.Terminate()
}

// Exit stop the main loop
func Exit() {
	if MainWindow != nil {
		MainWindow.SetShouldClose(true)
	}
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
