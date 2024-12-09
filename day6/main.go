package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/carlaurr/aoc-2024/utils"
)

var guardSymbols = "<>^v"

func printLabMap(labMap [][]rune) {
	for i := 0; i < len(labMap); i++ {
		fmt.Println(string(labMap[i]))
	}
}

func locateGuard(labMap [][]rune) ([2]int, rune) {
	var position [2]int

	for i := 0; i < len(labMap); i++ {
		for j := 0; j < len(labMap[i]); j++ {
			if strings.ContainsRune(guardSymbols, labMap[i][j]) {
				position[0] = i
				position[1] = j
				break
			}
		}
	}

	return position, labMap[position[0]][position[1]]
}

func turnRightGuard(currentDirection rune) rune {
	switch currentDirection {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}

	return ' '
}

// isStuck checks if the guard is stuck in a loop. It will happen if the guard visited the same positions the same number of times
func isStuck(movementsCount int, labMap [][]rune) bool {
	if (movementsCount > len(labMap) * len(labMap[0])) {
		return true
	}
	return false
}

func processGuardRoute(labMap [][]rune) (int, bool) {
	var guardPos [2]int
	var guardDir rune

	distinctPositions := make(map[[2]int]int)
	movementsCount := 0

	var directions = map[rune][2]int{
    '^': {-1, 0},
    '>': {0, 1},
    'v': {1, 0},
    '<': {0, -1},
	}

	// include the guard starting position
	positionsCount := 1

	// locate the guard starting position
	guardPos, guardDir = locateGuard(labMap)
	distinctPositions[guardPos] = 1

	// move the guard
	row := guardPos[0]
	column := guardPos[1]

	for {
		nextPosition := [2]int{row + directions[guardDir][0], column + directions[guardDir][1]}
		movementsCount++

		if nextPosition[0] < 0 || nextPosition[0] >= len(labMap) || nextPosition[1] < 0 || nextPosition[1] >= len(labMap[nextPosition[0]]) {
			// guard has leave the map area
			break
		}

		if (labMap[nextPosition[0]][nextPosition[1]] == '#') {
			guardDir = turnRightGuard(guardDir)
			
			if isStuck(movementsCount, labMap) {
				return positionsCount, true
			}
			continue
		}

		// update the guard position
		row = nextPosition[0]
		column = nextPosition[1]
		if _, ok  := distinctPositions[nextPosition]; !ok {
			distinctPositions[nextPosition] = 1
			positionsCount++
		} else {
			distinctPositions[nextPosition]++
		}
	}


	return positionsCount, false
}


func main() {	

	var fileName string
	var labMap [][]rune

	args := os.Args

	if len(args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = args[1]
	}

	labMap = utils.BuildMap(fileName)

	obsDifPositions := 0

	for i := 0; i < len(labMap); i++ {
		for j := 0; j < len(labMap[i]); j++ {
			if (labMap[i][j] == '.') {
				labMap[i][j] = '#'
				_, isStuck := processGuardRoute(labMap)

				if isStuck {
					obsDifPositions++;
				}

				labMap[i][j] = '.'
			}
		}
	}

	fmt.Printf("Number of different positions for the obstacle: %d\n", obsDifPositions)

	//positions, isStuck := processGuardRoute(labMap)
	//fmt.Printf("\nResult: %d, is stuck? %v\n", positions, isStuck)
}