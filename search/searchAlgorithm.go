package search

import (
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/goutils"
	"math/rand"
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

// An iterator that returns successors sequentially
func (node *teamNode) IncrementalSuccessorIterator(allWorkers []model.Worker) (func() (*teamNode, bool), bool) {
	return node.successorIterator(allWorkers, incrementalIterator(len(allWorkers)+len(node.Workers)))
}

func incrementalIterator(incrementRange int) func(int) (int, bool) {
	return func(num int) (int, bool) {
		num++
		return num, num < incrementRange
	}
}

// An iterator that returns a random successor
func (node *teamNode) RandomSuccessorIterator(allWorkers []model.Worker) (func() (*teamNode, bool), bool) {
	return node.successorIterator(allWorkers, randomIterator(len(allWorkers)+len(node.Workers)))
}

func randomIterator(incrementRange int) func(int) (int, bool) {
	return func(num int) (int, bool) {
		num = rand.Intn(incrementRange)
		return num, true
	}
}

// Iterator producing the successors of a given teamNode.
// First add all the workers (that are not in the team) one by one
// Then remove the workers in the team on by one
func (node *teamNode) successorIterator(allWorkers []model.Worker, indexIterator func(int) (int, bool)) (func() (*teamNode, bool), bool) {
	// Init iterator
	nextWorker, hasNext := indexIterator(-1)
	return func() (*teamNode, bool) {
		currentWorker := nextWorker
		nextWorker, hasNext = indexIterator(nextWorker)

		// Find next worker that is not in the team
		for nextWorker <= len(allWorkers) && node.WorkerMap[currentWorker] == true {
			currentWorker = nextWorker
			nextWorker, hasNext = indexIterator(nextWorker)
		}

		// Copy the team
		newTeam := node.copy()

		// If there are more workers to add, add
		if currentWorker < len(allWorkers) {
			newTeam.Workers = append(newTeam.Workers, allWorkers[currentWorker])
			newTeam.WorkerMap[currentWorker] = true

			return newTeam, hasNext
		}

		// Finished returning teams with added workers, start removing workers
		idToRemove := currentWorker - len(allWorkers)

		workerToRemove := newTeam.Workers[idToRemove]
		newTeam.Workers = append(newTeam.Workers[:idToRemove], newTeam.Workers[idToRemove+1:]...)
		newTeam.WorkerMap[workerToRemove.ID] = false

		return newTeam, hasNext

	}, hasNext
}

//// Iterator producing the successors of a given teamNode.
//// First add all the workers (that are not in the team) one by one
//// Then remove the workers in the team on by one
//func (node *teamNode) successorIterator(allWorkers []model.Worker) (func() (*teamNode, bool), bool) {
//	// Init iterator
//	nextWorker := 0
//	return func() (*teamNode, bool) {
//		currentWorker := nextWorker
//		nextWorker++
//
//		// Find next worker that is not in the team
//		for nextWorker <= len(allWorkers) && node.WorkerMap[currentWorker] == true {
//			currentWorker = nextWorker
//			nextWorker++
//		}
//
//		// Copy the team
//		newTeam := node.copy()
//
//		// If there are more workers to add, add
//		if currentWorker < len(allWorkers) {
//			newTeam.Workers = append(newTeam.Workers, allWorkers[currentWorker])
//			newTeam.WorkerMap[currentWorker] = true
//
//			return newTeam, (nextWorker < len(allWorkers)) || (len(node.Workers) > 0)
//		}
//
//		// Finished returning teams with added workers, start removing workers
//		idToRemove := currentWorker - len(allWorkers)
//
//		workerToRemove := newTeam.Workers[idToRemove]
//		newTeam.Workers = append(newTeam.Workers[:idToRemove], newTeam.Workers[idToRemove+1:]...)
//		newTeam.WorkerMap[workerToRemove.ID] = false
//
//		return newTeam, (idToRemove < len(node.Workers)-1)
//
//	}, (len(allWorkers) > 0)
//}
