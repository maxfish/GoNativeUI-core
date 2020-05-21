package gui

import (
	"image"
)

type ImageView struct {
	Widget
	image image.Image
}

func NewImageView(image image.Image) *ImageView {
	s := &ImageView{}
	widgetInit(s)
	s.image = image
	return s
}

func (i *ImageView) Measure() {
	if i.image != nil {
		i.measuredWidth = i.image.Bounds().Dx()
		i.measuredHeight = i.image.Bounds().Dy()
	}

	i.measuredFlex = i.flex
}

func (i *ImageView) Image() image.Image {
	return i.image
}

func (i *ImageView) SetImage(image image.Image) {
	i.image = image
	i.Measure()
}
