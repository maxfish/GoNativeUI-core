package tests

import gui "github.com/maxfish/GoNativeUI-Core"

func InitTestGui(width int, height int, theme *gui.Theme) *gui.Gui {
	if theme == nil {
		theme = gui.NewDefaultTheme()
	}
	return gui.NewGui(theme, width, height)
}
