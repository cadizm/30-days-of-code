package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func round(f float64) int {
	i, frac := math.Modf(f)

	if frac >= 0.5 {
		return int(i) + 1
	}

	return int(i)
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	mealCost, _ := strconv.ParseFloat(input[0], 64)
	tipPercent, _ := strconv.ParseFloat(input[1], 64)
	taxPercent, _ := strconv.ParseFloat(input[2], 64)

	tip := mealCost * (tipPercent / 100.0)
	tax := mealCost * (taxPercent / 100.0)

	totalCost := mealCost + tip + tax

	fmt.Printf("The total meal cost is %v dollars.\n", round(totalCost))
}
