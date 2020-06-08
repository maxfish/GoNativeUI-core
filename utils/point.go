package utils

import "fmt"

type Point [2]int

func (p Point) Point(x, y int) Point {
	return Point{x, y}
}

func (p Point) X() int {
	return p[0]
}

func (p Point) Y() int {
	return p[1]
}

func (p Point) Add(other Point) Point {
	return Point{p[0] + other[0], p[1] + other[1]}
}

func (p Point) EqualsTo(other Point) bool {
	return p[0] == other[0] && p[1] == other[1]
}

func (p Point) ToString() string {
	return fmt.Sprintf("{x:%d,y:%d}", p.X, p.Y)
}
