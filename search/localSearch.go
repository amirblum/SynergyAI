package search

import (
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct{}

func (alg HillClimbingAlgorithm) SearchTeam(world *model.World, task model.Task) model.Team {

	current := &teamNode{make(model.Team, 0), make(map[int]bool)}

	for {
		maxNeighbor := alg.getMaxNeighbor(current, world, task)

		// Check break condition
		if maxNeighbor != nil && world.CompareTeams(maxNeighbor.Workers, current.Workers, task) <= 0 {
			return current.Workers
		}

		// Continue iteration
		current = maxNeighbor
	}
}

// Find highest neighbor
func (HillClimbingAlgorithm) getMaxNeighbor(current *teamNode, world *model.World, task model.Task) *teamNode {
	var maxNeighbor *teamNode = nil

	// Get neighbors iterator
	if neighborsIterator, hasNext := current.successorIterator(world.Workers); hasNext {
		// Initiailize maxNeighbor to be the first successor (cuz no do-while)
		maxNeighbor, hasNext = neighborsIterator()

		// Find the max
		for hasNext {
			var currentNeighbor *teamNode
			currentNeighbor, hasNext = neighborsIterator()

			if world.CompareTeams(currentNeighbor.Workers, maxNeighbor.Workers, task) > 0 {
				maxNeighbor = currentNeighbor
			}
		}
	}

	return maxNeighbor
}
