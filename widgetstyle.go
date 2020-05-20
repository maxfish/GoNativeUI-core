package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

//type IWidgetStyle interface {
//	Theme() *Theme
//	SetTheme(t *Theme)
//	Font() IFont
//	FontSize() int
//	TextColor() utils.Color
//	BackgroundColor() utils.Color
//	SetBackgroundColor(color utils.Color)
//	Padding() utils.Insets
//	SetPadding(b utils.Insets)
//}

type WidgetStyle struct {
	//theme            *Theme
	Font             IFont
	FontSize         int
	TextColor        utils.Color
	BackgroundColor  utils.Color
	Padding          utils.Insets
	ContentAlignment utils.Alignment

	// TODO Mouse cursor
}

// Getters
//func (s *WidgetStyle) Theme() *Theme { return s.theme }

//func (s *WidgetStyle) Font() IFont                  { return s.font }
//func (s *WidgetStyle) FontSize() int                { return s.fontSize }
//func (s *WidgetStyle) TextColor() utils.Color       { return s.textColor }
//func (s *WidgetStyle) BackgroundColor() utils.Color { return s.backgroundColor }
//func (s *WidgetStyle) Padding() utils.Insets     { return s.padding }
//
//func (s *WidgetStyle) SetBackgroundColor(color utils.Color) {
//	s.backgroundColor = color
//}
//func (s *WidgetStyle) SetPadding(b utils.Insets) { s.padding = b }
