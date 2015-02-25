package search

import (
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/goutils"
	//    "fmt"
)

type SearchAlgorithm interface {
	SearchTeam(*model.World, model.Task) model.Team
}

type teamNode struct {
	Workers   model.Team
	WorkerMap map[int]bool
}

// Copy a teamNode
func (node *teamNode) copy() *teamNode {
	newTeam := new(teamNode)

	// Copy array
	newTeam.Workers = make(model.Team, len(node.Workers))
	copy(newTeam.Workers, node.Workers)

	// Copy map
	newTeam.WorkerMap = make(map[int]bool)
	goutils.CopyMap(node.WorkerMap, newTeam.WorkerMap)

	return newTeam
}

// Iterator producing the successors of a given teamNode.
// First add all the workers (that are not in the team) one by one
// Then remove the workers in the team on by one
func (node *teamNode) successorIterator(allWorkers []model.Worker) (func() (*teamNode, bool), bool) {
	// Init iterator
	nextWorker := 0
	return func() (*teamNode, bool) {
		currentWorker := nextWorker
		nextWorker++

		// Find next worker that is not in the team
		for nextWorker <= len(allWorkers) && node.WorkerMap[currentWorker] == true {
			currentWorker = nextWorker
			nextWorker++
		}

		// Copy the team
		newTeam := node.copy()

		// If there are more workers to add, add
		if currentWorker < len(allWorkers) {
			newTeam.Workers = append(newTeam.Workers, allWorkers[currentWorker])
			newTeam.WorkerMap[currentWorker] = true

			return newTeam, (nextWorker < len(allWorkers)) || (len(node.Workers) > 0)
		}

		// Finished returning teams with added workers, start removing workers
		idToRemove := currentWorker - len(allWorkers)

		workerToRemove := newTeam.Workers[idToRemove]
		newTeam.Workers = append(newTeam.Workers[:idToRemove], newTeam.Workers[idToRemove+1:]...)
		newTeam.WorkerMap[workerToRemove.ID] = false

		return newTeam, (idToRemove < len(node.Workers)-1)

	}, (len(allWorkers) > 0)
}
