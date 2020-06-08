package gui

type ComboBoxChangeCallback func(index int)

type ComboBox struct {
	Label
	options       []string
	selectedIndex int
	popup         IPopup

	open bool

	onChangeCallback ComboBoxChangeCallback
}

func NewComboBox(options []string, selectedIndex int, changeCallback ...ComboBoxChangeCallback) *ComboBox {
	c := &ComboBox{}
	widgetInit(c)
	c.style = CurrentGui().Theme().ComboBox
	c.options = options
	c.selectedIndex = selectedIndex
	c.text = options[selectedIndex]
	if len(changeCallback) == 1 {
		c.onChangeCallback = changeCallback[0]
	}

	c.initPopup()
	return c
}

func (c *ComboBox) initPopup() {
	model := NewStringListModel(c.options)
	c.popup = NewPopupMenu(model, func(source interface{}, index int) {
		if source == c.popup {
			c.SetSelectedIndex(index)
		}
	})
}

func (c *ComboBox) SelectedIndex() int {
	return c.selectedIndex
}

func (c *ComboBox) SetSelectedIndex(index int) {
	if index == c.selectedIndex {
		return
	}
	c.selectedIndex = index
	c.text = c.options[index]
	c.fireChangeEvent(index)
}

func (c *ComboBox) SetOnChangeCallback(callback ComboBoxChangeCallback) {
	c.onChangeCallback = callback
}

func (c *ComboBox) Open() bool {
	return c.popup.Visible()
}

func (c *ComboBox) fireChangeEvent(selectedIndex int) {
	if c.onChangeCallback != nil {
		c.onChangeCallback(selectedIndex)
	}
}

func (c *ComboBox) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if button != MouseButtonLeft {
		return false
	}
	if event == EventActionPress {
		if c.popup.Visible() {
			CurrentGui().DismissPopup()
		} else {
			absPosition := c.AbsolutePosition()
			c.popup.Widget().SetLeft(absPosition.X())
			c.popup.Widget().SetTop(absPosition.Y())
			CurrentGui().ShowPopup(c.popup)
		}
		return true
	}
	return false
}
