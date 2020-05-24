package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	BackgroundColor   utils.Color
	TextColor         utils.Color
	TextDisabledColor utils.Color
	TextFont          IFont
	TextFontSize      int

	ButtonColor         utils.Color
	ButtonDisabledColor utils.Color
	ButtonPadding       utils.Insets

	IconColor         utils.Color
	IconDisabledColor utils.Color

	InputPlaceholderColor utils.Color
	InputBackgroundColor  utils.Color
	InputTextColor        utils.Color
	InputSelectionColor   utils.Color
	InputNotValidColor    utils.Color
	InputCursorColor      utils.Color
	InputPadding          utils.Insets

	CheckboxPadding   utils.Insets
	CheckboxAlignment utils.Alignment
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{
		BackgroundColor:       utils.NewColorHex(0x3C3F41FF),
		TextColor:             utils.NewColorHex(0xAFB1B3FF),
		TextDisabledColor:     utils.TransparentColor,
		TextFont:              baseFont,
		TextFontSize:          15,
		ButtonColor:           utils.NewColorHex(0x4B5052FF),
		ButtonDisabledColor:   utils.TransparentColor,
		ButtonPadding:         utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10},
		IconColor:             utils.NewColorHex(0xAFB1B3FF),
		IconDisabledColor:     utils.TransparentColor,
		InputPlaceholderColor: utils.TransparentColor,
		InputBackgroundColor:  utils.NewColorHex(0x3C3C3CFF),
		InputTextColor:        utils.TransparentColor,
		InputSelectionColor:   utils.NewColorHex(0x164288FF),
		InputCursorColor:      utils.NewColorHex(0xDDDDDDFF),
		InputNotValidColor:    utils.NewColorHex(0x743A3AFF),
		InputPadding:          utils.Insets{Top: 2, Right: 4, Bottom: 2, Left: 4},
		CheckboxPadding:       utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2},
		CheckboxAlignment:     utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter},
	}
	//t.ButtonAlignment = utils.Alignment{Horizontal: utils.AlignmentHCenter, Vertical: utils.AlignmentVCenter}
	//t.InputFieldAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}

	return t
}
