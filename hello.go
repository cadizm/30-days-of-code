package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	line := scanner.Text()

	fmt.Println("Hello, World.")
	fmt.Println(line)
}
