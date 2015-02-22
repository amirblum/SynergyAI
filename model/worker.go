package model

import "fmt"

type Worker struct {
	ID         int
	Salary     int
	Components map[Ability]float32
}

func (w Worker) String() string {
	return fmt.Sprintf("ID: %d, Salary: %d", w.ID, w.Salary)
}
