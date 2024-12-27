package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Plant struct {
	letter    string
	cornerCount int
	checked   bool
}

type Position struct {
	X, Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-12)

	input, err := utils.FileToStringMatrix("../input.txt")
	// input test 1 -> 80
	// input test 2 -> 436
	// input test 3 -> 1206
	// input test 4 -> 236
	// input test 5 -> 368
	// input ?

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalFencingPrice(input)

	fmt.Println(total)
}

func getTotalFencingPrice(garden [][]string) int {

	// First store every plant in a map while checking its neighbors
	neighborsMap := make(map[Position][]Position)
	plantPositionMap := make(map[Position]Plant)

	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			pos := Position{i, j}
			plant := Plant{garden[i][j], 0, false}
			plantPositionMap[pos] = plant
			trackPlantsWithNeighbors(garden, neighborsMap, plantPositionMap, pos)
		}
	}

	// Then do the dfs
	total := findTheRegionsWithDfs(neighborsMap, plantPositionMap)

	return total
}

func trackPlantsWithNeighbors(garden [][]string, neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, currentPosition Position) {

	up := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X - 1, currentPosition.Y}, currentPosition, true)
	down := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X + 1, currentPosition.Y}, currentPosition, true)
	right := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X, currentPosition.Y + 1}, currentPosition, true)
	left := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X, currentPosition.Y - 1}, currentPosition, true)
	diagUpR := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X - 1, currentPosition.Y + 1}, currentPosition, false)
	diagUpL := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X - 1, currentPosition.Y - 1}, currentPosition, false)
	diagDownR := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X + 1, currentPosition.Y + 1}, currentPosition, false)
	diagDownL := checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X + 1, currentPosition.Y - 1}, currentPosition, false)
	
	cornerCount := 0
	plant := plantPositionMap[currentPosition]
	
	// Top left corner
	if (left && up && !diagUpL) || (!left && !up) {
		cornerCount++
	}

	// Top right corner
	if (right && up && !diagUpR) || (!right && !up) {
		cornerCount++
	}

	// Bottom right corner
	if (right && down && !diagDownR) || (!right && !down) {
		cornerCount++
	}

	// Bottom left corner
	if (left && down && !diagDownL) || (!left && !down) {
		cornerCount++
	}

	plant.cornerCount += cornerCount
	plantPositionMap[currentPosition] = plant

	// No matching neighbor
	if _, exists := neighborsMap[currentPosition]; !exists {
		neighborsMap[currentPosition] = nil
	}
}

func checkAndAddNeighbors(garden [][]string, neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, nextPosition Position, currentPosition Position, shouldAdd bool) bool {
	if !isPositionInMatrix(garden, nextPosition) ||
		garden[currentPosition.X][currentPosition.Y] != garden[nextPosition.X][nextPosition.Y] {
		plant := plantPositionMap[currentPosition]
		plantPositionMap[currentPosition] = plant
		return false
	} else if garden[currentPosition.X][currentPosition.Y] == garden[nextPosition.X][nextPosition.Y] {
		if shouldAdd {
			neighborsMap[currentPosition] = append(neighborsMap[currentPosition], nextPosition)
		}
		return true
	}
	return false
}

func findTheRegionsWithDfs(neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant) int {

	regionList := make([]int, 0)
	for plantPos, plantNeighborsPos := range neighborsMap {
		plant := plantPositionMap[plantPos]
		if !plant.checked {
			plant.checked = true
			plantPositionMap[plantPos] = plant

			regionArea := 1
			regionSides := plant.cornerCount
			for _, neighbor := range plantNeighborsPos {
				area, cornerCount := dfsRecursive(neighborsMap, plantPositionMap, neighbor)
				regionArea += area
				regionSides += cornerCount
			}
			// mul := regionArea*regionPerim
			// fmt.Printf("\nRegion of plant %v has price : %d ", plant, regionArea*regionPerim)

			regionList = append(regionList, regionArea*regionSides)
		}
	}

	total := 0
	for _, regionPrice := range regionList {
		total += regionPrice
	}
	return total
}

func dfsRecursive(neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, plantPos Position) (int, int) {
	plant := plantPositionMap[plantPos]
	if !plant.checked {
		plant.checked = true
		plantPositionMap[plantPos] = plant

		regionArea := 1
		regionSides := plant.cornerCount
		for _, neighbor := range neighborsMap[plantPos] {
			area, cornerCount := dfsRecursive(neighborsMap, plantPositionMap, neighbor)
			regionArea += area
			regionSides += cornerCount
		}
		return regionArea, regionSides
	}
	return 0, 0
}

func isPositionInMatrix(garden [][]string, pos Position) bool {
	return pos.X < len(garden) &&
		pos.X >= 0 &&
		pos.Y >= 0 &&
		pos.Y < len(garden[pos.X])
}
