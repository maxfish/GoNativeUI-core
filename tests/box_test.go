package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"testing"
)

func TestBox(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
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
	g := InitDummyGui(screenWidth, screenHeight, nil)
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
	c := box.FindChildAt(180, 10)
	assertStructEqual(t, c, button3)
}

func TestBoxMouseInput(t *testing.T) {
	theme := InitDummyTheme()
	g := InitDummyGui(screenWidth, screenHeight, theme)
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
	theme := InitDummyTheme()
	g := InitDummyGui(screenWidth, screenHeight, theme)

	l1 := gui.NewLabel(textStrings[2])
	l1.SetId("label1")
	l2 := gui.NewLabel(textStrings[3])
	l2.SetId("label2")
	c1 := gui.NewBoxContainer(theme, gui.BoxHorizontalOrientation, l1)
	c1.AddChild(l2)
	g.Screen().AddChild(c1)
	g.Screen().Layout()

	// Assert content wrapping
	assertStructEqual(t, c1.Bounds(), utils.Rect{X: 0, Y: 0, W: 380, H: 16})
	// Assert children placement
	assertStructEqual(t, l1.Bounds(), utils.Rect{X: 0, Y: 0, W: 180, H: 16})
	assertStructEqual(t, l2.Bounds(), utils.Rect{X: 180, Y: 0, W: 200, H: 16})

	// Spacing
	c1.SetSpacing(10)
	g.Screen().Layout()
	assertStructEqual(t, c1.Bounds(), utils.Rect{X: 0, Y: 0, W: 390, H: 16})
	assertStructEqual(t, l1.Bounds(), utils.Rect{X: 0, Y: 0, W: 180, H: 16})
	assertStructEqual(t, l2.Bounds(), utils.Rect{X: 190, Y: 0, W: 200, H: 16})

	// Removing children
	c1.RemoveChildById("label1")
	c1.RemoveChildById("label2")
	g.Screen().Layout()
	assertStructEqual(t, c1.Children(), []gui.IWidget{})
	// Content wrapping without children
	assertStructEqual(t, c1.Bounds(), utils.Rect{})
}

func TestBoxComplexLayout(t *testing.T) {
	theme := InitDummyTheme()
	g := InitDummyGui(screenWidth, screenHeight, theme)

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
	assertStructEqual(t, c1.Bounds(), utils.Rect{X: 0, Y: 0, W: 200, H: 500})
	assertStructEqual(t, c2.Bounds(), utils.Rect{X: 0, Y: 474, W: 200, H: 26})
	assertStructEqual(t, l.Bounds(), utils.Rect{X: 400, Y: 0, W: 200, H: 16})
	assertStructEqual(t, spacer1.Bounds(), utils.Rect{X: 0, Y: 68, H: 406})
	assertStructEqual(t, spacer2.Bounds(), utils.Rect{X: 200, Y: 0, W: 200, H: 0})

	// Minimum dimensions constraints
	g.Screen().SetDimension(256, 150)
	g.Screen().Layout()
	assertStructEqual(t, c1.Bounds(), utils.Rect{X: 0, Y: 0, W: 148, H: 150})
	assertStructEqual(t, c2.Bounds(), utils.Rect{X: 0, Y: 168, W: 148, H: 26})
	assertStructEqual(t, l.Bounds(), utils.Rect{X: 148, Y: 0, W: 200, H: 16})
	assertStructEqual(t, spacer1.Bounds(), utils.Rect{X: 0, Y: 68, W: 0, H: 100})
	assertStructEqual(t, spacer2.Bounds(), utils.Rect{X: 148, Y: 0, W: 0, H: 0})
}
