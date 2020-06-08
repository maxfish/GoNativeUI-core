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

// Mouse handling
func (g *Gui) OnMouseCursorMoved(x, y float32) bool {
	// TODO: Here the mouse pointer should change based on the component under it
	//log.Printf("[Gui] Mouse moved %.2f,%.2f\n", x, y)
	if len(g.popupStack) > 0 {
		topItem := g.popupStack[len(g.popupStack)-1]
		widget := topItem.Widget()
		if widget.Bounds().ContainsPoint(int(x), int(y)) {
			return widget.OnMouseCursorMoved(x-float32(widget.Bounds().X), y-float32(widget.Bounds().Y))
		}
	}
	return g.root.OnMouseCursorMoved(x, y)
}

func (g *Gui) OnMouseButtonEvent(x float32, y float32, buttonIndex ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	//log.Printf("[Gui] Mouse button #%d <%d> modifiers:%d\n", buttonIndex, event, modifiers)
	if len(g.popupStack) > 0 {
		topItem := g.popupStack[len(g.popupStack)-1]
		widget := topItem.Widget()
		if widget.Bounds().ContainsPoint(int(x), int(y)) {
			return widget.OnMouseButtonEvent(x-float32(widget.Bounds().X), y-float32(widget.Bounds().Y), buttonIndex, event, modifiers)
		} else {
			if buttonIndex == MouseButtonLeft && event == EventActionPress {
				g.DismissPopup()
				return true
			}
		}
	}
	return g.root.OnMouseButtonEvent(x, y, buttonIndex, event, modifiers)
}

func (g *Gui) OnMouseScrolled(x float32, y float32, scrollX, scrollY float32) bool {
	//log.Printf("[Gui] Mouse wheel scrolled %.2f,%.2f\n", scrollX, scrollY)
	if len(g.popupStack) > 0 {
		topItem := g.popupStack[len(g.popupStack)-1]
		widget := topItem.Widget()
		if widget.Bounds().ContainsPoint(int(x), int(y)) {
			return widget.OnMouseScrolled(x-float32(widget.Bounds().X), y-float32(widget.Bounds().Y), scrollX, scrollY)
		}
	}
	return g.root.OnMouseScrolled(x, y, scrollX, scrollY)
}

// Key events

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
