 package model

import (
	"fmt"
	"github.com/amirblum/goutils"
	"math/rand"
)

type Worker struct {
	ID         int
	Salary     float64
	Components map[Ability]float64
}

func (worker Worker) String() string {
	s := fmt.Sprintf("ID: %v | Salary: %v | ", worker.ID, worker.Salary)
	for component, amount := range worker.Components {
		s += fmt.Sprintf("%v: %v, ", component, amount)
	}
	return s
}

type Team struct {
	Workers   []Worker
	WorkerMap map[int]bool
}

func CreateTeamNode(workers []Worker, allWorkers []Worker) *Team {
	team := &Team{make([]Worker, len(workers)), make(map[int]bool, len(allWorkers))}

	copy(team.Workers, workers)

	for _, worker := range workers {
		team.WorkerMap[worker.ID] = true
	}

	return team
}

func CreateRandomTeam(allWorkers []Worker) *Team {
	team := &Team{WorkerMap: make(map[int]bool)}

	for _, worker := range allWorkers {
		if rand.Intn(1) == 0 {
			team.Workers = append(team.Workers, worker)
			team.WorkerMap[worker.ID] = true
		}
	}

	return team
}

func (team Team) Length() int {
	return len(team.Workers)
}

// Copy a team
func (node *Team) Copy() *Team {
	// Create new node of same size
	newTeam := Team{make([]Worker, len(node.Workers)), make(map[int]bool, len(node.WorkerMap))}

	// Copy array
	copy(newTeam.Workers, node.Workers)

	// Copy map
	goutils.CopyMap(node.WorkerMap, newTeam.WorkerMap)

	return &newTeam
}

func (team Team) String() string {
	if team.Length() == 0 {
		return "Empty"
	}

	var s string

	for i, worker := range team.Workers {
		s += fmt.Sprintf("%v", worker)
		if i < team.Length()-1 {
			s += "\n"
		}
	}

	return s
}
