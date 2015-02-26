package learning

import "github.com/amirblum/SynergyAI/model"

type LearningAlgorithm interface {
	LearnSynergy(world, realWorld *model.World, team *model.Team, task model.Task)
}
