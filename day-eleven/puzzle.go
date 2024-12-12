package main

import (
	"fmt"
	"log"
	"strconv"
	"runtime"


	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("%v days to XMASSS", 24-12)

	// Reused report utils that sends [][]int, but we only need [0]
	input, err := utils.FileToReportsLists("input.txt") 

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	var memStatsBefore, memStatsAfter runtime.MemStats

    // Capture memory statistics before execution
    runtime.ReadMemStats(&memStatsBefore)
	
	//total := getTotalStonesAfterXBlinks(input[0], 45) 		//PART ONE
	total := getTotalStonesAfterXBlinksPartTwo(input[0], 75)	//PART TWO

	// Capture memory statistics after execution
	runtime.ReadMemStats(&memStatsAfter)

	// Calculate and log memory usage
	allocBefore := memStatsBefore.Alloc
	allocAfter := memStatsAfter.Alloc
	fmt.Printf("Memory used by getTotalStonesAfterXBlinksPartTwo: %d bytes\n", allocAfter-allocBefore)

	bytes := allocAfter-allocBefore

	mb := float64(bytes) / (1024 * 1024)
	gb := float64(bytes) / (1024 * 1024 * 1024)

	fmt.Printf("Memory usage: %d bytes\n", bytes)
	fmt.Printf("Memory usage: %.2f MB\n", mb)
	fmt.Printf("Memory usage: %.6f GB\n", gb)

	fmt.Println(total)
}

func getTotalStonesAfterXBlinks(stoneList []int, blinks int) int{

	cacheMap := make(map[int][]int) // stores {stoneN : {list of splitted stones}
	for range blinks {
		stoneList = applyStoneRulesRecursive(stoneList, cacheMap) 	//PART ONE
		//stoneList = applyStoneRulesIteratively(stoneList, cacheMap) 	//PART TWO
	}
	return len(stoneList)
}

func applyStoneRulesRecursive(stoneList []int, cacheMap map[int][]int) []int{
	if len(stoneList) == 1 {
		return applyStoneRule(stoneList[0], cacheMap)
	} 
		
	middle := len(stoneList)/2
	return append(applyStoneRulesRecursive(stoneList[:middle], cacheMap), applyStoneRulesRecursive(stoneList[middle:], cacheMap)...)
}

func applyStoneRule(stone int, cacheMap map[int][]int) []int {
	
	if value, exists := cacheMap[stone];exists{
		return value
	}
	
	var finalArray []int
	
	if stone == 0 {
		finalArray = []int{1} 
	} else if isEven, result := verifyEvenDigitsAndSplit(stone); isEven {
		finalArray = result
	} else {
		finalArray = []int{stone*2024}
	}

	cacheMap[stone] = finalArray

	return finalArray
}

func verifyEvenDigitsAndSplit(number int) (bool,[]int) {
	s := strconv.Itoa(number)

	digits := len(s)
	
	if digits % 2 == 0 {
		n1,err := strconv.Atoi(s[:digits/2])
		if err != nil {
			log.Fatal(err)
		}

		s2 := s[digits/2:]
		n2,err := strconv.Atoi(s2)
		if err != nil {
			log.Fatal(err)
		}

		return true, []int{n1,n2}
	}
	return false, nil
}

// Part two

func getTotalStonesAfterXBlinksPartTwo(stoneList []int, blinks int) int {

    cacheMap := make(map[int][]int) // stores {stoneN : {list of splitted stones}
    total := 0
    tempStoneMap := make(map[int][]int)
    for _, stone := range stoneList {
		clear(tempStoneMap)
        tempStoneMap[stone] = nil

        for i := 0; i < blinks; i++ {
            applyStoneRulesIteratively(&tempStoneMap, cacheMap)
        }
		
		total += countMap(tempStoneMap)
    }

    return total
}


func applyStoneRulesIteratively(stoneMap *map[int][]int, cacheMap map[int][]int) {

    copyStoneMap := make(map[int]int)
    for stone, ocur := range *stoneMap {

        newStones := applyStoneRule(stone, cacheMap)

		count := 1
		if len(ocur) > 0 {
			count = ocur[0]
		}

		for _, newStone := range newStones{
			copyStoneMap[newStone] += count
		}
    }

    clear(*stoneMap)

	for stone, count := range copyStoneMap{
		(*stoneMap)[stone] = []int{count}
	}


}

func countMap(stoneMap map[int][]int) int{
	total := 0
	for _, value := range stoneMap {
		total += value[0]
	}
	return total
}
