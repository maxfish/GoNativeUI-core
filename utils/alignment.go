package utils

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
