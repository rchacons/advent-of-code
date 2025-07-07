package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rchacons/advent-of-code/utils"
)

func main(){
	log.Printf("%v days to XMASSS", 24-10)

	input, err := utils.FileToIntegerLists("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalCalibrationResults(input)

	fmt.Println(total)
}

func getTotalCalibrationResults(numberList [][]int) int {
	total := 0
	for _, numbers := range numberList {
		evaluatedNumbersMap := evaluateCalibrationEquation (numbers[1:])
		if _, exists := evaluatedNumbersMap[numbers[0]]; exists {
			total += numbers[0]
		}
	}
	return total
}

func evaluateCalibrationEquation(numbers []int) map[int]string{
	if len(numbers) == 1 {
		return map[int]string{numbers[0]: ""}
	}

	// We account the possibilities of all left part of array
	possibilities := evaluateCalibrationEquation(numbers[:len(numbers)-1])

	newPosibilitiesMap := make(map[int]string)
	for key, _ := range possibilities {
		sumNum := key + numbers[len(numbers)-1]
		mulNum := key * numbers[len(numbers)-1]
		joinedNumStr := strconv.Itoa(key) + strconv.Itoa(numbers[len(numbers)-1])
		joinedNum, err := strconv.Atoi(joinedNumStr)
		if err != nil {
			log.Fatalf("Error converting number : %v", err)
		}
		newPosibilitiesMap[sumNum] = "+"
		newPosibilitiesMap[mulNum] = "*"
		newPosibilitiesMap[joinedNum] = "||"
	}

	return newPosibilitiesMap
}