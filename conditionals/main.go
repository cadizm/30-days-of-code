package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func weird(n int64) string {
	if n%2 != 0 {
		return "Weird"
	}

	switch {
	case n < 6:
		return "Not Weird"
	case n < 21:
		return "Weird"
	}

	return "Not Weird"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)

	fmt.Println(weird(n))
}
