package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"github.com/carlaurr/aoc-2024/utils"
)

func processInput(fileName string) ([]int, []int, map[int]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil, nil, err
	}
	defer file.Close()

	var listA []int
	var listB []int
	repetitions := make(map[int]int)
	
	re := regexp.MustCompile("[^0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.Split(line, -1)

		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return nil, nil, nil, err
		}
		
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return nil, nil, nil, err
		}
		
		listA = append(listA, num1)
		listB = append(listB, num2)

		// update the repetitions map
		if _, exists := repetitions[num2]; exists {
			repetitions[num2]++;
		} else {
			repetitions[num2] = 1
		}
	}

	return listA, listB, repetitions, nil
}

func calculateTotalDistance(listA []int, listB []int) int {
	var totalDistance int
	totalDistance = 0

	for i := 0; i < len(listA); i++ {
		totalDistance += utils.Abs(listA[i] - listB[i])
	}

	return totalDistance
}

func calculateSimilarityScore(list []int, repetitions map[int]int) int {
	similarityScore := 0;

	for _, num := range list {
		reps, exists := repetitions[num]
		if exists {
			similarityScore += num * reps;
		}
	}

	return similarityScore;
}

func main() {
		listA, listB, repetitions, err := processInput("input.txt")

		if err != nil {
			fmt.Println("Error processing input: ", err)
			return
		}

		// check both list are valid
		if len(listA) != len(listB) {
			fmt.Println("Error: lists are not the same size")
			return
		}

		// sort the lists
		sort.Ints(listA)
		sort.Ints(listB)

		// calculate the total distance
		fmt.Println("Total distance: ", calculateTotalDistance(listA, listB))

		// calculate similarity score
		fmt.Println("Similarity score: ", calculateSimilarityScore(listA, repetitions))
}