package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type IContainer interface {
	IWidget

	//Update(deltaMs uint32) bool

	ContentRect() utils.Rect
	//adaptWidthToComponents()
	//adaptHeightToComponents()

	// Children components
	Children() []IWidget
	ChildrenCount() int
	AddChild(c IWidget)
	AddChildren(children ...IWidget)
	AddChildAtIndex(c IWidget, i int32)
	//RemoveChildByIs(i string)
	//GetMaximumComponentSize(c IWidget)
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
	c.visible = true
	c.enabled = true
	c.children = make([]IWidget, 0, 16)
	c.childrenMap = make(map[string]IWidget)
	c.theme = theme

	if children != nil {
		for _, child := range children {
			c.AddChild(child)
		}
	}
	return c
}

//
//func (c *Container) ContentWidth() int {
//	contentRect := utils.Rect{}
//	for _, child := range c.children {
//		contentRect = contentRect.UnionWith(child.Bounds())
//	}
//	return contentRect.W
//}
//
//func (c *Container) ContentHeight() int {
//	contentRect := utils.Rect{}
//	for _, child := range c.children {
//		contentRect = contentRect.UnionWith(child.Bounds())
//	}
//	return contentRect.H
//}

func (c *Container) ContentRect() utils.Rect {
	contentRect := utils.Rect{}
	for _, child := range c.children {
		switch childCast := child.(type) {
		case IContainer:
			contentRect = contentRect.UnionWith(childCast.ContentRect())
		case IWidget:
			contentRect = contentRect.UnionWith(child.Bounds())
		}
	}
	return contentRect
}

func (c *Container) WrapContent() {
	c.bounds = c.ContentRect()
}

// Children
func (c *Container) Children() []IWidget { return c.children }
func (c *Container) ChildrenCount() int  { return len(c.children) }
func (c *Container) AddChild(child IWidget) {
	c.children = append(c.children, child)
	child.SetParent(c)
	child.SetTheme(c.theme)
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
	for i := 0; i < len(c.children); i++ {
		if c.children[i].Id() == id {
			return c.children[i]
		}
	}
	return nil
}

func (c *Container) RemoveChildById(id string)             {}
func (c *Container) GetMaximumComponentSize(child IWidget) {}
func (c *Container) FindChildAt(x, y int) IWidget {
	for j := 0; j < len(c.children); j++ {
		child := c.children[j]
		if child.Bounds().ContainsPoint(x, y) {
			return child
		}
	}
	return nil
}

// Mouse handling
func (c *Container) OnMouseCursorMoved(x float32, y float32) bool {
	return false
}

func (c *Container) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	var child IWidget = nil
	for j := 0; j < len(c.children); j++ {
		oneChild := c.children[j]
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
