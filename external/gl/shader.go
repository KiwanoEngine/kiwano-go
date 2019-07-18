package gl

import (
	"fmt"
)

func CreateAndCompileShader(shaderType uint32, src string) (uint32, error) {
	shader := CreateShader(shaderType)
	ShaderSource(shader, src)
	CompileShader(shader)

	if ret := GetShader(shader, COMPILE_STATUS); ret != TRUE {
		log := GetShaderInfoLog(shader)
		return 0, fmt.Errorf("Failed to compile %v: %v", src, log)
	}
	return shader, nil
}

func LinkShaderProgram(program, vertexShader, fragmentShader uint32) error {
	AttachShader(program, vertexShader)
	AttachShader(program, fragmentShader)
	LinkProgram(program)

	if ret := GetProgram(program, LINK_STATUS); ret != TRUE {
		log := GetProgramInfoLog(program)
		return fmt.Errorf("Failed to link shader program: %v", log)
	}
	return nil
}
