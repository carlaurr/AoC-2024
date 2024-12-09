package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ID 			  int
	Children 	[]*Node
}

func (n *Node) Contains (id int) bool {
	for _, child := range n.Children {
		if child.ID == id {
			return true
		}
	}
	return false
}

var scanner *bufio.Scanner

func buildGraph(fileName string) map[int]*Node {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	graph := make(map[int]*Node)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// Rules processing is done
			return graph
		}
		
		parts := strings.Split(line, "|")
		
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return nil
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return nil
		}

		// update the graph
		if _, ok := graph[num1]; !ok {
			graph[num1] = &Node{ID: num1, Children: []*Node{}}
		}

		if _, ok := graph[num2]; !ok {
			graph[num2] = &Node{ID: num2, Children: []*Node{}}
		}

		// update the children
		graph[num1].Children = append(graph[num1].Children, graph[num2])
	}

	return graph
}

func processUpdate(graph map[int]*Node, pages []string) int {
	middlePageNumber := 0

	for i := 0; i < len(pages); i++ {
		pageNumber, err := strconv.Atoi(pages[i])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return 0
		}

		// update middle page number
		if i == len(pages)/2 {
			middlePageNumber = pageNumber
		}

		for j := i + 1; j < len(pages); j++ {
			nextPageNumber, err := strconv.Atoi(pages[j])
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
				return 0
			}

			if graph[nextPageNumber].Contains(pageNumber) {
				fmt.Printf("Line is invalid: %v, next page %d should be before %d\n", pages, nextPageNumber, pageNumber)
				return 0
			}
		}
	}

	fmt.Printf("Line is valid: %v, middle page number %d\n", pages, middlePageNumber)

	return middlePageNumber
}

func processUpdateV2(graph map[int]*Node, pages []string) int {
	i := 0
	corrected := false

	for i < len(pages) - 1 {
		pageNumber, err := strconv.Atoi(pages[i])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return 0
		}

		j := i + 1

		fmt.Printf("Checking page number: %d, index i: %d, index j: %d \n", pageNumber, i, j)

		for j < len(pages) {
			nextPageNumber, err := strconv.Atoi(pages[j])
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
				return 0
			}

			if graph[nextPageNumber].Contains(pageNumber) {
				// swap
				corrected = true
				aux := pages[i]
				pages[i] = pages[j]
				pages[j] = aux
				break
			} 

			j++
			if j == len(pages) {
				i++
			}
		}
	}

	if corrected {
		middlePage, err := strconv.Atoi(pages[len(pages)/2])
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return 0
		}

		fmt.Printf("Line corrected: %v, middle page number %d\n", pages, middlePage)
		return middlePage
	}

	// line originally correct
	return 0
}

func processUpdates(graph map[int]*Node) int {
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Processing line: ", line)
		pages := strings.Split(line, ",")

		//result += processUpdate(graph, pages)
		result += processUpdateV2(graph, pages)
	}

	return result
}

func printGraph(graph map[int]*Node) {
	for key, node := range graph {
		fmt.Printf("Node %d: ", key)
		for _, child := range node.Children {
			fmt.Printf("%d, ", child.ID)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	var fileName string
	var graph map[int]*Node

	args := os.Args

	if len(args) < 2 {
		fileName = "input.txt"
	} else {
		fileName = args[1]
	}

	graph = buildGraph(fileName)

	printGraph(graph)

	result := processUpdates(graph)

	fmt.Println("Result: ", result)
}