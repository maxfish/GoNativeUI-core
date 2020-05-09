package gui

type ModifierKey int

const (
	ModifierKeyShift    ModifierKey = 0x01
	ModifierKeyControl              = 0x02
	ModifierKeyAlt                  = 0x04
	ModifierKeyCapsLock             = 0x08
	ModifierKeySuper                = 0x10
)

type EventAction int

const (
	EventActionPress EventAction = iota
	EventActionRelease
	EventActionRepeat
)

type ButtonIndex int

type IMouseListener interface {
	OnMouseCursorMoved(x float32, y float32) bool
	OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool
	OnMouseScrolled(scrollX float32, scrollY float32) bool
}

type MouseData struct {
	currentPosX, currentPosY   float32
	previousPosX, previousPosY float32
	button                     ButtonIndex // Button currently pressed
}

type IKeyboardListener interface {
}
