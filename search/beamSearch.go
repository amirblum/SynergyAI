package search

import (
	"fmt"
	"github.com/amirblum/SynergyAI/model"
	"sync"
)

type LocalBeamAlgorithm struct {
	numBeams       int
	neighborPicker NeighborPicker
}

func CreateLocalBeamAlgorithm(numBeams int, neighborPicker NeighborPicker) *LocalBeamAlgorithm {
	return &LocalBeamAlgorithm{numBeams, neighborPicker}
}

func (alg LocalBeamAlgorithm) SearchTeam(world *model.World, task model.Task) *model.Team {
	// Cannot perform search with more beams than total workers
	if alg.numBeams > len(world.Workers) {
		fmt.Println("Too many beams!")
		return nil
	}

	// Every beam gets numBeams maximums
	globalMaxNeighbors := make([]*model.Team, alg.numBeams*alg.numBeams)

	// Init currentNodes to empty
	currentNodes := make([]*model.Team, alg.numBeams)
	for i := 0; i < len(currentNodes); i++ {
		currentNodes[i] = model.CreateTeamNode(make([]model.Worker, 0), world.Workers)
	}

	// Start searching
	for isNotConverged := true; isNotConverged; {
		var wg sync.WaitGroup
		for i := 0; i < alg.numBeams; i++ {
			numThread := i

			wg.Add(1)
			go func() {
				defer wg.Done()
				//                fmt.Println("Starting beam:", numThread)
				current := currentNodes[numThread]

				maxNeighbors := make([]*model.Team, alg.numBeams)

				// Get neighbors iterator
				if neighborsIterator, hasNext := IncrementalSuccessorIterator(current, world.Workers); hasNext {
					var currentNeighbor *model.Team

					// Initiailize maxNeighbors to be current + empty nodes
					maxNeighbors[0] = current
					for j := 1; j < len(maxNeighbors); j++ {
						maxNeighbors[j] = model.CreateTeamNode(make([]model.Worker, 0), world.Workers)
					}

					// Find the maxs
					for hasNext {
						currentNeighbor, hasNext = neighborsIterator()

						for j := 0; j < len(maxNeighbors); j++ {
							if world.CompareTeams(currentNeighbor, maxNeighbors[j], task) > 0 {
								maxNeighbors = append(maxNeighbors[:j], append([]*model.Team{currentNeighbor}, maxNeighbors[j:len(maxNeighbors)-1]...)...)
								break
							}
						}
					}

					// Transfer our maxs to the global maxs
					for j := 0; j < len(maxNeighbors); j++ {
						globalMaxNeighbors[numThread*alg.numBeams+j] = maxNeighbors[j]
					}
				}
			}()
		}

		// Wait for threads
		wg.Wait()

		// Initialize new currents
		newCurrentNodes := make([]*model.Team, alg.numBeams)
		// Set currentNodes to empty (so we can sort)
		for i := 0; i < len(currentNodes); i++ {
			newCurrentNodes[i] = model.CreateTeamNode(make([]model.Worker, 0), world.Workers)
		}

		// Extract the maximums
		for _, team := range globalMaxNeighbors {
			for j := 0; j < len(newCurrentNodes); j++ {
				if world.CompareTeams(team, newCurrentNodes[j], task) > 0 {
					newCurrentNodes = append(newCurrentNodes[:j], append([]*model.Team{team}, newCurrentNodes[j:len(currentNodes)-1]...)...)
					break
				}
			}
		}

		// Check convergence
		isNotConverged = false
		for i := 0; i < len(currentNodes); i++ {
			isNotConverged = isNotConverged || (currentNodes[i] != newCurrentNodes[i])
		}

		// Set new current nodes
		currentNodes = newCurrentNodes
	}

	return currentNodes[0]
}
