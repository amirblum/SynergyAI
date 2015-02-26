package search

import (
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct{}

func CreateHillClimbingAlgorithm() *HillClimbingAlgorithm {
	return &HillClimbingAlgorithm{}
}

func (alg HillClimbingAlgorithm) SearchTeam(world *model.World, task model.Task) *model.Team {

	current := new(model.Team)

	for {
		maxNeighbor := alg.getMaxNeighbor(current, world, task)

		// Check break condition
		if maxNeighbor != nil && world.CompareTeams(maxNeighbor, current, task) <= 0 {
			return current
		}

		// Continue iteration
		current = maxNeighbor
	}
}

// Find highest neighbor
func (HillClimbingAlgorithm) getMaxNeighbor(current *model.Team, world *model.World, task model.Task) *model.Team {
	var maxNeighbor *model.Team = nil

	// Get neighbors iterator
	if neighborsIterator, hasNext := IncrementalSuccessorIterator(current, world.Workers); hasNext {
		// Initiailize maxNeighbor to be the first successor (cuz no do-while)
		maxNeighbor, hasNext = neighborsIterator()

		// Find the max
		for hasNext {
			var currentNeighbor *model.Team
			currentNeighbor, hasNext = neighborsIterator()

			if world.CompareTeams(currentNeighbor, maxNeighbor, task) > 0 {
				maxNeighbor = currentNeighbor
			}
		}
	}

	return maxNeighbor
}
