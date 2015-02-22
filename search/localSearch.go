package search

import (
	"fmt"
	"github.com/amirblum/SynergyAI/model"
)

type HillClimbingAlgorithm struct{}

func (HillClimbingAlgorithm) SearchTeam(world *model.World, task model.Task) []model.Worker {
	var (
		current  *teamNode
		neighbor *teamNode
	)

	current = &teamNode{make([]model.Worker, 0), make(map[int]bool)}
	currentScore := world.ScoreTeam(current.Workers, task)

	for {
		// Find highest neighbor
		neighborsIterator, hasNext := current.successorIterator(world.Workers)
		fmt.Printf("%v\n", hasNext)
		maxNeighborScore := -1.

		// Loop on all neighbors
		for hasNext {
			fmt.Println("Hello")
			var currentNeighbor *teamNode
			currentNeighbor, hasNext = neighborsIterator()
			fmt.Printf("%v\n", hasNext)

			if neighborScore := world.ScoreTeam(currentNeighbor.Workers, task); neighborScore > maxNeighborScore {
				maxNeighborScore = neighborScore
				neighbor = currentNeighbor
			}
		}

		// Check break condition
		if maxNeighborScore <= currentScore {
			return current.Workers
		}

		// Continue iteration
		current = neighbor
	}
}
