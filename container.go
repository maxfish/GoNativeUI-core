package gui

type IContainer interface {
	IWidget

	//KeyEvent(eventType int32, key int32) bool
	//Update(deltaMs uint32) bool

	//RequestLayout()
	//layout()
	//layoutComponents()
	//adaptWidthToComponents()
	//adaptHeightToComponents()

	// Children components
	Children() []IWidget
	ChildrenCount() int
	AddChild(c IWidget)
	AddChildren(children ...IWidget)
	AddChildAtIndex(c IWidget, i int32)
	//DescendantById(Id string)
	//ComponentByPath(path string)
	//ComponentByRelativeId(Id string)
	//RemoveChildByIs(i string)
	//GetMaximumComponentSize(c IWidget)
	FindChildAt(x int, y int) IWidget

	//SetAlignment(alignmentH ui.AlignmentH, alignmentV ui.AlignmentV)
}

type Container struct {
	Widget

	// Children
	children    []IWidget
	childrenMap map[string]IWidget
}

func NewContainer() IContainer {
	c := &Container{}
	c.children = make([]IWidget, 0, 16)
	c.childrenMap = make(map[string]IWidget)
	return c
}

func (c *Container) SizeToContent() {
	panic("implement me")
}

// Children
func (c *Container) Children() []IWidget { return c.children }
func (c *Container) ChildrenCount() int  { return len(c.children) }
func (c *Container) AddChild(child IWidget) {
	c.children = append(c.children, child)
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

//func (c *Widget) DescendantById(Id string)
//func (c *Widget) ComponentByPath(path string)
//func (c *Widget) ComponentByRelativeId(Id string)
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
		// FIXME: This won't work on nested containers. It needs to translate to the parent's origin
		if oneChild.Bounds().ContainsPoint(int(x), int(y)) {
			child = oneChild
			break
		}
	}
	if child == nil {
		return false
	}
	consumed := child.OnMouseButtonEvent(x, y, button, event, modifiers)
	return consumed
}

func (c *Container) OnMouseScrolled(scrollX float32, scrollY float32) bool {
	return false
}
