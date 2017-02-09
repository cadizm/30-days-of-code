package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // throw away first line

	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")

	reverse(s)
	for _, c := range s {
		fmt.Printf("%s ", c)
	}
}
