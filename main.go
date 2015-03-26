package main

import (
	"fmt"

	"flag"
	"github.com/amirblum/SynergyAI/model"
	"math"
	"math/rand"
	"time"
)

// Command line flags
var configFile string

//var worldFile string
//var tasksFile string
//var taskAmount int

var realWorld *model.World

func init() {
	flag.StringVar(&configFile, "config", "", "Config JSON file")
	//
	//	flag.StringVar(&worldFile, "world", "", "World JSON file")
	//	flag.StringVar(&tasksFile, "tasks", "", "Tasks JSON file")
	//	flag.IntVar(&taskAmount, "taskAmount", 50, "Amount of tasks to run")
	//
	//	var seed int64
	//	flag.Int64Var(&seed, "seed", time.Now().Unix(), "The random seed")

	flag.Parse()
}

func main() {
	// Load config
	config := LoadConfig(configFile)
	if config.RandomSeed < 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(config.RandomSeed)
	}

	// Load real world from file
	realWorld = model.LoadWorld(config.World)
	fmt.Println(realWorld)
	// Init learned world
	world := model.CreateWorld(realWorld.Workers, config.Budget)

	synergyAgent := config.CreateAgent(realWorld)

	taskGenerator, hasNext := model.DummyTaskGenerator()
	if config.Tasks != "" {
		taskGenerator, hasNext = model.FileTaskGenerator(config.Tasks, config.TasksAmount)
	}

	for hasNext {
		// Log current worldview

		// Get the next task
		var task model.Task
		task, hasNext = taskGenerator()

		synergyAgent.GetTeam(world, task)

		//		Temporary print
		//        score, fulfill, _ := world.ScoreTeam(team, task)
		//
		//        realScore, realFulfill, _ := realWorld.ScoreTeam(team, task)
		//        fmt.Printf("For task:\n%v\nThe team:\n%v\nAppraised the score: %v, fulfillPercent: %v\nAnd the Real score: %v, fullfillPercent: %v\n", task, team, score, fulfill, realScore, realFulfill)
	}

	fmt.Printf("\nReal-World Synergy: \n%v\n", realWorld.Synergy)
	fmt.Printf("\nFinal Synergy: \n%v\n", world.Synergy)

	learnScore := 0.0
	for _, worker := range world.Workers {
		for _, otherWorker := range world.Workers {
			if worker.ID > otherWorker.ID {
				learnScore += math.Pow(realWorld.Synergy[worker.ID][otherWorker.ID]-world.Synergy[worker.ID][otherWorker.ID], 2)
			}
		}
	}

	fmt.Println("The learning receives a score of:", learnScore)
}
