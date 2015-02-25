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
	for _, row := range matrix {
		s += fmt.Sprintf("%v\n", row)
	}
	return s
}

type World struct {
	Workers []Worker
	Synergy SynergyMatrix
}

// Create a new world
func CreateWorld(workers []Worker) *World {
	var w *World = new(World)

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

func (w *World) CompareTeams(team1, team2 Team, task Task) int {
	score1, fulfillPercent1 := w.ScoreTeam(team1, task)
	score2, fulfillPercent2 := w.ScoreTeam(team2, task)

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

func (w *World) ScoreTeam(team Team, task Task) (score float64, fulfillPercent float64) {
	componentSum := 0.
	for _, component := range task.Components {
		componentSum += component
	}

	if componentSum == 0 {
		return 1., 1.
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

	return score, fulfillPercent
}

// Calculate the output of the team for each component.
func (w *World) teamOutput(team Team, task Task) map[Ability]float64 {
	output := make(map[Ability]float64, len(task.Components))

	// For each component
	for ability, _ := range task.Components {
		// For each worker
		for _, worker := range team {
			// Take the workers ability...
			workerOutput := worker.Components[ability]
			for _, otherWorker := range team {
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
