package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	X, Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-15)
	inputMap, inputMouvements, err := utils.FileToMapAndRobotMouvements("../input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	newRobotMap := transformMap(inputMap)
	total := simulateAndCalculateBoxesGpsCoordinatesSum(newRobotMap, inputMouvements)

	fmt.Println(total)

}

func transformMap(inputMap [][]string) [][]string{
	var newRobotMap [][]string
	for i := range inputMap{
		newLine := make([]string, 0)
		for j := range inputMap[i] {
			val := inputMap[i][j]
			if val == "#" || val == "." {
				newLine = append(newLine, val)
				newLine = append(newLine, val)
			} else if val == "O" {
				newLine = append(newLine, "[")
				newLine = append(newLine, "]")
			} else if val == "@"{
				newLine = append(newLine, "@")
				newLine = append(newLine, ".")
			}
		}
		newRobotMap = append(newRobotMap, newLine)
	}

	return newRobotMap
}

func simulateAndCalculateBoxesGpsCoordinatesSum(robotMap [][]string, robotMouvements []string) int {
	total := 0
	robotPos := locateRobot(robotMap);

	for mouvementIdx := range robotMouvements {
		simulateRobotMovement(&robotMap, &robotPos, robotMouvements[mouvementIdx])
	}

	utils.PrintStringMatrix(robotMap)
	total += getSumOfBoxesGpsCoordinates(robotMap)
	
	return total
}

func locateRobot(robotMap [][]string) Position{
	for i := range robotMap {
		for j := range robotMap[i] {
			if robotMap[i][j] == "@" {
				return Position{i,j}
			}
		}
	}
	return Position{}
}

func simulateRobotMovement(robotMap *[][]string, robotPos *Position, robotMouvement string) {
	objectsToMove := make([]Position, 0)
	nextPosition := getNextPositionBasedOnMouvement(*robotPos, robotMouvement)
	canMove := false


	alreadyMarked := make(map[Position]bool)
	// fmt.Printf("\nActual Robot Position : %v, moving to : %s \n", robotPos, robotMouvement)
	if (*robotMap)[nextPosition.X][nextPosition.Y] == "#"{
		return
	} else if (*robotMap)[nextPosition.X][nextPosition.Y] == "."{
		objectsToMove = append(objectsToMove, *robotPos)
		canMove = true
	} else {
		objectsToMove = append(objectsToMove, *robotPos)
		if robotMouvement == "^" || robotMouvement == "v"{
			if (*robotMap)[nextPosition.X][nextPosition.Y] == "[" {
				objectsToMove = append(objectsToMove, nextPosition)
				objectsToMove = append(objectsToMove, Position{nextPosition.X, nextPosition.Y+1})
				canMove = addPossibleBoxRecursively(robotMap, nextPosition, &objectsToMove, robotMouvement, alreadyMarked) && 
						addPossibleBoxRecursively(robotMap, Position{nextPosition.X, nextPosition.Y+1}, &objectsToMove, robotMouvement, alreadyMarked)
			} else {
				objectsToMove = append(objectsToMove, nextPosition)
				objectsToMove = append(objectsToMove, Position{nextPosition.X, nextPosition.Y-1})
				canMove = addPossibleBoxRecursively(robotMap, nextPosition, &objectsToMove, robotMouvement, alreadyMarked) && 
						addPossibleBoxRecursively(robotMap, Position{nextPosition.X, nextPosition.Y-1}, &objectsToMove, robotMouvement, alreadyMarked)
			}
		} else {
			objectsToMove = append(objectsToMove, nextPosition)
			canMove = addPossibleBoxRecursively(robotMap, nextPosition, &objectsToMove, robotMouvement, alreadyMarked)
		}
	}


	if canMove {
        matrixCopy := copyMatrix(*robotMap)

		// first mark them
        for i := len(objectsToMove) - 1; i >= 0; i-- {
            (*robotMap)[objectsToMove[i].X][objectsToMove[i].Y] = "."
        }

		// them move them
		for i := len(objectsToMove) - 1; i >= 0; i-- {
            nextPos := getNextPositionBasedOnMouvement(objectsToMove[i], robotMouvement)
            tmp1 := matrixCopy[objectsToMove[i].X][objectsToMove[i].Y]
            (*robotMap)[nextPos.X][nextPos.Y] = tmp1

        }

        *robotPos = nextPosition
    } 
	
}

func copyMatrix(matrix [][]string) [][]string {
    copy := make([][]string, len(matrix))
    for i := range matrix {
        copy[i] = make([]string, len(matrix[i]))
        for j := range matrix[i] {
            copy[i][j] = matrix[i][j]
        }
    }
    return copy
}

func addPossibleBoxRecursively(robotMap *[][]string, actualBoxPosition Position, objectsToMove *[]Position, mouvement string, alreadyMarked map[Position]bool) bool {
	nextPosition := getNextPositionBasedOnMouvement(actualBoxPosition, mouvement)
	
	if alreadyMarked[actualBoxPosition] {
		return true
	}
	
	if (*robotMap)[nextPosition.X][nextPosition.Y] == "#"{
		return false
	} else if (*robotMap)[nextPosition.X][nextPosition.Y] == "."{
		alreadyMarked[nextPosition] = true
		return true
	} else if mouvement == "<" || mouvement == ">" { // block
		*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y})
		return addPossibleBoxRecursively(robotMap, nextPosition, objectsToMove, mouvement, alreadyMarked)
	} else if mouvement == "^" || mouvement == "v"{
		if (*robotMap)[nextPosition.X][nextPosition.Y] == "[" {
			*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y})
			*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y+1})

			return 	addPossibleBoxRecursively(robotMap, nextPosition, objectsToMove, mouvement, alreadyMarked) && 
					addPossibleBoxRecursively(robotMap, Position{nextPosition.X, nextPosition.Y+1}, objectsToMove, mouvement, alreadyMarked)
		} else {
			*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y})
			*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y-1})

			return 	addPossibleBoxRecursively(robotMap, nextPosition, objectsToMove, mouvement, alreadyMarked) && 
					addPossibleBoxRecursively(robotMap, Position{nextPosition.X, nextPosition.Y-1}, objectsToMove, mouvement, alreadyMarked)
		} 
	} else {
		*objectsToMove =  append(*objectsToMove, Position{nextPosition.X, nextPosition.Y})
		return addPossibleBoxRecursively(robotMap, nextPosition, objectsToMove, mouvement, alreadyMarked)
	}
}

func getNextPositionBasedOnMouvement(actualPos Position, mouvement string) Position {
	switch mouvement {
	case "<" :
		return Position{actualPos.X, actualPos.Y-1}
	case ">" : 
		return Position{actualPos.X, actualPos.Y+1}
	case "^" : 
		return Position{actualPos.X-1, actualPos.Y}
	case "v" : 
		return Position{actualPos.X+1, actualPos.Y}
	default :
		return Position{}
	}
}

func getSumOfBoxesGpsCoordinates(robotMap [][]string) int {
	total := 0

	for i := range robotMap {
		for j := range robotMap[i] {
			if robotMap[i][j] == "[" {
				total += (i*100)+j
			}
		}
	}
	return total
}