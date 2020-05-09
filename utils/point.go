package utils

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) Point(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) EqualsTo(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p Point) ToString() string {
	return fmt.Sprintf("{x:%d,y:%d}", p.X, p.Y)
}
