package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type WidgetStyle struct {
	Font             IFont
	FontSize         int
	TextColor        utils.Color
	BackgroundColor  utils.Color
	Padding          utils.Insets
	ContentAlignment utils.Alignment

	// TODO Mouse cursor
}
