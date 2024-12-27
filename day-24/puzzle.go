package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("XMAS is over but we still have some work to do...")

	wiresMap, gatesAndWiresMap, err := utils.FileToFormulasMaps("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Part 1
	total := getZWiresValue(wiresMap, gatesAndWiresMap)
	fmt.Println("Total part 1 : ", total)

	// Part two
	getSwappedWires(wiresMap, gatesAndWiresMap)
}

type InstructionFunc func(wire1Value int, wire2Value int) (output int)

func getZWiresValue(wiresMap map[string]int, gatesAndWiresMap map[string][]string) int64 {

	for zWire, _ := range gatesAndWiresMap {
		wiresMap[zWire] = calculateWireValue(zWire, wiresMap, gatesAndWiresMap)
	}

	z := []string{}
	i := 0

	for {
		key := fmt.Sprintf("z%02d", i)
		if _, ok := wiresMap[key]; !ok {
			break
		}
		z = append(z, strconv.Itoa(wiresMap[key]))
		i++
	}
	reverseSlice(z)
	total, err := strconv.ParseInt(strings.Join(z, ""), 2, 64)
	if err != nil {
		log.Fatalf("Error converting binary to int: %v", err)
	}

	return total
}

func calculateWireValue(wire string, wiresMap map[string]int, gatesAndWiresMap map[string][]string) int {
	instructionFuncs := map[string]InstructionFunc{
		"AND": and,
		"OR":  or,
		"XOR": xor,
	}

	if value, exists := wiresMap[wire]; exists {
		return value
	}
	x, op, y := gatesAndWiresMap[wire][0], gatesAndWiresMap[wire][1], gatesAndWiresMap[wire][2]
	wiresMap[wire] = instructionFuncs[op](calculateWireValue(x, wiresMap, gatesAndWiresMap), calculateWireValue(y, wiresMap, gatesAndWiresMap))
	return wiresMap[wire]
}

func or(wire1Value int, wire2Value int) int {
	return wire1Value | wire2Value
}

func and(wire1Value int, wire2Value int) int {
	return wire1Value & wire2Value
}

func xor(wire1Value int, wire2Value int) int {
	return wire1Value ^ wire2Value
}

func reverseSlice(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Part 2
func getSwappedWires(wiresMap map[string]int, gatesAndWiresMap map[string][]string) {
	// identify the highest Z in the gatesAndWiresMap (for identifying the last sum in the ripple carry adder)
	highestZ := 0
	highestZWire := ""
	for zWire, _ := range gatesAndWiresMap {
		if strings.Index(zWire, "z") == 0 {
			z, err := strconv.Atoi(zWire[1:])
			if err != nil {
				log.Fatalf("Error converting wire number to int: %v", err)
			}
			if z > highestZ {
				highestZ = z
				highestZWire = zWire
			}
		}
	}

	wrongWires := make(map[string]bool)
	// identify the wires that are wrong (the ones does not respect the ripple carry adder logic)
	for zWire, operations := range gatesAndWiresMap {
		possibleX := operations[0]
		op := operations[1]
		possibleY := operations[2]

		// The last sum in the ripple carry adder should be XOR
		if strings.Index(zWire, "z") == 0 && op != "XOR" && zWire != highestZWire {
			wrongWires[zWire] = true
		}

		// There should not be any wire that is not a z,x,y wire with XOR operation
		if op == "XOR" &&
			!strings.HasPrefix(possibleX, "x") && !strings.HasPrefix(possibleX, "y") && !strings.HasPrefix(possibleX, "z") &&
			!strings.HasPrefix(possibleY, "x") && !strings.HasPrefix(possibleY, "y") && !strings.HasPrefix(possibleY, "z") &&
			!strings.HasPrefix(zWire, "x") && !strings.HasPrefix(zWire, "y") && !strings.HasPrefix(zWire, "z") {
			wrongWires[zWire] = true
		}

		// There should not be any AND output that is the carry bit and which its next operation is not a OR
		if op == "AND" && possibleX != "x00" && possibleY != "x00" {
			// Check if the output is properly used (for an OR operation)
			for _, nextOperations := range gatesAndWiresMap {
				possibleX2 := nextOperations[0]
				op2 := nextOperations[1]
				possibleY2 := nextOperations[2]
				if (possibleX2 == zWire || possibleY2 == zWire) && op2 != "OR" {
					wrongWires[zWire] = true
				}
			}
		}

		// There should not be any XOR output used to an OR gate -> XOR should be used only for the sum not in propagating of carry
		if op == "XOR" {
			for _, nextOperations := range gatesAndWiresMap {
				possibleX2 := nextOperations[0]
				op2 := nextOperations[1]
				possibleY2 := nextOperations[2]
				if (possibleX2 == zWire || possibleY2 == zWire) && op2 == "OR" {
					wrongWires[zWire] = true
				}
			}
		}
	}
	part2Wires := makeSortedSlice(wrongWires)
	fmt.Println("Part 2's wrong wires: ", strings.Join(part2Wires, ","))

}

func makeSortedSlice(uniqueMap map[string]bool) []string {

	uniqueSlice := make([]string, 0, len(uniqueMap))
	for item := range uniqueMap {
		uniqueSlice = append(uniqueSlice, item)
	}

	sort.Strings(uniqueSlice)
	return uniqueSlice
}
