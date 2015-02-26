package search

import (
	"github.com/amirblum/SynergyAI/model"
	"math/rand"
)

type SearchAlgorithm interface {
	SearchTeam(*model.World, model.Task) *model.Team
}

//type teamNode struct {
//	Workers   model.Team
//	WorkerMap map[int]bool
//}

// An iterator that returns successors sequentially
func IncrementalSuccessorIterator(team *model.Team, allWorkers []model.Worker) (func() (*model.Team, bool), bool) {
	return successorIterator(team, allWorkers, incrementalIterator(len(allWorkers)+team.Length()))
}

func incrementalIterator(incrementRange int) func(int) (int, bool) {
	return func(num int) (int, bool) {
		num++
		return num, num < incrementRange
	}
}

// An iterator that returns a random successor
func RandomSuccessorIterator(team *model.Team, allWorkers []model.Worker) (func() (*model.Team, bool), bool) {
	return successorIterator(team, allWorkers, randomIterator(len(allWorkers)+team.Length()))
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
func successorIterator(node *model.Team, allWorkers []model.Worker, indexIterator func(int) (int, bool)) (func() (*model.Team, bool), bool) {
	// Init iterator
	nextWorker, hasNext := indexIterator(-1)
	return func() (*model.Team, bool) {
		currentWorker := nextWorker
		nextWorker, hasNext = indexIterator(nextWorker)

		// Find next worker that is not in the team
		for nextWorker <= len(allWorkers) && node.WorkerMap[currentWorker] == true {
			currentWorker = nextWorker
			nextWorker, hasNext = indexIterator(nextWorker)
		}

		// Copy the team
		newTeam := node.Copy()

		// If there are more workers to add, add
		if currentWorker < len(allWorkers) {
			newTeam.Workers = append(newTeam.Workers, allWorkers[currentWorker])
			newTeam.WorkerMap[allWorkers[currentWorker].ID] = true

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
