package parser

import (
	"github.com/amirblum/SynergyAI/model"
)

func LoadWorld(filename string) *model.World {
	// Temp workers
	worker := model.Worker{0, 1, make(map[model.Ability]float64, 2)}
	worker.Components["Poop"] = 10.0

	workers := make([]model.Worker, 0)
	workers = append(workers, worker)

	return model.CreateWorld(workers)
}
