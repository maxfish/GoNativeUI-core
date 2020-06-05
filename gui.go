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
	screen    *Screen
	theme     *Theme
}

func NewGui(theme *Theme, w int, h int) *Gui {
	if currentGui != nil {
		panic("There can be only one instance of the Gui")
	}
	g := &Gui{}
	g.theme = theme
	g.screen = NewScreen(BoxHorizontalOrientation, w, h)
	currentGui = g
	return g
}

func (g *Gui) Free() {
	currentGui = nil
	g.theme = nil
	g.screen = nil
}

func (g *Gui) Screen() *Screen { return g.screen }
func (g *Gui) Theme() *Theme   { return g.theme }

// Mouse handling
func (g *Gui) OnMouseCursorMoved(x, y float32) bool {
	// TODO: Here the mouse pointer should change based on the component under it
	//log.Printf("[Gui] Mouse moved %.2f,%.2f\n", x, y)
	return g.screen.OnMouseCursorMoved(x, y)
}

func (g *Gui) OnMouseButtonEvent(x float32, y float32, buttonIndex ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	//log.Printf("[Gui] Mouse button #%d <%d> modifiers:%d\n", buttonIndex, event, modifiers)
	return g.screen.OnMouseButtonEvent(x, y, buttonIndex, event, modifiers)
}

func (g *Gui) OnMouseScrolled(x float32, y float32, scrollX, scrollY float32) bool {
	//log.Printf("[Gui] Mouse wheel scrolled %.2f,%.2f\n", scrollX, scrollY)
	return g.screen.OnMouseScrolled(x, y, scrollX, scrollY)
}

// Key events

func (g *Gui) OnKeyEvent(key Key, action EventAction, modifierKey ModifierKey) bool {
	return g.screen.OnKeyEvent(key, action, modifierKey)
}

func (g *Gui) OnCharEvent(char rune) bool {
	return g.screen.OnCharEvent(char)
}
