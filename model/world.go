package model

import (
	"fmt"
	"math"
)

type Ability string

type SynergyRow []float64
type SynergyMatrix []SynergyRow

func (matrix SynergyMatrix) String() string {
	s := ""
	for i, row := range matrix {
		s += fmt.Sprintf("%d: %.2f\n", i, row)
	}
	return s
}

type World struct {
	Workers []Worker
	Synergy SynergyMatrix
	budget  bool
}

// Create a new world
func CreateWorld(workers []Worker, budget bool) *World {
	var w *World = new(World)
	// Init budget
	w.budget = budget
	// Initialize workers
	w.Workers = workers

	// Initialize Synergy matrix
	w.Synergy = make(SynergyMatrix, len(workers))
	for i := range workers {
		w.Synergy[i] = make(SynergyRow, len(workers))
		for j := range workers {
			w.Synergy[i][j] = 1
		}
	}

	return w
}

func (w *World) CompareTeams(team1, team2 *Team, task Task) int {
	score1, fulfillPercent1, inBudget1 := w.ScoreTeam(team1, task)
	score2, fulfillPercent2, inBudget2 := w.ScoreTeam(team2, task)
	// Consider budget
	if w.budget {
		if inBudget1 != inBudget2 {
			if inBudget1 {
				return 1
			} else {
				return -1
			}
		}
	}
	if fulfillPercent1 == fulfillPercent2 {
		if score1 < score2 {
			return -1
		} else if score1 == score2 {
			return 0
		} else {
			return 1
		}
	} else {
		if fulfillPercent1 < fulfillPercent2 {
			return -1
		} else {
			return 1
		}
	}
}

func (w *World) ScoreTeam(team *Team, task Task) (score float64, fulfillPercent float64, inBudget bool) {
	// Calculate budget
	budgetSum := 0.
	for _, worker := range team.Workers {
		budgetSum += worker.Salary
	}

	componentSum := 0.
	for _, component := range task.Components {
		componentSum += component
	}

	if componentSum == 0 {
		return 1., 1., budgetSum <= task.Budget
	}

	outputMap := w.teamOutput(team, task)

	for outputComponent, output := range outputMap {
		componentPercent := task.Components[outputComponent] / componentSum

		outputRelation := 0.
		// Protect from divide-by-zero
		demand := task.Components[outputComponent]
		if demand > 0 {
			outputRelation = output / task.Components[outputComponent]
		}

		score += outputRelation * componentPercent
		fulfillPercent += math.Min(1., outputRelation) * componentPercent
	}

	return score, fulfillPercent, budgetSum <= task.Budget
}

// Calculate the output of the team for each component.
func (w *World) teamOutput(team *Team, task Task) map[Ability]float64 {
	output := make(map[Ability]float64, len(task.Components))

	// For each component
	for ability, _ := range task.Components {
		// For each worker
		for _, worker := range team.Workers {
			// Take the workers ability...
			workerOutput := worker.Components[ability]
			for _, otherWorker := range team.Workers {
				x, y := worker.ID, otherWorker.ID
				if x != y {
					if x < y {
						x, y = y, x
					}
					// ...and multiply it with his dynamic with other worker in team
					workerOutput *= w.Synergy[x][y]
				}
			}

			// Add the output
			output[ability] += workerOutput
		}
	}

	return output
}
