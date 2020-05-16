package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"testing"
)

func TestLabel(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
	label := gui.NewLabel(textStrings[0])
	g.Screen().AddChild(label)

	// Check initial state
	assertBoolEqual(t, label.Visible(), true)
	assertBoolEqual(t, label.Enabled(), true)

	// Measured size == Text size + Padding
	assertIntEqual(t, label.MeasuredWidth(), textLengths[0].W()+label.Padding().Left+label.Padding().Right)
	assertIntEqual(t, label.MeasuredHeight(), textLengths[0].H()+label.Padding().Top+label.Padding().Bottom)

	assertStructEqual(t, label.TextColor(), g.Theme().LabelTextColor)
	assertStructEqual(t, label.Font(), g.Theme().LabelFont)
	assertStructEqual(t, label.FontSize(), g.Theme().LabelFontSize)

	assertStringEqual(t, label.Text(), textStrings[0])

	label.SetText(textStrings[2])
	assertStringEqual(t, label.Text(), textStrings[2])

	for i := 0; i < len(textStrings); i++ {
		label.SetText(textStrings[i])
		assertStructEqual(t, utils.Size{label.ContentWidth(), label.ContentHeight()}, textLengths[i])
	}
}
