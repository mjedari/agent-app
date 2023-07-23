package pkg

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const NumberOfAgents = 50

type AgentPool struct {
	cap  int
	pool []*Agent
}

func NewAgentPool(cap int) *AgentPool {
	return &AgentPool{cap: cap}
}

func (ap *AgentPool) SelectAgent(target *Target) *Agent {
	freeAgents := ap.getFreeAgents()

	stepsToTargetMap := make(map[int]Agent)
	smallestAgentIndex := 0

	for i, agent := range freeAgents {
		steps := agent.findStepsToTarget(target)
		if i == 0 {
			smallestAgentIndex = steps
		}

		stepsToTargetMap[steps] = *agent
	}

	// choose the smallest one
	for index, _ := range stepsToTargetMap {
		if smallestAgentIndex >= index {
			smallestAgentIndex = index
		}
	}

	agent := stepsToTargetMap[smallestAgentIndex]

	fmt.Printf("%v with location %v were selected. \n\n", agent.name, agent.location.ToString())

	return &agent

}

func (ap *AgentPool) Initiate() {
	for i := 0; i < NumberOfAgents; i++ {
		newAgent := NewAgent(fmt.Sprintf("Agent #%v", i+1), createRandomLocation())
		if err := ap.addAgent(newAgent); err != nil {
			log.Println(err)
			break
		}
	}
	fmt.Printf("%v random located agents were created.", NumberOfAgents)
}

func (a *Agent) numberOfStepsToTarget() int {
	// calculate
	return 0
}

func (a *Agent) findStepsToTarget(target *Target) int {
	a.SetTarget(target)
	defer func() {
		a.unSetTarget()
	}()

	return a.findSteps()
}

func (ap *AgentPool) addAgent(agent *Agent) error {
	if ap.cap < len(ap.pool) {
		return errors.New("pool is full")
	}
	ap.pool = append(ap.pool, agent)
	return nil
}

func (ap *AgentPool) getFreeAgents() []*Agent {
	var agents []*Agent

	for _, agent := range ap.pool {
		if agent.free {
			agents = append(agents, agent)
		}
	}

	return agents
}

func createRandomLocation() Location {
	var location Location
	rand.Seed(time.Now().UnixNano())

	location = Location{
		X: rand.Intn(50) + 20,
		Y: rand.Intn(50) + 20,
	}
	return location
}
