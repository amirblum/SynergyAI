package agents

import (
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"
)

type SynergyAgent interface {
	GetTeam(world *model.World, task model.Task) *model.Team
}

type AgentChooser func(search.SearchAlgorithm, *model.World, model.Task) *model.Team

func RandomTeam(searchAlg search.SearchAlgorithm, world *model.World, task model.Task) *model.Team {
	return model.CreateRandomTeam(world.Workers)
}

func AugmentedOptimalTeam(searchAlg search.SearchAlgorithm, world *model.World, task model.Task) *model.Team {
	optimalTeam := searchAlg.SearchTeam(world, task)
	// Travel 30% of the size of the team
	distance := (optimalTeam.Length()*3)/10 + 1

	team := model.CreateTeamNode(optimalTeam.Workers, world.Workers)
	randomIterator, hasNext := search.RandomSuccessorIterator(optimalTeam, world.Workers)
	for i := 0; i < distance && hasNext; i++ {
		if hasNext {
			team, hasNext = randomIterator()
			randomIterator, hasNext = search.RandomSuccessorIterator(team, world.Workers)
		}
	}

	return team
}
