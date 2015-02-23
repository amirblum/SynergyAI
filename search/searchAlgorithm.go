package search

import (
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/goutils"
)

type SearchAlgorithm interface {
	SearchTeam(*model.World, model.Task) []model.Worker
}

type teamNode struct {
	Workers   []model.Worker
	WorkerMap map[int]bool
}

// Copy a teamNode
func (node *teamNode) copy() *teamNode {
	newTeam := new(teamNode)

	// Copy array
	newTeam.Workers = make([]model.Worker, len(node.Workers))
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
	currentWorker := 0
	return func() (*teamNode, bool) {
		prevWorker := currentWorker
		currentWorker++

		// Find next worker that is not in the team
		for currentWorker < len(allWorkers) && node.WorkerMap[prevWorker] == true {
			prevWorker = currentWorker
			currentWorker++
			break
		}

		// Copy the team
		newTeam := node.copy()

		// If there are more workers to add, add
		if prevWorker < len(allWorkers) {
			newTeam.Workers = append(newTeam.Workers, allWorkers[prevWorker])
			newTeam.WorkerMap[prevWorker] = true

			return newTeam, (currentWorker < len(allWorkers)) || (len(node.Workers) > 0)
		}

		// Finished returning teams with added workers, start removing workers
		idToRemove := prevWorker - len(allWorkers)

		workerToRemove := newTeam.Workers[idToRemove]
		newTeam.Workers = append(newTeam.Workers[:idToRemove], newTeam.Workers[idToRemove+1:]...)
		newTeam.WorkerMap[workerToRemove.ID] = false

		return newTeam, (idToRemove < len(node.Workers)-1)

	}, (len(allWorkers) > 0)
}
