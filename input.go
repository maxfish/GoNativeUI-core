package gui

type ModifierKey uint

const (
	ModifierKeyShift    ModifierKey = 0x01
	ModifierKeyControl              = 0x02
	ModifierKeyAlt                  = 0x04
	ModifierKeyCapsLock             = 0x08
)

type EventAction uint

const (
	EventActionPress EventAction = iota
	EventActionRelease
	EventActionRepeat
)

type IMouseListener interface {
	OnMouseMoved(x, y float32) bool
	OnMouseButtonEvent(buttonIndex uint, event EventAction, modifiers ModifierKey) bool
	OnMouseScrolled(x, y float32) bool
}

type IKeyboardListener interface {
}
