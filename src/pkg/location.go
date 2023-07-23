package pkg

import "fmt"

type Location struct {
	X int
	Y int
}

func (l Location) ToString() string {
	return fmt.Sprintf("(%v, %v)", l.X, l.Y)
}
func NewLocation(x int, y int) *Location {
	return &Location{X: x, Y: y}
}

type Target Location

func NewTarget(x int, y int) *Target {
	return &Target{X: x, Y: y}
}
