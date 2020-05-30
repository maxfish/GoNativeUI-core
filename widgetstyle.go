package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type WidgetStyle struct {
	Font              IFont
	FontSize          int
	BackgroundColor   utils.Color
	BorderColor       utils.Color
	Padding           utils.Insets
	ContentAlignment  utils.Alignment
	TextColor         utils.Color
	CursorColor       utils.Color
	IconColor         utils.Color
	TextDisabledColor utils.Color
	IconDisabledColor utils.Color
	SelectionColor    utils.Color
	ErrorColor        utils.Color
	// TODO Mouse cursor
}
