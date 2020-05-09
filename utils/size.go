package utils

import "fmt"

type Size struct {
	W int
	H int
}

func (s Size) ToString() string {
	return fmt.Sprintf("{w:%d,h:%d}", s.W, s.H)
}
