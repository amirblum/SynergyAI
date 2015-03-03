package search

import (
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct {
	neighborPicker NeighborPicker
}

type NeighborPicker func(*model.Team, *model.World, model.Task) *model.Team

func CreateHillClimbingAlgorithm(neighborPicker NeighborPicker) *HillClimbingAlgorithm {
	return &HillClimbingAlgorithm{neighborPicker}
}

func (alg HillClimbingAlgorithm) SearchTeam(world *model.World, task model.Task) *model.Team {

	current := new(model.Team)

	for {
		nextNeighbor := alg.neighborPicker(current, world, task)

		// Check break condition
		if nextNeighbor != nil && world.CompareTeams(nextNeighbor, current, task) <= 0 {
			return current
		}

		// Continue iteration
		current = nextNeighbor
	}
}

// Find highest neighbor
func MaxNeighbor(current *model.Team, world *model.World, task model.Task) *model.Team {
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

// Find first neighbor
func FirstChoiceNeighbor(current *model.Team, world *model.World, task model.Task) *model.Team {
	// Get neighbors iterator
	if neighborsIterator, hasNext := RandomSuccessorIterator(current, world.Workers); hasNext {
		// Find the first bigger neighbor (if exists)
		for hasNext {
			var nextNeighbor *model.Team
			nextNeighbor, hasNext = neighborsIterator()

			if world.CompareTeams(nextNeighbor, current, task) > 0 {
				return nextNeighbor
			}
		}
	}

	return model.CreateTeamNode(make([]model.Worker, 0), world.Workers)
}
