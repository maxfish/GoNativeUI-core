package gui

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
