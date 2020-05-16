package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	ScreenBackgroundColor utils.Color

	// Label
	LabelFont      IFont
	LabelFontSize  int
	LabelTextColor utils.Color
	LabelPadding   utils.Insets
	LabelAlignment utils.Alignment

	// Button
	ButtonFont      IFont
	ButtonFontSize  int
	ButtonTextColor utils.Color
	ButtonFillColor utils.Color
	ButtonPadding   utils.Insets
	ButtonAlignment utils.Alignment

	// Checkbox
	CheckboxFont      IFont
	CheckboxFontSize  int
	CheckboxTextColor utils.Color
	CheckboxFillColor utils.Color
	CheckboxPadding   utils.Insets
	CheckboxAlignment utils.Alignment
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}

	t.ScreenBackgroundColor = utils.NewColorGrayInt(56, 255)
	defaultFont := baseFont

	t.LabelFont = defaultFont
	t.LabelTextColor = utils.NewColorGrayInt(225, 255)
	t.LabelFontSize = 15
	t.LabelPadding = utils.HomogeneousInsets(0)
	t.LabelAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	t.ButtonFont = defaultFont
	t.ButtonFillColor = utils.NewColorGrayInt(105, 255)
	t.ButtonTextColor = utils.NewColorGrayInt(232, 255)
	t.ButtonFontSize = 15
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	t.ButtonAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	t.CheckboxFont = defaultFont
	t.CheckboxFillColor = utils.NewColorTransparent()
	t.CheckboxTextColor = utils.NewColorGrayInt(232, 255)
	t.CheckboxFontSize = 15
	t.CheckboxPadding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 25}
	t.CheckboxAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	return t
}
