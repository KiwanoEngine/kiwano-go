package node

type Sprite struct {
	NodeProperties
	image string
}

func NewSprite(image string) *Sprite {
	return &Sprite{
		image: image,
	}
}

func (s *Sprite) OnRender() {
}
