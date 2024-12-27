package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Robot struct {
	name string
	position Position
	rowV int
	colV int
}

type Position struct{
	X,Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-15)

	input, err := utils.FileToIntegerListsWithNegatives("input.txt")
	// input test -> 480
	// input test 1 -> 14

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	//initialSpace := initiateSpace(11,7) // Test case
	initialSpace := initiateSpace(101,103) // Input
	robots := initiateRobots(input)
	total := simulateAndCalculateSafetyFactor(initialSpace, robots, 100)

	fmt.Println(total)
}

func initiateSpace(spaceWidth int, spaceHeight int) [][]int {
	space := make([][]int, spaceHeight)

	for i:=0; i<spaceHeight; i++{
		space[i] = make([]int, spaceWidth)
		for j := 0; j<spaceWidth; j++ {
			space[i][j] = 0
		}
	}
	return space
}

func initiateRobots(input [][]int) []Robot{
	robotList := make([]Robot,0)
	for i:=0;i<len(input);i++{
		robot := Robot{fmt.Sprintf("robot-%d",i), Position{input[i][1],input[i][0]}, input[i][3], input[i][2]}
		robotList = append(robotList, robot)
	}
	return robotList
}

func simulateAndCalculateSafetyFactor(space [][]int, robots []Robot, seconds int) int {

	for _, robot := range robots {
		// Add robot to the position
		space[robot.position.X][robot.position.Y]++
		for range seconds{
			// fmt.Printf("\n\n********** ROBOT %s simulation********\n", robot.name)
			simulateRobotMovement(&space, &robot)
		}
	}
	return calculateSafetyFactorWithoutMiddle(space)
}

// for debugging
func printSpace(space [][]int) {
    for _, row := range space {
        for _, val := range row {
            fmt.Printf("%d ", val)
        }
        fmt.Println()
    }
}
func simulateRobotMovement(space *[][]int, robot *Robot){
	// First remove the position from the space
	(*space)[robot.position.X][robot.position.Y]--

	newPosition := Position{robot.position.X+robot.rowV, robot.position.Y+robot.colV}
	// Then deal with overlapping
	newPosition = dealWithPositionOverlapping(*space, newPosition)

	robot.position.X = newPosition.X
	robot.position.Y = newPosition.Y

	// Finally add new robot position
	(*space)[robot.position.X][robot.position.Y]++
	
}

func dealWithPositionOverlapping(space [][]int, pos Position) Position{
	fixedPosition := pos
	// upper border
	if pos.X < 0 {fixedPosition.X = len(space) + fixedPosition.X}

	// lower border
	if pos.X >= len(space) {fixedPosition.X = fixedPosition.X - len(space)}

	// right border
	if pos.Y >= len(space[0]) {fixedPosition.Y = fixedPosition.Y - len(space[0])}

	// left border
	if pos.Y < 0 {fixedPosition.Y = len(space[0]) + fixedPosition.Y}

	return fixedPosition
}

func calculateSafetyFactorWithoutMiddle(space [][]int) int {
	middleHorizontal := (len(space)/2)
	middleVertical := (len(space[0])/2)

	// Define quadrants
	spaceQ1 := make([][]int, middleHorizontal)
	spaceQ2 := make([][]int, middleHorizontal)
	spaceQ3 := make([][]int, len(space)-middleHorizontal-1)
	spaceQ4 := make([][]int, len(space)-middleHorizontal-1)

	for i := 0; i < middleHorizontal; i++ {
        spaceQ1[i] = space[i][:middleVertical]
        spaceQ2[i] = space[i][middleVertical+1:]
    }
    for i := middleHorizontal+1; i < len(space); i++ {
        spaceQ3[i-middleHorizontal-1] = space[i][:middleVertical]
        spaceQ4[i-middleHorizontal-1] = space[i][middleVertical+1:]
    }

	// fmt.Println(spaceQ1)
	// fmt.Println(spaceQ2)
	// fmt.Println(spaceQ3)
	// fmt.Println(spaceQ4)

	return 	calculateSafetyForQuadrant(spaceQ1) * 
			calculateSafetyForQuadrant(spaceQ2) * 
			calculateSafetyForQuadrant(spaceQ3) *
			calculateSafetyForQuadrant(spaceQ4)
}

func calculateSafetyForQuadrant(quadrant [][]int) int {
	total := 0
	for _, row := range quadrant{
		for _, val := range row {
			total += val
		}
	}
	return total
}

