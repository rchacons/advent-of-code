package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("%v days to XMASSS", 24-21)
	patterns, wordList, err := utils.FileToPatternSlices("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	patternMap := patternToMap(patterns)

	// Part one
	//total := getTotalPossibleDesigns(patternMap, wordList)
	totalPartTwo := getTotalPossibleDesignsPartTwo(patternMap, wordList)

	//fmt.Println("Total part 1 : ", total)
	fmt.Println("Total part 2 : ", totalPartTwo)

}

func getTotalPossibleDesigns(patterns map[string]bool, wordList []string) int {
	total := 0
	cacheMap := make(map[string]bool)
	for _, wordNeeded := range wordList {
		if checkDesign(patterns, wordNeeded, cacheMap) {
			total++
		}
	}
	return total
}


func checkDesign(patterns map[string]bool, design string, cache map[string]bool) bool {

	if design == "" {
		return true
	}
	if val, found := cache[design]; found {
        return val
    }
		
	for i := 1; i<= len(design); i++{

		if patterns[design[:i]] && checkDesign(patterns, design[i:], cache) {
			cache[design] = true
			return true
		}
	}
	cache[design] = false
	return false
}


func patternToMap(patterns []string) map[string]bool {
	patternMap := make(map[string]bool)
	for _, pattern := range patterns {
		patternMap[pattern] = true
	}
	return patternMap
}

// Part two 
func getTotalPossibleDesignsPartTwo(patterns map[string]bool, wordList []string) int {
	total := 0
	cacheMap := make(map[string]int)
	for _, wordNeeded := range wordList {
		total += checkDesignPartTwo(patterns, wordNeeded, cacheMap)
	}
	return total
}


func checkDesignPartTwo(patterns map[string]bool, design string, cache map[string]int) int {

	if design == "" {return 1}
	if val, found := cache[design]; found {
        return val
    }
		
	count := 0
	for i := 1; i<= len(design); i++{
		if patterns[design[:i]] {
			count += checkDesignPartTwo(patterns, design[i:], cache) 
			cache[design]= count
		}
	}
	return count
}
