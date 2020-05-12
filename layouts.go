package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
	"log"
)

type DistributionH uint32

const (
	DistributionHFill DistributionH = iota
	DistributionHLeft
	DistributionHCenter
	DistributionHRight
)

func HBoxLayout(container IContainer, alignH utils.AlignmentH, spacing int) {
	numChildren := container.ChildrenCount()
	left := container.InnerBounds().X

	switch alignH {
	case utils.AlignmentHLeft:
		for i := 0; i < numChildren; i++ {
			c := container.Children()[i]
			c.SetLeft(left)
			left += c.Bounds().W + spacing
		}
	case utils.AlignmentHRight:
		right := container.InnerBounds().W
		for i := numChildren - 1; i >= 0; i-- {
			c := container.Children()[i]
			//c.SetRight(right)
			right -= c.Bounds().W + spacing
		}
	}
}

// TODO: this needs to be optimized
func HGridLayout(container IContainer, numColumns int, alignH []utils.AlignmentH, percentages []float32) {
	if numColumns != len(alignH) || numColumns != len(percentages) {
		log.Panicf("HGridLayout parameters don't have all the same length")
	}

	numChildren := container.ChildrenCount()
	containerY := container.InnerBounds().Y
	containerW := float32(container.InnerBounds().W)
	columnWidths := make([]int, numColumns)
	maxContentWidths := make([]int, numColumns)
	totalPercentages := float32(0)

	for i := 0; i < numChildren; i++ {
		column := i % numColumns
		maxContentWidths[column] = utils.MaxI(container.Children()[i].Bounds().W, maxContentWidths[column])
	}
	for i := 0; i < numColumns; i++ {
		totalPercentages += percentages[i]
		columnWidths[i] = int((containerW / totalPercentages) * percentages[i])
	}
	totalWidth := 0
	for i := 0; i < numColumns; i++ {
		columnWidths[i] = utils.MaxI(maxContentWidths[i], columnWidths[i])
		totalWidth += columnWidths[i]
	}
	// Enforces the total width of the parent
	if float32(totalWidth) > containerW {
		scaleFactor := containerW / float32(totalWidth)
		for i := 0; i < numColumns; i++ {
			columnWidths[i] = int(float32(columnWidths[i]) * scaleFactor)
		}
	}

	top := containerY
	numRows := numChildren/numColumns + 1
	for row := 0; row < numRows; row++ {
		rowHeight := 0
		left := 0
		for col := 0; col < numColumns; col++ {
			if row*numColumns+col >= numChildren {
				return
			}
			//cellRect := utils.Rect{left, top, columnWidths[col], 1}
			c := container.Children()[row*numColumns+col]
			//c.SetBounds(c.Bounds().AlignIn(cellRect, utils.Alignment{alignH[col], utils.AlignmentVTop}))
			rowHeight = utils.MaxI(rowHeight, c.Bounds().H)
			left += columnWidths[col]
		}
		top += rowHeight
	}
}

func HorizontalLayout(container IContainer, ) {

}
