package gui

type IContainer interface {
	IWidget

	Children() []IWidget
	ChildrenCount() int
	AddChild(c IWidget)
	AddChildren(children ...IWidget)
	AddChildAtIndex(c IWidget, i int32)
	RemoveChildById(id string)
	RemoveChildAtIndex(index int)
	ChildById(id string) IWidget
	IndexOfChild(aChild IWidget) int
	ChildAtIndex(index int) IWidget
	FindChildAt(x int, y int) IWidget
}

type Container struct {
	Widget

	children []IWidget
}

func (c *Container) initStyle() {
	c.style = WidgetStyle{}
}

// Children
func (c *Container) Children() []IWidget { return c.children }
func (c *Container) ChildrenCount() int  { return len(c.children) }

func (c *Container) AddChild(child IWidget) {
	c.children = append(c.children, child)
	child.setParent(c)
}

func (c *Container) AddChildren(children ...IWidget) {
	for _, child := range children {
		c.AddChild(child)
	}
}

func (c *Container) AddChildAtIndex(child IWidget, i int32) {
	child.setParent(c)

	s := append(c.children, nil)
	copy(s[i+1:], s[i:])
	s[i] = child
	c.children = s
}

func (c *Container) ChildAtIndex(index int) IWidget {
	if index < 0 || index >= len(c.children) {
		return nil
	}
	return c.children[index]
}

func (c *Container) ChildById(id string) IWidget {
	for _, child := range c.children {
		if child.Id() == id {
			return child
		}
	}
	return nil
}

func (c *Container) IndexOfChild(aChild IWidget) int {
	for index, child := range c.children {
		if child == aChild {
			return index
		}
	}
	return -1
}

func (c *Container) RemoveChildById(id string) {
	toBeRemovedIndex := -1
	for i, child := range c.children {
		if child.Id() == id {
			toBeRemovedIndex = i
		}
	}
	if toBeRemovedIndex >= 0 {
		c.RemoveChildAtIndex(toBeRemovedIndex)
	}
}

func (c *Container) RemoveChildAtIndex(index int) {
	CurrentGui().RemoveFocusFrom(c.children[index])
	ret := make([]IWidget, 0)
	ret = append(ret, c.children[:index]...)
	c.children = append(ret, c.children[index+1:]...)
}

func (c *Container) FindChildAt(x, y int) IWidget {
	for _, child := range c.children {
		if child.Bounds().ContainsPoint(x, y) {
			return child
		}
	}
	return nil
}

// Mouse handling
func (c *Container) OnMouseCursorMoved(x float32, y float32) bool {
	for _, oneChild := range c.children {
		if !(oneChild.Visible() && oneChild.Enabled()) {
			continue
		}
		if oneChild.Bounds().ContainsPoint(int(x), int(y)) {
			return oneChild.OnMouseCursorMoved(x-float32(oneChild.Bounds().X), y-float32(oneChild.Bounds().Y))
		}
	}

	return false
}

func (c *Container) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, action EventAction, modifierKey ModifierKey) bool {
	for _, oneChild := range c.children {
		if !(oneChild.Visible() && oneChild.Enabled()) {
			continue
		}
		if oneChild.Bounds().ContainsPoint(int(x), int(y)) {
			return oneChild.OnMouseButtonEvent(x-float32(oneChild.Bounds().X), y-float32(oneChild.Bounds().Y), button, action, modifierKey)
		}
	}

	CurrentGui().RemoveFocus()
	return false
}

func (c *Container) OnMouseScrolled(x float32, y float32, scrollX float32, scrollY float32) bool {
	for _, oneChild := range c.children {
		if !(oneChild.Visible() && oneChild.Enabled()) {
			continue
		}
		if oneChild.Bounds().ContainsPoint(int(x), int(y)) {
			return oneChild.OnMouseScrolled(x-float32(oneChild.Bounds().X), y-float32(oneChild.Bounds().Y), scrollX, scrollY)
		}
	}
	return false
}

func containerInit(c IContainer) {
	c.SetEnabled(true)
	c.SetVisible(true)
}
