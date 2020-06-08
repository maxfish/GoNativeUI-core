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
	EventActionNone EventAction = iota
	EventActionPress
	EventActionRelease
	EventActionRepeat
)

type ButtonIndex int

const (
	MouseButtonNone ButtonIndex = iota
	MouseButtonLeft
	MouseButtonMiddle
	MouseButtonRight
)

// Mouse
type MouseEventType int

const (
	MouseEventPosition MouseEventType = iota
	MouseEventButton
	MouseEventWheel
)

type MouseEvent struct {
	Type             MouseEventType
	Action           EventAction
	X, Y             float32
	ScrollX, ScrollY float32
	Button           ButtonIndex
	Modifiers        ModifierKey
}

type IMouseListener interface {
	OnMouseEvent(event MouseEvent) IWidget
}

// Keyboard

type Key int

const (
	KeyUnknown Key = -1

	// Printable keys
	KeySpace        = 32
	KeyApostrophe   = 39  // '
	KeyComma        = 44  // ','
	KeyMinus        = 45  // '-'
	KeyPeriod       = 46  // '.'
	KeySlash        = 47  // '/'
	Key0            = 48
	Key1            = 49
	Key2            = 50
	Key3            = 51
	Key4            = 52
	Key5            = 53
	Key6            = 54
	Key7            = 55
	Key8            = 56
	Key9            = 57
	KeySemicolon    = 59  // ';'
	KeyEqual        = 61  // '='
	KeyA            = 65
	KeyB            = 66
	KeyC            = 67
	KeyD            = 68
	KeyE            = 69
	KeyF            = 70
	KeyG            = 71
	KeyH            = 72
	KeyI            = 73
	KeyJ            = 74
	KeyK            = 75
	KeyL            = 76
	KeyM            = 77
	KeyN            = 78
	KeyO            = 79
	KeyP            = 80
	KeyQ            = 81
	KeyR            = 82
	KeyS            = 83
	KeyT            = 84
	KeyU            = 85
	KeyV            = 86
	KeyW            = 87
	KeyX            = 88
	KeyY            = 89
	KeyZ            = 90
	KeyLeftBracket  = 91  // '['
	KeyBackslash    = 92  // '\'
	KeyRightBracket = 93  // ']'
	KeyGraveAccent  = 96  // '`'
	KeyWorld1       = 161 // non-US #1
	KeyWorld2       = 162 // non-US #2

	// Function keys
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyF13          = 302
	KeyF14          = 303
	KeyF15          = 304
	KeyF16          = 305
	KeyF17          = 306
	KeyF18          = 307
	KeyF19          = 308
	KeyF20          = 309
	KeyF21          = 310
	KeyF22          = 311
	KeyF23          = 312
	KeyF24          = 313
	KeyF25          = 314
	KeyKP0          = 320
	KeyKP1          = 321
	KeyKP2          = 322
	KeyKP3          = 323
	KeyKP4          = 324
	KeyKP5          = 325
	KeyKP6          = 326
	KeyKP7          = 327
	KeyKP8          = 328
	KeyKP9          = 329
	KeyKPDecimal    = 330
	KeyKPDivide     = 331
	KeyKPMultiply   = 332
	KeyKPSubtract   = 333
	KeyKPAdd        = 334
	KeyKPEnter      = 335
	KeyKPEqual      = 336
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyMenu         = 348
)

type IKeyboardListener interface {
	OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool
	OnCharEvent(char rune) bool
}

type IFocusable interface {
	IKeyboardListener
	FocusGained()
	FocusLost()
	Focused() bool
}
