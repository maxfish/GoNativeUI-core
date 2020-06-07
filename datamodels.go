package gui

type ListModel interface {
	ItemAt(view *ListView, row int) string
	ItemsMaxWidth(view *ListView) int
	ItemHeight(view *ListView) int
	NumItems(view *ListView) int
}

type StringListModel struct {
	items []string
}

func NewStringListModel(items []string) *StringListModel {
	return &StringListModel{items: items}
}

func (d *StringListModel) ItemAt(view *ListView, row int) string {
	return d.items[row]
}

func (d *StringListModel) ItemsMaxWidth(view *ListView) int {
	// This is slow but accurate
	maxWidth := 0
	for _, s := range d.items {
		size := view.style.Font.TextSize(view.style.FontSize, s)
		if size.W() > maxWidth {
			maxWidth = size.W()
		}
	}
	return maxWidth
}

func (d *StringListModel) ItemHeight(view *ListView) int {
	return view.style.Font.LineHeight(view.style.FontSize)
}

func (d *StringListModel) NumItems(view *ListView) int {
	return len(d.items)
}
