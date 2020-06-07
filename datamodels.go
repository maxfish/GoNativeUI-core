package gui

type ListModel interface {
	ItemAt(widget IWidget, row int) string
	ItemsMaxWidth(widget IWidget) int
	ItemHeight(widget IWidget) int
	NumItems(widget IWidget) int
}

type StringListModel struct {
	items []string
}

func NewStringListModel(items []string) *StringListModel {
	return &StringListModel{items: items}
}

func (d *StringListModel) ItemAt(widget IWidget, row int) string {
	return d.items[row]
}

func (d *StringListModel) ItemsMaxWidth(widget IWidget) int {
	// This is slow but accurate
	maxWidth := 0
	for _, s := range d.items {
		size := widget.Style().Font.TextSize(widget.Style().FontSize, s)
		if size.W() > maxWidth {
			maxWidth = size.W()
		}
	}
	return maxWidth
}

func (d *StringListModel) ItemHeight(widget IWidget) int {
	return widget.Style().Font.LineHeight(widget.Style().FontSize)
}

func (d *StringListModel) NumItems(widget IWidget) int {
	return len(d.items)
}
