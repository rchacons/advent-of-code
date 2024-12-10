package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type guard struct {
	direction string
	posX      int
	posY      int
}

type State struct {
	direction string
	visited bool
}

type Position struct {
	X,Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-6)

	input, err := utils.FileToStringMatrix("../input.txt")
	// input_test -> expected 6
	// input two -> expected 1
	// input three -> expected 1
	// input four -> expected 0
	// input five -> expected 4
	// input six -> expected 0 
	// input seven -> expected 1
	// input eight -> expected 1
	// input nine -> expected 1

	// The block should be only on places not visited by the X
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	inputGuard := identifyGuard(input)
	total := getPossibleObstacles(&inputGuard, input)

	fmt.Println(total)

}

func identifyGuard(strMatrix [][]string) guard {
	for i := range strMatrix {
		for j := range strMatrix[i] {
			val := strMatrix[i][j]
			if val == "^" || val == ">" ||
				val == "v" || val == "<" {
				return guard{direction: val, posX: i, posY: j}
			}
		}
	}
	return guard{}
}

func simulateGuardsPath(inputGuard *guard, inputMatrix [][]string, shouldMarkX bool) bool{
	stateMap := make(map[Position]State)


    for inputGuard.posX < len(inputMatrix) &&
        inputGuard.posX >= 0 &&
        inputGuard.posY >= 0 &&
        inputGuard.posY < len(inputMatrix[inputGuard.posX]) {
			currentPos := Position{inputGuard.posX, inputGuard.posY}

			if state, exists := stateMap[currentPos]; exists {
				fmt.Printf("State exists for position (%d, %d): direction=%s, visited=%t\n", currentPos.X, currentPos.Y, state.direction, state.visited)
				if state.direction == inputGuard.direction && state.visited {
					fmt.Println("Detected a loop. Exiting simulation.")
					return true
				}
			} else {
				stateMap[currentPos] = State{
					direction: inputGuard.direction,
					visited:   true,
				}
			}

			switch inputGuard.direction {
				case "^":
					fmt.Println("Moving up")
					moveUp(inputGuard, inputMatrix, shouldMarkX)
				case ">":
					fmt.Println("Moving right")
					moveRight(inputGuard, inputMatrix, shouldMarkX)
				case "v":
					fmt.Println("Moving down")
					moveDown(inputGuard, inputMatrix, shouldMarkX)
				case "<":
					fmt.Println("Moving left")
					moveLeft(inputGuard, inputMatrix, shouldMarkX)
			}
    }
	return false
}

func moveUp(inputGuard *guard, inputMatrix [][]string, shouldMarkX bool) int {
	nextPosition := []int{inputGuard.posX - 1, inputGuard.posY}
	turnCount := 0
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = ">"
	} else {
		if shouldMarkX {
			inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		}
		inputGuard.posX--
	}
	return turnCount
}

func moveDown(inputGuard *guard, inputMatrix [][]string, shouldMarkX bool) int {
	nextPosition := []int{inputGuard.posX + 1, inputGuard.posY}
	turnCount := 0

	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "<"
	} else {
		if shouldMarkX {
			inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		}
		inputGuard.posX++
	}

	return turnCount
}
func moveRight(inputGuard *guard, inputMatrix [][]string, shouldMarkX bool) int {
	nextPosition := []int{inputGuard.posX, inputGuard.posY + 1}
	turnCount := 0
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "v"
	} else {
		if shouldMarkX {
			inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		}
		inputGuard.posY++
	}
	return turnCount
}

func moveLeft(inputGuard *guard, inputMatrix [][]string, shouldMarkX bool) int{
	nextPosition := []int{inputGuard.posX, inputGuard.posY - 1}
	turnCount := 0
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "^"
	} else {
		if shouldMarkX {
			inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		}
		inputGuard.posY--
	}
	return turnCount
}

func guardIsBlocked(inputMatrix [][]string, inputGuard guard, nextPosition []int) bool {

	if !(nextPosition[0] < len(inputMatrix) && nextPosition[0] >= 0) {
		return false
	} else if !(nextPosition[1] < len(inputMatrix[inputGuard.posX]) && nextPosition[1] >= 0) {
		return false
	} else if inputMatrix[nextPosition[0]][nextPosition[1]] == "#" {
		return true
	} else {
		return false
	}
}


// PART TWO METHODS

func getPossibleObstacles(inputGuard *guard, inputMatrix [][]string) int {
	backupGuard := guard{inputGuard.direction, inputGuard.posX, inputGuard.posY}

	fmt.Printf("\n\n********************STEP 1********************\n\n")
	simulateGuardsPath(inputGuard, inputMatrix, true)
	
	return calculateObstacles(&backupGuard, inputMatrix)
}

func calculateObstacles(inputGuard *guard, inputMatrix [][]string) int {
	var obstacles int
	visited := fillVisitedMatrix(inputMatrix)

	fmt.Printf("\n\n********************STEP 2********************\n\n")
	
	for inputGuard.posX < len(inputMatrix) &&
	inputGuard.posX >= 0 &&
	inputGuard.posY >= 0 &&
	inputGuard.posY < len(inputMatrix[inputGuard.posX]) {
		switch inputGuard.direction {
		case "^":
			obstacles += moveUpX(inputGuard, inputMatrix, visited)
		case ">":
			obstacles += moveRightX(inputGuard, inputMatrix, visited)
		case "v":
			obstacles += moveDownX(inputGuard, inputMatrix, visited)
		case "<":
			obstacles += moveLeftX(inputGuard, inputMatrix, visited)
		}
	}
	return obstacles
}

