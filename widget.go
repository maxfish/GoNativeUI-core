package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
	"log"
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
	PreferredWidth() int
	PreferredHeight() int
	SetBounds(b utils.Rect)
	SetLeft(left int)
	SetRight(right int)
	SetTop(top int)
	SetBottom(bottom int)
	SetPadding(b utils.Insets)
	DimensionH() utils.Dimension
	SetDimensionH(h utils.Dimension)
	DimensionV() utils.Dimension
	SetDimensionV(v utils.Dimension)

	// Content
	ContentWidth() int
	ContentHeight() int
	ContentAlignmentH() utils.AlignmentH
	ContentAlignmentV() utils.AlignmentV

	Layout()
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

	contentAlignmentH utils.AlignmentH
	contentAlignmentV utils.AlignmentV
	contentWidth      int
	contentHeight     int
	contentSizeValid  bool
}

func (w *Widget) Id() string         { return w.id }
func (w *Widget) SetId(id string)    { w.id = id }
func (w *Widget) Parent() IContainer { return w.parent }
func (w *Widget) SetParent(container IContainer) {
	w.parent = container
	w.theme = container.Theme()
}

func (w *Widget) Theme() *Theme {
	return w.theme
}
func (w *Widget) SetTheme(t *Theme) {
	w.theme = t
}

// Getters
func (w *Widget) Enabled() bool                       { return w.enabled }
func (w *Widget) Visible() bool                       { return w.visible }
func (w *Widget) Bounds() utils.Rect                  { return w.bounds }
func (w *Widget) Padding() utils.Insets               { return w.padding }
func (w *Widget) InnerBounds() utils.Rect             { return w.bounds.ShrinkByInsets(w.padding) }
func (w *Widget) ContentWidth() int                   { return w.contentWidth }
func (w *Widget) ContentHeight() int                  { return w.contentHeight }
func (w *Widget) ContentAlignmentH() utils.AlignmentH { return w.contentAlignmentH }
func (w *Widget) ContentAlignmentV() utils.AlignmentV { return w.contentAlignmentV }
func (w *Widget) DimensionH() utils.Dimension         { return w.dimensionH }
func (w *Widget) DimensionV() utils.Dimension         { return w.dimensionV }

// Setters
func (w *Widget) SetVisible(v bool)         { w.visible = v }
func (w *Widget) SetEnabled(e bool)         { w.enabled = e }
func (w *Widget) SetPadding(b utils.Insets) { w.padding = b }
func (w *Widget) SetBounds(b utils.Rect)    { w.bounds = b }
func (w *Widget) SetLeft(left int)          { w.bounds.X = left }
func (w *Widget) SetRight(right int)        { w.bounds.X = right - w.bounds.W }
func (w *Widget) SetTop(top int)            { w.bounds.Y = top }
func (w *Widget) SetBottom(bottom int)      { w.bounds.X = bottom - w.bounds.H }

func (w *Widget) SetDimensionH(h utils.Dimension) {
	w.dimensionH = h
	if h.Unit == utils.PixelUnit {
		w.bounds.W = h.Value
	}
}
func (w *Widget) SetDimensionV(v utils.Dimension) {
	w.dimensionV = v
	if v.Unit == utils.PixelUnit {
		w.bounds.H = v.Value
	}
}

// Layout
func (w *Widget) PreferredWidth() int {
	switch w.dimensionH.Unit {
	case utils.PixelUnit:
		return w.dimensionH.Value
	case utils.NoUnit:
		// Wraps the content
		return w.padding.Left + w.contentWidth + w.padding.Right
	case utils.PercentageUnit:
		if w.Parent() == nil {
			// Wraps the content
			return w.padding.Left + w.contentWidth + w.padding.Right
		} else {
			return (w.Parent().InnerBounds().W * w.dimensionH.Value) / 100
		}
	}
	log.Panicf("Dimension unit is unknown")
	return -1
}

func (w *Widget) PreferredHeight() int {
	switch w.dimensionV.Unit {
	case utils.PixelUnit:
		return w.dimensionV.Value
	case utils.NoUnit:
		// Wraps the content
		return w.padding.Top + w.contentHeight + w.padding.Bottom
	case utils.PercentageUnit:
		if w.Parent() == nil {
			// Wraps the content
			return w.padding.Top + w.contentHeight + w.padding.Bottom
		} else {
			return (w.Parent().InnerBounds().H * w.dimensionV.Value) / 100
		}
	}
	log.Panicf("Dimension unit is unknown")
	return -1
}

func (w *Widget) Layout() {
	// NOP
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

func widgetLayout(w IWidget) {
	bounds := w.Bounds()
	bounds.W = w.PreferredWidth()
	bounds.H = w.PreferredHeight()
	w.SetBounds(bounds)
}
