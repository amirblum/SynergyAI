package agents

import (
	"fmt"
	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"
)

// This is our #1 agent
type Bond struct {
	searchAlg   search.SearchAlgorithm
	learningAlg learning.LearningAlgorithm
	realWorld   *model.World
}

func CreateBond(searchAlg search.SearchAlgorithm, learningAlg learning.LearningAlgorithm, realWorld *model.World) *Bond {
	fmt.Println("The name is Bond, James Bond")
	return &Bond{searchAlg, learningAlg, realWorld}
}

func (james *Bond) GetTeam(world *model.World, task model.Task) *model.Team {
	// Find the optimal team
	team := james.searchAlg.SearchTeam(world, task)

	// Learn from experience
	james.learningAlg.LearnSynergy(world, james.realWorld, team, task)

	return team
}
