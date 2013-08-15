package main

// Population contains a slice of Chromosomes
type Population []*Chromosome

// Len returns the number of elements in the population
func (pop *Population) Len() int {
	return len(*pop)
}

// Less returns whether the ith element compares less than the jth element
func (pop *Population) Less(i, j int) bool {
	return (*pop)[i].Fitness < (*pop)[j].Fitness
}

// Swap swaps the position of the ith and jth elements of the population
func (pop *Population) Swap(i, j int) {
	tmp := (*pop)[i]
	(*pop)[i] = (*pop)[j]
	(*pop)[j] = tmp
}

func (pop *Population) Deletion() {
	
}