package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

var currentGui *Gui

func CurrentGui() *Gui {
	return currentGui
}

type IFont interface {
	FaceName() string
	TextSize(fontSize int, text string, numGlyphs ...int) utils.Size
	LineHeight(fontSize int) int
	IndexFromCoords(fontSize int, text string, x int, y int) int
}

type IRenderer interface {
	Init(screen *Gui, dpiScale float32, uiScale float32)
	Draw()
}

type Gui struct {
	root  *BoxContainer
	theme *Theme

	focusedDescendant IFocusable
	popupStack        []IPopup
}

func NewGui(theme *Theme, w int, h int) *Gui {
	if currentGui != nil {
		panic("There can be only one instance of the Gui")
	}
	g := &Gui{
		theme:      theme,
		root:       NewBoxContainer(BoxHorizontalOrientation),
		popupStack: make([]IPopup, 0, 16),
	}
	g.root.SetDimension(w, h)
	currentGui = g
	return g
}

func (g *Gui) Free() {
	currentGui = nil
	g.theme = nil
	g.root = nil
}

func (g *Gui) Screen() *BoxContainer { return g.root }
func (g *Gui) Theme() *Theme         { return g.theme }
func (g *Gui) Popups() []IPopup      { return g.popupStack }

func (g *Gui) ShowPopup(popup IPopup) {
	popup.SetVisible(true)
	g.popupStack = append(g.popupStack, popup)
}

func (g *Gui) DismissPopup() {
	g.popupStack[len(g.popupStack)-1].SetVisible(false)
	g.popupStack = g.popupStack[:len(g.popupStack)-1]
}

func (g *Gui) RemoveFocusFrom(child IWidget) {
	focusable, ok := child.(IFocusable)
	if ok && focusable == g.focusedDescendant {
		g.focusedDescendant.FocusLost()
		g.focusedDescendant = nil
	}
}

func (g *Gui) RemoveFocus() {
	if g.focusedDescendant != nil {
		g.focusedDescendant.FocusLost()
		g.focusedDescendant = nil
	}
}

func (g *Gui) RequestFocusFor(child IWidget) {
	focusable, ok := child.(IFocusable)
	if ok {
		if focusable == g.focusedDescendant {
			return
		}
		if g.focusedDescendant != nil {
			previousFocused, _ := g.focusedDescendant.(IFocusable)
			previousFocused.FocusLost()
		}
		g.focusedDescendant = focusable
		focusable.FocusGained()
	}
}

// Mouse events

func (g *Gui) OnMouseEvent(event MouseEvent) IWidget {
	// Popups
	// TODO: If the top popup is modal then the root controller shouldn't get the event
	if len(g.popupStack) > 0 {
		topItem := g.popupStack[len(g.popupStack)-1]
		widget := topItem.Widget()
		if widget.Bounds().ContainsPoint(int(event.X), int(event.Y)) {
			listener, ok := widget.(IMouseListener)
			if ok {
				e := event
				e.X -= float32(widget.Bounds().X)
				e.Y -= float32(widget.Bounds().Y)
				return listener.OnMouseEvent(e)
			}
		}
	}

	// Root controller
	return g.root.OnMouseEvent(event)
}

// Keyboard events

func (g *Gui) OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool {
	// Sends the key events only to the focusedDescendant
	if g.focusedDescendant == nil {
		return false
	}
	return g.focusedDescendant.OnKeyEvent(key, action, modifierKey)
}

func (g *Gui) OnCharEvent(char rune) bool {
	// Sends the key events only to the focusedDescendant
	if g.focusedDescendant != nil {
		return g.focusedDescendant.OnCharEvent(char)
	}
	return false
}
