package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"reflect"
	"testing"
)

var textStrings = [4]string{"Text 1", "Text 2 $$", "Longer test text 3", "Another text string!"}
var textLengths = []utils.Size{{34, 16}, {52, 16}, {99, 16}, {108, 16}}

const (
	screenW = 800
	screenH = 500
)

func assertIntEqual(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("expected: %d\nreceived: %d", expected, actual)
	}
}

func assertStringEqual(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected: %s\nreceived: %s", expected, actual)
	}
}

func assertStructEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %s\nreceived: %s", expected, actual)
	}
}

func InitTestGui(width int, height int, theme *gui.Theme) *gui.Gui {
	if theme == nil {
		theme = gui.NewDefaultTheme()
	}
	return gui.NewGui(theme, width, height)
}
