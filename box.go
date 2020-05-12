package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type BoxOrientation int

const (
	BoxHorizontalOrientation BoxOrientation = iota
	BoxVerticalOrientation
)

type BoxContainer struct {
	Container
	orientation    BoxOrientation
	widgetSpacing  int
	contentLength  int
	tempFlexValues []int
}

func NewBoxContainer(theme *Theme, orientation BoxOrientation, children ...IWidget) *BoxContainer {
	b := &BoxContainer{}
	b.Container.Init()
	b.tempFlexValues = make([]int, 0)
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
	for _, child := range c.children {
		_, ok := child.(IContainer)
		if ok {
			child.Layout()
		}
	}
	fixedSpace, flexSpace := c.computeContentSize()
	c.layoutChildren(fixedSpace, flexSpace)
	containerLayout(c)
}

func (c *BoxContainer) layoutChildren(fixedSpace int, flexSpace int) {
	if c.ChildrenCount() == 0 {
		return
	}

	// Cache for the flex values
	c.tempFlexValues = c.tempFlexValues[:0]
	for i := 0; i < c.ChildrenCount(); i++ {
		c.tempFlexValues = append(c.tempFlexValues, i)
		c.tempFlexValues[i] = c.children[i].Flex()
	}

	switch c.orientation {
	case BoxHorizontalOrientation:
		// TODO: It should consider the container's padding
		hasToRestart := true
		for hasToRestart {
			hasToRestart = false
			pos := 0
			spaceFree := utils.MaxI(c.bounds.W-fixedSpace, 0)
			maxHeight := 0
			for i, child := range c.children {
				if !child.Visible() {
					continue
				}
				child.SetLeft(pos)
				if c.tempFlexValues[i] > 0 {
					width := (c.tempFlexValues[i] * spaceFree) / flexSpace
					child.SetDimension(width, child.Bounds().H)
					if width <= child.MinimumWidth() || width >= child.MaximumWidth() {
						// A flexible component has reach one of its limits
						fixedSpace += child.Bounds().W
						flexSpace -= c.tempFlexValues[i]
						c.tempFlexValues[i] = 0
						hasToRestart = true
						break
					}
				} else {
					// NOP
				}
				maxHeight = utils.MaxI(maxHeight, child.Bounds().H)
				pos += child.Bounds().W
			}
			c.contentWidth = pos
			c.contentHeight = maxHeight
		}
	case BoxVerticalOrientation:
		// TODO: It should consider the container's padding
		pos := 0
		spaceFree := utils.MaxI(c.bounds.H-fixedSpace, 0)
		maxWidth := 0
		for _, child := range c.children {
			if !child.Visible() {
				continue
			}
			child.SetTop(pos)
			if child.Flex() > 0 {
				child.SetDimension(child.Bounds().W, (child.Flex()*spaceFree)/flexSpace)
			} else {
				// NOP
			}
			maxWidth = utils.MaxI(maxWidth, child.Bounds().W)
			pos += child.Bounds().H
		}
		c.contentWidth = maxWidth
		c.contentHeight = pos
	}
}

func (c *BoxContainer) computeContentSize() (int, int) {
	switch c.orientation {
	case BoxHorizontalOrientation:
		totalFixedSize := 0
		totalFlexSize := 0
		for _, child := range c.children {
			if !child.Visible() {
				continue
			}
			if child.Flex() > 0 {
				totalFlexSize += child.Flex()
			} else {
				totalFixedSize += child.Bounds().W
			}
		}
		return totalFixedSize, totalFlexSize
	case BoxVerticalOrientation:
		totalFixedSize := 0
		totalFlexSize := 0
		for _, child := range c.children {
			if !child.Visible() {
				continue
			}
			if child.Flex() > 0 {
				totalFlexSize += child.Flex()
			} else {
				totalFixedSize += child.Bounds().H
			}
		}
		return totalFixedSize, totalFlexSize
	}

	panic("computeContentSize() called with unknown orientation")
}
