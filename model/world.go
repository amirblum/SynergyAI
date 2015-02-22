package model

type Ability string

type SynergyRow []float32
type SynergyMatrix []SynergyRow
type World struct {
	Workers []Worker
	Synergy SynergyMatrix
}

// Create a new world
func CreateWorld(workers []Worker) *World {
	var w *World = new(World)

	// Initialize workers
	w.Workers = workers

	// Initialize Synergy matrix
	w.Synergy = make(SynergyMatrix, len(workers))
	for i := range workers {
		w.Synergy[i] = make(SynergyRow, len(workers))
		for j := range workers {
			w.Synergy[i][j] = 1
		}
	}

	return w
}

func (w *World) ScoreTeam(team []Worker, task Task) float32 {
	return 2.
}
