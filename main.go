package main

import (
	"fmt"

	"flag"
	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"
	"math/rand"
	"time"
)

// Command line flags
var worldFile string
var tasksFile string
var taskAmount int

var realWorld *model.World

func init() {
	flag.StringVar(&worldFile, "world", "", "World JSON file")
	flag.StringVar(&tasksFile, "tasks", "", "Tasks JSON file")
	flag.IntVar(&taskAmount, "taskAmount", 50, "Amount of tasks to run")

	var seed int64
	flag.Int64Var(&seed, "seed", time.Now().Unix(), "The random seed")
	rand.Seed(seed)

	flag.Parse()
}

func main() {
	// Load real world from file
	realWorld = model.LoadWorld(worldFile)
	fmt.Println(realWorld)
	// Init learned world
	world := model.CreateWorld(realWorld.Workers)

	// Init Search agent
	searchAgent := search.CreateSearchAgent(search.HillClimbingAlgorithm{})

	// Init Learning agent
	learningAgent := learning.CreateLearningAgent(learning.TemporalDifferenceAlgorithm{0.1})

	taskGenerator, hasNext := model.DummyTaskGenerator()
	if tasksFile != "" {
		taskGenerator, hasNext = model.FileTaskGenerator(tasksFile, taskAmount)
	}
	for hasNext {
		// Log current worldview

		// Get the next task
		var task model.Task
		task, hasNext = taskGenerator()

		// Find the optimal team
		team := searchAgent.SearchTeam(world, task)
		// Temporary print
		score, fulfill := world.ScoreTeam(team, task)

		realScore, realFulfill := realWorld.ScoreTeam(team, task)
		fmt.Printf("For task:\n%v\nThe team:\n%v\nAppraised the score: %v, fulfillPercent: %v\nAnd the Real score: %v, fullfillPercent: %v\n", task, team, score, fulfill, realScore, realFulfill)

		// Learn from experience
		learningAgent.LearnSynergy(world, realWorld, team, task)
	}

	fmt.Printf("\nFinal Synergy: \n%v\n", world.Synergy)
}
