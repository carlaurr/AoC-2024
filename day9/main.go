package main

import (
	"fmt"
	"os"
	"bufio"
)

func processDiskMap(fileName string) ([]rune, []int, error) {
	var disk []rune
	var input []int

	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	isFile := true
	fileID := 0

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, nil, err
		}

		input = append(input, int(r - '0'))
		
		if isFile {
			for i := 0; i < int(r - '0'); i++ {
				disk = append(disk, rune(fileID + '0'))
			}

			fileID++
		} else {
			for i := 0; i <int(r - '0'); i++ {
				disk = append(disk, '.')
			}
		}
		isFile = !isFile

		fmt.Printf("%c", r)
	}

	return disk, input, nil
}

func calculateChecksum(disk []rune) uint64 {
	var checksum uint64 = 0

	for i, r := range disk {
		if r == '.' {
			continue
		}
		checksum += uint64(i) * uint64(r - '0')
	}

	return checksum
}

func compactProcess(disk []rune) uint64 {
	endIndex := len(disk) - 1

	for startIndex := 0; startIndex < endIndex; startIndex++ {
		if disk[startIndex] != '.' {
			continue
		}

		// swap free space with a block
		for endIndex > startIndex && disk[endIndex] == '.' {
			endIndex--
		}

		disk[startIndex] = disk[endIndex]
		disk[endIndex] = '.'
	}

	return calculateChecksum(disk)
}

func compactProcessV2(disk []rune, input []int) uint64 {
	startIndex := 0
	endIndex := len(disk) - 1
	blocksChecked := make(map[int]bool)

	for endIndex >= 0 && disk[endIndex] != '0' {
		for endIndex >= 0 && disk[endIndex] == '.' {
			endIndex--
		}
		if endIndex < 0 || disk[endIndex] == '0' {
			break
		}

		currentBlock := int(disk[endIndex] - '0')
		blocksOccupied := input[currentBlock*2]

		// skip blocks already checked
		if blocksChecked[currentBlock] {
			endIndex -= blocksOccupied
			continue
		}
		blocksChecked[currentBlock] = true

		for startIndex < endIndex-blocksOccupied+1 {
			freeStart := startIndex
			for freeStart < endIndex && disk[freeStart] != '.' {
				freeStart++
			}

			freeSpace := 0
			for freeStart+freeSpace < len(disk) && disk[freeStart+freeSpace] == '.' {
				freeSpace++
			}

			if freeSpace >= blocksOccupied {
				fmt.Printf("--> Swapping block %d\n", currentBlock)
				for i := 0; i < blocksOccupied; i++ {
					disk[freeStart+i] = disk[endIndex-i]
					disk[endIndex-i] = '.'
				}
				endIndex -= blocksOccupied
				break
			} else {
				startIndex = freeStart + freeSpace
			}
		}

		startIndex = 0
	}

	return calculateChecksum(disk)
}

 
func printDisk(disk []rune) {
	fmt.Printf("\n")
	for _, r := range disk {
		fmt.Printf("%c", r)
	}
	fmt.Printf("\n")
}

func main() {
	fileName := "input.txt"
	args := os.Args
	if len(args) >= 2 {
		fileName = args[1]
	}

	disk, intput, err := processDiskMap(fileName)
	if err != nil {
		fmt.Println("Error processing disk map: ", err)
		return
	}

	printDisk(disk)

	// checksum := compactProcess(disk)
	checksum := compactProcessV2(disk, intput)
	
	printDisk(disk)

	fmt.Printf("\nChecksum: %d\n", checksum)
}