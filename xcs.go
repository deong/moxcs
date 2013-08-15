package main

import (
//	"github.com/deong/multixcs/conf"
)

var (
	// maximum size of the population (in micro-classifiers)
	maxNumClassifiers int

	// learning rate for updating p, epsilon, f, and as
	beta float64

	// used in calculating the fitness of a classifier
	alpha    float64
	epsilon0 float64
	v        float64

	// discount factor for multi-step learning problems
	gamma float64

	// GA threshold; when the average time since the last GA run in the action set
	// exceeds this value, the GA is run
	thetaGA int

	// crossover rate
	crossoverRate float64

	// mutation rate
	mutationRate float64
	
	// thetaDel is the deletion threshold. If the experience of a classifier is greater 
	// than thetaDel, its fitness may be considered in its probability of deletion
	thetaDel float64

	// delta specifies the fraction of the mean fitness in the population below which
	// the fitness of a classifier may be considered in its deletion probability
	delta float64

	// thetaSubsumption is the subsumption threshold. The experience of a classifier 
	// must exceed thetaSubsumption in order to be able to subsume another classifier
	thetaSubsumption float64
	
	// starProbability is the probability of using a * in one attribute in C when covering
	starProbability float64

	// used as initial values in new classifiers
	pI, epsilonI, fI float64

	// explorationProbability specifies the probability of choosing a random action during 
	// action selection
	explorationProbability float64

	// thetaMNA specifies the minimum number of unique actions that must be present in 
	// an action set to prevent covering from occuring
	thetaMNA float64

	// doGASubsumption specifies whether to check offspring for subsumption by their parents
	doGASubsumption bool

	// doActionSetSubsumption specifies whether action sets are to be tested for subsuming
	// classifiers
	doActionSetSubsumption bool
)

func GenerateMatchSet(pop Population, cond string, thetaMNA int) (ms Population) {
	ms = Population{}
	actions := make(map[int]int)
	for _, chr := range pop {
		if chr.Matches(cond) {
			ms = append(ms, chr)
			actions[chr.Action]++
		}
	}

	// if len(actions) < thetaMNA {
	// 	cl := generateCoveringClassifier(ms, cond)
	// 	pop = append(pop, cl)
	// 	pop.Deletion()
	// }
	return
}
