package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	ContainerBackgroundColor utils.Color
	BaseFont                 IFont
	BaseFontSize             int
	BaseTextColor            utils.Color
	BaseBackgroundColor      utils.Color

	// Label
	LabelFont            IFont
	LabelFontSize        int
	LabelTextColor       utils.Color
	LabelBackgroundColor utils.Color
	LabelPadding         utils.Insets
	LabelAlignment       utils.Alignment

	// Button / ToggleButton
	ButtonFont            IFont
	ButtonFontSize        int
	ButtonTextColor       utils.Color
	ButtonBackgroundColor utils.Color
	ButtonPadding         utils.Insets
	ButtonAlignment       utils.Alignment

	// Checkbox
	CheckboxFont            IFont
	CheckboxFontSize        int
	CheckboxTextColor       utils.Color
	CheckboxBackgroundColor utils.Color
	CheckboxPadding         utils.Insets
	CheckboxAlignment       utils.Alignment

	// TextField
	InputFieldFont            IFont
	InputFieldFontSize        int
	InputFieldTextColor       utils.Color
	InputFieldBackgroundColor utils.Color
	InputFieldPadding         utils.Insets
	InputFieldAlignment       utils.Alignment
	InputFieldSelectionColor  utils.Color
	InputFieldCursorColor     utils.Color
	InputFieldNotValidColor   utils.Color
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}

	// Common
	t.ContainerBackgroundColor = utils.NewColorHex(0x3B3F41FF)
	t.BaseFont = baseFont
	t.BaseFontSize = 15
	t.BaseTextColor = utils.NewColorHex(0xAFB1B3FF)
	t.BaseBackgroundColor = utils.NewColorHex(0x4B5052FF)

	// Label
	t.LabelFont = t.BaseFont
	t.LabelBackgroundColor = utils.NewColorTransparent()
	t.LabelTextColor = t.BaseTextColor
	t.LabelFontSize = t.BaseFontSize
	t.LabelPadding = utils.HomogeneousInsets(0)
	t.LabelAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	// Button / ToggleButton
	t.ButtonFont = t.BaseFont
	t.ButtonBackgroundColor = t.BaseBackgroundColor
	t.ButtonTextColor = t.BaseTextColor
	t.ButtonFontSize = t.BaseFontSize
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	t.ButtonAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}

	// Checkbox
	t.CheckboxFont = t.BaseFont
	t.CheckboxBackgroundColor = utils.NewColorTransparent()
	t.CheckboxTextColor = t.BaseTextColor
	t.CheckboxFontSize = t.BaseFontSize
	t.CheckboxPadding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	t.CheckboxAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}

	// TextField
	t.InputFieldFont = t.BaseFont
	t.InputFieldFontSize = t.BaseFontSize
	t.InputFieldTextColor = t.BaseTextColor
	t.InputFieldBackgroundColor = utils.NewColorHex(0x45494AFF)
	t.InputFieldPadding = utils.Insets{Top: 4, Right: 5, Bottom: 4, Left: 5}
	t.InputFieldAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	t.InputFieldSelectionColor = utils.NewColorHex(0x164288FF)
	t.InputFieldCursorColor = utils.NewColorHex(0xBBBBBBFF)
	t.InputFieldNotValidColor = utils.NewColorHex(0x743A3AFF)

	return t
}
