package gui

import (
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
	LabelAlignment utils.Alignment

	// Button
	ButtonFont      IFont
	ButtonFontSize  int
	ButtonTextColor utils.ColorF
	ButtonFillColor utils.ColorF
	ButtonPadding   utils.Insets
	ButtonAlignment utils.Alignment

	// Checkbox
	CheckboxFont      IFont
	CheckboxFontSize  int
	CheckboxTextColor utils.ColorF
	CheckboxFillColor utils.ColorF
	CheckboxPadding   utils.Insets
	CheckboxAlignment utils.Alignment
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}

	t.ScreenBackgroundColor = utils.ColorGrayi(56, 255)
	defaultFont := baseFont//NewFontFromData(assets.FontRobotoRegularDefinition, []string{assets.FontRobotoRegularImage})

	t.LabelFont = defaultFont
	t.LabelTextColor = utils.ColorGrayi(225, 255)
	t.LabelFontSize = 15
	t.LabelPadding = utils.HomogeneousInsets(0)
	t.LabelAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	t.ButtonFont = defaultFont
	t.ButtonFillColor = utils.ColorGrayi(105, 255)
	t.ButtonTextColor = utils.ColorGrayi(232, 255)
	t.ButtonFontSize = 15
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	t.ButtonAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	t.CheckboxFont = defaultFont
	t.CheckboxFillColor = utils.ColorTransparent()
	t.CheckboxTextColor = utils.ColorGrayi(232, 255)
	t.CheckboxFontSize = 15
	t.CheckboxPadding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 25}
	t.CheckboxAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	return t
}
