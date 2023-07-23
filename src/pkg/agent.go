package pkg

import (
	"fmt"
	"time"
)

type Agent struct {
	name     string
	location Location
	target   Target
	free     bool
}

func NewAgent(name string, location Location) *Agent {
	return &Agent{name: name, location: location, free: true}
}

func (a *Agent) SetTarget(target *Target) {
	a.target = *target
}

func (a *Agent) Move() {
	a.makeUnFree()
	defer func() {
		a.makeFree()
	}()

	steps := 0
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		a.location = a.nextLocation()

		fmt.Printf("agent %v went to the point (%v, %v) \n", a.name, a.location.X, a.location.Y)

		steps++
		if a.arriveToTarget() {
			fmt.Printf("agent %v arrived to the target in %v step(s). \n", a.name, steps)
			break
		}
	}
}

func (a *Agent) nextLocation() Location {
	return Location{
		X: a.location.X + a.nextStep(a.location.X, a.target.X),
		Y: a.location.Y + a.nextStep(a.location.Y, a.target.Y),
	}
}

func (a *Agent) unSetTarget() {
	a.target = Target{}
}

func (a *Agent) findSteps() int {
	realLocation := a.location
	defer func() {
		a.location = realLocation
	}()
	steps := 0
	for {
		a.location = Location{
			X: a.location.X + a.nextStep(a.location.X, a.target.X),
			Y: a.location.Y + a.nextStep(a.location.Y, a.target.Y),
		}
		steps++

		if a.arriveToTarget() {
			break
		}
	}

	return steps
}

func (a *Agent) arriveToTarget() bool {
	return a.location.X == a.target.X && a.location.Y == a.target.Y
}

func (a *Agent) nextStep(origin, target int) int {
	if origin < target {
		return 1
	}

	if origin > target {
		return -1
	}

	return 0
}

func (a *Agent) makeFree() {
	a.free = true
}

func (a *Agent) makeUnFree() {
	a.free = false
}
