package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("XMAS is over but we still have some work to do...")
	input, err := utils.FileToIntegerList("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// ******** FIRST PART **********
	nbChanges := 2000
	
	totalPart1 := calculateSecretNumber(input, nbChanges)

	// ******** SECOND PART **********
	totalPart2 := figureOutBestSequenceAndBananas(input, nbChanges)

	fmt.Printf("\nTotal part one is : %d\n", totalPart1)
	fmt.Printf("\nTotal part two is : %d\n", totalPart2)
}

func calculateSecretNumber(input []int, nb int) int {

	total := 0

	for _, number := range input {

		for range nb {
			number = computeSecretNumber(number)
		}

		total += number
	}
	return total
}


// Part two
type Sequence struct {
	n1,n2,n3,n4 int
}
func figureOutBestSequenceAndBananas(numbers []int, nb int) int {
	sequences := make(map[Sequence]int)

	for _, number := range numbers {
		sequence := make([]int, 4)

		currentLastDigit := number % 10
		seen := make(map[Sequence]bool)
		for index := range nb-1 {
			
			number = computeSecretNumber(number)
	
			// Extract last digit
			lastDigit := number % 10
			diff := lastDigit - currentLastDigit
			currentLastDigit = lastDigit
	
			if index < 3 {
				sequence = append(sequence, diff)
			} else {
				sequence = sequence[len(sequence)-3:]
				sequence = append(sequence, diff)
				
				if seen[Sequence{sequence[0], sequence[1], sequence[2], sequence[3]}] { continue }
				seen[Sequence{sequence[0], sequence[1], sequence[2], sequence[3]}] = true
				sequences[Sequence{sequence[0], sequence[1], sequence[2], sequence[3]}] += currentLastDigit
			
			}
		}
	}

	maxValue := 0
	for _, value := range sequences {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue

}

func computeSecretNumber(number int) int {
	// 1st step
	subResult1 := number * 64
	mixedResult := subResult1 ^ number
	prunedResult := mixedResult % 16777216

	// 2nd step
	subResult2 := prunedResult / 32
	mixedResult = subResult2 ^ prunedResult
	prunedResult = mixedResult % 16777216

	// 3rd step
	subResult3 := prunedResult * 2048
	mixedResult = subResult3 ^ prunedResult
	prunedResult = mixedResult % 16777216

	return prunedResult
}