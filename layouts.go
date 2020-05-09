package gui

import (
	"github.com/maxfish/GoNativeUI-Core/utils"
)

type AlignmentH uint32

const (
	AlignmentHLeft AlignmentH = iota
	AlignmentHCenter
	AlignmentHRight
)

type AlignmentV uint32

const (
	AlignmentVTop AlignmentV = iota
	AlignmentVCenter
	AlignmentVBottom
)

type Alignment struct {
	horizontal AlignmentH
	vertical   AlignmentV
}

type DistributionH uint32

const (
	DistributionHFill DistributionH = iota
	DistributionHLeft
	DistributionHCenter
	DistributionHRight
)

func HBoxLayout(container IContainer, alignH AlignmentH, spacing int) {
	numChildren := container.ChildrenCount()
	left := container.InnerBounds().X

	switch alignH {
	case AlignmentHLeft:
		for i := 0; i < numChildren; i++ {
			c := container.Children()[i]
			c.SetLeft(left)
			left += c.Bounds().W + spacing
		}
	case AlignmentHRight:
		right := container.InnerBounds().W
		for i := numChildren - 1; i >= 0; i-- {
			c := container.Children()[i]
			c.SetRight(right)
			right -= c.Bounds().W + spacing
		}
	}
}

// TODO: this needs to be optimized
func HGridLayout(container IContainer, numColumns int, alignH []AlignmentH, percentages []float32) {
	numChildren := container.ChildrenCount()
	containerY := container.InnerBounds().Y
	containerW := float32(container.InnerBounds().W)
	columnWidths := make([]int, numColumns)
	maxContentWidths := make([]int, numColumns)
	totalPercentages := float32(0)

	for i := 0; i < numColumns; i++ {
		totalPercentages += percentages[i]
	}
	for i := 0; i < numColumns; i++ {
		columnWidths[i] = int((containerW / totalPercentages) * percentages[i])
	}
	for i := 0; i < numChildren; i++ {
		column := i % numColumns
		maxContentWidths[column] = utils.MaxI(container.Children()[i].Bounds().W, maxContentWidths[column])
	}
	for i := 0; i < numChildren; i++ {
		column := i % numColumns
		maxContentWidths[column] = utils.MaxI(container.Children()[i].Bounds().W, maxContentWidths[column])
	}
	for col := 0; col < numColumns; col++ {
		columnWidths[0] = utils.MaxI(maxContentWidths[col], columnWidths[col])
	}

	top := containerY
	numRows := numChildren / numColumns + 1
	for row := 0; row < numRows; row++ {
		rowHeight := 0
		left := 0
		for col := 0; col < numColumns; col++ {
			if row*numColumns+col >= numChildren {
				return
			}
			c := container.Children()[row*numColumns+col]
			c.SetLeft(left)
			c.SetTop(top)
			rowHeight = utils.MaxI(rowHeight, c.Bounds().H)
			left += columnWidths[col]
		}
		top += rowHeight
	}
}

func HorizontalLayout(container IContainer, ) {

}
