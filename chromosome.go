package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Chromosome is the basic XCS rule data structure
type Chromosome struct {
	Condition  []byte
	Action     int
	Prediction []float64
	Error      float64
	Fitness    float64
	Experience int
	TimeStamp  int
	MeanASSize float64
	Numerosity int
}

// NewChromosome constructs and initializes a chromosome
func NewChromosome(nBits, nTasks int) *Chromosome {
	return &Chromosome{make([]byte, nBits), 0, make([]float64, nTasks), 0, 0, 0, 0, 0, 0}
}

// Randomize sets the condition to a random ternary string
func (chr *Chromosome) Randomize() {
	for i := 0; i < len(chr.Condition); i++ {
		switch b := rand.Intn(3); b {
		case 0:
			chr.Condition[i] = '0'
		case 1:
			chr.Condition[i] = '1'
		case 2:
			chr.Condition[i] = '*'
		}
	}
	// TODO: delete this
	chr.Action = rand.Intn(8)
	chr.Fitness = float64(rand.Intn(50))
}

// Matches returns true if the chromosome's condition matches the current
// environment string
func (chr *Chromosome) Matches(cond string) bool {
	for i := range chr.Condition {
		if chr.Condition[i] != '*' && chr.Condition[i] != cond[i] {
			return false
		}
	}
	return true
}

// String converts a chromosome into a human readable representation
func (chr *Chromosome) String() (output string) {
	output = strings.Join([]string{
		fmt.Sprintf("condition: %s", string(chr.Condition)),
		fmt.Sprintf("action: %v", chr.Action),
		fmt.Sprintf("prediction: %v", chr.Prediction),
		fmt.Sprintf("error: %v", chr.Error),
		fmt.Sprintf("fitness: %v", chr.Fitness),
		fmt.Sprintf("experience: %v", chr.Experience),
		fmt.Sprintf("timestamp: %v", chr.TimeStamp),
		fmt.Sprintf("meanassize: %v", chr.MeanASSize),
		fmt.Sprintf("numerosity: %v", chr.Numerosity),
	}, "\n")
	return 
}
