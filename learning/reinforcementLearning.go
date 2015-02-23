package learning

import "github.com/amirblum/SynergyAI/model"

type TemporalDifferenceAlgorithm struct {
	Eta float64
}

func (alg TemporalDifferenceAlgorithm) LearnSynergy(world, realWorld *model.World, team []model.Worker, task model.Task) {
	// Create a "boring world", where no-one affects anyone elses work. This gives us a normalizing factor.
	boringWorld := model.CreateWorld(world.Workers)
	normalizingFactor, _ := boringWorld.ScoreTeam(team, task)

	myScore, _ := world.ScoreTeam(team, task)
	realScore, _ := realWorld.ScoreTeam(team, task)

	// We normalize the scores to reduce influence from the scale of the ability.
	// This way, the resulting difference is on a similar scale to the synergies, and can be used to
	// learn the matrix.
	myScore /= normalizingFactor
	realScore /= normalizingFactor

	difference := (realScore - myScore)

	// In addition, we further reduce the difference because bigger teams give a bigger score.
	difference /= float64(len(team))

	// Update the matrix
	for _, worker := range team {
		for _, otherWorker := range team {
			if worker.ID != otherWorker.ID {
				world.Synergy[worker.ID][otherWorker.ID] += difference * alg.Eta
			}
		}
	}
}
