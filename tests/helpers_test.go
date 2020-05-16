package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"reflect"
	"testing"
)

const (
	screenWidth    = 800
	screenHeight   = 500
	fontCharWidth  = 10
	fontCharHeight = 16
)

var textStrings = [4]string{"Text 1", "Text 2 $$", "Longer test text 3", "Another text string!"}
var textLengths = [4]utils.Size{
	{len(textStrings[0]) * fontCharWidth, fontCharHeight},
	{len(textStrings[1]) * fontCharWidth, fontCharHeight},
	{len(textStrings[2]) * fontCharWidth, fontCharHeight},
	{len(textStrings[3]) * fontCharWidth, fontCharHeight},
}

func assertIntEqual(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("expected: %d\nreceived: %d", expected, actual)
	}
}

func assertBoolEqual(t *testing.T, actual bool, expected bool) {
	if actual != expected {
		t.Errorf("expected: %t\nreceived: %t", expected, actual)
	}
}

func assertStringEqual(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected: %s\nreceived: %s", expected, actual)
	}
}

func assertStructEqual(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %s\nreceived: %s", expected, actual)
	}
}

type DummyFont struct{}

func (f *DummyFont) FaceName() string { return "Test font" }
func (f *DummyFont) TextSize(size int, text string) utils.Size {
	return utils.Size{len(text) * fontCharWidth, fontCharHeight}
}

func InitDummyTheme() *gui.Theme {
	return gui.NewDefaultTheme(&DummyFont{})
}

func InitDummyGui(width int, height int, theme *gui.Theme) *gui.Gui {
	if theme == nil {
		theme = InitDummyTheme()
	}
	return gui.NewGui(theme, width, height)
}
