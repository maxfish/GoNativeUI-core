package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IWidget interface {
	IMouseListener

	Id() string
	SetId(id string)
	Parent() IContainer
	SetParent(container IContainer)

	// Status
	Visible() bool
	SetVisible(v bool)
	Enabled() bool
	SetEnabled(e bool)

	// Theme
	Theme() *Theme
	SetTheme(t *Theme)

	// Dimensions
	Bounds() utils.Rect
	InnerBounds() utils.Rect
	Padding() utils.Insets
	SetLeft(left int)
	SetTop(top int)
	SetDimension(width int, height int)
	SetPadding(b utils.Insets)

	// Layout
	Flex() int
	SetFlex(flex int)
	MinimumWidth() int
	SetMinimumWidth(int)
	MaximumWidth() int
	SetMaximumWidth(int)
	MinimumHeight() int
	SetMinimumHeight(int)
	MaximumHeight() int
	SetMaximumHeight(int)
	Layout()

	// Content
	ContentWidth() int
	ContentHeight() int
	ContentAlignmentH() utils.AlignmentH
	ContentAlignmentV() utils.AlignmentV
}

type Widget struct {
	id     string
	parent IContainer
	theme  *Theme

	enabled    bool
	visible    bool
	bounds     utils.Rect
	dimensionH utils.Dimension
	dimensionV utils.Dimension
	padding    utils.Insets

	minimumWidth  int
	maximumWidth  int
	minimumHeight int
	maximumHeight int
	flex          int

	contentAlignmentH utils.AlignmentH
	contentAlignmentV utils.AlignmentV
	contentWidth      int
	contentHeight     int
}

// Getters / Setters
func (w *Widget) Id() string                          { return w.id }
func (w *Widget) SetId(id string)                     { w.id = id }
func (w *Widget) Parent() IContainer                  { return w.parent }
func (w *Widget) SetParent(container IContainer)      { w.parent = container; w.theme = container.Theme() }
func (w *Widget) Theme() *Theme                       { return w.theme }
func (w *Widget) SetTheme(t *Theme)                   { w.theme = t }
func (w *Widget) Enabled() bool                       { return w.enabled }
func (w *Widget) SetEnabled(e bool)                   { w.enabled = e }
func (w *Widget) Visible() bool                       { return w.visible }
func (w *Widget) SetVisible(v bool)                   { w.visible = v }
func (w *Widget) Padding() utils.Insets               { return w.padding }
func (w *Widget) SetPadding(b utils.Insets)           { w.padding = b }
func (w *Widget) InnerBounds() utils.Rect             { return w.bounds.ShrinkByInsets(w.padding) }
func (w *Widget) ContentWidth() int                   { return w.contentWidth }
func (w *Widget) ContentHeight() int                  { return w.contentHeight }
func (w *Widget) ContentAlignmentH() utils.AlignmentH { return w.contentAlignmentH }
func (w *Widget) ContentAlignmentV() utils.AlignmentV { return w.contentAlignmentV }
func (w *Widget) Flex() int                           { return w.flex }
func (w *Widget) SetFlex(flex int)                    { w.flex = flex }
func (w *Widget) Bounds() utils.Rect                  { return w.bounds }
func (w *Widget) SetTop(top int)                      { w.bounds.Y = top }
func (w *Widget) SetLeft(left int)                    { w.bounds.X = left }
func (w *Widget) MaximumHeight() int                  { return w.maximumHeight }
func (w *Widget) SetMaximumHeight(maximumHeight int)  { w.maximumHeight = maximumHeight }
func (w *Widget) MinimumHeight() int                  { return w.minimumHeight }
func (w *Widget) SetMinimumHeight(minimumHeight int)  { w.minimumHeight = minimumHeight }
func (w *Widget) MinimumWidth() int                   { return w.minimumWidth }
func (w *Widget) SetMinimumWidth(minimumWidth int)    { w.minimumWidth = minimumWidth }
func (w *Widget) MaximumWidth() int                   { return w.maximumWidth }
func (w *Widget) SetMaximumWidth(maximumWidth int)    { w.maximumWidth = maximumWidth }

func (w *Widget) SetDimension(width int, height int) {
	if width < w.minimumWidth {
		w.bounds.W = w.minimumWidth
	} else if w.maximumWidth > 0 && width > w.maximumWidth {
		w.bounds.W = w.maximumWidth
	} else {
		w.bounds.W = width
	}
	if height < w.minimumHeight {
		w.bounds.H = w.minimumHeight
	} else if w.maximumHeight > 0 && height > w.maximumHeight {
		w.bounds.H = w.maximumHeight
	} else {
		w.bounds.H = height
	}
}

// Layout
func (w *Widget) Layout() {
	panic("Cannot be called on Widget")
}

// Mouse handling
func (w *Widget) OnMouseCursorMoved(x, y float32) bool {
	return false
}
func (w *Widget) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	return false
}
func (w *Widget) OnMouseScrolled(scrollX, scrollY float32) bool {
	return false
}

func widgetInit(w IWidget) {
	w.SetEnabled(true)
	w.SetVisible(true)
}

func widgetLayout(widget IWidget) {
	w := 0
	h := 0
	//if widget.Flex() > 0 {
	//	NOP here. The container will take care of resizing the widget
	//} else {
	wDimension := widget.Padding().Left + widget.ContentWidth() + widget.Padding().Right
	w = utils.MaxI(wDimension, widget.MinimumWidth())
	if widget.MaximumWidth() > 0 {
		w = utils.MinI(w, widget.MaximumWidth())
	}
	hDimension := widget.Padding().Top + widget.ContentHeight() + widget.Padding().Top
	h = utils.MaxI(hDimension, widget.MinimumHeight())
	if widget.MaximumHeight() > 0 {
		h = utils.MinI(h, widget.MaximumHeight())
	}
	//}
	widget.SetDimension(w, h)
}
