// Defines a convenient interface for configuration information.
//
// Hides the details of parsing out the parameters and provides additional 
// functionality (e.g., parsing an array of int or float64 values.

package conf

import (
	"fmt"
	gc "github.com/dlintw/goconf"
	"strings"
	"strconv"
	"os"
)

// private variable that stores the parsed configuration info
var (
	rc *gc.ConfigFile
)

// InitConfig parses the configuration file, returning any error encountered
func Init(file string) (err error) {
	rc, err = gc.ReadConfigFile(file)
	return 
}

// StringParameter returns the value of a given parameter as a string
func StringParameter(sec, name string) (val string, err error) {
	val, err = rc.GetString(sec, name)
	return
}

// IntParameter returns the value of the parameter as an int
func IntParameter(sec, name string) (val int, err error) {
	val, err = rc.GetInt(sec, name)
	return 
}

// UintParameter returns the value of the parameter as an uint
func UintParameter(sec, name string) (val uint, err error) {
	// TODO: replace conversion with GetUint operator
	var ival int
	ival, err = rc.GetInt(sec, name)
	val = uint(ival)
	return 
}

// Float64Parameter returns the value of the parameter as a float64
func Float64Parameter(sec, name string) (val float64, err error) {
	val, err = rc.GetFloat64(sec, name)
	return 
}

// IntArrayParameter returns the value of the parameter as a slice of int64s
func IntArrayParameter(sec, name string) (val []int, err error) {
	var s string
	if s, err = rc.GetString(sec, name); err != nil {
		val = []int{}
	} else {
		val = parseIntVector(s, " ")
	}
	return 
}

// Float64ArrayParameter returns the value of the parameter as a slice of float64s
func Float64ArrayParameter(sec, name string) (val []float64, err error) {
	var s string
	if s, err = rc.GetString(sec, name); err != nil {
		val = []float64{}
	} else {
		val = parseFloat64Vector(s, " ")
	}
	return 
}

// parseFloat64Vector parses a string of numbers separated by spaces into a slice of float64s
func parseFloat64Vector(str, sep string) []float64 {
	tokens := strings.Split(str, sep)
	vals := make([]float64, len(tokens))
	for i := range tokens {
		val, err := strconv.ParseFloat(tokens[i], 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			vals[i] = val
		}
	}
	return vals
}

// parseIntVector parses a string of numbers separated by spaces into a slice of ints
func parseIntVector(str, sep string) []int {
	tokens := strings.Split(str, sep)
	vals := make([]int, len(tokens))
	for i := range tokens {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			vals[i] = val
		}
	}
	return vals
}
