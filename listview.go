package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

const (
	DefaultRowWidth    = 150
	DefaultRowHeight   = 20
	DefaultVisibleRows = 20
)

type ListView struct {
	Widget
	dataModel     ListModel
	offset        float32
	visibleRows   int
	selectedIndex int
}

func NewListView() *ListView {
	l := &ListView{}
	widgetInit(l)
	l.style = CurrentGui().Theme().ListView
	l.visibleRows = DefaultVisibleRows
	l.selectedIndex = -1
	return l
}

func (l *ListView) Measure() {
	l.measuredWidth = DefaultRowWidth
	l.measuredFlex = l.flex

	if l.dataModel != nil {
		l.measuredHeight = l.visibleRows * l.dataModel.ItemHeight(l)
	} else {
		l.measuredHeight = l.visibleRows * DefaultRowHeight
	}
}

func (l *ListView) Offset() int                      { return int(l.offset) }
func (l *ListView) DataModel() ListModel             { return l.dataModel }
func (l *ListView) SetDataModel(dataModel ListModel) { l.dataModel = dataModel }
func (l *ListView) SelectedIndex() int               { return l.selectedIndex }
func (l *ListView) SetSelectedIndex(index int)       { l.selectedIndex = index }

func (l *ListView) totalHeight() int {
	if l.dataModel == nil {
		return 0
	}
	return l.dataModel.NumItems(l) * l.dataModel.ItemHeight(l)
}

func (l *ListView) OnMouseCursorMoved(x, y float32) bool {
	return false
}

func (l *ListView) OnMouseButtonEvent(x float32, y float32, button ButtonIndex, event EventAction, modifiers ModifierKey) bool {
	if button != MouseButtonLeft {
		return false
	}
	if event == EventActionPress {
		if l.dataModel != nil {
			l.selectedIndex = utils.ClampI((int(y+l.offset))/l.dataModel.ItemHeight(l), 0, l.dataModel.NumItems(l)-1)
		}
		return true
	} else if event == EventActionRelease {
		return true
	}
	return false
}

func (l *ListView) OnMouseScrolled(x float32, y float32, scrollX, scrollY float32) bool {
	l.offset = utils.Clamp(l.offset-scrollY, 0, float32(l.totalHeight()-l.InnerBounds().H))
	return true
}
