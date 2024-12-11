package main

import (
	"fmt"
	"os"
)

var memo = make(map[int]int)

func getInitialStones(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	defer file.Close()

	var stones []int
	var stone int
	for {
		_, err := fmt.Fscanf(file, "%d", &stone)
		if err != nil {
			break
		}
		stones = append(stones, stone)
	}

	return stones
}

func pow10(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
			result *= 10
	}
	return result
}

func blink(memo map[int]int) map[int]int {
	newMemo := make(map[int]int)
	
	for k, v := range memo {
		if k == 0 {
			newMemo[1] += v
		} else {
			num := k
			digitCount := 0

			// Count the number of digits (by dividing by 10)
			for num > 0 {
				num /= 10
				digitCount++
			}

			if digitCount%2 == 0 {
				// Even number of digits, split in the middle
				left := k / pow10(digitCount/2)
				right := k % pow10(digitCount/2)
				
				newMemo[left] += v
				newMemo[right] += v
			} else {
				newMemo[k*2024] += v
			}
		}
	}

	return newMemo
}

func countStones(memo map[int]int) int {
	count := 0
	for _, v := range memo {
		count += v
	}
	return count
}

func main() {
	fileName := "input.txt"
	blinks := 75

	args := os.Args
	if len(args) > 1 {
		fileName = args[1]
	}

	stones := getInitialStones(fileName)
	

	memo = make(map[int]int)

	for _, stone := range stones {
		memo[stone]++
	}

	for i := 0; i < blinks; i++ {
		memo = blink(memo)
		fmt.Printf("\nNlinks num. %d", i+1)
	}

	fmt.Printf("\nTotal stones: %d\n", countStones(memo))
}
