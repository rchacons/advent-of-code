package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	xmasMatrix, err := utils.FileToStringMatrix("../input.txt")
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
			if xmasMatrix[row][column] == "A"{
				total += scanForXmasCross(xmasMatrix, row, column)
			}
		}
	}
	return total
}

func scanForXmasCross(xmasMatrix [][]string, currentRow int, currentColumn int) int {
	total := 0
	
	if scanCross(xmasMatrix, currentRow, currentColumn, true) && 
	scanCross(xmasMatrix, currentRow, currentColumn, false) { 
		total++ 
	}

	return total
}

func scanCross(xmasMatrix [][]string, currentRow int, currentColumn int, scanRight bool) bool {
	uncheckedXmas := map[string]bool{
        "M": true,
        "A": false, // By default
        "S": true,
    }

    // Check boundary conditions
    if currentRow+1 >= len(xmasMatrix) || currentRow-1 < 0 ||
        currentColumn+1 >= len(xmasMatrix[currentRow]) || currentColumn-1 < 0 {
        return false
    }

    // Determine the upper and lower letters based on scan direction
    var upperLetter, lowerLetter string
    if scanRight {
        upperLetter = xmasMatrix[currentRow-1][currentColumn+1]
        lowerLetter = xmasMatrix[currentRow+1][currentColumn-1]
    } else {
        upperLetter = xmasMatrix[currentRow-1][currentColumn-1]
        lowerLetter = xmasMatrix[currentRow+1][currentColumn+1]
    }

	// Check and update the map for upper and lower letters
    if !checkAndUpdate(uncheckedXmas, upperLetter) || !checkAndUpdate(uncheckedXmas, lowerLetter) {
        return false
    }
	
	return true
}

// Helper function to check and update the map
func checkAndUpdate(uncheckedXmas map[string]bool, letter string) bool {
    if uncheckedXmas[letter] {
        uncheckedXmas[letter] = false // Mark as checked
        return true
    }
    return false
}