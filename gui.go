package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IFont interface {
	FaceName() string
	TextSize(size int, text string) (int, int)
}

type IRenderer interface {
	Init(screen *Gui, dpiScale float32, uiScale float32)
	Draw()
}

type Gui struct {
	screen    IContainer
	mouseData MouseData
}

func NewGui(theme *Theme, w int, h int) *Gui {
	g := &Gui{}
	g.screen = NewContainer(theme)
	g.screen.SetBounds(utils.Rect{W: w, H: h})
	return g
}

func (g *Gui) Screen() IContainer { return g.screen }
func (g *Gui) Theme() *Theme      { return g.screen.Theme() }

//func (g *Gui) Scale() float32             { return g.scale }
//func (g *Gui) SetScale(scale float32)        { g.scale = scale }

// Mouse handling
func (g *Gui) OnMouseCursorMoved(x, y float32) bool {
	g.mouseData.previousPosX, g.mouseData.previousPosY = g.mouseData.currentPosX, g.mouseData.currentPosY
	g.mouseData.currentPosX, g.mouseData.currentPosY = x, y
	//log.Printf("[Gui] Mouse moved %.2f,%.2f\n", x, y)
	return g.screen.OnMouseCursorMoved(x, y)
}

// NOTE: We don't expect to receive the cursor coordinates here, that's why they are stored when the cursor moves
func (g *Gui) OnMouseButtonEvent(x float32, y float32, buttonIndex ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	//log.Printf("[Gui] Mouse button #%d <%d> modifiers:%d\n", buttonIndex, event, modifiers)
	return g.screen.OnMouseButtonEvent(g.mouseData.currentPosX, g.mouseData.currentPosY, buttonIndex, event, modifiers)
}

func (g *Gui) OnMouseScrolled(scrollX, scrollY float32) bool {
	//log.Printf("[Gui] Mouse wheel scrolled %.2f,%.2f\n", scrollX, scrollY)
	return g.screen.OnMouseScrolled(scrollX, scrollY)
}
