package main

import (
	"fmt"

	"flag"
	//	"github.com/amirblum/SynergyAI/agents"
	//	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	//	"github.com/amirblum/SynergyAI/search"
	"math/rand"
	"time"
	//    "github.com/amirblum/SynergyAI"
)

// Command line flags
var configFile string
var worldFile string
var tasksFile string
var taskAmount int

var realWorld *model.World

func init() {
	flag.StringVar(&configFile, "config", "", "Config JSON file")

	flag.StringVar(&worldFile, "world", "", "World JSON file")
	flag.StringVar(&tasksFile, "tasks", "", "Tasks JSON file")
	flag.IntVar(&taskAmount, "taskAmount", 50, "Amount of tasks to run")

	var seed int64
	flag.Int64Var(&seed, "seed", time.Now().Unix(), "The random seed")
	rand.Seed(seed)

	flag.Parse()
}

func main() {
	// Load config
	config := LoadConfig(configFile)

	// Load real world from file
	realWorld = model.LoadWorld(config.World)
	fmt.Println(realWorld)
	// Init learned world
	world := model.CreateWorld(realWorld.Workers)

	//	// Init Search algorithm
	//	searchAlgorithm := search.HillClimbingAlgorithm{}
	//
	//	// Init Learning algorithm
	//	learningAgent := learning.CreateLearningAgent(learning.CreateTemporalDifferenceAlgorithm(learning.CreateAverageDelta(0.1, 30)))
	//
	//	// Init the agent
	//	//    synergyAgent := agents.CreateBond(searchAlgorithm, learningAgent, realWorld)
	//	synergyAgent := agents.CreatePowers(searchAlgorithm, learningAgent, realWorld, agents.RandomTeam)

	synergyAgent := config.CreateAgent(realWorld)

	taskGenerator, hasNext := model.DummyTaskGenerator()
	if tasksFile != "" {
		taskGenerator, hasNext = model.FileTaskGenerator(tasksFile, taskAmount)
	}

	for hasNext {
		// Log current worldview

		// Get the next task
		var task model.Task
		task, hasNext = taskGenerator()

		//		// Find the optimal team
		//		team := searchAlgorithm.SearchTeam(world, task)
		//
		//		// Learn from experience
		//		learningAgent.LearnSynergy(world, realWorld, team, task)

		//        team := synergyAgent.GetTeam(world, task)
		synergyAgent.GetTeam(world, task)

		// Temporary print
		//		score, fulfill := world.ScoreTeam(team, task)
		//
		//		realScore, realFulfill := realWorld.ScoreTeam(team, task)
		//		fmt.Printf("For task:\n%v\nThe team:\n%v\nAppraised the score: %v, fulfillPercent: %v\nAnd the Real score: %v, fullfillPercent: %v\n", task, team, score, fulfill, realScore, realFulfill)
	}

	fmt.Printf("\nFinal Synergy: \n%v\n", world.Synergy)
}
