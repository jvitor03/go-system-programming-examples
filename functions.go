/*
	Functions declarations in Go
*/

package main

import (
	"fmt"
)

// Declares a unnamed function
func unnamedMinMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

// Declares a named function
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

// Creates a named function and returns it(the default is to return the named vars)
func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

// creates a unnamed vars, and returns anything int, int
func sort(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}

func main() {
	y := 4
	square := func(s int) int {
		return s * s
	}
	fmt.Println("the square of ", y, "is ", square(y))

	// Create an anonymous function with an int param and int return
	square = func(s int) int {
		return s + s
	}
	fmt.Println("the double of ", y, "is ", square(y))
	fmt.Println(minMax(16, 6))
	fmt.Println(namedMinMax(15, 6))
	min, max := namedMinMax(12, -1)
	fmt.Println(min, max)
}
