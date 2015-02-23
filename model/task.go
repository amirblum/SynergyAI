package model

//import "fmt"

type Task struct {
	Budget     int
	Components map[Ability]float64
}

func DummyTaskGenerator() (func() (Task, bool), bool) {
	tasks := make([]Task, 2)
	for i, _ := range tasks {
		tasks[i] = Task{10, make(map[Ability]float64, 2)}
		tasks[i].Components["Poop"] = float64(5 * (i + 1))
	}

	currTask := 0
	return func() (Task, bool) {
		prevTask := currTask
		currTask++
		return tasks[prevTask], (currTask < len(tasks))
	}, currTask < len(tasks)
}

//func FileTaskGenerator(taskFile string) (func() (Task, bool), bool) {
//    // Open file
////    var tasks []Task
////    var currTask int
//    return func() (Task, bool) {
//        // Return next task
//        return Task{10, make(map[Ability]float32, 2)}
//    }
//}

//func SmartTaskGenerator(world *World) (func() (Task, bool), bool) {
//
//}
