package main

import (
	"fmt"
	"os"
	"github.com/carlaurr/aoc-2024/utils"
)

func printMap(labMap [][]rune) {
	for i := 0; i < len(labMap); i++ {
		fmt.Println(string(labMap[i]))
	}
}

type Antenna struct {
	x int
	y int
}

func getAntennasPositions(input [][]rune) map[rune][]Antenna {
	antennas := make(map[rune][]Antenna)

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '.' {
				continue
			}

			if _, ok := antennas[input[i][j]]; ok {
				antennas[input[i][j]] = append(antennas[input[i][j]], Antenna{x: i, y: j})
			} else {
				antennas[input[i][j]] = []Antenna{Antenna{x: i, y: j}}
			}
		}
	}
	return antennas
}

func printAntennas(antennas map[rune][]Antenna) {
	for k, v := range antennas {
		fmt.Printf("Antenna %c at positions: ", k)
		for _, antenna := range v {
			fmt.Printf("(%d, %d) ", antenna.x, antenna.y)
		}
		fmt.Println()
	}
}

func updateAntinode(input [][]rune, x, y int, antinodes map[[2]int]bool) {
	if x < 0 || y < 0 || x >= len(input) || y >= len(input[x]) {
		return
	}

	if _, ok := antinodes[[2]int{x, y}]; ok {
		return
	}

	antinodes[[2]int{x, y}] = true

	fmt.Printf("Antinode at (%d, %d)\n", x, y)

	if input[x][y] == '.' {
		// update the map
		input[x][y] = '#'
	}

}

func processAntennas(input [][]rune, antennas []Antenna, resonateHarmonics bool, antinodes map[[2]int]bool) {
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			difX := antennas[i].x - antennas[j].x
			difY := antennas[i].y - antennas[j].y

			antinodeAX := antennas[i].x + difX
			antinodeAY := antennas[i].y + difY

			antinodeBX := antennas[j].x - difX
			antinodeBY := antennas[j].y - difY

			for antinodeAX >= 0 && antinodeAY >= 0 && antinodeAX < len(input) && antinodeAY < len(input[i]) {
				updateAntinode(input, antinodeAX, antinodeAY, antinodes)

				if !resonateHarmonics {
					break
				}

				antinodeAY += difY;
				antinodeAX += difX;
			}

			for antinodeBX >= 0 && antinodeBY >= 0 && antinodeBX < len(input) && antinodeBY < len(input[i]) {
				updateAntinode(input, antinodeBX, antinodeBY, antinodes)

				if !resonateHarmonics {
					break
				}

				antinodeBY -= difY;
				antinodeBX -= difX;
			}

			if resonateHarmonics {
				// include the antinode at the antenna positions
				updateAntinode(input, antennas[i].x, antennas[i].y, antinodes)
				updateAntinode(input, antennas[j].x, antennas[j].y, antinodes)
			}
		}
	}
}

func findAntinodes(antennas map[rune][]Antenna, input [][]rune) map[[2]int]bool {
	antinodes := make(map[[2]int]bool)

	for k, v := range antennas {
		if len(v) == 1 {
			// no antinodes are possible for this antenna frequency
			continue
		}

		fmt.Printf("Processing antenna %c\n", k)

		processAntennas(input, v, true, antinodes)
	}

	return antinodes
}

func main() {
	var fileName string

	args := os.Args
	if len(args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = args[1]
	}

	inputMap := utils.BuildMap(fileName)
	antennas := getAntennasPositions(inputMap)
	printAntennas(antennas)

	antinodes := findAntinodes(antennas, inputMap)
	printMap(inputMap)

	fmt.Printf("\nAntinodes found: %d\n", len(antinodes))
}