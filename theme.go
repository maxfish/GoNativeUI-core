package gui

import (
	"github.com/maxfish/GoNativeUI-Core/assets"
	"github.com/maxfish/GoNativeUI-Core/utils"
)

const defaultAssetsPath = "github.com/maxfish/GoNativeUI-Core/assets/"

type Theme struct {
	ScreenBackgroundColor utils.ColorF

	// Label
	LabelFont      IFont
	LabelFontSize  int
	LabelTextColor utils.ColorF
	LabelPadding   utils.Insets

	// Button
	ButtonFont      IFont
	ButtonFontSize  int
	ButtonTextColor utils.ColorF
	ButtonFillColor utils.ColorF
	ButtonPadding   utils.Insets
	ButtonAlignment Alignment
}

func NewDefaultTheme() *Theme {
	t := &Theme{}

	t.ScreenBackgroundColor = utils.ColorGrayi(56, 255)

	//defaultFont := NewFontFromFile("assets/Roboto-Regular.fnt")
	defaultFont := NewFontFromData(assets.FontRobotoRegularDefinition, []string{assets.FontRobotoRegularImage})

	t.LabelFont = defaultFont
	t.LabelTextColor = utils.ColorGrayi(225, 255)
	t.LabelFontSize = 15
	t.LabelPadding = utils.HomogeneousInsets(0)

	t.ButtonFont = defaultFont
	t.ButtonFillColor = utils.ColorGrayi(105, 255)
	t.ButtonTextColor = utils.ColorGrayi(232, 255)
	t.ButtonFontSize = 15
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	t.ButtonAlignment = Alignment{AlignmentHCenter, AlignmentVCenter}
	return t
}
