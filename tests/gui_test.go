package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"testing"
)

const (
	screenW = 800
	screenH = 500
	text1   = "LabelText"
	text2   = "LabelText_2"
	text3   = "Longer text to test"
	text4   = "Another text"
)

func assertIntEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("expected: %d\nreceived: %d", expected, value)
	}
}

func assertStringEqual(t *testing.T, value string, expected string) {
	if value != expected {
		t.Errorf("expected: %s\nreceived: %s", expected, value)
	}
}

func InitTestGui(width int, height int, theme *gui.Theme) *gui.Gui {
	if theme == nil {
		theme = gui.NewDefaultTheme()
	}
	return gui.NewGui(theme, width, height)
}
