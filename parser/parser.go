package parser

import "github.com/amirblum/SynergyAI/model"

func LoadWorld(filename string) *model.World {
	// Temp workers
	worker := model.Worker{1, 1, make(map[model.Ability]float32, 2)}
	worker.Components["basic"] = 1.0
	workers := make([]model.Worker, 2)
	workers = append(workers, worker)

	return model.CreateWorld(workers)
}
