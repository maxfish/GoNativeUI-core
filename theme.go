package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type Theme struct {
	BaseStyle WidgetStyle
	Label     WidgetStyle
	Button    WidgetStyle
	Checkbox  WidgetStyle
	TextField WidgetStyle
	ImageView WidgetStyle
	ListView  WidgetStyle
}

func NewDefaultTheme(baseFont IFont) *Theme {
	t := &Theme{}
	t.BaseStyle = WidgetStyle{
		Font:              baseFont,
		FontSize:          15,
		BackgroundColor:   utils.NewColorHex(0x3C3F41FF),
		BorderColor:       utils.NewColorHex(0x59595AFF),
		Padding:           utils.Insets{},
		ContentAlignment:  utils.Alignment{},
		TextColor:         utils.NewColorHex(0xAFB1B3FF),
		CursorColor:       utils.NewColorHex(0xDDDDDDFF),
		TextDisabledColor: utils.NewColorHex(0xAFB1B3FF).Scaled(0.6),
		IconColor:         utils.NewColorHex(0xAFB1B3FF),
		IconDisabledColor: utils.NewColorHex(0xAFB1B3FF).Scaled(0.6),
		SelectionColor:    utils.NewColorHex(0x1A65D1FF),
		ErrorColor:        utils.NewColorHex(0x743A3AFF),
	}

	// Label
	t.Label = t.BaseStyle
	t.Label.BackgroundColor = utils.TransparentColor
	// Button
	t.Button = t.BaseStyle
	t.Button.BackgroundColor = utils.NewColorHex(0x4B5052FF)
	t.Button.Padding = utils.Insets{Top: 5, Right: 10, Bottom: 5, Left: 10}
	// Checkbox
	t.Checkbox = t.BaseStyle
	t.Checkbox.Padding = utils.Insets{Top: 2, Right: 2, Bottom: 2, Left: 2}
	t.Checkbox.ContentAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	// TextField
	t.TextField = t.BaseStyle
	t.TextField.Padding = utils.Insets{Top: 2, Right: 4, Bottom: 2, Left: 4}
	t.TextField.ContentAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	t.TextField.BackgroundColor = utils.NewColorHex(0x44494AFF)
	t.TextField.BorderColor = utils.NewColorHex(0x646464FF)
	// ImageView
	t.ImageView = t.BaseStyle
	t.ImageView.BackgroundColor = utils.TransparentColor
	// ListView
	t.ListView = t.BaseStyle
	t.ListView.BackgroundColor = utils.NewColorHex(0x4B5052FF)
	t.ListView.Padding = utils.HomogeneousInsets(4)
	t.ListView.ContentAlignment = utils.Alignment{Horizontal: utils.AlignmentHLeft, Vertical: utils.AlignmentVCenter}
	return t
}
