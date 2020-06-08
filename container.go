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

func (c *Container) OnMouseEvent(event MouseEvent) IWidget {
	for _, oneChild := range c.children {
		if !(oneChild.Visible() && oneChild.Enabled()) {
			continue
		}
		if oneChild.Bounds().ContainsPoint(int(event.X), int(event.Y)) {
			listener, ok := oneChild.(IMouseListener)
			if ok {
				e := event
				e.X -= float32(oneChild.Bounds().X)
				e.Y -= float32(oneChild.Bounds().Y)
				return listener.OnMouseEvent(e)
			}
		}
	}
	return nil
}

func containerInit(c IContainer) {
	c.SetEnabled(true)
	c.SetVisible(true)
}
