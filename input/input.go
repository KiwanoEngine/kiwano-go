package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"

	"kiwanoengine.com/kiwano"
)

// Pressed ...
func Pressed(key Key) bool {
	if kiwano.MainWindow == nil {
		return false
	}
	return kiwano.MainWindow.GetKey(glfw.Key(key)) == glfw.Press
}
