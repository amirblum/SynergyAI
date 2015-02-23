package search

import (
	"github.com/alexcesaro/log/stdlog"
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct{}

func (alg HillClimbingAlgorithm) SearchTeam(world *model.World, task model.Task) []model.Worker {
	logger := stdlog.GetFromFlags()

	current := &teamNode{make([]model.Worker, 0), make(map[int]bool)}

	for {
		maxNeighbor := alg.getMaxNeighbor(current, world, task)
		logger.Debugf("maxNeighbor: %v", maxNeighbor)

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
	neighborsIterator, hasNext := current.successorIterator(world.Workers)
	if hasNext {
		// Initiailize maxNeighbor to be the first successor (cuz no do-while)
		maxNeighbor, hasNext = neighborsIterator()

		// Find the max
		for hasNext {
			var currentNeighbor *teamNode
			currentNeighbor, hasNext = neighborsIterator()

			if world.CompareTeams(maxNeighbor.Workers, currentNeighbor.Workers, task) > 0 {
				maxNeighbor = currentNeighbor
			}
		}
	}

	return maxNeighbor
}
