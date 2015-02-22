package main

import (
	"fmt"
	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	synergyParser "github.com/amirblum/SynergyAI/parser"
	"github.com/amirblum/SynergyAI/search"
)

var realWorld *model.World

func main() {
	// Load real world from file
	realWorld = synergyParser.LoadWorld("temp.world")

	// Init learned world
	world := model.CreateWorld(realWorld.Workers)

	taskGenerator, hasNext := model.DummyTaskGenerator()
	for hasNext {
		// Log current worldview

		// Get the next task
		var task model.Task
		task, hasNext = taskGenerator()

		// Find the optimal team
		team := search.SearchTeam(world, task)

		// Score the team
		score := realWorld.ScoreTeam(team, task)

		// Learn from experience
		learning.LearnSynergy(world, team, task, score)
	}

	fmt.Printf("%v\n", world.Synergy)
}
