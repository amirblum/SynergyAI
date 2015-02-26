package agents

import (
	"fmt"
	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"
	"math/rand"
)

// This agent fools around sometimes...
type Powers struct {
	searchAlg    search.SearchAlgorithm
	learningAlg  learning.LearningAlgorithm
	realWorld    *model.World
	agentChooser AgentChooser
}

func CreatePowers(searchAlg search.SearchAlgorithm, learningAlg learning.LearningAlgorithm, realWorld *model.World, chooser AgentChooser) *Powers {
	fmt.Println("Shagadelic baby, yeah!")
	return &Powers{searchAlg, learningAlg, realWorld, chooser}
}

func (austin *Powers) GetTeam(world *model.World, task model.Task) *model.Team {
	// With probability 20%, choose a team other than the optimal
	teamProp := rand.Intn(5)

	var team *model.Team

	if teamProp == 0 {
		team = austin.agentChooser(austin.searchAlg, world, task)
	} else {
		// Find the optimal team
		team = austin.searchAlg.SearchTeam(world, task)
	}

	// Learn from experience
	austin.learningAlg.LearnSynergy(world, austin.realWorld, team, task)

	return team
}
