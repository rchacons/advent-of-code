package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Println("%d days to XMASSS", 24-5)

	input, err := utils.FileToTextWithSpaces("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	rules := utils.TextToRules(input)

	ordering := utils.TextToListOfNumbers(input)

	total := getTotalValidMiddles(rules, ordering)

	fmt.Println("Total is: ", total)
}

func getTotalValidMiddles(rules [][]int, ordering [][]int) int {
	total := 0
	rulesMap := saveRulesToMap(rules)

	for _, orderLine := range ordering {
		//total += getMiddleIfValid(rulesMap, orderLine) 		// PART ONE
		total += getMiddleOfSortedInvalid(rulesMap, orderLine) 	// PART TWO
	}

	return total
}

func saveRulesToMap(rules [][]int) map[int]map[int]bool {
	m := make(map[int]map[int]bool)

	for _, ruleNumbers := range rules {
		if v, exists := m[ruleNumbers[0]]; exists {
			v[ruleNumbers[1]] = true
		} else {
			m[ruleNumbers[0]] = map[int]bool{ruleNumbers[1]: true}
		}
	}
	return m
}

func getMiddleIfValid(rulesMap map[int]map[int]bool, orderLine []int) int {
	if isValidOrder(rulesMap, orderLine) {
		middle := orderLine[len(orderLine)/2]
		return middle
	}
	return 0
}

func isValidOrder(rulesMap map[int]map[int]bool, orderLine []int) bool {
	if len(orderLine) < 2 {
		return true
	} else if len(orderLine) == 2 {
		valid := listInMapContains(rulesMap, orderLine[0], orderLine[1])
		return valid
	}

	// The error was in here. I mistakenly splitted it, and by saving the pivotindex it was better
	pivotIndex := len(orderLine)/2
	pivot := orderLine[pivotIndex]

	left := orderLine[:pivotIndex]
	right := orderLine[pivotIndex+1:]

    return isValidOrder(rulesMap, left) &&
        listInMapContains(rulesMap, left[len(left)-1], pivot) &&
        isValidOrder(rulesMap, right) &&
        listInMapContains(rulesMap, pivot, right[0])
}

func listInMapContains(rulesMap map[int]map[int]bool, key int, value int) bool {
	if v, exists := rulesMap[key]; exists {
		contains := v[value]
		return contains
	} else {
		return false
	}
}

// Lets begin again as it didnt work lol -> let's try brutforcing.. 
// Update, with brutforcing it works, but i found out the problem on my algorithm, i solved it and it was ok
func getTotalValidMiddlesTwo(rules [][]int, ordering [][]int) int {
	total := 0
	rulesMap := saveRulesToMap(rules)

	for _, orderLine := range ordering {
		total += checkIfValidBrutForce(rulesMap, orderLine)
	}

	return total
}

func checkIfValidBrutForce(rulesMap map[int]map[int]bool, orderLine []int) int {

	for i := 1; i < len(orderLine); i++ {
		num1 := orderLine[i-1]
		num2 := orderLine[i]
		if !listInMapContains(rulesMap, num1, num2) {
			return 0
		}
	}
	return orderLine[len(orderLine)/2]
}


// PART TWO : Check the incorrect file
func getMiddleOfSortedInvalid(rulesMap map[int]map[int]bool, orderLine []int) int {
	if !isValidOrder(rulesMap, orderLine) {
		newSorted := sortInvalidList(rulesMap, orderLine)
		middle := newSorted[len(newSorted)/2]
		return middle
	}
	return 0
}

// Using quicksort
func sortInvalidList(rulesMap map[int]map[int]bool, orderLine []int) []int {
	if len(orderLine) < 2 {
		return orderLine
	} 

	pivot := orderLine[0]
	var left []int
	for _, num := range orderLine[1:] {
		if listInMapContains(rulesMap, num, pivot){
			left = append(left, num)
		}
	}

	var right []int
	for _, num := range orderLine[1:] {
		if listInMapContains(rulesMap, pivot, num){
			right = append(right, num)
		}
	}
	pivotAsList := []int{pivot}
	finalList := append(append(sortInvalidList(rulesMap, left), pivotAsList...), sortInvalidList(rulesMap, right)...)
    return finalList
}
