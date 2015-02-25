package model

import (
	"fmt"
	"math"
	"math/rand"
)

type Task struct {
	Name       string
	Budget     int
	Components map[Ability]float64
}

func (task Task) String() string {
	s := fmt.Sprintf("Name: %v\nBudget: %v\n", task.Name, task.Budget)
	counter := 0
	for component, amount := range task.Components {
		counter++
		s += fmt.Sprintf("%v: %v", component, amount)

		if counter < len(task.Components)-1 {
			s += "\n"
		}
	}
	return s
}

type normalDistribution struct {
	Expect   float64
	Variance float64
}

type taskTemplate struct {
	Name       string
	Weight     int
	Budget     normalDistribution
	Components map[Ability]normalDistribution
}

func (template taskTemplate) generateTask() Task {

	// Set budget
	budget := int(math.Max(0, rand.NormFloat64()*math.Sqrt(template.Budget.Variance)+template.Budget.Expect))

	// Set components
	components := make(map[Ability]float64)
	for component, distribution := range template.Components {
		components[component] = math.Max(0, rand.NormFloat64()*math.Sqrt(distribution.Variance)+distribution.Expect)
	}

	return Task{template.Name, budget, components}
}

func DummyTaskGenerator() (func() (Task, bool), bool) {
	tasks := make([]Task, 1)
	for i, _ := range tasks {
		tasks[i] = Task{"Dummy", 10, make(map[Ability]float64, 2)}
		tasks[i].Components["basic"] = float64(5 * (i + 1))
	}

	return sequentialTaskGenerator(tasks)
}

func FileTaskGenerator(taskFile string, count int) (func() (Task, bool), bool) {
	// Read file
	tasks := LoadTasks(taskFile)

	return taskGenerator(tasks, count)
}

func sequentialTaskGenerator(tasks []Task) (func() (Task, bool), bool) {
	currTask := 0
	return func() (Task, bool) {
		prevTask := currTask
		currTask++
		return tasks[prevTask], (currTask < len(tasks))
	}, currTask < len(tasks)
}

func taskGenerator(templates []taskTemplate, count int) (func() (Task, bool), bool) {
	// Calculate total weights
	totalWeights := 0
	for _, template := range templates {
		totalWeights += template.Weight
	}

	// Populate the template picker with template pointers
	templatePicker := make([]*taskTemplate, totalWeights)
	for _, template := range templates {
		// Add template.Weight pointers
		for i := 0; i < template.Weight; i++ {
			templatePicker = append(templatePicker, &template)
		}
	}

	currentTask := 0
	return func() (Task, bool) {
		currentTask++

		// Choose a random template (with consideration for the weights)
		templateIndex := rand.Intn(totalWeights)
		template := templates[templateIndex]

		return template.generateTask(), (currentTask < count)
	}, count > 0
}

//func SmartTaskGenerator(world *World) (func() (Task, bool), bool) {
//
//}
