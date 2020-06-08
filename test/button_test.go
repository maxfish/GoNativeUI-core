package test

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"testing"
)

func TestButton(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
	defer FreeGui(g)

	button := gui.NewButton(textStrings[0])
	g.Screen().AddChild(button)

	// Button pressed
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), true)
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), false)
}

func TestToggleButton(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
	defer FreeGui(g)

	button := gui.NewToggleButton(textStrings[0])
	g.Screen().AddChild(button)

	// Button pressed
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), true)
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), true)

	// Button released
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button.Pressed(), false)
}

func TestToggleButtonGrouped(t *testing.T) {
	g := InitDummyGui(screenWidth, screenHeight, nil)
	defer FreeGui(g)

	button1 := gui.NewToggleButton(textStrings[0])
	button2 := gui.NewToggleButton(textStrings[1])
	g.Screen().AddChildren(button1, button2)

	group := []*gui.ToggleButton{button1, button2}
	button1.SetButtonGroup(group)
	button2.SetButtonGroup(group)

	// Initial state
	assertBoolEqual(t, button1.Pressed(), false)
	assertBoolEqual(t, button2.Pressed(), false)

	// Button1 pressed
	button1.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button1.Pressed(), true)
	button1.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button1.Pressed(), true)

	// Another click doesn't change the button status
	button1.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button1.Pressed(), true)
	button1.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button1.Pressed(), true)

	// Button2 pressed -> disabled Button1
	button2.OnMouseEvent(MouseButtonTestEvent(gui.EventActionPress, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button2.Pressed(), true)
	assertBoolEqual(t, button1.Pressed(), false)
	button2.OnMouseEvent(MouseButtonTestEvent(gui.EventActionRelease, gui.MouseButtonLeft, 0, 0))
	assertBoolEqual(t, button2.Pressed(), true)
	assertBoolEqual(t, button1.Pressed(), false)
}
