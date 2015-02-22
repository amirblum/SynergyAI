package search

import "github.com/amirblum/SynergyAI/model"

type SearchAgent struct {
	alg SearchAlgorithm
}

func CreateSearchAgent(alg SearchAlgorithm) *SearchAgent {
	return &SearchAgent{alg}
}

func (agent SearchAgent) SearchTeam(world *model.World, task model.Task) []model.Worker {
	return agent.alg.SearchTeam(world, task)
}
