package main // making package for standalone executable

import (
	"fmt"
	"math"
	"studygo/print"
)

const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
const BEEF, TWO, C = "meat", 2, "veg"

func main() { // making an entry point
	// printing using fmt functionality
	fmt.Println("Hello World Go")

	// import print.go
	print.ToConsole("Hello World Go")
	print.ToConsole("Single return values: ", sum(1, 5))
	fmt.Print("Multiple return values: ")
	print.ToConsole(doMath(1, 5))
	print.ToConsole("Constants: ", MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY)
	print.ToConsole("Constants: ", BEEF, TWO, C)
	print.ToConsole("Enum: ", MALE, FEMALE, UNKNOWN)
	print.ToConsole(KB, MB, GB, TB, PB, EB)
	defaultVariableValues()
}

// function with return type
func sum(a float64, b float64) float64 {
	return a + b
}

// function with multiple return type
func doMath(a float64, b float64) (plus float64, minus float64, abs int, cst string) {
	var p = a + b
	m := a - b
	const CONST string = "random constant"
	return p, m, int(math.Abs(m)), CONST
}

// enumeration
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

// iota to assign values
const (
	UNKNOWN = iota
	FEMALE  = iota
	MALE    = iota
)

// variable declaration
func defaultVariableValues() {
	var number int    // Declaring  an integer variable
	var decision bool // Declaring a boolean variable
	var name string   // Declaring a string variable
	print.ToConsole("Variables: ", number, decision, name == "")
}
