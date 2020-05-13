package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"reflect"
	"testing"
)

const (
	screenW = 800
	screenH = 500
	text1   = "LabelText"
	text2   = "LabelText_2"
)

func TestLabel(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	label := gui.NewLabel(text1)
	g.Screen().AddChild(label)

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

	value := label.Text()
	expected := text1
	if value != expected {
		t.Errorf("expected: %s\nreceived: %s", expected, value)
	}

	label.SetText(text2)
	value = label.Text()
	expected = text2
	if value != expected {
		t.Errorf("expected: %s\nreceived: %s", expected, value)
	}

	valueW, valueH := g.Theme().LabelFont.TextSize(label.FontSize(), label.Text())
	expectedW, expectedH := 68, 16
	if valueW != expectedW || valueH != expectedH {
		t.Errorf("expected: %d,%d\nreceived: %d,%d", expectedW, expectedH, valueW, valueH)
	}
}