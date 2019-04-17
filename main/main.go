package main

import (
	"log"
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"kiwanoengine.com/kiwano"
	"kiwanoengine.com/kiwano/core"
	"kiwanoengine.com/kiwano/render"
)

var vertexShaderSource = `
#version 330 core
layout (location = 0) in vec3 aPos;

void main()
{
    gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}
`

var fragmentShaderSource = `
#version 330 core
out vec4 FragColor;

void main()
{
    FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}
`

var vertices = []float32{
	-0.5, -0.5, 0.0, // left
	0.5, -0.5, 0.0, // right
	0.0, 0.5, 0.0, // top
}

type MainScene struct {
	VAO           uint32
	VBO           uint32
	shaderProgram uint32
}

func (s *MainScene) OnEnter() {
	// Create shader program
	var err error
	s.shaderProgram, err = render.CreateShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		log.Fatalln(err)
	}

	// Set up vertex array
	gl.GenVertexArrays(1, &s.VAO)
	gl.BindVertexArray(s.VAO)

	// Set up vertex buffer
	gl.GenBuffers(1, &s.VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func (s *MainScene) OnUpdate(dt time.Duration) {
	if kiwano.Window.GLFWWindow.GetKey(glfw.KeyEscape) == glfw.Press {
		kiwano.Window.GLFWWindow.SetShouldClose(true)
	}

	gl.UseProgram(s.shaderProgram)
	gl.BindVertexArray(s.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func (s *MainScene) OnExit() {
	gl.DeleteProgram(s.shaderProgram)
	gl.DeleteVertexArrays(1, &s.VAO)
	gl.DeleteBuffers(1, &s.VBO)
}

func main() {
	// Init kiwano engine
	if err := kiwano.Init(&core.Option{
		Width:           640,
		Height:          480,
		Title:           "LearnOpenGL",
		BackgroundColor: core.ColorRGB(0.2, 0.3, 0.3),
		NoTitleBar:      false,
		Resizable:       true,
		Fullscreen:      false,
		Vsync:           true,
	}); err != nil {
		log.Fatalln(err)
	}
	defer kiwano.Destroy()

	// Enter scene
	kiwano.EnterScene(&MainScene{})

	// Start game
	kiwano.Run()
}
