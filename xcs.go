package main

import (
	"fmt"
	"github.com/deong/multixcs/conf"
	"math/rand"
	"os"
)

// parameters for the XCS learning algorithm
type XCS struct {
	// N is maximum size of the population (in micro-classifiers)
	N int

	// learning rate for updating p, epsilon, f, and as
	Beta float64

	// used in calculating the fitness of a classifier
	Alpha, Epsilon0, V float64

	// discount factor for multi-step learning problems
	Gamma float64

	// GAThreshold is the value at which if the average time since the
	// last GA run in the action set exceeds this value, the GA is run
	GAThreshold int

	// crossover rate
	CrossoverRate float64

	// mutation rate
	MutationRate float64

	// DeletionThreshold sets the minimum value for the experience of a classifier,
	// above which its fitness may be considered in its probability of deletion
	DeletionThreshold float64

	// Delta specifies the fraction of the mean fitness in the population below which
	// the fitness of a classifier may be considered in its deletion probability
	Delta float64

	// SubsumptionThreshold is the minimum experience of a classifier
	// to be eligible to subsume another classifier
	SubsumptionThreshold int

	// ProbDC is the probability of generating don't care symbols in covering
	ProbDC float64

	// used as initial values in new classifiers
	PredictionI, EpsilonI, FitnessI float64

	// ProbExploration specifies the probability of choosing a random action during
	// action selection
	ProbExploration float64

	// ThetaMNA specifies the minimum number of unique actions that must be present in
	// an action set to prevent covering from occuring
	ThetaMNA int

	// DoGASubsumption specifies whether to check offspring for subsumption by their parents
	DoGASubsumption bool

	// DoActionSetSubsumption specifies whether action sets are to be tested for subsuming
	// classifiers
	DoActionSetSubsumption bool
}

// Initialize reads the XCS parameters from a configuration file
func (xcs *XCS) Initialize() {
	if val, err := conf.IntParameter("xcs", "N"); err != nil {
		fmt.Println("required parameter 'N' not specified")
		os.Exit(1)
	} else {
		xcs.N = val
	}

	if val, err := conf.Float64Parameter("xcs", "Beta"); err != nil {
		fmt.Println("required parameter 'Beta' not specified")
		os.Exit(1)
	} else {
		xcs.Beta = val
	}

	if val, err := conf.Float64Parameter("xcs", "Alpha"); err != nil {
		fmt.Println("required parameter 'Alpha' not specified")
		os.Exit(1)
	} else {
		xcs.Alpha = val
	}

	if val, err := conf.Float64Parameter("xcs", "Epsilon0"); err != nil {
		fmt.Println("required parameter 'Epsilon0' not specified")
		os.Exit(1)
	} else {
		xcs.Epsilon0 = val
	}

	if val, err := conf.Float64Parameter("xcs", "V"); err != nil {
		fmt.Println("required parameter 'V' not specified")
		os.Exit(1)
	} else {
		xcs.V = val
	}

	if val, err := conf.Float64Parameter("xcs", "Gamma"); err != nil {
		fmt.Println("required parameter 'Gamma' not specified")
		os.Exit(1)
	} else {
		xcs.Gamma = val
	}

	if val, err := conf.IntParameter("xcs", "GAThreshold"); err != nil {
		fmt.Println("required parameter 'GAThreshold' not specified")
		os.Exit(1)
	} else {
		xcs.GAThreshold = val
	}

	if val, err := conf.Float64Parameter("xcs", "CrossoverRate"); err != nil {
		fmt.Println("required parameter 'CrossoverRate' not specified")
		os.Exit(1)
	} else {
		xcs.CrossoverRate = val
	}

	if val, err := conf.Float64Parameter("xcs", "MutationRate"); err != nil {
		fmt.Println("required parameter 'MutationRate' not specified")
		os.Exit(1)
	} else {
		xcs.MutationRate = val
	}

	if val, err := conf.Float64Parameter("xcs", "DeletionThreshold"); err != nil {
		fmt.Println("required parameter 'DeletionThreshold' not specified")
		os.Exit(1)
	} else {
		xcs.DeletionThreshold = val
	}

	if val, err := conf.Float64Parameter("xcs", "Delta"); err != nil {
		fmt.Println("required parameter 'Delta' not specified")
		os.Exit(1)
	} else {
		xcs.Delta = val
	}

	if val, err := conf.IntParameter("xcs", "SubsumptionThreshold"); err != nil {
		fmt.Println("required parameter 'SubsumptionThreshold' not specified")
		os.Exit(1)
	} else {
		xcs.SubsumptionThreshold = val
	}

	if val, err := conf.Float64Parameter("xcs", "ProbDC"); err != nil {
		fmt.Println("required parameter 'ProbDC' not specified")
		os.Exit(1)
	} else {
		xcs.ProbDC = val
	}

	if val, err := conf.Float64Parameter("xcs", "PredictionI"); err != nil {
		fmt.Println("required parameter 'PredictionI' not specified")
		os.Exit(1)
	} else {
		xcs.PredictionI = val
	}

	if val, err := conf.Float64Parameter("xcs", "EpsilonI"); err != nil {
		fmt.Println("required parameter 'EpsilonI' not specified")
		os.Exit(1)
	} else {
		xcs.EpsilonI = val
	}

	if val, err := conf.Float64Parameter("xcs", "FitnessI"); err != nil {
		fmt.Println("required parameter 'FitnessI' not specified")
		os.Exit(1)
	} else {
		xcs.FitnessI = val
	}

	if val, err := conf.Float64Parameter("xcs", "ProbExploration"); err != nil {
		fmt.Println("required parameter 'ProbExploration' not specified")
		os.Exit(1)
	} else {
		xcs.ProbExploration = val
	}

	if val, err := conf.IntParameter("xcs", "ThetaMNA"); err != nil {
		fmt.Println("required parameter 'ThetaMNA' not specified")
		os.Exit(1)
	} else {
		xcs.ThetaMNA = val
	}

	if val, err := conf.BoolParameter("xcs", "DoGASubsumption"); err != nil {
		fmt.Println("required parameter 'DoGASubsumption' not specified")
		os.Exit(1)
	} else {
		xcs.DoGASubsumption = val
	}

	if val, err := conf.BoolParameter("xcs", "DoActionSetSubsumption"); err != nil {
		fmt.Println("required parameter 'DoActionSetSubsumption' not specified")
		os.Exit(1)
	} else {
		xcs.DoActionSetSubsumption = val
	}
}

