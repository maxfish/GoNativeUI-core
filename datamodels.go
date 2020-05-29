package gui

type ListModel interface {
	ItemAt(view *ListView, row int) string
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

func (d *StringListModel) ItemHeight(view *ListView) int {
	return view.style.Font.LineHeight(view.style.FontSize)
}

func (d *StringListModel) NumItems(view *ListView) int {
	return len(d.items)
}
