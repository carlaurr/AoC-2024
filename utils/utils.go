package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BuildMap(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	var inputMap [][]rune

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			return inputMap
		}
		
		inputMap = append(inputMap, []rune(line))
	}

	return inputMap
}