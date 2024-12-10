package main

import (
	"fmt"
	"os"

	"github.com/carlaurr/aoc-2024/utils"
)

func printMap(trailMap [][]rune) {
	for _, row := range trailMap {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func exploreTrail(trailMap [][]rune, i, j, expectedHeight int, positionExplored [][]int, part2 bool) int {
	trailsCount := 0

	if i < 0 || i >= len(trailMap) || j < 0 || j >= len(trailMap[i]) {
		return 0
	}
	
	height := int(trailMap[i][j] - '0')
	if  height != expectedHeight {
		return 0
	}

	if part2 {
		if positionExplored[i][j] != -1 {
	 		return positionExplored[i][j]
	 }
	}

	if (height == 9) {
		if part2 {
			positionExplored[i][j] = 1
		} else {
			trailMap[i][j] = 'X'
		}
		return 1
	}

	trailsCount += exploreTrail(trailMap, i-1, j, height+1, positionExplored, part2)
	trailsCount += exploreTrail(trailMap, i+1, j, height+1, positionExplored, part2)
	trailsCount += exploreTrail(trailMap, i, j-1, height+1, positionExplored, part2)
	trailsCount += exploreTrail(trailMap, i, j+1, height+1, positionExplored, part2)

	positionExplored[i][j] = trailsCount

	return trailsCount
}

func copyTrailMap(trailMap [][]rune) [][]rune {
	copyMap := make([][]rune, len(trailMap))

	for i := 0; i < len(trailMap); i++ {
		copyMap[i] = make([]rune, len(trailMap[i]))
		for j := 0; j < len(trailMap[i]); j++ {
			copyMap[i][j] = trailMap[i][j]
		}
	}

	return copyMap
}


func getHikingTrails(trailMap [][]rune, positionExplored [][]int, memoryEnabled bool) int {
	hikingTrails := 0

	for i := 0; i < len(trailMap); i++ {
		for j := 0; j < len(trailMap[i]); j++ {
			if trailMap[i][j] == '0' {
				copyMap := copyTrailMap(trailMap)
				positionExplored = initPositionExplored(len(trailMap), len(trailMap[0]))
				hikingTrails += exploreTrail(copyMap, i, j, 0, positionExplored, memoryEnabled)
			}
		}
	}

	return hikingTrails
}

func initPositionExplored(rows, cols int) [][]int {
	positionExplored := make([][]int, rows)
	for i := 0; i < rows; i++ {
		positionExplored[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			positionExplored[i][j] = -1
		}
	}

	return positionExplored
}

func main() {
	fileName := "input.txt"

	args := os.Args
	if len(args) > 1 {
		fileName = args[1]
	}

	trailMap := utils.BuildMap(fileName)

	positionExplored := initPositionExplored(len(trailMap), len(trailMap[0]))

	// Part 1
	//hikingTrails := getHikingTrails(trailMap, positionExplored, false)

	// Part 2
	hikingTrails := getHikingTrails(trailMap, positionExplored, true)

	fmt.Println("\nHiking trails: ", hikingTrails)
}