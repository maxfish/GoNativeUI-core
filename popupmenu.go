package gui

type PopupMenu struct {
	Popup
	listView             *ListView
	ItemSelectedCallback ItemSelectedCallback
}

func NewPopupMenu(dataModel ListModel, itemSelectedCallback ItemSelectedCallback) *PopupMenu {
	m := &PopupMenu{}
	m.listView = NewListView(dataModel, m.onItemSelected)
	// TODO: Assign a menu style to the list
	m.listView.SetStyle(CurrentGui().Theme().ListView)
	container := NewBoxContainer(BoxVerticalOrientation, m.listView)
	container.Layout()
	m.widget = container
	m.ItemSelectedCallback = itemSelectedCallback
	return m
}

func (m *PopupMenu) onItemSelected(source interface{}, index int) {
	// The ListView selected item is not useful in this context
	m.listView.selectedIndex = -1
	CurrentGui().DismissPopup()
	m.ItemSelectedCallback(m, index)
}
