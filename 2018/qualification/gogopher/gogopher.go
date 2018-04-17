package main

import (
	"fmt"
	"math"
	"os"
)

func findSquareWMaxOpenAroundIt(theMap [][]bool) (int, int) {
	lmy := len(theMap)
	lmx := len(theMap[0])
	maxSoFar := 0
	maxX := 0
	maxY := 0
	for yy := 0; yy < lmy; yy++ {
		for xx := 0; xx < lmx; xx++ {
			count := 0
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					ty := yy + ii
					tx := xx + jj
					if (ty >= 0) && (ty < lmy) && (tx >= 0) && (tx < lmx) {
						if !theMap[ty][tx] {
							count++
						}
					}
				}
			}
			if count > maxSoFar {
				maxSoFar = count
				maxY = yy
				maxX = xx
			}
		}
	}
	return maxY, maxX
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 1; i <= t; i++ {
		var a, x, y int
		fmt.Scanf("%d", &a)
		stpl := math.Sqrt(float64(a))
		mapx := int(math.Floor(stpl))
		for (a % mapx) != 0 {
			mapx--
		}
		mapy := a / mapx
		// we have to have at least 3 so the opponent can't pick coordinates off our map
		if mapy < 3 {
			mapy = 3
		}
		if mapx < 3 {
			mapx = 3
		}
		theMap := make([][]bool, mapy)
		for i := 0; i < mapy; i++ {
			theMap[i] = make([]bool, mapx)
		}
		keepGoing := true
		for keepGoing {

			ourY, ourX := findSquareWMaxOpenAroundIt(theMap)
			// make sure we don't hit out of bounds
			if ourY == 0 {
				ourY++
			}
			if ourX == 0 {
				ourX++
			}
			if ourY == (mapy - 1) {
				ourY--
			}
			if ourX == (mapx - 1) {
				ourX--
			}

			fmt.Println(ourX+2, ourY+2)
			fmt.Scanf("%d %d", &x, &y)
			if (x == -1) && (y == -1) {
				os.Exit(1) // we screwed up, so we have to crash
			}
			if (x == 0) && (y == 0) {
				keepGoing = false // go to next test case
			} else {
				replyX := x - 2
				replyY := y - 2
				theMap[replyY][replyX] = true
			}
		}
	}
}
