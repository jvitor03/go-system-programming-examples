/*
 Arguments handling in Go
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0
	for i := 1; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i]) // _ is for ignoring error return value
		sum = sum + temp
	}
	fmt.Println("sum:", sum)
}
