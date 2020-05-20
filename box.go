package gui

import "github.com/maxfish/GoNativeUI-Core/utils"

type BoxOrientation int

const (
	BoxHorizontalOrientation BoxOrientation = 0
	BoxVerticalOrientation                  = 1
)

type BoxContainer struct {
	Container
	orientation              BoxOrientation
	spacing                  int
	contentLength            int
	childrenTotalFlexAmount  int
	childrenTotalFixedLength int
}

func NewBoxContainer(orientation BoxOrientation, children ...IWidget) *BoxContainer {
	b := &BoxContainer{}
	containerInit(b)
	b.children = make([]IWidget, 0, 16)
	b.orientation = orientation

	if children != nil {
		for _, child := range children {
			b.AddChild(child)
		}
	}
	return b
}

func (c *BoxContainer) Spacing() int           { return c.spacing }
func (c *BoxContainer) SetSpacing(spacing int) { c.spacing = spacing }

func (c *BoxContainer) Measure() {
	var mainLength, oppositeLength int
	mainIndex := c.orientation
	oppositeIndex := 1 - c.orientation
	var totalFlex, totalFixed int
	for _, child := range c.children {
		if !child.Visible() {
			continue
		}
		child.Measure()
		measuredLength := []int{child.MeasuredWidth(), child.MeasuredHeight()}
		minimumLength := []int{child.MinimumWidth(), child.MinimumHeight()}

		// From the second widget adds the container spacing
		if mainLength > 0 {
			mainLength += c.spacing
			totalFixed += c.spacing
		}

		mainLength += utils.MaxI(measuredLength[mainIndex], minimumLength[mainIndex])
		oppositeLength = utils.MaxI(oppositeLength, utils.MaxI(measuredLength[oppositeIndex], minimumLength[oppositeIndex]))
		if child.MeasuredFlex() == 0 {
			totalFixed += measuredLength[mainIndex]
		} else {
			totalFlex += child.MeasuredFlex()
		}
	}
	c.measuredFlex = c.flex
	c.childrenTotalFixedLength = totalFixed
	c.childrenTotalFlexAmount = totalFlex

	lengths := []int{mainLength, oppositeLength}
	c.measuredWidth = lengths[c.orientation] + c.style.Padding.Left + c.style.Padding.Right
	c.measuredHeight = lengths[1-c.orientation] + c.style.Padding.Top + c.style.Padding.Bottom
}

func (c *BoxContainer) Layout() {
	c.Measure()
	c.layoutChildren()
}

func (c *BoxContainer) layoutChildren() {
	if c.ChildrenCount() == 0 {
		return
	}

	flexSpace := c.childrenTotalFlexAmount
	switch c.orientation {
	case BoxHorizontalOrientation:
		fixedSpace := c.childrenTotalFixedLength
		totalSpace := c.InnerBounds().W
		hasToRestart := true
		for hasToRestart {
			hasToRestart = false
			offset := c.style.Padding.Left
			spaceFree := utils.MaxI(totalSpace-fixedSpace, 0)
			maxHeight := 0
			for _, child := range c.children {
				if !child.Visible() {
					continue
				}
				child.SetLeft(offset)
				if child.MeasuredFlex() > 0 {
					width := (child.MeasuredFlex() * spaceFree) / flexSpace
					child.SetWidth(width)
					if width <= child.MinimumWidth() || (child.MaximumWidth() > 0 && width >= child.MaximumWidth()) {
						// A flexible component has reached one of its limits
						fixedSpace += child.Bounds().W
						flexSpace -= child.MeasuredFlex()
						child.SetMeasuredFlex(0)
						child.SetMeasuredWidth(child.Bounds().W)
						hasToRestart = true
						break
					}
				} else {
					child.SetWidth(child.MeasuredWidth())
				}

				if child.Stretch() > 0 {
					child.SetHeight(c.Bounds().H)
				} else {
					child.SetHeight(child.MeasuredHeight())
				}

				maxHeight = utils.MaxI(maxHeight, child.Bounds().H)
				offset += child.Bounds().W + c.spacing
			}
		}
	case BoxVerticalOrientation:
		fixedSpace := c.childrenTotalFixedLength
		totalSpace := c.InnerBounds().H
		hasToRestart := true
		for hasToRestart {
			hasToRestart = false
			offset := c.style.Padding.Top
			spaceFree := utils.MaxI(totalSpace-fixedSpace, 0)
			for _, child := range c.children {
				if !child.Visible() {
					continue
				}
				child.SetTop(offset)
				if child.MeasuredFlex() > 0 {
					height := (child.MeasuredFlex() * spaceFree) / flexSpace
					child.SetHeight(height)
					if height <= child.MinimumHeight() || (child.MaximumHeight() > 0 && height >= child.MaximumHeight()) {
						// A flexible component has reached one of its limits
						fixedSpace += child.Bounds().H
						flexSpace -= child.MeasuredFlex()
						child.SetMeasuredFlex(0)
						child.SetMeasuredHeight(child.Bounds().H)
						hasToRestart = true
						break
					}
				} else {
					child.SetHeight(child.MeasuredHeight())
				}

				if child.Stretch() > 0 {
					child.SetWidth(c.Bounds().W)
				} else {
					child.SetWidth(child.MeasuredWidth())
				}

				offset += child.Bounds().H + c.spacing
			}
		}
	}

	for _, child := range c.children {
		child.Layout()
	}
}
