package learning

import "github.com/amirblum/SynergyAI/model"

type LearningAgent struct {
	alg LearningAlgorithm
}

func CreateLearningAgent(alg LearningAlgorithm) *LearningAgent {
	return &LearningAgent{alg}
}

func (agent LearningAgent) LearnSynergy(world, realWorld *model.World, team *model.Team, task model.Task) {
	agent.alg.LearnSynergy(world, realWorld, team, task)
}
