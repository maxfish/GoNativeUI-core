package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"testing"
)

func TestBox(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	box := gui.NewBoxContainer(g.Theme(), gui.BoxHorizontalOrientation)
	label := gui.NewLabel(textStrings[0])
	box.AddChild(label)
	g.Screen().AddChild(box)
	g.Screen().Layout()

	// Wrap content
	dimensions := utils.Size{box.Bounds().W, box.Bounds().H}
	assertStructEqual(t, dimensions, textLengths[0])

	// FIXME: This fails because of a bug happening when the only child of a Box is a simple widget with constraints
	//label.SetMinimumWidth(100)
	//box.Layout()
	//dimensions = utils.Size{W: box.Bounds().W, H: box.Bounds().H}
	//expected := utils.Size{W: 100, H: box.Bounds().H}
	//assertStructEqual(t, dimensions, expected)
}

func TestBoxComplexLayout(t *testing.T) {
	theme := gui.NewDefaultTheme()
	g := InitTestGui(screenW, screenH, theme)

	// Prepare the layout
	spacer1 := gui.NewSpacer()
	spacer1.SetMinimumHeight(100)
	spacer2 := gui.NewSpacer()
	invisibleLabel := gui.NewLabel(textStrings[1])
	invisibleLabel.SetVisible(false)

	c2 := gui.NewBoxContainer(theme, gui.BoxHorizontalOrientation,
		gui.NewButton(textStrings[1]),
		gui.NewSpacer(),
		gui.NewButton(textStrings[2]),
	)
	c2.SetStretch(1)
	c1 := gui.NewBoxContainer(theme, gui.BoxVerticalOrientation,
		gui.NewLabel(textStrings[1]),
		gui.NewButton(textStrings[2]),
		invisibleLabel,
		gui.NewToggleButton(textStrings[3]),
		spacer1,
		c2,
	)
	c1.SetFlex(1)
	c1.SetStretch(1)
	c1.SetMinimumWidth(148)
	l := gui.NewLabel(textStrings[3])
	g.Screen().AddChildren(c1, spacer2, l, gui.NewSpacer())

	// Main layout
	g.Screen().Layout()
	assertStringEqual(t, c1.Bounds().ToString(), "{x:0,y:0,w:230,h:500}")
	assertStringEqual(t, c2.Bounds().ToString(), "{x:0,y:484,w:230,h:16}")
	assertStringEqual(t, l.Bounds().ToString(), "{x:460,y:0,w:108,h:16}")
	assertStringEqual(t, spacer1.Bounds().ToString(), "{x:0,y:48,w:0,h:436}")
	assertStringEqual(t, spacer2.Bounds().ToString(), "{x:230,y:0,w:230,h:0}")

	// Minimum dimensions constraints
	g.Screen().SetDimension(256, 150)
	g.Screen().Layout()
	assertStringEqual(t, c1.Bounds().ToString(), "{x:0,y:0,w:148,h:150}")
	assertStringEqual(t, c2.Bounds().ToString(), "{x:0,y:148,w:148,h:16}")
	assertStringEqual(t, l.Bounds().ToString(), "{x:148,y:0,w:108,h:16}")
	assertStringEqual(t, spacer1.Bounds().ToString(), "{x:0,y:48,w:0,h:100}")
	assertStringEqual(t, spacer2.Bounds().ToString(), "{x:148,y:0,w:0,h:0}")

}
