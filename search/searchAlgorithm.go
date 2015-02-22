package search

import (
	"fmt"
	"github.com/amirblum/SynergyAI/model"
)

type SearchAlgorithm interface {
	SearchTeam(*model.World, model.Task) []model.Worker
}

type teamNode struct {
	Workers   []model.Worker
	WorkerMap map[int]bool
}

func (node *teamNode) copy() *teamNode {
	newTeam := new(teamNode)
	newTeam.Workers = node.Workers[:]
	newTeam.WorkerMap = make(map[int]bool)
	for k, v := range node.WorkerMap {
		newTeam.WorkerMap[k] = v
	}
	return newTeam
}

func (node *teamNode) successorIterator(allWorkers []model.Worker) (func() (*teamNode, bool), bool) {
	currentWorker := 0
	return func() (*teamNode, bool) {
		prevWorker := currentWorker
		currentWorker++

		for currentWorker < len(allWorkers) && node.WorkerMap[prevWorker] == false {
			prevWorker = currentWorker
			currentWorker++
		}

		if prevWorker == len(allWorkers) {
			// TODO: Iterate removing workers
			return nil, false
		} else {
			newTeam := node.copy()
			newTeam.Workers = append(newTeam.Workers, allWorkers[prevWorker])
			newTeam.WorkerMap[prevWorker] = true
			// TODO: hasNext true if len(node.Workers) > 0
			fmt.Printf("currWorker: %v, len(allWorkers): %v\n", currentWorker, len(allWorkers))
			return newTeam, (currentWorker < len(allWorkers))
		}
	}, (len(allWorkers) > 0)
}
