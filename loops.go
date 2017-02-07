package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const limit = 10

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)

	for i := int64(1); i <= limit; i++ {
		fmt.Printf("%v x %v = %v\n", n, i, n*i)
	}
}
