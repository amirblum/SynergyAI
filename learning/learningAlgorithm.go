package learning

import "github.com/amirblum/SynergyAI/model"

type LearningAlgorithm interface {
	LearnSynergy(world, realWorld *model.World, team []model.Worker, task model.Task)
}
