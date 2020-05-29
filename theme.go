package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	// Common
	BackgroundColor   utils.Color
	SelectionColor    utils.Color
	TextColor         utils.Color
	TextDisabledColor utils.Color
	TextFont          IFont
	TextFontSize      int
	// Icons
	IconColor         utils.Color
	IconDisabledColor utils.Color
	// Button
	ButtonColor         utils.Color
	ButtonDisabledColor utils.Color
	ButtonPadding       utils.Insets
	// Checkbox
	CheckboxPadding   utils.Insets
	CheckboxAlignment utils.Alignment
	// TextField
	TextFieldPlaceholderColor utils.Color
	TextFieldBackgroundColor  utils.Color
	TextFieldTextColor        utils.Color
	TextFieldSelectionColor   utils.Color
	TextFieldNotValidColor    utils.Color
	TextFieldCursorColor      utils.Color
	TextFieldPadding          utils.Insets
	// ListView
	ListViewTextColor       utils.Color
	ListViewSelectionColor  utils.Color
	ListViewBackgroundColor utils.Color
	ListViewPadding         utils.Insets
	ListViewAlignment       utils.Alignment
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}

	// Common
	t.BackgroundColor = utils.NewColorHex(0x3C3F41FF)
	t.SelectionColor = utils.NewColorHex(0x1A65D1FF)
	t.TextColor = utils.NewColorHex(0xAFB1B3FF)
	t.TextDisabledColor = t.TextColor.Scaled(0.6)
	t.TextFont = baseFont
	t.TextFontSize = 15
	// Icons
	t.IconColor = utils.NewColorHex(0xAFB1B3FF)
	t.IconDisabledColor = t.IconColor.Scaled(0.6)
	// Button
	t.ButtonColor = utils.NewColorHex(0x4B5052FF)
	t.ButtonDisabledColor = t.TextDisabledColor
	t.ButtonPadding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	// Checkbox
	t.CheckboxPadding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	t.CheckboxAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	// TextField
	t.TextFieldPlaceholderColor = utils.TransparentColor
	t.TextFieldBackgroundColor = utils.NewColorHex(0x3C3C3CFF)
	t.TextFieldTextColor = t.TextColor
	t.TextFieldSelectionColor = t.SelectionColor
	t.TextFieldCursorColor = utils.NewColorHex(0xDDDDDDFF)
	t.TextFieldNotValidColor = utils.NewColorHex(0x743A3AFF)
	t.TextFieldPadding = utils.Insets{Top: 2, Right: 4, Bottom: 2, Left: 4}
	// ListView
	t.ListViewTextColor = t.TextColor
	t.ListViewSelectionColor = t.SelectionColor
	t.ListViewBackgroundColor = utils.NewColorHex(0x4B5052FF)
	t.ListViewPadding = utils.HomogeneousInsets(4)
	t.ListViewAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}

	return t
}
