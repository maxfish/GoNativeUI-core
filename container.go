package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IContainer interface {
	IWidget

	//Update(deltaMs uint32) bool

	//ContentRect() utils.Rect
	//adaptWidthToComponents()
	//adaptHeightToComponents()

	// Children components
	Children() []IWidget
	ChildrenCount() int
	AddChild(c IWidget)
	AddChildren(children ...IWidget)
	AddChildAtIndex(c IWidget, i int32)
	FindChildAt(x int, y int) IWidget
}

type Container struct {
	Widget

	// Children
	children    []IWidget
	childrenMap map[string]IWidget
}

func NewContainer(theme *Theme, children ...IWidget) IContainer {
	c := &Container{}
	c.Init()
	c.theme = theme

	if children != nil {
		for _, child := range children {
			c.AddChild(child)
		}
	}
	return c
}

func (c *Container) Init() {
	widgetInit(c)
	c.children = make([]IWidget, 0, 16)
	c.childrenMap = make(map[string]IWidget)
}

// Children
func (c *Container) Children() []IWidget { return c.children }
func (c *Container) ChildrenCount() int  { return len(c.children) }
func (c *Container) AddChild(child IWidget) {
	c.children = append(c.children, child)
	child.SetParent(c)
	child.SetTheme(c.theme)
	//child.Layout()
	c.contentSizeValid = false
}
func (c *Container) AddChildren(children ...IWidget) {
	for _, child := range children {
		c.AddChild(child)
	}
}

func (c *Container) AddChildAtIndex(child IWidget, i int32) {
	s := append(c.children, nil)
	copy(s[i+1:], s[i:])
	s[i] = child
	c.children = s
}

func (c *Container) ChildById(id string) IWidget {
	for _, child := range c.children {
		if child.Id() == id {
			return child
		}
	}
	return nil
}

func (c *Container) RemoveChildById(id string)             {}
func (c *Container) GetMaximumComponentSize(child IWidget) {}
func (c *Container) FindChildAt(x, y int) IWidget {
	for _, child := range c.children {
		if child.Bounds().ContainsPoint(x, y) {
			return child
		}
	}
	return nil
}

// Layout
func (c *Container) Layout() {
	c.computeContentSize()
	widgetLayout(c)
	c.layoutChildren()
}

func (c *Container) layoutChildren() {
	for _, child := range c.children {
		child.Layout()
		// Force container alignment
		child.SetBounds(child.Bounds().AlignIn(
			c.InnerBounds(),
			utils.Alignment{Horizontal: c.contentAlignmentH, Vertical: c.contentAlignmentV},
		))
	}
}

// Mouse handling
func (c *Container) OnMouseCursorMoved(x float32, y float32) bool {
	return false
}

func (c *Container) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	var child IWidget = nil
	for _, oneChild := range c.children {
		if !(oneChild.Visible() && oneChild.Enabled()) {
			continue
		}
		if oneChild.Bounds().ContainsPoint(int(x), int(y)) {
			child = oneChild
			break
		}
	}
	if child == nil {
		return false
	}
	consumed := false
	_, ok := child.(*Container)
	if ok {
		consumed = child.OnMouseButtonEvent(x-float32(child.Bounds().X), y-float32(child.Bounds().Y), button, event, modifiers)
	} else {
		consumed = child.OnMouseButtonEvent(x, y, button, event, modifiers)
	}
	return consumed
}

func (c *Container) OnMouseScrolled(scrollX float32, scrollY float32) bool {
	return false
}

func (c *Container) computeContentSize() {
	contentRect := utils.Rect{}
	for _, child := range c.children {
		switch childCast := child.(type) {
		case IContainer:
			contentRect = contentRect.UnionWith(childCast.Bounds())
		case IWidget:
			contentRect = contentRect.UnionWith(child.Bounds())
		}
	}
	c.contentWidth = contentRect.W
	c.contentHeight = contentRect.H
}

func containerLayout(c IContainer) {
	bounds := c.Bounds()
	bounds.W = utils.MaxI(c.PreferredWidth(), c.ContentWidth())
	bounds.H = c.PreferredHeight()
	c.SetBounds(bounds)
}
