package main

import (
	"github.com/mjedari/agent-app/src/cmd"
)

/*
(n) agents
(1) target per agent

---
* every second go forward to agent's target (x, y)
* they can go just every second: (0, 0)-> (1, 1)

when:
---
* is the nearest agents to the target
* free to move
* if we get two agents we choose the smallest index one
*/

func main() {
	cmd.Execute()
}
