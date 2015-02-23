package model

import (
	"math"
)

type Ability string

type SynergyRow []float64
type SynergyMatrix []SynergyRow
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

func (w *World) CompareTeams(team1, team2 []Worker, task Task) int {
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

func (w *World) ScoreTeam(team []Worker, task Task) (score float64, fulfillPercent float64) {
	componentSum := 0.
	for _, component := range task.Components {
		componentSum += component
	}

	outputMap := w.teamOutput(team, task)

	for outputComponent, output := range outputMap {
		componentPercent := task.Components[outputComponent] / componentSum

		outputRelation := output / task.Components[outputComponent]

		score += outputRelation * componentPercent
		fulfillPercent += math.Min(1., outputRelation) * componentPercent
	}

	return score, fulfillPercent
}

func (w *World) teamOutput(team []Worker, task Task) map[Ability]float64 {
	output := make(map[Ability]float64, len(task.Components))

	for ability, _ := range task.Components {
		for _, worker := range team {
			workerOutput := worker.Components[ability]
			for _, otherWorker := range team {
				workerOutput *= w.Synergy[worker.ID][otherWorker.ID]
			}
			output[ability] += workerOutput
		}
	}

	return output
}
