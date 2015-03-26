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

var LogScore bool = false

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

		if LogScore {
			fmt.Println("output:", output, "demand:", demand, "componentPercent:", componentPercent, "score:", score)
		}
	}

	return score, fulfillPercent, budgetSum <= task.Budget
}

// Calculate the output of the team for each component.
func (w *World) teamOutput(team *Team, task Task) map[Ability]float64 {
	output := make(map[Ability]float64, len(task.Components))

	// For each component, take the workers ability in that component and multiply it
	// with the dynamics between that worker and all other workers in the team
	for ability, _ := range task.Components {
		for _, worker := range team.Workers {
			workerOutput := worker.Components[ability]
			for _, otherWorker := range team.Workers {
				x, y := worker.ID, otherWorker.ID
				// Don't calculate for dynamic with itself...
				if x != y {
					// Our matrix is triangular (easier to config)
					if x < y {
						x, y = y, x
					}
					workerOutput *= w.Synergy[x][y]
					if LogScore {
						fmt.Println("worker", worker.ID, "score:", workerOutput, "*", w.Synergy[x][y], "w.Synergy[x]:", w.Synergy[x])
					}
				}
			}

			// Add the output
			output[ability] += workerOutput
		}
	}

	return output
}
