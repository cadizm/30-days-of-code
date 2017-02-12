package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	curRun, maxRun := 0, 0

	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			curRun += 1
			if curRun > maxRun {
				maxRun = curRun
			}
		} else {
			curRun = 0
		}
	}

	fmt.Println(maxRun)
}
