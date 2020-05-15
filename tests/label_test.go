package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"reflect"
	"testing"
)

func TestLabel(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	label := gui.NewLabel(textStrings[0])
	g.Screen().AddChild(label)

	// Check initial state
	assertBoolEqual(t, label.Visible(), true)
	assertBoolEqual(t, label.Enabled(), true)

	// Measured size == Text size + Padding
	assertIntEqual(t, label.MeasuredWidth(), textLengths[0].W()+label.Padding().Left+label.Padding().Right)
	assertIntEqual(t, label.MeasuredHeight(), textLengths[0].H()+label.Padding().Top+label.Padding().Bottom)

	valueColor := label.TextColor()
	expectedColor := g.Theme().LabelTextColor
	if !reflect.DeepEqual(valueColor, expectedColor) {
		t.Errorf("expected: %s\nreceived: %s", gui.ColorToString(expectedColor), gui.ColorToString(valueColor))
	}

	valueFont := label.Font()
	expectedFont := g.Theme().LabelFont
	if !reflect.DeepEqual(valueFont, expectedFont) {
		t.Errorf("expected: %s\nreceived: %s", expectedFont.FaceName(), valueFont.FaceName())
	}

	valueFontSize := label.FontSize()
	expectedFontSize := g.Theme().LabelFontSize
	if valueFontSize != expectedFontSize {
		t.Errorf("expected: %d\nreceived: %d", expectedFontSize, valueFontSize)
	}

	assertStringEqual(t, label.Text(), textStrings[0])

	label.SetText(textStrings[2])
	assertStringEqual(t, label.Text(), textStrings[2])

	for i := 0; i < len(textStrings); i++ {
		label.SetText(textStrings[i])
		assertStructEqual(t, utils.Size{label.ContentWidth(), label.ContentHeight()}, textLengths[i])
		//label.Measure()
		//expectedW, expectedH := 68, 16
		//if label.ContentWidth() != expectedW || label.ContentHeight() != expectedH {
		//	t.Errorf("expected: %d,%d\nreceived: %d,%d", expectedW, expectedH, label.ContentWidth(), label.ContentHeight())
		//}
	}
}
