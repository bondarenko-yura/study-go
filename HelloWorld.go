package main // making package for standalone executable
import (
	"fmt"
	"math"
) // importing a package

const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int = 1, 2, 3, 4, 5, 6
const BEEF, TWO, C = "meat", 2, "veg"

func main() { // making an entry point
	// printing using fmt functionality
	fmt.Println("Hello World Go")

	toConsole("Single return values: ", sum(1, 5))
	fmt.Print("Multiple return values: ")
	toConsole(doMath(1, 5))
	toConsole("Constants: ", MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY)
	toConsole("Constants: ", BEEF, TWO, C)
	toConsole("Enum: ", MALE, FEMALE, UNKNOWN)
} // exiting the program

func toConsole(v ...any) {
	fmt.Println(v)
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
	ZB
	YB
)

// iota to assign values
const (
	UNKNOWN = iota
	FEMALE  = iota
	MALE    = iota
)
