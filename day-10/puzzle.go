package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	X,Y int
}


func main() {
	log.Printf("%v days to XMASSS", 24-12)

	input, err := utils.FileToIntMatrix("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalTrailheadsScore(input)

	fmt.Println(total)
}

func getTotalTrailheadsScore(numberMatrix [][]int) int{
	total := 0

	for i:=0; i < len(numberMatrix) ; i++{
		for j:=0; j < len(numberMatrix[i]); j++{
			if numberMatrix[i][j] == 0{
				// Trailhead
				newMap := make(map[Position]bool)
				//total += findTotalScore(numberMatrix, Position{i,j}, newMap)	// Part one
				total += findTotalRating(numberMatrix, Position{i,j}, newMap) 	// Part two

			}
		}
	}
	return total
}

func findTotalScore(numberMatrix [][]int, currentNumberPosition Position, posMap map[Position]bool) int{
	total := 0
	currentNumber := numberMatrix[currentNumberPosition.X][currentNumberPosition.Y]
	nextNumber := currentNumber+1

	if currentNumber == 9 {
		if alreadyChecked, exists := posMap[currentNumberPosition]; exists{
			if alreadyChecked {
				return 0
			} else {
				return 1
			}
		}
		posMap[currentNumberPosition] = true
		return 1
	}
	positions := make([]Position, 0, 4)
	positions = append(positions, Position{currentNumberPosition.X-1, currentNumberPosition.Y}) // Up
	positions = append(positions, Position{currentNumberPosition.X+1, currentNumberPosition.Y}) // Down
	positions = append(positions, Position{currentNumberPosition.X, currentNumberPosition.Y-1}) // Left
	positions = append(positions, Position{currentNumberPosition.X, currentNumberPosition.Y+1}) // Right

	for _, pos := range positions {
		if isPositionInMatrixAndFollowsCurrent(numberMatrix, pos, nextNumber){ 
			total += findTotalScore(numberMatrix, pos, posMap)
		}
	}
	return total
}

func isPositionInMatrixAndFollowsCurrent(numberMatrix [][]int, pos Position, nextNumber int) bool{
	return pos.X < len(numberMatrix) && 
			pos.X >= 0 && 
			pos.Y >= 0 && 
			pos.Y < len(numberMatrix[pos.X]) &&
			numberMatrix[pos.X][pos.Y] == nextNumber
}

// Second part
func findTotalRating(numberMatrix [][]int, currentNumberPosition Position, posMap map[Position]bool) int{
	total := 0
	currentNumber := numberMatrix[currentNumberPosition.X][currentNumberPosition.Y]
	nextNumber := currentNumber+1

	if currentNumber == 9 {
		return 1
	}

	positions := make([]Position, 0, 4)
	positions = append(positions, Position{currentNumberPosition.X-1, currentNumberPosition.Y}) // Up
	positions = append(positions, Position{currentNumberPosition.X+1, currentNumberPosition.Y}) // Down
	positions = append(positions, Position{currentNumberPosition.X, currentNumberPosition.Y-1}) // Left
	positions = append(positions, Position{currentNumberPosition.X, currentNumberPosition.Y+1}) // Right

	for _, pos := range positions {
		if isPositionInMatrixAndFollowsCurrent(numberMatrix, pos, nextNumber){ 
			total += findTotalRating(numberMatrix, pos, posMap)
		}
	}
	return total
}
