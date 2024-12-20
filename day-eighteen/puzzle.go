package main

import (
	"fmt"
	"log"

	//	"strings"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	X,Y int
}

type Node struct {
	pos Position
	currentLevel int
}

func main() {
	log.Printf("%v days to XMASSS", 24-18)
	input, err := utils.FileToTextWithSpaces("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	bytesPos := utils.TextToBytePosition(input)
	
	//memoryMap := initiateMapAndPlaceFallenBytes(7,7, bytesPos, 12)

	// ******** FIRST PART **********
	memoryMap := initiateMapAndPlaceFallenBytes(71,71, bytesPos, 1024)  
	//_, numberOfSteps := getNumberOfStepsToEnd(memoryMap)

	// ******** SECOND PART **********
	firstByte := bruteForceFindBlockingByte(memoryMap, bytesPos[1024:])
	//fmt.Printf("\nNumber of steps %d \n", numberOfSteps)
	fmt.Printf("\nBlocking byte is %d,%d \n", firstByte.Y, firstByte.X)

}

func initiateMapAndPlaceFallenBytes(height int, width int, bytesPos [][] int, numberOfBytes int) [][]string {
	memoryMap := make([][]string,height)
	for i := range height {
		memoryMap[i] = make([]string, width)
		for j := range width {
			memoryMap[i][j] = "."
		}
	}

	for i := range numberOfBytes {
		memoryMap[bytesPos[i][1]][bytesPos[i][0]] = "#"
	}

	return memoryMap
}

func addByteToMap(memoryMap [][]string, bytePos Position) {
	memoryMap[bytePos.X][bytePos.Y] = "#"
}


func getNumberOfStepsToEnd(memoryMap [][]string) (bool, int) {
	start := Node{Position{0,0},0}
	end := Node{Position{len(memoryMap)-1, len(memoryMap[0])-1}, 0}
	
	searchQueue := append([]Node{}, getNeighbors(memoryMap, start)...)

	searchedMap := make(map[Position]bool)
	searchedMap[start.pos] = true

	for len(searchQueue) > 0 {
		var tmpNeighbor Node
		tmpNeighbor, searchQueue = popleft(searchQueue)
		if !searchedMap[tmpNeighbor.pos] {
			if tmpNeighbor.pos == end.pos {
				return true, tmpNeighbor.currentLevel
			} else {
				searchQueue = append(searchQueue, getNeighbors(memoryMap, tmpNeighbor)...)
				searchedMap[tmpNeighbor.pos] = true
			}		
		}
	}

	return false, 0
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
    neighbors := make([]Node,0)

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
				neighbors = append(neighbors, Node{Position{x, y}, currentNode.currentLevel+1})
        }
    }
    return neighbors
}

// Par two
func bruteForceFindBlockingByte(memoryMap [][]string, bytesPos [][]int) Position {
	var blockingPos Position

	for _,byteP := range bytesPos {
		blockingPos := Position{byteP[1], byteP[0]}
		addByteToMap(memoryMap, blockingPos)
		isValid,_ := getNumberOfStepsToEnd(memoryMap)
		if !isValid {
			return blockingPos
		}
	}
	utils.PrintStringMatrix(memoryMap)


	return blockingPos
}