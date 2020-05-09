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
	OnMouseCursorMoved(x, y float32) bool
	OnMouseButtonEvent(button ButtonIndex, event EventAction, modifiers ModifierKey) bool
	OnMouseScrolled(scrollX, scrollY float32) bool
}

type IKeyboardListener interface {
}
