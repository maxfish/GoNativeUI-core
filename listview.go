package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

const (
	DefaultRowWidth    = 150
	DefaultRowHeight   = 20
	DefaultVisibleRows = 10
)

type ItemSelectedCallback func(source interface{}, index int)

type ListView struct {
	Widget
	dataModel            ListModel
	offset               float32
	visibleRows          int
	hoveredIndex         int
	selectedIndex        int
	itemSelectedCallback ItemSelectedCallback
}

func NewListView(dataModel ListModel, itemSelectedCallback ...ItemSelectedCallback) *ListView {
	l := &ListView{}
	widgetInit(l)
	l.style = CurrentGui().Theme().ListView
	l.visibleRows = DefaultVisibleRows
	l.hoveredIndex = -1
	l.selectedIndex = -1
	l.dataModel = dataModel
	if len(itemSelectedCallback) == 1 {
		l.itemSelectedCallback = itemSelectedCallback[0]
	}
	return l
}

func (l *ListView) Offset() int                      { return int(l.offset) }
func (l *ListView) DataModel() ListModel             { return l.dataModel }
func (l *ListView) SetDataModel(dataModel ListModel) { l.dataModel = dataModel }
func (l *ListView) SelectedIndex() int               { return l.selectedIndex }
func (l *ListView) HoveredIndex() int                { return l.hoveredIndex }
func (l *ListView) SetSelectedIndex(index int)       { l.selectedIndex = index }

func (l *ListView) SetOnItemSelectedCallback(f ItemSelectedCallback) {
	l.itemSelectedCallback = f
}

func (l *ListView) fireSelectionChangedEvent(index int) {
	if l.itemSelectedCallback != nil {
		l.itemSelectedCallback(l, index)
	}
}

func (l *ListView) Measure() {
	l.computeContentSize()
	l.measuredWidth = l.contentWidth + l.style.Padding.Left + l.style.Padding.Right
	l.measuredHeight = utils.MinI(l.visibleRows*l.dataModel.ItemHeight(l), l.contentHeight) + l.style.Padding.Top + l.style.Padding.Bottom
	l.measuredFlex = l.flex
}

func (l *ListView) computeContentSize() {
	if l.dataModel != nil {
		l.contentWidth = l.dataModel.ItemsMaxWidth(l)
		l.contentHeight = l.dataModel.NumItems(l) * l.dataModel.ItemHeight(l)
	} else {
		l.contentWidth = DefaultRowWidth
		l.contentHeight = l.visibleRows * DefaultRowHeight
	}
}

func (l *ListView) hoveredIndexFromCoords(x, y float32) {
	l.hoveredIndex = utils.ClampI((int(y+l.offset))/l.dataModel.ItemHeight(l), 0, l.dataModel.NumItems(l)-1)
}

func (l *ListView) OnMouseEvent(event MouseEvent) IWidget {
	if event.Type == MouseEventPosition {
		// Updates the hovered item index
		l.hoveredIndexFromCoords(event.X, event.Y)
		return l
	}
	if event.Type == MouseEventWheel {
		l.offset = utils.Clamp(l.offset-event.ScrollY, 0, float32(l.contentHeight-l.InnerBounds().H))
		l.hoveredIndexFromCoords(event.X, event.Y)
		return l
	}
	if event.Type == MouseEventButton {
		if event.Button != MouseButtonLeft {
			return nil
		}
		if event.Action == EventActionPress {
			if l.dataModel != nil {
				newIndex := l.hoveredIndex
				if newIndex != l.selectedIndex {
					l.selectedIndex = newIndex
					l.fireSelectionChangedEvent(l.selectedIndex)
				}
			}
			return l
		} else if event.Action == EventActionRelease {
			return l
		}
	}

	return nil
}
