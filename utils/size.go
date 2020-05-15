package utils

import "fmt"

type Size [2]int

func (s Size)W() int {
	return s[0]
}

func (s Size)H() int {
	return s[1]
}

func (s Size)SetW(w int) {
	s[0] = w
}

func (s Size)SetH(h int) {
	s[1] = h
}

func (s Size) ToString() string {
	return fmt.Sprintf("{w:%d,h:%d}", s[0], s[1])
}
