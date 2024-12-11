package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	X,Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-10)

	input, err := utils.FileToIntList("input_test.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	fmt.Println(input)
	fmt.Println(len(input))


	//total := getTotalAntinodeLocations(input)
	total := 0
	fmt.Println(total)
}