package tests

import (
	gui "github.com/maxfish/GoNativeUI-Core"
	"testing"
)

func TestButton(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	button := gui.NewButton(textStrings[0])
	g.Screen().AddChild(button)

	// Button pressed
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button.Pressed(), true)
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button.Pressed(), false)
}

func TestToggleButton(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
	button := gui.NewToggleButton(textStrings[0])
	g.Screen().AddChild(button)

	// Button pressed
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button.Pressed(), true)
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button.Pressed(), true)

	// Button released
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button.Pressed(), false)
	button.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button.Pressed(), false)
}

func TestToggleButtonGrouped(t *testing.T) {
	g := InitTestGui(screenW, screenH, nil)
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
	button1.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button1.Pressed(), true)
	button1.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button1.Pressed(), true)

	// Another click doesn't change the button status
	button1.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button1.Pressed(), true)
	button1.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button1.Pressed(), true)

	// Button2 pressed -> disabled Button1
	button2.OnMouseButtonEvent(0, 0, 0, gui.EventActionPress, 0)
	assertBoolEqual(t, button2.Pressed(), true)
	button2.OnMouseButtonEvent(0, 0, 0, gui.EventActionRelease, 0)
	assertBoolEqual(t, button2.Pressed(), true)
	assertBoolEqual(t, button1.Pressed(), false)
}
