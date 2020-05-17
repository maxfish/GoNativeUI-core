package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	ScreenBackgroundColor utils.Color
	BaseFont              IFont
	BaseFontSize          int
	BaseTextColor         utils.Color
	BaseFillColor         utils.Color

	// Label
	LabelFont      IFont
	LabelFontSize  int
	LabelTextColor utils.Color
	LabelFillColor utils.Color
	LabelPadding   utils.Insets
	LabelAlignment utils.Alignment

	// Button / ToggleButton
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

	// InputField
	InputFieldFont           IFont
	InputFieldFontSize       int
	InputFieldTextColor      utils.Color
	InputFieldFillColor      utils.Color
	InputFieldPadding        utils.Insets
	InputFieldAlignment      utils.Alignment
	InputFieldSelectionColor utils.Color
	InputFieldCursorColor    utils.Color
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}

	// Common
	t.ScreenBackgroundColor = utils.NewColorHex(0x3B3F41FF)
	t.BaseFont = baseFont
	t.BaseFontSize = 15
	t.BaseTextColor = utils.NewColorHex(0xAFB1B3FF)
	t.BaseFillColor = utils.NewColorHex(0x4B5052FF)

	// Label
	t.LabelFont = t.BaseFont
	t.LabelFillColor = utils.NewColorTransparent()
	t.LabelTextColor = t.BaseTextColor
	t.LabelFontSize = t.BaseFontSize
	t.LabelPadding = utils.HomogeneousInsets(0)
	t.LabelAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	// Button / ToggleButton
	t.ButtonFont = t.BaseFont
	t.ButtonFillColor = t.BaseFillColor
	t.ButtonTextColor = t.BaseTextColor
	t.ButtonFontSize = t.BaseFontSize
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	t.ButtonAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	// Checkbox
	t.CheckboxFont = t.BaseFont
	t.CheckboxFillColor = utils.NewColorTransparent()
	t.CheckboxTextColor = t.BaseTextColor
	t.CheckboxFontSize = t.BaseFontSize
	t.CheckboxPadding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	t.CheckboxAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}

	// InputField
	t.InputFieldFont = t.BaseFont
	t.InputFieldFontSize = t.BaseFontSize
	t.InputFieldTextColor = t.BaseTextColor
	t.InputFieldFillColor = t.BaseFillColor
	t.InputFieldPadding = utils.Insets{Top: 4, Right: 5, Bottom: 4, Left: 5}
	t.InputFieldAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	t.InputFieldSelectionColor = utils.NewColorHex(0x164288FF)
	t.InputFieldCursorColor = utils.NewColorHex(0xBBBBBBFF)

	return t
}
