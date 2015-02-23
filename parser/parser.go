package parser

import (
	"fmt"
	"github.com/amirblum/SynergyAI/model"
)

func LoadWorld(filename string) *model.World {
	// Temp workers
	workers := make([]model.Worker, 2)

	for i, _ := range workers {
		worker := model.Worker{i, 1, make(map[model.Ability]float64, 1)}
		worker.Components["Poop"] = float64((2 - i) * 5)

		workers[i] = worker
	}

	realWorld := model.CreateWorld(workers)
	realWorld.Synergy[0][1], realWorld.Synergy[1][0] = 0.1, 0.1

	fmt.Printf("Real synergy: %v\n", realWorld.Synergy)

	return realWorld
}
