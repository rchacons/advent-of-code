package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("XMAS is over but we still have some work to do...")

	input, err := utils.FileToMapOfLocksAndKeysMatrix("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Part 1
	total := getPossibleCombinations(input)
	fmt.Println("Total part 1: ", total)
}

func getPossibleCombinations(locksAndKeysMap map[string][][][]string) int {
    total := 0
    heightMap := computeHeights(locksAndKeysMap)

    for _, keyHeight := range heightMap["key"] {
        valid := true
        for _, lockHeight := range heightMap["lock"] {
            for i := range keyHeight {
                if keyHeight[i]+lockHeight[i] > 7 {
                    valid = false
                }
            }
            if valid {
                total++
            }
			valid = true
        }

    }

    return total
}

// Compute the heights of the keys
func computeHeights(locksAndKeysMap map[string][][][]string) map[string][][]int {
	heights := make(map[string][][]int)

	for _, lockMatrix := range locksAndKeysMap["lock"] {
		lockHeights := []int{}
		for col := range lockMatrix[0] {
			colHeight := 0
			for row := range lockMatrix {
				if lockMatrix[row][col] == "#" {
					colHeight++
				} else {
					break
				}
			}
			lockHeights = append(lockHeights, colHeight)
		}
		heights["lock"] = append(heights["lock"], lockHeights)
	}

	for _, keyMatrix := range locksAndKeysMap["key"] {
		keyHeights := []int{}
		for col := range keyMatrix[len(keyMatrix)-1] {
			colHeight := 0
			for i := len(keyMatrix) - 1; i >= 0; i-- {
				if keyMatrix[i][col] == "#" {
					colHeight++
				} else {
					break
				}
			}
			keyHeights = append(keyHeights, colHeight)
		}
		heights["key"] = append(heights["key"], keyHeights)
	}

	return heights
}