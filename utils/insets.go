package utils

import "fmt"

type Insets struct {
	Top, Right, Bottom, Left int
}

func HomogeneousInsets(inset int) Insets {
	return Insets{Top: inset, Right: inset, Bottom: inset, Left: inset}
}

func (i Insets) ToString() string {
	return fmt.Sprintf("{top:%d,right:%d,bottom:%d,left:%d}", i.Top, i.Right, i.Bottom, i.Left)
}
