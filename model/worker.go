package model

//import "fmt"

type Worker struct {
	ID         int
	Salary     int
	Components map[Ability]float64
}
