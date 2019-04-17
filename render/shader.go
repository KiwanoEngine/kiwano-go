package render

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func CreateShader(shaderType uint32, src string) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	cstrs, free := gl.Strs(src)
	defer free()

	length := int32(len(src))
	gl.ShaderSource(shader, 1, cstrs, &length)
	gl.CompileShader(shader)

	if ok, err := GetShaderStatus(shader); !ok {
		gl.DeleteShader(shader)
		return 0, fmt.Errorf("Failed to compile shader: %v", err)
	}
	return shader, nil
}

func GetShaderStatus(shader uint32) (ok bool, log string) {
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

func LinkShaderProgram(program, vertexShader, fragmentShader uint32) error {
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	if ok, err := GetProgramStatus(program); !ok {
		gl.DeleteProgram(program)
		return fmt.Errorf("Failed to link shader program: %v", err)
	}
	return nil
}

func GetProgramStatus(program uint32) (ok bool, log string) {
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

func CreateShaderProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {

	vertexShader, err := CreateShader(gl.VERTEX_SHADER, vertexShaderSource)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(vertexShader)

	fragmentShader, err := CreateShader(gl.FRAGMENT_SHADER, fragmentShaderSource)
	if err != nil {
		return 0, err
	}
	defer gl.DeleteShader(fragmentShader)

	program := gl.CreateProgram()
	if err := LinkShaderProgram(program, vertexShader, fragmentShader); err != nil {
		return 0, err
	}
	return program, nil
}
