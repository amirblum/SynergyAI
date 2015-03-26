package agents

import (
	"fmt"
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"
)

// This is our stupid agent. We use him for testing the success of the search algorithms.
type English struct {
	searchAlg search.SearchAlgorithm
	realWorld *model.World
}

func CreateEnglish(searchAlg search.SearchAlgorithm, realWorld *model.World) *English {
	fmt.Println("The word 'mistake' is not one that appears in my dictionary.")
	return &English{searchAlg, realWorld}
}

func (johnny *English) GetTeam(world *model.World, task model.Task) *model.Team {
	// Find the optimal team using the real world. No learning for this agent
	team := johnny.searchAlg.SearchTeam(johnny.realWorld, task)

	model.LogScore = true
	score, _, _ := johnny.realWorld.ScoreTeam(team, task)
	fmt.Println("For task:", task, "\nFound the team:\n", team, "With a score of:", score)

	return team
}
