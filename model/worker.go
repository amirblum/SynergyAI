package model

import "fmt"

type Worker struct {
	ID         int
	Salary     int
	Components map[Ability]float64
}

func (worker Worker) String() string {
	s := fmt.Sprintf("ID: %v | Salary: %v | ", worker.ID, worker.Salary)
	for component, amount := range worker.Components {
		s += fmt.Sprintf("%v: %v, ", component, amount)
	}
	return s
}

type Team []Worker

func (team Team) String() string {
	var s string
	for i, worker := range team {
		s += fmt.Sprintf("%v", worker)
		if i < len(team)-1 {
			s += "\n"
		}
	}
	return s
}
