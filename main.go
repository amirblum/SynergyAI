package main

import (
	"fmt"

	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	synergyParser "github.com/amirblum/SynergyAI/parser"
	"github.com/amirblum/SynergyAI/search"

	"github.com/alexcesaro/log/stdlog"
)

var realWorld *model.World

func main() {
	logger := stdlog.GetFromFlags()

	// Load real world from file
	realWorld = synergyParser.LoadWorld("temp.world")

	// Init learned world
	world := model.CreateWorld(realWorld.Workers)

	// Init Search agent
	searchAgent := search.CreateSearchAgent(search.HillClimbingAlgorithm{})

	taskGenerator, hasNext := model.DummyTaskGenerator()
	for hasNext {
		// Log current worldview

		// Get the next task
		var task model.Task
		task, hasNext = taskGenerator()

		// Find the optimal team
		team := searchAgent.SearchTeam(world, task)

		// Temporary print
		score, fulfill := world.ScoreTeam(team, task)
		logger.Debugf("The team %v received the score %v for task %v, with fulfillPercent %v", team, score, task, fulfill)

		// Learn from experience
		learning.LearnSynergy(world, realWorld, team, task)
	}

	fmt.Printf("%v\n", world.Synergy)
}
