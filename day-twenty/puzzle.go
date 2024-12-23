package main

import (
	"fmt"
	"log"
	"math"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	X, Y int
}

type Node struct {
	pos          Position
	currentLevel int
}

func main() {
	log.Printf("%v days to XMASSS", 24-20)
	input, err := utils.FileToStringMatrix("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	fmt.Println("hello")
	utils.PrintStringMatrix(input)
	// ******** FIRST AND SECOND PART **********
	getNumberOfSavingCheats(input)

}

func getNumberOfSavingCheats(raceMap [][]string) int {
	start := findNode(raceMap, "S")
	end := findNode(raceMap, "E")
	searchedMap := make(map[Position]Node)

	getNumberOfStepsToEnd(raceMap, start, end, searchedMap)

	numberOfCheats := doPathAgainAndGetCheats(raceMap, start, end, searchedMap)
	numberOfCheatsPartTwo := doPathAgainAndGetCheatsPartTwo(searchedMap)
	// part two 
	fmt.Printf("\nNumber of cheats that saves more than 100 are : %d \n", numberOfCheats)
	fmt.Printf("\nNumber of cheats part two : %d\n", numberOfCheatsPartTwo )

	return numberOfCheats
}

func getNumberOfStepsToEnd(raceMap [][]string, start Node, end Node, searchedMap map[Position]Node) int {

	searchQueue := append([]Node{}, getNeighbors(raceMap, start)...)
	searchedMap[start.pos] = start

	for len(searchQueue) > 0 {
		var tmpNeighbor Node
		tmpNeighbor, searchQueue = popleft(searchQueue)
		if _, exists := searchedMap[tmpNeighbor.pos]; !exists {
			if tmpNeighbor.pos == end.pos {
				searchedMap[tmpNeighbor.pos] = tmpNeighbor
				return tmpNeighbor.currentLevel
			} else {
				searchQueue = append(searchQueue, getNeighbors(raceMap, tmpNeighbor)...)
				searchedMap[tmpNeighbor.pos] = tmpNeighbor
			}
		}
	}

	return 0
}

// Part one
func doPathAgainAndGetCheats(raceMap [][]string, start Node, end Node, searchedMap map[Position]Node) int {
	total := 0
	//	cheatsMap := make(map[Position]bool)

	for _, node := range searchedMap {
		for _, neighbor := range getNeighborsDiamond(raceMap, node) {
			difference := searchedMap[neighbor.pos].currentLevel - node.currentLevel - 2
			if difference >= 100 {
				total++
			}
		}
	}

	return total
}

// Part two: Uses manhattan distances and calculates it for every pair of positions in the searched map
func doPathAgainAndGetCheatsPartTwo(searchedMap map[Position]Node) int {
	total := 0
	//	cheatsMap := make(map[Position]bool)

	positions := make([]Position, 0, len(searchedMap))
	for pos := range searchedMap {
		positions = append(positions, pos)
	}

    // Iterate over all pairs of positions
    for i := 0; i < len(positions); i++ {
        for j := 0; j < len(positions); j++ {
            p, q := positions[i], positions[j]
            d := manhattanDistance(p, q)
            if d < 21 && searchedMap[q].currentLevel-searchedMap[p].currentLevel-d >= 100 {
                total++
            }
        }
    }

	return total
}

// popleft removes the first element from a slice and returns the remaining slice
func popleft[T any](slice []T) (T, []T) {
	var zeroValue T
	if len(slice) == 0 {
		return zeroValue, slice
	}
	return slice[0], slice[1:]
}

func getNeighbors(memoryMap [][]string, currentNode Node) []Node {
	neighbors := make([]Node, 0)

	directions := []struct {
		dx, dy int
	}{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for _, d := range directions {
		x, y := currentNode.pos.X+d.dx, currentNode.pos.Y+d.dy
		if x >= 0 && x < len(memoryMap) && y >= 0 && y < len(memoryMap[0]) && memoryMap[x][y] != "#" {
			neighbors = append(neighbors, Node{Position{x, y}, currentNode.currentLevel + 1})
		}
	}
	return neighbors
}

func getNeighborsDiamond(memoryMap [][]string, currentNode Node) []Node {
	neighbors := make([]Node, 0)

	directions := []struct {
		dx, dy int
	}{
		{-2, 0},
		{2, 0},
		{0, -2},
		{0, 2},
	}

	for _, d := range directions {
		x, y := currentNode.pos.X+d.dx, currentNode.pos.Y+d.dy
		if x >= 0 && x < len(memoryMap) && y >= 0 && y < len(memoryMap[0]) && memoryMap[x][y] != "#" {
			neighbors = append(neighbors, Node{Position{x, y}, currentNode.currentLevel + 1})
		}
	}
	return neighbors
}
func manhattanDistance(p1, p2 Position) int {
	return int(math.Abs(float64(p1.X-p2.X)) + math.Abs(float64(p1.Y-p2.Y)))
}
func findNode(maze [][]string, mark string) Node {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == mark {
				return Node{Position{i, j}, 0}
			}
		}
	}
	return Node{}
}
