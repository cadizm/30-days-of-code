package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // throwaway first line

	for scanner.Scan() {
		var evenBuf bytes.Buffer
		var oddBuf bytes.Buffer

		for i, c := range scanner.Text() {
			if i%2 == 0 {
				evenBuf.WriteRune(c)
			} else {
				oddBuf.WriteRune(c)
			}
		}
		fmt.Printf("%v %v\n", evenBuf.String(), oddBuf.String())
	}
}
