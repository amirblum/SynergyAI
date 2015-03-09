package search

import (
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct {
	neighborPicker NeighborPicker
}

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