// GenerateMatchSet returns a Population consisting only of the members of the current
// population that match the given condition
func (xcs *XCS) GenerateMatchSet(pop Population, cond string) (ms Population) {
	ms = Population{}
	actions := make(map[int]int)
	for _, chr := range pop {
		if chr.Matches(cond) {
			ms = append(ms, chr)
			actions[chr.Action]++
		}
	}

	// TODO: finish GenerateMatchSet
	//
	// if len(actions) < thetaMNA {
	// 	cl := generateCoveringClassifier(ms, cond)
	// 	pop = append(pop, cl)
	// 	pop.Deletion()
	// }
	return
}

// GenerateCoveringClassifier returns a new classifier that
func (xcs *XCS) GenerateCoveringClassifier(ms Population, cond string) (cl *Chromosome) {
	// TODO: pass a correct value for numTasks in somehow
	// TODO: write test for GenerateCoveringClassifier
	numTasks := 2
	cl = NewChromosome(len(cond), numTasks)
	for i := range cl.Condition {
		if rand.Float64() < xcs.ProbDC {
			cl.Condition[i] = '*'
		} else {
			cl.Condition[i] = cond[i]
		}
	}

	// TODO: finish GeneratecoveringClassifier

	// generate an action at random from the set of actions not present in ms

	// TODO: make the pI parameter into a []float64

	cl.Error = xcs.EpsilonI
	cl.Fitness = xcs.FitnessI
	cl.Experience = 0

	// TODO: set up a globally accessible timestamp to be copied into the new Chromosome

	cl.MeanASSize = 1
	cl.Numerosity = 1
	return
}

// GeneratePredictionArray returns a slice of prediction values for each action
func (xcs *XCS) GeneratePredictionArray(ms Population) (pa map[int][]float64) {
	// TODO: write test for GeneratePredictionArray
	pa = make(map[int][]float64)
	fsa := make(map[int]float64)
	for _, cl := range ms {
		for task, pred := range cl.Prediction {
			pa[cl.Action][task] += pred * cl.Fitness
			fsa[cl.Action] += cl.Fitness
		}
	}
	for action, _ := range pa {
		if fsa[action] != 0 {
			for task := range pa[action] {
				pa[action][task] /= fsa[action]
			}
		}
	}
	return
}
