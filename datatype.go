package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var j uint64
	var f float64
	var t string

	var input []string

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i, v := range input {
		switch i {
		case 0:
			j, _ = strconv.ParseUint(v, 0, 64) // ignore errors
		case 1:
			f, _ = strconv.ParseFloat(v, 64) // ignore errors
		case 2:
			t = v
		default:
			// do nothing
		}
	}

	// Print the sum of both integer variables on a new line.
	fmt.Printf("%v\n", i+j)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+f)

	// Concatenate and print the String variables on a new line
	fmt.Printf("%v\n", s+t)
}