func moveUpX(inputGuard *guard, inputMatrix [][]string, visitedMatrix [][]bool) int{
	validObstacle := 0
	nextPosition := []int{inputGuard.posX - 1, inputGuard.posY}

	if nextIsX(inputMatrix, *inputGuard, nextPosition) {
		if !isAlreadyVisited(visitedMatrix, nextPosition) {
			validObstacle += blockAndSimulatePath(inputMatrix, *inputGuard, nextPosition)
		}
		visitedMatrix[inputGuard.posX][inputGuard.posY] = true
		inputGuard.posX--
	} else if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = ">"
	} else {
		inputGuard.posX--
	}
	return validObstacle
}

func moveDownX(inputGuard *guard, inputMatrix [][]string, visitedMatrix [][]bool) int {
    validObstacle := 0
    nextPosition := []int{inputGuard.posX + 1, inputGuard.posY}

    if nextIsX(inputMatrix, *inputGuard, nextPosition) {
		if !isAlreadyVisited(visitedMatrix, nextPosition) {
			validObstacle += blockAndSimulatePath(inputMatrix, *inputGuard, nextPosition)
		}
		visitedMatrix[inputGuard.posX][inputGuard.posY] = true
		inputGuard.posX++
    } else if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
        inputGuard.direction = "<"
    } else {
		inputGuard.posX++    
	}

    return validObstacle
}

func moveRightX(inputGuard *guard, inputMatrix [][]string, visitedMatrix [][]bool) int{
	validObstacle := 0
	nextPosition := []int{inputGuard.posX, inputGuard.posY + 1}

	if nextIsX(inputMatrix, *inputGuard, nextPosition){
		if !isAlreadyVisited(visitedMatrix, nextPosition) {
			validObstacle += blockAndSimulatePath(inputMatrix, *inputGuard, nextPosition)
			}		
		visitedMatrix[inputGuard.posX][inputGuard.posY] = true
		inputGuard.posY++
	} else if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "v"
	} else {
		inputGuard.posY++
	}
	return validObstacle
}

func moveLeftX(inputGuard *guard, inputMatrix [][]string, visitedMatrix [][]bool) int{
	validObstacle := 0
	nextPosition := []int{inputGuard.posX, inputGuard.posY - 1}

	if nextIsX(inputMatrix, *inputGuard, nextPosition){
		if !isAlreadyVisited(visitedMatrix, nextPosition) {
			validObstacle += blockAndSimulatePath(inputMatrix, *inputGuard, nextPosition)
		} 
		visitedMatrix[inputGuard.posX][inputGuard.posY] = true
		inputGuard.posY--
	} else if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "^"
	} else {
		inputGuard.posY--
	}
	return validObstacle
}

func blockAndSimulatePath(inputMatrix [][]string, inputGuard guard, nextPosition []int) int {
	validObstacle := 0
	backupValue := blockPath(inputMatrix, nextPosition)
	simulationGuard := guard{inputGuard.direction, inputGuard.posX, inputGuard.posY}
	
	if isPathALoop(inputMatrix, simulationGuard) {
		validObstacle++
	}

	unblockPath(inputMatrix, nextPosition, backupValue)
	return validObstacle
}
func blockPath(inputMatrix [][]string, nextPosition []int) string{
	backupValue := inputMatrix[nextPosition[0]][nextPosition[1]]
	inputMatrix[nextPosition[0]][nextPosition[1]] = "#"
	return backupValue
}

func unblockPath(inputMatrix [][]string, nextPosition []int, value string){
	inputMatrix[nextPosition[0]][nextPosition[1]] = value
}


func nextIsX(inputMatrix [][]string, inputGuard guard, nextPosition []int) bool {
	return nextPosition[0] < len(inputMatrix) && nextPosition[0] >= 0 && 
	nextPosition[1] < len(inputMatrix[inputGuard.posX]) && nextPosition[1] >= 0 &&
	inputMatrix[nextPosition[0]][nextPosition[1]] == "X"
}

func isAlreadyVisited(booleanMatrix [][]bool, nextPosition []int) bool {
	return booleanMatrix[nextPosition[0]][nextPosition[1]]
}

func isPathALoop(inputMatrix [][]string, inputGuard guard) bool {
	fmt.Printf("Starting loop detection for guard at initial position (%d, %d) with direction %s\n", inputGuard.posX, inputGuard.posY, inputGuard.direction)
	for inputGuard.posX < len(inputMatrix) &&
		inputGuard.posX >= 0 &&
		inputGuard.posY >= 0 &&
		inputGuard.posY < len(inputMatrix[inputGuard.posX]) {

		return simulateGuardsPath(&inputGuard, inputMatrix, false)
	}

	fmt.Println("Guard did not return to the start position, path is not a loop")
	return false
}


func fillVisitedMatrix(inputMatrix [][]string) [][]bool {
	visitedMatrix := make([][]bool, len(inputMatrix))
	for i := range inputMatrix {
		visitedMatrix[i] = make([]bool, len(inputMatrix[i]))
		for j := range inputMatrix[i] {
			visitedMatrix[i][j] = false
		}
	}
	return visitedMatrix
}
