package utils

type MeasureUnit int

const (
	NoUnit MeasureUnit = iota
	PixelUnit
	PercentageUnit
)

type Dimension struct {
	Value int
	Unit  MeasureUnit
}

func Pixels(value int) Dimension {
	return Dimension{value, PixelUnit}
}
func Percentage(value int) Dimension {
	return Dimension{value, PercentageUnit}
}

type AlignmentH uint32

const (
	AlignmentHNone AlignmentH = iota
	AlignmentHLeft
	AlignmentHCenter
	AlignmentHRight
)

type AlignmentV uint32

const (
	AlignmentVNone AlignmentV = iota
	AlignmentVTop
	AlignmentVCenter
	AlignmentVBottom
)

type Alignment struct {
	Horizontal AlignmentH
	Vertical   AlignmentV
}
