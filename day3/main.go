package main

import (
	"fmt"
	"os"
	"bufio"
)

var char rune

func expectRune(reader *bufio.Reader, expected rune) error {
	var err error

	char, _, err = reader.ReadRune()
	if err != nil {
		return err
	}
	if char != expected {
		return fmt.Errorf("Expected '%c' but got %c", expected, char)
	}
	return nil
}

func parseNumber(reader *bufio.Reader, end rune) (int, error) {
	var err error

	num := 0
	digits := 0
	for {
		char, _, err = reader.ReadRune()
		if err != nil {
			return 0, err
		}
		if char == end && digits > 0 {
			break
		}

		if char == end && digits == 0 {
			return 0, fmt.Errorf("No number found")
		}

		if (digits >= 3) {
			return 0, fmt.Errorf("Number is too big")
		}

		if char < '0' || char > '9' {
			return 0, fmt.Errorf("Invalid character '%c'", char)
		}
		num = num * 10 + int(char - '0')
		digits++
	}

	return num, nil
}

func parseMul(reader *bufio.Reader) (uint64, error) {
	if err := expectRune(reader, 'u'); err != nil {
		return 0, err
	}
	if err := expectRune(reader, 'l'); err != nil {
		return 0, err
	}
	if err := expectRune(reader, '('); err != nil {
		return 0, fmt.Errorf("Expected '('")
	}

	// parse first number
	numA, err := parseNumber(reader, ',')
	if err != nil {
		return 0, err
	}
	// parse second number
	numB, err := parseNumber(reader, ')')
	if err != nil {
		return 0, err
	}
	
	return uint64(numA) * uint64(numB), nil
}

func parseInstruction(reader *bufio.Reader) (string, error) {
	var err error

	if err := expectRune(reader, 'o'); err != nil {	
		return "", err
	}

	char, _, err = reader.ReadRune()

	if err != nil {
		return "", err
	}

	if char == '(' {
		// check if it's a valid do() instruction
		if err := expectRune(reader, ')'); err != nil {
			return "", fmt.Errorf("Expected ')' but found '%c'", char)
		}

		return "do()", nil

	} else if char == 'n' {
		// check if it's a valid don't() instruction
		if err := expectRune(reader, '\''); err != nil {
			return "", fmt.Errorf("Expected '\\'' but found '%c'", char)
		}

		if err := expectRune(reader, 't'); err != nil {
			return "", fmt.Errorf("Expected 't' but found '%c'", char)
		}

		if err := expectRune(reader, '('); err != nil {
			return "", fmt.Errorf("Expected '(' but found '%c'", char)
		}

		if err := expectRune(reader, ')'); err != nil {
			return "", fmt.Errorf("Expected ')' but found '%c'", char)
		}

		return "don't()", nil

	} else {
		return "", err
	}
}

func main() {
	fileName := "input.txt"
	var resultTotal uint64 = 0
	isMulEnabled := true

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	char, _, err = reader.ReadRune()

	for {
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading the input file: ", err)
			return
		}

		if char == 'm' {
			result, err := parseMul(reader);
			if err == nil {
				if isMulEnabled {
					resultTotal += result
				}
				char, _, err = reader.ReadRune()
			}
		} else if char == 'd' {
			instruction, err := parseInstruction(reader);
			if err == nil {
				if instruction == "do()" {
					fmt.Printf("\nMultiplication enabled\n")
					isMulEnabled = true
					char, _, err = reader.ReadRune()
				} else if instruction == "don't()" {
					fmt.Printf("\nMultiplication disabled\n")
					isMulEnabled = false
					char, _, err = reader.ReadRune()
				}
			}
		} else {
			char, _, err = reader.ReadRune()
		}
	}

	fmt.Printf("\nTotal result: %d\n", resultTotal)
}
