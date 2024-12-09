package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"github.com/carlaurr/aoc-2024/utils"
)

func removeElement(original []string, index int) []string {
	copy := make([]string, len(original) - 1)
	j :=  0
	for i, v := range original {
		if (i != index) {
			copy[j] = v
			j ++
		}
	}

	return copy;
}

func isSafe(records []string) bool {
	if len(records) == 1 {
		return true
	}

	prevRecord, err := strconv.Atoi(records[0])
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return false
	}

	isAscending := false

	for i := 1; i < len(records); i++ {
		num, err := strconv.Atoi(records[i])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return false
		}

		dif:= num - prevRecord

		// initialize the isAscending flag if it's the first iteration
		if i == 1 {
			isAscending = dif > 0
		}

		if (isAscending && dif < 0) || (!isAscending && dif > 0) {
			return false
		}

		if utils.Abs(dif) < 1 || utils.Abs(dif) > 3 {
			return false
		}

		prevRecord = num
	}

	return true
}


func isSafeV2(records []string, tolerationUsed bool) bool {
	if len(records) == 1 {
		return true
	}

	prevRecord, err := strconv.Atoi(records[0])
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return false
	}

	isAscending := false

	for i := 1; i < len(records); i++ {
		num, err := strconv.Atoi(records[i])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return false
		}

		dif:= num - prevRecord

		// initialize the isAscending flag if it's the first iteration
		if i == 1 {
			isAscending = dif > 0
		}

		if (isAscending && dif < 0) || (!isAscending && dif > 0) || 
				utils.Abs(dif) < 1 || utils.Abs(dif) > 3 {
			if !tolerationUsed {
				// Brute force. remove level by level and check if it's safe
				for j := 0; j < len(records); j++ {
					newRecords := removeElement(records, j)
					if isSafeV2(newRecords, true) {
						return true
					}
				}
			}

			return false
		}

		prevRecord = num
	}

	return true
}

func processInput(fileName string) (int, error) {
	safeRecords := 0

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return -1, err
	}
	defer file.Close()

	
	re := regexp.MustCompile("[^0-9]+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		records := re.Split(line, -1)

		// if (isSafe(records)) {
		if (isSafeV2(records, false)) {
			safeRecords++
		}
	}

	return safeRecords, nil
}


func main() {
	fileName := "input.txt"
	safeRecords, err := processInput(fileName)
	if err != nil {
		fmt.Println("Error processing input: ", err)
		return
	}

	fmt.Println("Safe records: ", safeRecords)
}