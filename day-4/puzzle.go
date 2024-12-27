package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	xmasMatrix, err := utils.FileToStringMatrix("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Println("hello")
	}	
	total := countXmasOccurences(xmasMatrix)
	fmt.Println(total)

}

func countXmasOccurences(xmasMatrix [][]string) int{
	total := 0

	for row := range xmasMatrix {
		for column := range xmasMatrix[row] {
			if xmasMatrix[row][column] == "X"{
				total += scanForXmasWord(xmasMatrix, row, column)
			}
		}
	}
	return total
}

func scanForXmasWord(xmasMatrix [][]string, currentRow int, currentColumn int) int {
	total := 0
	if scanBackward(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanForward(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanUpward(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanDownward(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanCrosswardUpFwd(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanCrosswardUpBwd(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanCrosswardDownFwd(xmasMatrix, currentRow, currentColumn) { total++ }
	if scanCrosswardDownBwd(xmasMatrix, currentRow, currentColumn) { total++ }

	return total
}

func scanBackward(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentColumn - 3 < 0 {
		return false
	}
	letterCount := 0 // currently at X
	for j := currentColumn; j >= currentColumn - 3; j-- {
		if (!isCorrectXmasWord(letterCount, xmasMatrix[currentRow][j])) {
			return false
		}
		letterCount ++
	}
	return true
}

func scanForward(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentColumn + 3 >= len(xmasMatrix[currentRow]) {
		return false
	}
	
	letterCount := 0 // currently at X
	for j := currentColumn; j <= currentColumn + 3; j++ {
		if !isCorrectXmasWord(letterCount, xmasMatrix[currentRow][j]) {
			return false
		}
		letterCount ++
	}
	return true
}

func scanUpward(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow - 3 < 0 {
		return false
	}
	
	letterCount := 0 // currently at X
	for i := currentRow; i >= currentRow - 3; i--{
		if !isCorrectXmasWord(letterCount, xmasMatrix[i][currentColumn]) {
			return false
		}
		letterCount ++
	}
	return true
}

func scanDownward(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow + 3 >= len(xmasMatrix) {
		return false
	}
	
	letterCount := 0 // currently at X
	for i := currentRow; i <= currentRow + 3; i++{
		if !isCorrectXmasWord(letterCount, xmasMatrix[i][currentColumn]) {
			return false
		}
		letterCount ++
	}
	return true
}

func scanCrosswardUpFwd(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow-3 < 0 || currentColumn + 3 >= len(xmasMatrix[currentRow]) {
		return false
	}
	letterCount := 0 // currently at X
	for range 4 {
		if !isCorrectXmasWord(letterCount, xmasMatrix[currentRow][currentColumn]) {
			return false
		}
		currentRow--
		currentColumn++
		letterCount++
	}

	return true
}

func scanCrosswardUpBwd(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow - 3 < 0 || currentColumn - 3 < 0 {
		return false
	}
	letterCount := 0 // currently at X
	
	for range 4 {

		if !isCorrectXmasWord(letterCount, xmasMatrix[currentRow][currentColumn]) {
			return false
		}
		currentRow--
		currentColumn--
		letterCount++
	}

	return true
}

func scanCrosswardDownFwd(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow+3 >= len(xmasMatrix) || currentColumn + 3 >= len(xmasMatrix[currentRow]) {
		return false
	}

	letterCount := 0 // currently at X	
	for range 4 {
		if !isCorrectXmasWord(letterCount, xmasMatrix[currentRow][currentColumn]) {
			return false
		}
		currentRow++
		currentColumn++
		letterCount++
	}
	return true

 }

func scanCrosswardDownBwd(xmasMatrix [][]string, currentRow int, currentColumn int) bool {
	if currentRow+3 >= len(xmasMatrix) || currentColumn - 3 < 0 {
		return false
	}

	letterCount := 0 // currently at X	
	for range 4 {
		if !isCorrectXmasWord(letterCount, xmasMatrix[currentRow][currentColumn]) {
			return false
		}
		currentRow++
		currentColumn--
		letterCount++
	}
	return true
}

func isCorrectXmasWord(letterCount int, currentLetter string) bool {
	xmasWord := []string{"X", "M", "A", "S"}
	return xmasWord[letterCount] == currentLetter
}
