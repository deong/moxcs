package main

import (
	"fmt"
	"flag"
	"github.com/deong/multixcs/conf"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	configFile := flag.String("conf", "", "A configuration file defining the parameters of the run.")
	flag.Parse()

	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := conf.Init(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	popSize, err := conf.IntParameter("xcs", "populationSize")
	if err != nil {
		fmt.Printf("required parameter '%v' missing\n", "populationSize")
		os.Exit(1)
	}

	pop := Population(make([]*Chromosome, popSize))
	for i := 0; i < popSize; i++ {
		chr := NewChromosome(10, 2)
		chr.Randomize()
		pop[i] = chr
	}
	sort.Sort(&pop)
	for i := 0; i < popSize; i++ {
		fmt.Println(pop[i], "\n")
	}

	thetaMNA, err := conf.IntParameter("xcs", "thetaMNA")
	if err != nil {
		fmt.Printf("required parameter '%v' missing", "thetaMNA")
		os.Exit(1)
	}
	env := "1000101110"
	fmt.Println("finding match set for state:", env)
	ms := GenerateMatchSet(pop, env, thetaMNA)
	fmt.Println(ms)
}
