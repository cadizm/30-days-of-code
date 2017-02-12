package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func factorial(n int64) int64 {
	switch {
	case n < 0:
		panic("factorial: n must be nonnegative")
	case n < 2:
		return n
	}

	return n * factorial(n-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)
	fmt.Println(factorial(n))
}
