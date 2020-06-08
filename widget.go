package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IWidget interface {
	Id() string
	SetId(id string)
	Parent() IContainer

	setParent(container IContainer)
	Style() *WidgetStyle
	SetStyle(style WidgetStyle)

	// Status
	Visible() bool
	SetVisible(v bool)
	Enabled() bool
	SetEnabled(e bool)

	// Dimensions
	Bounds() utils.Rect
	InnerBounds() utils.Rect
	SetLeft(left int)
	SetTop(top int)
	SetDimension(width int, height int)
	SetWidth(width int)
	SetHeight(height int)
	AbsolutePosition() utils.Point

	// Layout
	Flex() int
	SetFlex(flex int)
	Stretch() int
	SetStretch(stretch int)
	MinimumWidth() int
	SetMinimumWidth(int)
	MaximumWidth() int
	SetMaximumWidth(int)
	MinimumHeight() int
	SetMinimumHeight(int)
	MaximumHeight() int
	SetMaximumHeight(int)
	Measure()
	Layout()

	MeasuredFlex() int
	SetMeasuredFlex(measuredFlex int)
	MeasuredHeight() int
	SetMeasuredHeight(measuredHeight int)
	MeasuredWidth() int
	SetMeasuredWidth(measuredWidth int)

	// Content
	ContentWidth() int
	ContentHeight() int
}

type Widget struct {
	id     string
	parent IContainer

	enabled bool
	visible bool
	bounds  utils.Rect

	style WidgetStyle

	minimumWidth  int
	maximumWidth  int
	minimumHeight int
	maximumHeight int
	flex          int
	stretch       int

	measuredWidth  int
	measuredHeight int
	measuredFlex   int

	contentWidth  int
	contentHeight int
}

// Getters / Setters
func (w *Widget) Id() string          { return w.id }
func (w *Widget) SetId(id string)     { w.id = id }
func (w *Widget) Parent() IContainer  { return w.parent }
func (w *Widget) Style() *WidgetStyle { return &w.style }
func (w *Widget) SetStyle(style WidgetStyle) { w.style = style }

func (w *Widget) setParent(container IContainer)       { w.parent = container }
func (w *Widget) Enabled() bool                        { return w.enabled }
func (w *Widget) SetEnabled(e bool)                    { w.enabled = e }
func (w *Widget) Visible() bool                        { return w.visible }
func (w *Widget) SetVisible(v bool)                    { w.visible = v }
func (w *Widget) InnerBounds() utils.Rect              { return w.bounds.ShrinkByInsets(w.style.Padding) }
func (w *Widget) ContentWidth() int                    { return w.contentWidth }
func (w *Widget) ContentHeight() int                   { return w.contentHeight }
func (w *Widget) Flex() int                            { return w.flex }
func (w *Widget) SetFlex(flex int)                     { w.flex = flex }
func (w *Widget) Stretch() int                         { return w.stretch }
func (w *Widget) SetStretch(stretch int)               { w.stretch = stretch }
func (w *Widget) Bounds() utils.Rect                   { return w.bounds }
func (w *Widget) SetTop(top int)                       { w.bounds.Y = top }
func (w *Widget) SetLeft(left int)                     { w.bounds.X = left }
func (w *Widget) MaximumHeight() int                   { return w.maximumHeight }
func (w *Widget) SetMaximumHeight(maximumHeight int)   { w.maximumHeight = maximumHeight }
func (w *Widget) MinimumHeight() int                   { return w.minimumHeight }
func (w *Widget) SetMinimumHeight(minimumHeight int)   { w.minimumHeight = minimumHeight }
func (w *Widget) MinimumWidth() int                    { return w.minimumWidth }
func (w *Widget) SetMinimumWidth(minimumWidth int)     { w.minimumWidth = minimumWidth }
func (w *Widget) MaximumWidth() int                    { return w.maximumWidth }
func (w *Widget) SetMaximumWidth(maximumWidth int)     { w.maximumWidth = maximumWidth }
func (w *Widget) MeasuredFlex() int                    { return w.measuredFlex }
func (w *Widget) SetMeasuredFlex(measuredFlex int)     { w.measuredFlex = measuredFlex }
func (w *Widget) MeasuredHeight() int                  { return w.measuredHeight }
func (w *Widget) SetMeasuredHeight(measuredHeight int) { w.measuredHeight = measuredHeight }
func (w *Widget) MeasuredWidth() int                   { return w.measuredWidth }
func (w *Widget) SetMeasuredWidth(measuredWidth int)   { w.measuredWidth = measuredWidth }

func (w *Widget) AbsolutePosition() utils.Point {
	p := utils.Point{w.bounds.X, w.bounds.Y}
	if w.parent != nil {
		p = p.Add(w.parent.AbsolutePosition())
	}
	return p
}

func (w *Widget) SetWidth(width int) {
	if width < w.minimumWidth {
		w.bounds.W = w.minimumWidth
	} else if w.maximumWidth > 0 && width > w.maximumWidth {
		w.bounds.W = w.maximumWidth
	} else {
		w.bounds.W = width
	}
}

func (w *Widget) SetHeight(height int) {
	if height < w.minimumHeight {
		w.bounds.H = w.minimumHeight
	} else if w.maximumHeight > 0 && height > w.maximumHeight {
		w.bounds.H = w.maximumHeight
	} else {
		w.bounds.H = height
	}
}

func (w *Widget) SetDimension(width int, height int) {
	w.SetWidth(width)
	w.SetHeight(height)
}

// Methods which need to be implemented by the widgets

func (w *Widget) Layout() { /* NOP */ }

func (w *Widget) Measure() {
	w.measuredFlex = w.flex
}

func widgetInit(w IWidget) {
	w.SetEnabled(true)
	w.SetVisible(true)
}
