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
	assertStructEqual(t, dimensions, utils.Size{label.MeasuredWidth(), label.MeasuredHeight()})

	// FIXME: This fails because of a bug happening when the only child of a Box is a simple widget with constraints
	//label.SetMinimumWidth(100)
	//box.Layout()
	//dimensions = utils.Size{W: box.Bounds().W, H: box.Bounds().H}
	//expected := utils.Size{W: 100, H: box.Bounds().H}
	//assertStructEqual(t, dimensions, expected)
}

func TestBoxChildren(t *testing.T) {
	theme := gui.NewDefaultTheme()
	g := InitTestGui(screenW, screenH, theme)
	box := gui.NewBoxContainer(g.Theme(), gui.BoxHorizontalOrientation)
	label1 := gui.NewLabel(textStrings[0])
	label1.SetId("label1")
	label2 := gui.NewLabel(textStrings[0])
	label2.SetId("label2")
	button3 := gui.NewButton(textStrings[0])
	button3.SetId("label3")
	label4 := gui.NewLabel(textStrings[0])
	label4.SetId("label4")
	box.AddChildren(label1, button3)
	g.Screen().AddChild(box)
	g.Screen().Layout()

	assertStructEqual(t, box.Children(), []gui.IWidget{label1, button3})

	box.AddChildAtIndex(label2, 1)
	assertStructEqual(t, box.Children(), []gui.IWidget{label1, label2, button3})

	box.AddChild(label4)
	assertStructEqual(t, box.Children(), []gui.IWidget{label1, label2, button3, label4})

	g.Screen().Layout()
	c := box.FindChildAt(100, 10)
	assertStructEqual(t, c, button3)
}

func TestBoxInput(t *testing.T) {
	theme := gui.NewDefaultTheme()
	g := InitTestGui(screenW, screenH, theme)
	box := gui.NewBoxContainer(g.Theme(), gui.BoxHorizontalOrientation)
	button := gui.NewButton(textStrings[0])
	box.AddChildren(button)
	g.Screen().AddChild(box)
	g.Screen().Layout()

	// Button pressed
	assertBoolEqual(t, button.Pressed(), false)
	box.OnMouseButtonEvent(10, 10, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button.Pressed(), true)
	button.OnMouseButtonEvent(10, 10, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button.Pressed(), false)
}

func TestBoxBasics(t *testing.T) {
	theme := gui.NewDefaultTheme()
	g := InitTestGui(screenW, screenH, theme)

	l1 := gui.NewLabel(textStrings[2])
	l1.SetId("label1")
	l2 := gui.NewLabel(textStrings[3])
	l2.SetId("label2")
	c1 := gui.NewBoxContainer(theme, gui.BoxHorizontalOrientation, l1)
	c1.AddChild(l2)
	g.Screen().AddChild(c1)
	g.Screen().Layout()

	// Assert content wrapping
	assertStringEqual(t, c1.Bounds().ToString(), "{x:0,y:0,w:207,h:16}")
	// Assert children placement
	assertStringEqual(t, l1.Bounds().ToString(), "{x:0,y:0,w:99,h:16}")
	assertStringEqual(t, l2.Bounds().ToString(), "{x:99,y:0,w:108,h:16}")

	c1.RemoveChildById("label1")
	c1.RemoveChildById("label2")
	g.Screen().Layout()
	assertStructEqual(t, c1.Children(), []gui.IWidget{})
	// Content wrapping without children
	assertStringEqual(t, c1.Bounds().ToString(), "{x:0,y:0,w:0,h:0}")
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
	assertStringEqual(t, c2.Bounds().ToString(), "{x:0,y:474,w:230,h:26}")
	assertStringEqual(t, l.Bounds().ToString(), "{x:460,y:0,w:108,h:16}")
	assertStringEqual(t, spacer1.Bounds().ToString(), "{x:0,y:68,w:0,h:406}")
	assertStringEqual(t, spacer2.Bounds().ToString(), "{x:230,y:0,w:230,h:0}")

	// Minimum dimensions constraints
	g.Screen().SetDimension(256, 150)
	g.Screen().Layout()
	assertStringEqual(t, c1.Bounds().ToString(), "{x:0,y:0,w:148,h:150}")
	assertStringEqual(t, c2.Bounds().ToString(), "{x:0,y:168,w:148,h:26}")
	assertStringEqual(t, l.Bounds().ToString(), "{x:148,y:0,w:108,h:16}")
	assertStringEqual(t, spacer1.Bounds().ToString(), "{x:0,y:68,w:0,h:100}")
	assertStringEqual(t, spacer2.Bounds().ToString(), "{x:148,y:0,w:0,h:0}")
}
