package render

import (
	"kiwanoengine.com/kiwano/external/gl"
)

var (
	shaders map[uint32]*Shader
)

// CreateShader ...
func CreateShader(vertexShaderSource, fragmentShaderSource string) (*Shader, error) {
	var (
		vertexShader   uint32
		fragmentShader uint32
		shaderProgram  uint32
		err            error
	)

	if vertexShader, err = gl.CreateAndCompileShader(gl.VERTEX_SHADER, vertexShaderSource); err != nil {
		return nil, err
	}
	defer gl.DeleteShader(vertexShader)

	if fragmentShader, err = gl.CreateAndCompileShader(gl.FRAGMENT_SHADER, fragmentShaderSource); err != nil {
		return nil, err
	}
	defer gl.DeleteShader(fragmentShader)

	shaderProgram = gl.CreateProgram()
	if err = gl.LinkShaderProgram(shaderProgram, vertexShader, fragmentShader); err != nil {
		return nil, err
	}

	shader := &Shader{shaderProgram}
	saveShader(shader.ID, shader)
	return shader, nil
}

// DestroyAllShaders ...
func DestroyAllShaders() {
	if shaders != nil {
		for _, s := range shaders {
			s.Destroy()
		}
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

// Destroy delete the shader
func (s *Shader) Destroy() {
	gl.DeleteProgram(s.ID)
}

// SetInt ...
func (s *Shader) SetInt(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(s.ID, name), value)
}

// SetInt2 ...
func (s *Shader) SetInt2(name string, v0, v1 int32) {
	gl.Uniform2i(gl.GetUniformLocation(s.ID, name), v0, v1)
}

// SetInt3 ...
func (s *Shader) SetInt3(name string, v0, v1, v2 int32) {
	gl.Uniform3i(gl.GetUniformLocation(s.ID, name), v0, v1, v2)
}

// SetInt4 ...
func (s *Shader) SetInt4(name string, v0, v1, v2, v3 int32) {
	gl.Uniform4i(gl.GetUniformLocation(s.ID, name), v0, v1, v2, v3)
}

// SetFloat ...
func (s *Shader) SetFloat(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(s.ID, name), value)
}

// SetFloat2 ...
func (s *Shader) SetFloat2(name string, v0, v1 float32) {
	gl.Uniform2f(gl.GetUniformLocation(s.ID, name), v0, v1)
}

// SetFloat3 ...
func (s *Shader) SetFloat3(name string, v0, v1, v2 float32) {
	gl.Uniform3f(gl.GetUniformLocation(s.ID, name), v0, v1, v2)
}

// SetFloat4 ...
func (s *Shader) SetFloat4(name string, v0, v1, v2, v3 float32) {
	gl.Uniform4f(gl.GetUniformLocation(s.ID, name), v0, v1, v2, v3)
}
