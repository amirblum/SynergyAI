package learning

import (
	"github.com/amirblum/SynergyAI/model"
)

type DeltaCalcer interface {
	CalcDelta(float64) (float64, bool)
}

// Just returns the difference
type SimpleDelta struct {
	eta float64
}

func CreateSimpleDelta(eta float64) *SimpleDelta {
	return &SimpleDelta{eta}
}

func (delta *SimpleDelta) CalcDelta(diff float64) (float64, bool) {
	return delta.eta * diff, true
}

// Average the difference and return the average
type AverageDelta struct {
	eta             float64
	lastDifferences []float64
	frequency       int
	count           int
}

func CreateAverageDelta(eta float64, frequency int) *AverageDelta {
	return &AverageDelta{eta, make([]float64, frequency), frequency, 0}
}

func (delta *AverageDelta) CalcDelta(diff float64) (float64, bool) {
	if delta.count == delta.frequency {
		sum := 0.
		for i := 0; i < delta.frequency; i++ {
			sum += delta.lastDifferences[i]
		}
		delta.count = 0
		delta.lastDifferences = make([]float64, delta.frequency)
		return delta.eta * (sum / float64(delta.frequency)), true
	}

	delta.lastDifferences[delta.count] = diff
	delta.count++
	return 0, false
}

type TemporalDifferenceAlgorithm struct {
	deltaCalcer DeltaCalcer
}

func CreateTemporalDifferenceAlgorithm(calcer DeltaCalcer) *TemporalDifferenceAlgorithm {
	return &TemporalDifferenceAlgorithm{calcer}
}

func (alg TemporalDifferenceAlgorithm) LearnSynergy(world, realWorld *model.World, team *model.Team, task model.Task) {
	// Nothing to learn from teams smaller than 2
	if team.Length() < 2 {
		return
	}

	// Create a "boring world", where no-one affects anyone elses work. This gives us a normalizing factor.
	boringWorld := model.CreateWorld(world.Workers, false)
	normalizingFactor, _, _ := boringWorld.ScoreTeam(team, task)
	// Cover our asses in case of bad team
	if normalizingFactor == 0 {
		return
	}

	myScore, _, _ := world.ScoreTeam(team, task)
	realScore, _, _ := realWorld.ScoreTeam(team, task)

	// We normalize the scores to reduce influence from the scale of the ability.
	// This way, the resulting difference is on a similar scale to the synergies, and can be used to
	// learn the matrix.
	myScore /= normalizingFactor
	realScore /= normalizingFactor

	difference := (realScore - myScore)

	// In addition, we further reduce the difference because bigger teams give a bigger score.
	difference /= float64(team.Length())
	if delta, toChange := alg.deltaCalcer.CalcDelta(difference); toChange {
		// Update the matrix
		for _, worker := range team.Workers {
			for _, otherWorker := range team.Workers {
				if worker.ID > otherWorker.ID {
					world.Synergy[worker.ID][otherWorker.ID] += delta
				}
			}
		}
	}
}
