/*
	Deferred functions wait the end of the caller function to run
*/

package main

import (
	"fmt"
)

// Calling function with defer, passes the value of i as parameter in the moment of the call
func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

// Defer in anonymous function, using the value of i from a2, when the function is executed, the value of i is already 3, which is the end of the loop
func a2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

// Defer in anonymous function using a parameter that takes the value from i, at each iteration
func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()
}
