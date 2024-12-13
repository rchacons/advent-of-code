package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Plant struct {
	letter string
	perimeter int
	checked bool
}

type Position struct{
	X,Y int
}

func main() {
	log.Printf("%v days to XMASSS", 24-12)

	input, err := utils.FileToStringMatrix("input.txt")
	// input test 1 -> 140
	// input test 2 -> 772
	// input test 3 -> 1930
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

	for i:=0;i<len(garden);i++{
		for j:=0;j<len(garden[i]);j++{
			pos := Position{i,j}
			plant := Plant{garden[i][j], 0, false}
			plantPositionMap[pos] = plant
			trackPlantsWithNeighbors(garden,neighborsMap,plantPositionMap, pos)
		}
	}

	// Then do the dfs 
	total := findTheRegionsWithDfs(neighborsMap, plantPositionMap)

	return total
}

func trackPlantsWithNeighbors(garden [][]string, neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, currentPosition Position){
	
	// Check upper neighbor
	checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X-1, currentPosition.Y}, currentPosition)

	// Check lower neighbor
	checkAndAddNeighbors(garden, neighborsMap, plantPositionMap,Position{currentPosition.X+1, currentPosition.Y}, currentPosition)

	// Check right neighbor
	checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X, currentPosition.Y+1}, currentPosition)

	// Check left neighbor
	checkAndAddNeighbors(garden, neighborsMap, plantPositionMap, Position{currentPosition.X, currentPosition.Y-1}, currentPosition)

	// No matching neighbor
	if _,exists := neighborsMap[currentPosition]; !exists{
		neighborsMap[currentPosition] = nil
	}
}

func checkAndAddNeighbors(garden [][]string, neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, nextPosition Position, currentPosition Position){
	if !isPositionInMatrix(garden, nextPosition) || 
	garden[currentPosition.X][currentPosition.Y] != garden[nextPosition.X][nextPosition.Y] { 
		plant := plantPositionMap[currentPosition]
		plant.perimeter++
		plantPositionMap[currentPosition] = plant
	} else if garden[currentPosition.X][currentPosition.Y] == garden[nextPosition.X][nextPosition.Y] {
		neighborsMap[currentPosition] = append(neighborsMap[currentPosition], nextPosition)
	}
}

func findTheRegionsWithDfs(neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant) int{

	regionList := make([]int,0)
	for plantPos, plantNeighborsPos := range neighborsMap{
		plant := plantPositionMap[plantPos]
		if !plant.checked {
			plant.checked = true 
			plantPositionMap[plantPos] = plant

			regionArea := 1
			regionPerim := plant.perimeter
			for _, neighbor := range plantNeighborsPos{
				area,perim := dfsRecursive(neighborsMap, plantPositionMap, neighbor)
				regionArea += area
				regionPerim += perim
			}
			// mul := regionArea*regionPerim 
			// fmt.Printf("\nRegion of plant %v has price : %d ", plant, regionArea*regionPerim)

			regionList = append(regionList, regionArea*regionPerim)
		}
	}

	total := 0
	for _, regionPrice := range regionList {
		total += regionPrice
	}
	return total
}

func dfsRecursive(neighborsMap map[Position][]Position, plantPositionMap map[Position]Plant, plantPos Position) (int,int) {
	plant := plantPositionMap[plantPos]
	if !plant.checked {
		plant.checked = true 
		plantPositionMap[plantPos] = plant

		regionArea := 1
		regionPerim := plant.perimeter
		for _, neighbor := range neighborsMap[plantPos]{
			area,perim := dfsRecursive(neighborsMap, plantPositionMap, neighbor)
			regionArea += area
			regionPerim += perim
		}
		return regionArea,regionPerim
	}
	return 0,0
}


func isPositionInMatrix(garden [][]string, pos Position) bool{
	return pos.X < len(garden) && 
			pos.X >= 0 && 
			pos.Y >= 0 && 
			pos.Y < len(garden[pos.X])
}