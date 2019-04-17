package node

type Node interface {
	OnRender()
}

type NodeProperties struct {
	Position struct {
		X, Y float32
	}
	Anchor struct {
		X, Y float32
	}
}
