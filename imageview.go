package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
	"image"
)

type ImageView struct {
	Widget
	image        image.Image
	contentMode  utils.FitMode
	contentScale float32
}

func NewImageView(image image.Image) *ImageView {
	s := &ImageView{}
	widgetInit(s)
	s.contentMode = utils.FitModeAlign
	s.contentScale = 1
	s.image = image
	return s
}

func (i *ImageView) initStyle() {
	//t := CurrentGui().Theme()
	i.style = &WidgetStyle{
		BackgroundColor:  utils.TransparentColor,
		ContentAlignment: utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter},
	}
}

func (i *ImageView) Measure() {
	if i.image != nil {
		i.measuredWidth = i.image.Bounds().Dx()
		i.measuredHeight = i.image.Bounds().Dy()
	}

	i.measuredFlex = i.flex
}

func (i *ImageView) Image() image.Image            { return i.image }
func (i *ImageView) FitMode() utils.FitMode        { return i.contentMode }
func (i *ImageView) SetFitMode(c utils.FitMode)    { i.contentMode = c }
func (i *ImageView) ContentScale() float32         { return i.contentScale }
func (i *ImageView) SetContentScale(scale float32) { i.contentScale = scale }

func (i *ImageView) SetImage(image image.Image) {
	i.image = image
	i.Measure()
}
