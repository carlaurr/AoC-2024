package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateValue(eqValues []string, operators []rune) uint64 {
	value, err := strconv.ParseUint(eqValues[0], 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return 0
	}

	fmt.Printf("\nCalculating %d ", value)

	if len(operators) == 0 {
		num, err := strconv.ParseUint(eqValues[0], 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return 0
		}
		fmt.Printf("= %d", num)
		return num
	}

	for i := 0; i < len(operators); i++ {
		num, err := strconv.ParseUint(eqValues[i + 1], 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return 0
		}

		if operators[i] == '+' {
			value += num
		} else if operators[i] == '*' {
			value *= num
		} else if operators[i] == '|' {
			if i == 0 {
				value, err = strconv.ParseUint(eqValues[i] + eqValues[i + 1], 10, 64)
			} else {
				strValue := strconv.FormatUint(value, 10)
				value, err = strconv.ParseUint(strValue + eqValues[i + 1], 10, 64)
			}

			if err != nil {
				fmt.Println("Error converting string to int: ", err)
				return 0
			}
 		}

		fmt.Printf("%c %d ", operators[i], num)
	}

	fmt.Printf("= %d", value)

	return value
}

func copyArray(original []rune) []rune {
	copy := make([]rune, len(original))
	for i, v := range original {
		copy[i] = v
	}

	return copy
}

func checkIfEqValid(eqValues []string, operators []rune, value uint64, index int) (uint64, bool) {
	var eqResult uint64 = 0

	// Base case: if we have reach the end of the operators without finding a valid equation
	if index == len(operators) {
		return 0, false
	}

	if len(operators) == 0 {
		eqResult = calculateValue(eqValues, operators)
		if eqResult == value {
			return eqResult, true
		}
		return 0, false
	}
	
	operatorsCopy := copyArray(operators)

	eqResult = calculateValue(eqValues, operatorsCopy)
	fmt.Printf(" (%d)\n", value)
	
	if eqResult == value {
		return eqResult, true
	}

	eqResult, valid := checkIfEqValid(eqValues, operatorsCopy, value, index + 1)
	if valid {
		return eqResult, true
	}
	

	// Try changing the operator
	operatorsCopy[index] = '*'
	eqResult = calculateValue(eqValues, operatorsCopy)
	fmt.Printf(" (%d)\n", value)

	if eqResult == value {
		return eqResult, true
	}

	eqResult, valid = checkIfEqValid(eqValues, operatorsCopy, value, index + 1)
	if valid {
		return eqResult, true
	}

	// Try concatenating
	operatorsCopy = copyArray(operators)
	operatorsCopy[index] = '|'

	eqResult = calculateValue(eqValues, operatorsCopy)
	fmt.Printf(" (%d)\n", value)

	if eqResult == value {
		return eqResult, true
	}

	return checkIfEqValid(eqValues, operatorsCopy, value, index	+ 1)
}

func isTestValueValid(test string) (uint64, bool) {
	parts := strings.Split(test, ": ")

	value, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return 0, false
	}

	eqValues := strings.Split(parts[1], " ")
	operators := make ([]rune, len(eqValues) - 1)

	// initialize operators
	for i := 0; i < len(operators); i++ {
		operators[i] = '+'
	}

	result, valid := checkIfEqValid(eqValues, operators, value, 0)

	if valid {
		return result, true
	}

	return 0, false
}

func processCalibrations(fileName string) uint64 {
	file, err := os.Open(fileName)
	var result uint64 = 0

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, valid := isTestValueValid(scanner.Text())
		if valid {
			result += value
		}
	}

	return result
}

func main() {
	var fileName string
	args := os.Args

	if len(args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = args[1]
	}

	
	result := processCalibrations(fileName)

	fmt.Printf("\nResult: %d\n", result)
}