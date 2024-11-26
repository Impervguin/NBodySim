package conveyer

import "image"

type RenderConveyer interface {
	Convey() error
	GetImage() image.Image
}
