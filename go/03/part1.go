package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {

	input := getInput("input.txt")
	getVisitedHouses(input)

}

func makeMapKey(x int, y int) string {
	return "x" + strconv.Itoa(x) + ":y" + strconv.Itoa(y)
}

func getVisitedHouses(input []string) {

	x := 0
	y := 0

	visitedHouses := make(map[string]bool)

	visitedHouses[makeMapKey(x,y)] = true

	for _, line := range input {
		for _, char := range line {
			switch string(char){
			case "<":
				x--
			case ">":
				x++
			case "^":
				y++
			case "v":
				y--
			default:
				fmt.Println("unknown input char %s", string(char))
			}

			visitedHouses[makeMapKey(x,y)] = true

		}
	}

	// Each unique key is a house that santa visited
	fmt.Printf("Santa visited %d houses\n", len(visitedHouses))
}

func getInput(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)

	// Read the lines of the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}