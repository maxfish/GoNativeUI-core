package utils

type AlignmentH uint32

const (
	AlignmentHCenter AlignmentH = iota
	AlignmentHLeft
	AlignmentHRight
)

type AlignmentV uint32

const (
	AlignmentVCenter AlignmentV = iota
	AlignmentVTop
	AlignmentVBottom
)

type Alignment struct {
	Horizontal AlignmentH
	Vertical   AlignmentV
}

type FitMode int

const (
	FitModeAlign FitMode = iota
	FitModeFill
	FitModeAspectFit
	FitModeAspectFill
)
