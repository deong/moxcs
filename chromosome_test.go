package main

import (
	"testing"
)

// TestInit tests the chromosome initialization method
func TestInit(t *testing.T) {
	chr := NewChromosome(20, 2)
	if len(chr.Condition) != 20 || len(chr.Prediction) != 2 {
		t.Error("chromosome initialization failed")
	}

	chr = NewChromosome(100, 5)
	if len(chr.Condition) != 100 || len(chr.Prediction) != 5 {
		t.Error("chromosome initialization failed")
	}
}

func TestMatch(t *testing.T) {
	chr := NewChromosome(20, 1)
	chr.Condition = []byte("*00001111100**01**1*")
	str1 := "00000111110000011111"
	str2 := "10000111110001010010"
	str3 := "11001010010100101010"
	str4 := "00101101101110101001"
	if !chr.Matches(str1) {
		t.Errorf("error matching rule condition %v with state %v\n", string(chr.Condition), str1)
	}
	if !chr.Matches(str2) {
		t.Errorf("error matching rule condition %v with state %v\n", string(chr.Condition), str2)
	}
	if chr.Matches(str3) {
		t.Errorf("error matching rule condition %v with state %v\n", string(chr.Condition), str3)
	}
	if chr.Matches(str3) {
		t.Errorf("error matching rule condition %v with state %v\n", string(chr.Condition), str4)
	}
}

