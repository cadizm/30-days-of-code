package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)

	phonebook := make(map[string]string)

	for i := int64(0); i < n; i++ {
		scanner.Scan()
		namephone := strings.Split(scanner.Text(), " ")
		phonebook[namephone[0]] = namephone[1]
	}

	for scanner.Scan() {
		name := scanner.Text()
		if phone, ok := phonebook[name]; ok {
			fmt.Printf("%v=%v\n", name, phone)
		} else {
			fmt.Println("Not found")
		}
	}
}
