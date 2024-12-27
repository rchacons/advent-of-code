package main

import (
	"fmt"
	"log"
	"strconv"
//	"strings"

	"github.com/rchacons/advent-of-code/utils"
)


func main() {
	log.Printf("%v days to XMASSS", 24-17)
	input, err := utils.FileToProgramInstructions("input_reverse.txt")

	// input 1 -> 4,6,3,5,6,3,5,2,1,0.
	// input 2 -> B should be 1
	// input 3 -> 0,1,2
	// input 4 -> 4,2,5,6,7,7,7,7,3,1,0 and A should be 0
	// input 5 -> B should be 26
	// input 6 -> B should be 44354
	// input 7 (part two) -> A should be 117440

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	
	registers := map[string]int{"A": input["A"].(int), "B": input["B"].(int), "C": input["C"].(int)}
	program := input["P"].([]int)
	// resultArray := performInstructions(registers, program)
	// result := strings.Join(resultArray, ",")

	result2 := reverseEngineerAndfindLowestA(registers, program)

	//fmt.Println("Part 1 result : ", result)
	fmt.Println("Part 2 result : ", result2)

}

type InstructionFunc func(operand int, registers map[string]int) (output string, jump bool)

func performInstructions(registers map[string]int, program []int) []string {

	instructionFuncs := map[int]InstructionFunc{
        0: adv,
        1: bxl,
        2: bst,
        3: jnz,
        4: bxc,
        5: out,
        6: bdv,
        7: cdv,
    }

	var outputValues []string
	
	ip := 0

	for ip < len(program) {
		opcode := program[ip] 
		operand := program[ip+1]
		ip+=2  // attention point, the jump

		if fn, ok := instructionFuncs[opcode]; ok {
			result, shouldJump := fn(operand, registers)
			if result != ""{ 
				outputValues = append(outputValues, result)
			}
			if shouldJump {
				ip = operand
			}
		}
		// fmt.Println("Registers : ",registers)
		// fmt.Println("Program : ",program)
		// fmt.Printf("Output values = %v",outputValues)
	}
	return outputValues
}

// Stocks value into A
func adv(operand int, registers map[string]int) (string, bool) {
    numerator := registers["A"]
    denominator := (1 << getComboValue(operand, registers))
    registers["A"] = int(numerator / denominator)
    //fmt.Printf("adv: operand=%d, numerator=%d, denominator=%d, registers[A]=%d", operand, numerator, denominator, registers["A"])
    return "", false
}

func bxl(operand int, registers map[string]int) (string, bool) {
    registers["B"] = registers["B"] ^ operand
    //fmt.Printf("bxl: operand=%d, registers[B]=%d", operand, registers["B"])
    return "", false
}

func bst(operand int, registers map[string]int) (string, bool) {
    registers["B"] = getComboValue(operand, registers) % 8
    //fmt.Printf("bst: operand=%d, registers[B]=%d", operand, registers["B"])
    return "", false
}

func jnz(operand int, registers map[string]int) (string, bool) {
    if registers["A"] == 0 {
        //fmt.Printf("jnz: operand=%d, registers[A]=%d, jump=false", operand, registers["A"])
        return "", false
    } else {
        //fmt.Printf("jnz: operand=%d, registers[A]=%d, jump=true", operand, registers["A"])
        return "", true
    }
}

func bxc(operand int, registers map[string]int) (string, bool) {
    registers["B"] = registers["B"] ^ registers["C"]
    //fmt.Printf("bxc: operand=%d, registers[B]=%d, registers[C]=%d", operand, registers["B"], registers["C"])
    return "", false
}

func out(operand int, registers map[string]int) (string, bool) {
    result := getComboValue(operand, registers) % 8
    resultStr := strconv.Itoa(result)
    //fmt.Printf("out: operand=%d, result=%d", operand, result)
    return resultStr, false
}

func bdv(operand int, registers map[string]int) (string, bool) {
    numerator := registers["A"]
    denominator := (1 << getComboValue(operand, registers))
    registers["B"] = int(numerator / denominator)
    //fmt.Printf("bdv: operand=%d, numerator=%d, denominator=%d, registers[B]=%d", operand, numerator, denominator, registers["B"])
    return "", false
}

func cdv(operand int, registers map[string]int) (string, bool) {
    numerator := registers["A"]
    denominator := (1 << getComboValue(operand, registers))
    registers["C"] = int(numerator / denominator)
    //fmt.Printf("cdv: operand=%d, numerator=%d, denominator=%d, registers[C]=%d", operand, numerator, denominator, registers["C"])
    return "", false
}


func getComboValue(operand int, registers map[string]int) int {
	switch operand {
    case 0, 1, 2, 3:
        return operand
    case 4:
        return registers["A"]
    case 5:
        return registers["B"]
    case 6:
        return registers["C"]
    case 7:
        // Operand 7 is reserved and should not appear in valid programs
        panic("Invalid operand: 7 is reserved")
    default:
        panic("Invalid operand")
	}
}


// Part two 
func reverseEngineerAndfindLowestA(registers map[string]int, program []int) int {
	minA := 0
	programToLoop := make([]int, len(program))
	copy(programToLoop, program)

	programToLoop = []int{2, 4, 1, 1, 7, 5, 1, 5, 0, 3, 4, 4, 5, 5, 3, 0}
	//programToLoop = []int{5, 5, 3, 0}
	
	minA, _ = recursiveFind(registers, program, programToLoop, 0)
	return minA
}


func recursiveFind(registers map[string]int, program []int, programToLoop []int, answerUpToThisPoint int) (int, bool) {

    if len(programToLoop) == 0 {
        return answerUpToThisPoint, true
    }

    for possible := 0; possible < 8; possible++ {
        possibleAnswer := (answerUpToThisPoint << 3) + possible
        registers["A"] = possibleAnswer
        result := performInstructions(registers, program)

		subset := getSubset(program, len(result))
        //if registers["B"] == programToLoop[len(programToLoop)-1] {
		if compareArrays(result, subset) {

            val, found := recursiveFind(registers, program, programToLoop[:len(programToLoop)-1], possibleAnswer)
            if !found {
                continue
            }
            return val, true
        }
    }
    return 0, false
}


// compareArrays compares if two arrays, one string and one int, are the same
func compareArrays(strArray []string, intArray []int) bool {

    // Check if the lengths are different
    if len(strArray) != len(intArray) {
        return false
    }

    // Convert intArray to a string array
    convertedArray := make([]string, len(intArray))
    for i, v := range intArray {
        convertedArray[i] = strconv.Itoa(v)
    }

    // Compare the two arrays element by element
    for i := range strArray {
        if strArray[i] != convertedArray[i] {
            return false
        }
    }

    return true
}

func getSubset(arr []int, n int) []int {
    if n > len(arr) {
        return arr
    }
    return arr[len(arr)-n:]
}