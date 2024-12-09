package main

import (
	"bufio"
	"fmt"
	"os"
)

func buildMatrix(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	return matrix, nil
}

func printMatrix(matrix [][]rune) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Print(string(matrix[row][col]))
		}
		fmt.Println()
	}
}

func searchXmas(matrix [][]rune, row int, col int, rowLength int, columnLength int) int {
	xmas := []rune{'X', 'M', 'A', 'S'}
	xmasCount := 0

	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Diagonal up left
		{-1, 1},  // Diagonal up right
		{1, -1},  // Diagonal down left
		{1, 1},   // Diagonal down right
	}

	// helper function to check a single direction
	checkDirection := func(row, col, rowDelta, colDelta int) bool {
		for x := 1; x < len(xmas); x++ {
			row += rowDelta
			col += colDelta
			if row < 0 || row >= rowLength || col < 0 || col >= columnLength || matrix[row][col] != xmas[x] {
				return false
			}
		}
		return true
	}

	// Check all directions
	for _, dir := range directions {
		if checkDirection(row, col, dir[0], dir[1]) {
			fmt.Printf("Found XMAS starting at row: %d, col: %d, direction: (%d, %d)\n", row, col, dir[0], dir[1])
			xmasCount++
		}
	}

	return xmasCount
}

func checkChar(i int, j int, matrix [][]rune, expected rune) bool {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return false
	}

	return matrix[i][j] == expected
}

func searchXmasV2(matrix [][]rune, row int, col int, rowLength int, columnLength int) int {
	var i, j int

	// check left-right diagonal
	i = row - 1
	j = col - 1
	
	if checkChar(i, j, matrix, 'M') {
		i = row + 1
		j = col + 1
		if !checkChar(i, j, matrix, 'S') {
			return 0
		}
	} else if checkChar(i, j, matrix, 'S') {
		i = row + 1
		j = col + 1
		if !checkChar(i, j, matrix, 'M') {
			return 0
		}
	} else {
		return 0;
	}

	// check right-left diagonal
	i = row - 1
	j = col + 1
	if checkChar(i, j, matrix, 'M') {
		i = row + 1
		j = col - 1
		if !checkChar(i, j, matrix, 'S') {
			return 0
		}
	} else if checkChar(i, j, matrix, 'S') {
		i = row + 1
		j = col - 1
		if !checkChar(i, j, matrix, 'M') {
			return 0
		}
	} else {
		return 0;
	}

	return 1
}

func processMatrix(matrix [][]rune) int {
	totalXmasCount := 0
	rowLength := len(matrix)
	columnLength := len(matrix[0])

	fmt.Println("Row length: ", rowLength, " Column length: ", columnLength)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			// if matrix[row][col] == 'X' {
			// 	totalXmasCount += searchXmas(matrix, row, col, rowLength, columnLength)
			// }
			if matrix[row][col] == 'A' {
				//fmt.Println("Found X at row: ", row, " col: ", col)
				totalXmasCount += searchXmasV2(matrix, row, col, rowLength, columnLength)
			}
		}
	}

	return totalXmasCount
}

func main() {
  // build the character matrix
	matrix, err := buildMatrix("input.txt")
	if err != nil {
		fmt.Println("Error building matrix: ", err)
		return
	}

	// printMatrix(matrix)

	fmt.Printf("\nXMAS appears %d times\n", processMatrix(matrix))
}