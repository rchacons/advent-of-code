package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type guard struct {
	direction string
	posX int
	posY int
}


func main() {
	log.Println("%d days to XMASSS", 24-6)

	input, err := utils.FileToStringMatrix("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	inputGuard := identifyGuard(input)
	total := getDistinctPositionsOfGuard(&inputGuard, input)

	fmt.Println(total)

}

func identifyGuard(strMatrix [][]string) guard {
	for i := range strMatrix{
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

func getDistinctPositionsOfGuard(inputGuard *guard, inputMatrix [][]string) int {
	var total int
	
	simulateGuardsPath(inputGuard, inputMatrix)
	
	total += checkForXInMatrix(inputMatrix)

	return total;
}

func simulateGuardsPath(inputGuard *guard, inputMatrix [][]string) {
	for inputGuard.posX < len(inputMatrix) && 
		inputGuard.posX >= 0 &&
		inputGuard.posY >= 0 &&
		inputGuard.posY < len(inputMatrix[inputGuard.posX]) {
		switch inputGuard.direction {
		case "^" : moveUp(inputGuard, inputMatrix)
		case ">": moveRight(inputGuard, inputMatrix)
		case "v": moveDown(inputGuard, inputMatrix)
		case "<": moveLeft(inputGuard, inputMatrix)
		}
	}
}

func moveUp(inputGuard *guard, inputMatrix [][]string) {
	nextPosition := []int{inputGuard.posX-1, inputGuard.posY}
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = ">"
	} else {
		inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		inputGuard.posX--
	}
}

func moveDown(inputGuard *guard, inputMatrix [][]string) {
	nextPosition := []int{inputGuard.posX+1, inputGuard.posY}
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "<"
	} else {
		inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		inputGuard.posX++
	}
}

func moveRight(inputGuard *guard, inputMatrix [][]string) {
	nextPosition := []int{inputGuard.posX, inputGuard.posY+1}
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "v"
	} else {
		inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		inputGuard.posY++
	}
}

func moveLeft(inputGuard *guard, inputMatrix [][]string) {
	nextPosition := []int{inputGuard.posX, inputGuard.posY-1}
	if guardIsBlocked(inputMatrix, *inputGuard, nextPosition) {
		inputGuard.direction = "^"
	} else {
		inputMatrix[inputGuard.posX][inputGuard.posY] = "X"
		inputGuard.posY--
	}
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

func checkForXInMatrix(inputMatrix [][]string) int{
	total := 0
	for i := range inputMatrix{
		for j := range inputMatrix[i] {			
			if inputMatrix[i][j] == "X" {
				total++
			}
			
		}
	}
	return total
}
