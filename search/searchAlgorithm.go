package search

import (
	"github.com/amirblum/SynergyAI/model"
	"math/rand"
)

type SearchAlgorithm interface {
	SearchTeam(*model.World, model.Task) *model.Team
}

// An iterator that returns successors sequentially
func IncrementalSuccessorIterator(team *model.Team, allWorkers []model.Worker) (func() (*model.Team, bool), bool) {
	return successorIterator(team, allWorkers, incrementalIterator(len(allWorkers)+team.Length()))
}

func incrementalIterator(incrementRange int) func(int) (int, bool) {
	return func(num int) (int, bool) {
		num++
		return num, num < incrementRange-1
	}
}

// An iterator that returns a random successor
func RandomSuccessorIterator(team *model.Team, allWorkers []model.Worker) (func() (*model.Team, bool), bool) {
	return successorIterator(team, allWorkers, randomIterator(len(allWorkers)+team.Length()))
}

func randomIterator(incrementRange int) func(int) (int, bool) {
	permutation := rand.Perm(incrementRange)
	current := -1
	return func(num int) (int, bool) {
		current++
		return permutation[current], current < incrementRange-1
	}
}

// Iterator producing the successors of a given teamNode.
// First add all the workers (that are not in the team) one by one
// Then remove the workers in the team on by one
func successorIterator(node *model.Team, allWorkers []model.Worker, indexIterator func(int) (int, bool)) (func() (*model.Team, bool), bool) {
	// Init iterator
	nextWorker, nextHasNext := findNextIndex(indexIterator, -1, allWorkers, node)

	return func() (*model.Team, bool) {
		currentWorker := nextWorker
		hasNext := nextHasNext

		if nextHasNext {
			nextWorker, nextHasNext = findNextIndex(indexIterator, nextWorker, allWorkers, node)
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

	}, nextHasNext
}

func findNextIndex(indexIterator func(int) (int, bool), currentIndex int, allWorkers []model.Worker, node *model.Team) (int, bool) {
	nextIndex, hasNext := indexIterator(currentIndex)

	// Find next worker that is not in the team
	for nextIndex <= len(allWorkers) && node.WorkerMap[nextIndex] == true && hasNext {
		nextIndex, hasNext = indexIterator(nextIndex)
	}

	return nextIndex, hasNext
}
