package gui

import (
	"fmt"
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IWidget interface {
	IMouseListener

	Id() string
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
	SetBounds(b utils.Rect)
	SetPadding(b utils.Insets)
	InnerBounds() utils.Rect // Bounds - Padding
	ContentWidth() int
	ContentHeight() int

	// Layout
	SizeToContent()
	ContentAlignmentH() AlignmentH
	ContentAlignmentV() AlignmentV

	SetLeft(left int)
	SetRight(right int)
	SetTop(top int)
	SetBottom(bottom int)
}

type Widget struct {
	id     string
	parent IContainer
	theme  *Theme

	enabled bool
	visible bool
	bounds  utils.Rect
	padding utils.Insets

	contentAlignmentH AlignmentH
	contentAlignmentV AlignmentV
}

func (w *Widget) Id() string         { return w.id }
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
func (w *Widget) Enabled() bool                 { return w.enabled }
func (w *Widget) Visible() bool                 { return w.visible }
func (w *Widget) Bounds() utils.Rect            { return w.bounds }
func (w *Widget) Padding() utils.Insets         { return w.padding }
func (w *Widget) InnerBounds() utils.Rect       { return w.bounds.ShrinkByInsets(w.padding) }
func (w *Widget) ContentWidth() int             { return 0 }
func (w *Widget) ContentHeight() int            { return 0 }
func (w *Widget) ContentAlignmentH() AlignmentH { return w.contentAlignmentH }
func (w *Widget) ContentAlignmentV() AlignmentV { return w.contentAlignmentV }

// Setters
func (w *Widget) SetVisible(v bool)         { w.visible = v }
func (w *Widget) SetEnabled(e bool)         { w.enabled = e }
func (w *Widget) SetBounds(b utils.Rect)    { w.bounds = b }
func (w *Widget) SetPadding(b utils.Insets) { w.padding = b }

func (w *Widget) SetLeft(left int)     { w.bounds.X = left }
func (w *Widget) SetRight(right int)   { w.bounds.X = right - w.bounds.W }
func (w *Widget) SetTop(top int)       { w.bounds.Y = top }
func (w *Widget) SetBottom(bottom int) { w.bounds.X = bottom - w.bounds.H }

func (w *Widget) SizeToContent() {
	w.bounds.W = w.ContentWidth() + w.padding.Left + w.padding.Right
	w.bounds.H = w.ContentHeight() + w.padding.Top + w.padding.Bottom
}

// Mouse handling
func (w *Widget) OnMouseCursorMoved(x, y float32) bool {
	return false
}
func (w *Widget) OnMouseButtonEvent(button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	return false
}
func (w *Widget) OnMouseScrolled(scrollX, scrollY float32) bool {
	return false
}

func (w *Widget) ToString() string {
	flags := ""
	if w.enabled {
		flags += "E"
	}
	if w.visible {
		flags += "V"
	}
	s := "Id='%s' bounds=%s padding=%s [%s]"
	return fmt.Sprintf(s, w.id, w.bounds.ToString(), w.padding.ToString(), flags)
}
