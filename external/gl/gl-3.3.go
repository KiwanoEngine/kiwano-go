package gl

import (
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Init initialize OpenGL
func Init() error {
	return gl.Init()
}

// Ptr takes a slice or pointer and returns its GL-compatible address.
func Ptr(data interface{}) unsafe.Pointer {
	return gl.Ptr(data)
}

// PtrOffset takes a pointer offset and returns a GL-compatible pointer.
func PtrOffset(offset int) unsafe.Pointer {
	return gl.PtrOffset(offset)
}

// Str takes a null-terminated Go string and returns its GL-compatible address.
func Str(str string) *uint8 {
	return gl.Str(str)
}

// GoStr takes a null-terminated string returned by OpenGL and constructs a
// corresponding Go string.
func GoStr(cstr *uint8) string {
	return gl.GoStr(cstr)
}

// Strs takes a list of Go strings (with or without null-termination) and
// returns their C counterpart.
func Strs(strs ...string) (cstrs **uint8, free func()) {
	return gl.Strs(strs...)
}

// GetString return a string describing the current GL connection
func GetString(name uint32) string {
	return gl.GoStr(gl.GetString(name))
}

// ClearColor specify clear values for the color buffers
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.ClearColor(red, green, blue, alpha)
}

// Clear clear buffers to preset values
func Clear(mask uint32) {
	gl.Clear(mask)
}

// Viewport set the viewport
func Viewport(x int32, y int32, width int32, height int32) {
	gl.Viewport(x, y, width, height)
}

// CreateProgram creates a program object
func CreateProgram() uint32 {
	return gl.CreateProgram()
}

// LinkProgram links a program object
func LinkProgram(program uint32) {
	gl.LinkProgram(program)
}

// UseProgram installs a program object as part of current rendering state
func UseProgram(program uint32) {
	gl.UseProgram(program)
}

// DeleteProgram deletes a program object
func DeleteProgram(program uint32) {
	gl.DeleteProgram(program)
}

// GetProgram returns a parameter from a program object
func GetProgram(program uint32, pname uint32) int32 {
	var params int32
	gl.GetProgramiv(program, pname, &params)
	return params
}

// GetProgramInfoLog returns the information log for a program object
func GetProgramInfoLog(program uint32) string {
	logLength := GetProgram(program, INFO_LOG_LENGTH)
	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))
	return log
}

// CreateShader creates a shader object
func CreateShader(xtype uint32) uint32 {
	return gl.CreateShader(xtype)
}

// CompileShader compiles a shader object
func CompileShader(shader uint32) {
	gl.CompileShader(shader)
}

// AttachShader attaches a shader object to a program object
func AttachShader(program uint32, shader uint32) {
	gl.AttachShader(program, shader)
}

// DeleteShader deletes a shader object
func DeleteShader(shader uint32) {
	gl.DeleteShader(shader)
}

// ShaderSource replaces the source code in a shader object
func ShaderSource(shader uint32, xstring string) {
	cstrs, free := gl.Strs(xstring)
	defer free()

	length := int32(len(xstring))
	gl.ShaderSource(shader, 1, cstrs, &length)
}

// GetShader returns a parameter from a shader object
func GetShader(shader uint32, pname uint32) int32 {
	var params int32
	gl.GetShaderiv(shader, pname, &params)
	return params
}

// GetShaderInfoLog returns the information log for a shader object
func GetShaderInfoLog(shader uint32) string {
	logLength := GetShader(shader, INFO_LOG_LENGTH)
	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
	return log
}

// GenBuffers generate buffer object names
func GenBuffers(n int32, buffers *uint32) {
	gl.GenBuffers(n, buffers)
}

// BindBuffer bind a named buffer object
func BindBuffer(target uint32, buffer uint32) {
	gl.BindBuffer(target, buffer)
}

// BufferData creates and initializes a buffer object's data
func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	gl.BufferData(target, size, data, usage)
}

// BindVertexBuffer bind a buffer to a vertex buffer bind point
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int, stride int32) {
	gl.BindVertexBuffer(bindingindex, buffer, offset, stride)
}

// DeleteBuffers delete named buffer objects
func DeleteBuffers(n int32, buffers *uint32) {
	gl.DeleteBuffers(n, buffers)
}

// GenVertexArrays generate vertex array object names
func GenVertexArrays(n int32, arrays *uint32) {
	gl.GenVertexArrays(n, arrays)
}

// BindVertexArray bind a vertex array object
func BindVertexArray(array uint32) {
	gl.BindVertexArray(array)
}

// DeleteVertexArrays delete vertex array objects
func DeleteVertexArrays(n int32, arrays *uint32) {
	gl.DeleteVertexArrays(n, arrays)
}

// VertexAttribPointer define an array of generic vertex attribute data
func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	gl.VertexAttribPointer(index, size, xtype, normalized, stride, pointer)
}

// EnableVertexAttribArray enable or disable a generic vertex attribute
func EnableVertexAttribArray(index uint32) {
	gl.EnableVertexAttribArray(index)
}

// GenTextures generate texture names
func GenTextures(n int32, textures *uint32) {
	gl.GenTextures(n, textures)
}

// ActiveTexture select active texture unit
func ActiveTexture(texture uint32) {
	gl.ActiveTexture(texture)
}

// BindTexture bind a named texture to a texturing target
func BindTexture(target uint32, texture uint32) {
	gl.BindTexture(target, texture)
}

// TexParameteri establish the data storage, format, dimensions,
// and number of samples of a multisample texture's image
func TexParameteri(target uint32, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

// TexImage2D specify a two-dimensional texture image
func TexImage2D(target uint32, level int32, internalformat int32, width int32,
	height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(target, level, internalformat, width, height, border, format, xtype, pixels)
}

// GenerateMipmap generate mipmaps for a specified texture object
func GenerateMipmap(target uint32) {
	gl.GenerateMipmap(target)
}

// GetUniformLocation returns the location of a uniform variable
func GetUniformLocation(program uint32, name string) int32 {
	return gl.GetUniformLocation(program, gl.Str(name))
}

// Uniform1i specify the value of a uniform variable for the current program object
func Uniform1i(location int32, v0 int32) {
	gl.Uniform1i(location, v0)
}

// Uniform2i ...
func Uniform2i(location int32, v0, v1 int32) {
	gl.Uniform2i(location, v0, v1)
}

// Uniform3i ...
func Uniform3i(location int32, v0, v1, v2 int32) {
	gl.Uniform3i(location, v0, v1, v2)
}

// Uniform4i ...
func Uniform4i(location int32, v0, v1, v2, v3 int32) {
	gl.Uniform4i(location, v0, v1, v2, v3)
}

// Uniform1f ...
func Uniform1f(location int32, v0 float32) {
	gl.Uniform1f(location, v0)
}

// Uniform2f ...
func Uniform2f(location int32, v0, v1 float32) {
	gl.Uniform2f(location, v0, v1)
}

// Uniform3f ...
func Uniform3f(location int32, v0, v1, v2 float32) {
	gl.Uniform3f(location, v0, v1, v2)
}

// Uniform4f ...
func Uniform4f(location int32, v0, v1, v2, v3 float32) {
	gl.Uniform4f(location, v0, v1, v2, v3)
}

// DrawElements render primitives from array data
func DrawElements(mode uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	gl.DrawElements(mode, count, xtype, indices)
}
