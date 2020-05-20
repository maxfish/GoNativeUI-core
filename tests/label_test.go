package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"testing"
)

func TestLabel(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
	defer FreeGui(g)
	label := gui.NewLabel(textStrings[0])

	// Check initial state
	assertBoolEqual(t, label.Visible(), true)
	assertBoolEqual(t, label.Enabled(), true)
	assertStructEqual(t, label.Style().TextColor, g.Theme().TextColor)
	assertStructEqual(t, label.Style().Font, g.Theme().TextFont)
	assertStructEqual(t, label.Style().FontSize, g.Theme().TextFontSize)
	assertStringEqual(t, label.Text(), textStrings[0])

	// Change the text
	label.SetText(textStrings[2])
	assertStringEqual(t, label.Text(), textStrings[2])

	// Measured size == Text size + Padding
	label.Measure()
	assertIntEqual(t, label.MeasuredWidth(), textSizes[2].W()+label.Style().Padding.Left+label.Style().Padding.Right)
	assertIntEqual(t, label.MeasuredHeight(), textSizes[2].H()+label.Style().Padding.Top+label.Style().Padding.Bottom)

	for i := 0; i < len(textStrings); i++ {
		label.SetText(textStrings[i])
		assertStructEqual(t, utils.Size{label.ContentWidth(), label.ContentHeight()}, textSizes[i])
	}
}
