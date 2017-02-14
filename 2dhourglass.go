package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int64
type Hourglass []int64

const stride = 3
const hourglassLen = 7

func readGrid() Grid {
	var grid Grid
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; scanner.Scan(); i++ {
		strArr := strings.Split(scanner.Text(), " ")
		arr := make([]int64, len(strArr))

		for i, v := range strArr {
			arr[i], _ = strconv.ParseInt(v, 0, 0)
		}

		grid = append(grid, arr)
	}

	return grid
}

func (h Hourglass) sum() int64 {
	var sum int64

	for i := 0; i < len(h); i++ {
		sum += h[i]
	}

	return sum
}

func (g Grid) hourglass(i, j int) Hourglass {
	h := make(Hourglass, hourglassLen)

	h[0] = g[i+0][j+0]
	h[1] = g[i+0][j+1]
	h[2] = g[i+0][j+2]
	h[3] = g[i+1][j+1]
	h[4] = g[i+2][j+0]
	h[5] = g[i+2][j+1]
	h[6] = g[i+2][j+2]

	return h
}

func main() {
	grid := readGrid()

	// init with min since grid can have negative values
	// and we want the smallest hourglass in the grid
	var max int64 = math.MinInt64

	for i := 0; i <= stride; i++ {
		for j := 0; j <= stride; j++ {
			h := grid.hourglass(i, j)
			if sum := h.sum(); sum > max {
				max = sum
			}
		}
	}

	fmt.Println(max)
}
