package parser

import (
	"github.com/amirblum/SynergyAI/model"
)

func LoadWorld(filename string) *model.World {
	// Temp workers
	workers := make([]model.Worker, 2)

	for i, _ := range workers {
		worker := model.Worker{i, 1, make(map[model.Ability]float64, 1)}
		worker.Components["Poop"] = float64((i + 1) * 5)

		workers[i] = worker
	}

	return model.CreateWorld(workers)
}
