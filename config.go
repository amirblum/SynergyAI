package main

import (
	"github.com/amirblum/SynergyAI/agents"
	"github.com/amirblum/SynergyAI/learning"
	"github.com/amirblum/SynergyAI/model"
	"github.com/amirblum/SynergyAI/search"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type SearchConfig struct {
	SearchAlgorithm string
	NeighborPicker  string
}

type LearningConfig struct {
	LearningAlgorithm string
	DeltaCalcer       string
	Eta               float64
	Frequency         int
}

type AgentConfig struct {
	SynergyAgent string
	AgentChooser string
}

type Config struct {
	World          string
	Budget         bool
	Tasks          string
	TasksAmount    int
	RandomSeed     int64
	SearchConfig   SearchConfig
	LearningConfig LearningConfig
	AgentConfig    AgentConfig
}

func LoadConfig(filename string) *Config {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("Error loading config file: %v\n", e)
		os.Exit(1)
	}

	var config = new(Config)
	json.Unmarshal(file, config)

	fmt.Printf("config: %v\n", config)

	return config
}

func (config *Config) CreateSearchAlgorithm() search.SearchAlgorithm {
	var neighborPicker search.NeighborPicker
	switch config.SearchConfig.NeighborPicker {
	case "Max":
		neighborPicker = search.MaxNeighbor
		break
	case "FirstChoice":
		neighborPicker = search.FirstChoiceNeighbor
	}

	var alg search.SearchAlgorithm
	switch config.SearchConfig.SearchAlgorithm {
	case "HillClimbing":
		alg = search.CreateHillClimbingAlgorithm(neighborPicker)
		break
	}

	return alg
}

func (config *Config) CreateLearningAlgorithm() learning.LearningAlgorithm {
	var calcer learning.DeltaCalcer
	switch config.LearningConfig.DeltaCalcer {
	case "Simple":
		calcer = learning.CreateSimpleDelta(config.LearningConfig.Eta)
		break
	case "Average":
		calcer = learning.CreateAverageDelta(config.LearningConfig.Eta, config.LearningConfig.Frequency)
	}

	var alg learning.LearningAlgorithm
	switch config.LearningConfig.LearningAlgorithm {
	case "TemporalDifference":
		alg = learning.CreateTemporalDifferenceAlgorithm(calcer)
		break
	}

	return alg
}

func (config *Config) CreateAgent(realWorld *model.World) agents.SynergyAgent {
	var chooser agents.AgentChooser
	switch config.AgentConfig.AgentChooser {
	case "Random":
		chooser = agents.RandomTeam
		break
	case "Augmented":
		chooser = agents.AugmentedOptimalTeam
		break
	}

	var agent agents.SynergyAgent
	switch config.AgentConfig.SynergyAgent {
	case "Bond":
		agent = agents.CreateBond(config.CreateSearchAlgorithm(), config.CreateLearningAlgorithm(), realWorld)
		break
	case "Powers":
		agent = agents.CreatePowers(config.CreateSearchAlgorithm(), config.CreateLearningAlgorithm(), realWorld, chooser)
		break
	}

	return agent
}
