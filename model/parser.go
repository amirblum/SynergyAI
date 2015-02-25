package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadWorld(filename string) *World {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("Error loading world file: %v\n", e)
		os.Exit(1)
	}

	var realWorld = new(World)
	json.Unmarshal(file, realWorld)

	//    fmt.Printf("realWorld: %v\n", realWorld)

	return realWorld
}

func LoadTasks(filename string) []taskTemplate {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("Error loading world file: %v\n", e)
		os.Exit(1)
	}

	var tasks = make([]taskTemplate, 0)
	json.Unmarshal(file, &tasks)

	//    fmt.Printf("tasks: %v\n", tasks)

	return tasks
}

//func ParseJSON(filename string, object interface{}) *interface{} {
//
//}
