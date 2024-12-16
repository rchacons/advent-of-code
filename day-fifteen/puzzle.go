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
	inputMap, inputMouvements, err := utils.FileToMapAndRobotMouvements("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := simulateAndCalculateBoxesGpsCoordinatesSum(inputMap, inputMouvements)

	fmt.Println(total)
}

func simulateAndCalculateBoxesGpsCoordinatesSum(robotMap [][]string, robotMouvements []string) int {
	total := 0
	robotPos := locateRobot(robotMap);
	fmt.Println(robotPos)

	for mouvementIdx := range robotMouvements {
		simulateRobotMovement(&robotMap, &robotPos, robotMouvements[mouvementIdx])
	}

	total += getSumOfBoxesGpsCoordinates(robotMap)
	utils.PrintStringMatrix(robotMap)
	
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

	// fmt.Printf("\nActual Robot Position : %v, moving to : %s \n", robotPos, robotMouvement)
	if (*robotMap)[nextPosition.X][nextPosition.Y] == "#"{
		return
	} else if (*robotMap)[nextPosition.X][nextPosition.Y] == "."{
		objectsToMove = append(objectsToMove, *robotPos)
		canMove = true
	} else {
		objectsToMove = append(objectsToMove, *robotPos)
		canMove = addPossibleBoxRecursively(robotMap, nextPosition, &objectsToMove, robotMouvement)
	}

	if canMove {
		for i:=len(objectsToMove)-1;i>=0;i--{
			nextPos := getNextPositionBasedOnMouvement(objectsToMove[i], robotMouvement)
			tmp := (*robotMap)[objectsToMove[i].X][objectsToMove[i].Y]
			(*robotMap)[objectsToMove[i].X][objectsToMove[i].Y] = "."
			(*robotMap)[nextPos.X][nextPos.Y] = tmp
		}
		*robotPos = nextPosition
	}
	
}

func addPossibleBoxRecursively(robotMap *[][]string, actualBoxPosition Position, objectsToMove *[]Position, mouvement string) bool {
	nextPosition := getNextPositionBasedOnMouvement(actualBoxPosition, mouvement)
	*objectsToMove =  append(*objectsToMove, Position{actualBoxPosition.X, actualBoxPosition.Y})

	if (*robotMap)[nextPosition.X][nextPosition.Y] == "#"{
		return false
	} else if (*robotMap)[nextPosition.X][nextPosition.Y] == "."{
		return true
	} else { // block
		return addPossibleBoxRecursively(robotMap, nextPosition, objectsToMove, mouvement)
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
			if robotMap[i][j] == "O" {
				total += (i*100)+j
			}
		}
	}
	return total
}