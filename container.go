package gui

type IContainer interface {
	IWidget

	Children() []IWidget
	ChildrenCount() int
	AddChild(c IWidget)
	AddChildren(children ...IWidget)
	AddChildAtIndex(c IWidget, i int32)
	FindChildAt(x int, y int) IWidget
	RemoveChildById(id string)
	RemoveChildAtIndex(index int)
}

type Container struct {
	Widget

	children []IWidget
}

func (c *Container) Init() {
	widgetInit(c)
	c.children = make([]IWidget, 0, 16)
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
	child.SetParent(c)
	child.SetTheme(c.theme)

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

func (c *Container) SetTheme(t *Theme) {
	for _, child := range c.children {
		child.SetTheme(t)
	}
}

// Layout
func (c *Container) Layout() {
	panic("Layout() shouldn't be called on the base Controller")
}

func (c *Container) layoutChildren() {
	panic("layoutChildren() shouldn't be called on the base Controller")
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
