package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type BoxOrientation int

const (
	BoxHorizontalOrientation BoxOrientation = iota
	BoxVerticalOrientation
)

type BoxContainer struct {
	Container
	orientation   BoxOrientation
	widgetSpacing int
	offset        int
	contentLength int
}

func NewBoxContainer(theme *Theme, orientation BoxOrientation, children ...IWidget) *BoxContainer {
	b := &BoxContainer{}
	b.Container.Init()
	b.theme = theme
	b.orientation = orientation

	if children != nil {
		for _, child := range children {
			b.AddChild(child)
		}
	}
	return b
}

func (c *BoxContainer) Layout() {
	c.computeContentSize()
	containerLayout(c)
	c.layoutChildren()
}

func (c *BoxContainer) layoutChildren() {
	if c.ChildrenCount() == 0 {
		return
	}
	// TODO: It should consider the container's padding
	pos := c.offset
	for _, child := range c.children {
		if child.Visible() {
			child.Layout()
			switch c.orientation {
			case BoxHorizontalOrientation:
				child.SetLeft(pos)
				pos += child.Bounds().W + c.widgetSpacing
			case BoxVerticalOrientation:
				child.SetTop(pos)
				pos += child.Bounds().H + c.widgetSpacing
			}
		}
	}

	c.contentLength = pos - c.offset - c.widgetSpacing
}

func (c *BoxContainer) computeContentSize() {
	if c.contentSizeValid {
		return
	}

	w := 0
	h := 0
	switch c.orientation {
	case BoxHorizontalOrientation:
		for _, child := range c.children {
			if child.Visible() {
				w += child.PreferredWidth() + c.widgetSpacing
				h = utils.MaxI(child.PreferredHeight(), h)
			}
		}
		if w > 0 {
			w -= c.widgetSpacing
		}
	case BoxVerticalOrientation:
		for _, child := range c.children {
			if child.Visible() {
				h += child.PreferredHeight()
			}
		}
		if h > 0 {
			h = h + (c.ChildrenCount()-1)*c.widgetSpacing
		}
		maxWidth := 0
		for _, child := range c.children {
			maxWidth = utils.MaxI(child.PreferredWidth(), maxWidth)
		}
		w = maxWidth
	}

	c.contentWidth = w
	c.contentHeight = h
	c.contentSizeValid = true
}
