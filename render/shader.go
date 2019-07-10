package render

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	shaders map[uint32]*Shader
)

// CreateShader ...
func CreateShader(vertexShaderSource, fragmentShaderSource string) (shader *Shader, err error) {
	var (
		vertexShader   uint32
		fragmentShader uint32
		shaderProgram  uint32
	)

	if vertexShader, err = createShader(gl.VERTEX_SHADER, vertexShaderSource); err != nil {
		return
	}
	defer gl.DeleteShader(vertexShader)

	if fragmentShader, err = createShader(gl.FRAGMENT_SHADER, fragmentShaderSource); err != nil {
		return
	}
	defer gl.DeleteShader(fragmentShader)

	shaderProgram = gl.CreateProgram()
	if err = linkShaderProgram(shaderProgram, vertexShader, fragmentShader); err != nil {
		return
	}
	shader = &Shader{shaderProgram}
	saveShader(shader.ID, shader)
	return
}

// DeleteAllShaders ...
func DeleteAllShaders() {
	for _, s := range shaders {
		s.Delete()
	}
}

func saveShader(ID uint32, shader *Shader) {
	if shaders == nil {
		shaders = make(map[uint32]*Shader)
	}
	shaders[shader.ID] = shader
}

// Shader ...
type Shader struct {
	ID uint32
}

// Use activate the shader
func (s *Shader) Use() {
	gl.UseProgram(s.ID)
}

// Delete delete the shader
func (s *Shader) Delete() {
	gl.DeleteProgram(s.ID)
}

// SetInt ...
func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(s.ID, gl.Str(name)), value)
}

// SetInt2 ...
func (s *Shader) SetInt2(name string, v0, v1 int32) {
	gl.Uniform2i(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1)
}

// SetInt3 ...
func (s *Shader) SetInt3(name string, v0, v1, v2 int32) {
	gl.Uniform3i(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1, v2)
}

// SetInt4 ...
func (s *Shader) SetInt4(name string, v0, v1, v2, v3 int32) {
	gl.Uniform4i(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1, v2, v3)
}

// SetFloat ...
func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(s.ID, gl.Str(name)), value)
}

// SetFloat2 ...
func (s *Shader) SetFloat2(name string, v0, v1 float32) {
	gl.Uniform2f(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1)
}

// SetFloat3 ...
func (s *Shader) SetFloat3(name string, v0, v1, v2 float32) {
	gl.Uniform3f(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1, v2)
}

// SetFloat4 ...
func (s *Shader) SetFloat4(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(gl.GetUniformLocation(s.ID, gl.Str(name)), v0, v1, v2, v3)
}

func createShader(shaderType uint32, src string) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cstrs, free := gl.Strs(src)
	defer free()

	length := int32(len(src))
	gl.ShaderSource(shader, 1, cstrs, &length)
	gl.CompileShader(shader)

	if ok, err := getShaderStatus(shader); !ok {
		gl.DeleteShader(shader)
		return 0, fmt.Errorf("Failed to compile shader: %v", err)
	}
	return shader, nil
}

func getShaderStatus(shader uint32) (ok bool, log string) {
	var success int32
	if gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success); success != gl.TRUE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log = strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return
	}
	ok = true
	return
}

func linkShaderProgram(program, vertexShader, fragmentShader uint32) error {
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	if ok, err := getProgramStatus(program); !ok {
		gl.DeleteProgram(program)
		return fmt.Errorf("Failed to link shader program: %v", err)
	}
	return nil
}

func getProgramStatus(program uint32) (ok bool, log string) {
	var success int32
	if gl.GetProgramiv(program, gl.LINK_STATUS, &success); success != gl.TRUE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log = strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
		return
	}
	ok = true
	return
}
