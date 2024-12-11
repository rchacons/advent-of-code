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

	input, err := utils.FileToStringMatrix("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalAntinodeLocations(input)

	fmt.Println(total)
}

func getTotalAntinodeLocations(antennaMatrix [][]string) int {
	antennaMap := getAntennasPositions(antennaMatrix)
	total := 0
	antinodesMap := make(map[Position]bool)
	for antenna, posList := range antennaMap {
		log.Printf("\n\n**************** CHECKING ANTENNA : %v*********************\n\n",antenna)
		total += checkAntinodesForAntennaPartTwo(posList, antennaMatrix, antinodesMap)
	}
	return total;
}

func getAntennasPositions(antennaMatrix [][]string) map[string][]Position{
	antennaMap := make(map[string][]Position)
	for i := range antennaMatrix{
		for j := range antennaMatrix[i]{
			val := antennaMatrix[i][j]
			if (val != "."){
				if _, exists := antennaMap[val]; exists {
					pos := Position{i,j}
					antennaMap[val] = append(antennaMap[val], pos)
				} else {
					antennaMap[val] = []Position{{i,j}}
				}
			}
		}
	}
	return antennaMap
}

func checkAntinodesForAntenna(posList []Position, antennaMatrix [][]string, antinodesMap map[Position]bool) int {
	total := 0
	for i:=0; i< len(posList); i++ {
		pos := posList[i]
		for _, pos2 := range posList[i+1:]{
			
			distance := Position{pos.X-pos2.X, pos.Y-pos2.Y}

			if distance.X != 0 && distance.Y != 0{

				tempPos1 := Position{pos.X+distance.X, pos.Y+distance.Y}
				tempPos2 := Position{pos2.X-distance.X, pos2.Y-distance.Y}
				if isPositionInMatrixAndTrack (tempPos1, antennaMatrix, antinodesMap) {
					if (!antinodesMap[tempPos1]) {
						antinodesMap[tempPos1] = true
						total++
					}
				}
				if isPositionInMatrixAndTrack(tempPos2, antennaMatrix, antinodesMap){
					if (!antinodesMap[tempPos2]) {
						antinodesMap[tempPos2] = true
						total++
					}
				}
			}
		}
	}
	return total
}

func checkAntinodesForAntennaPartTwo(posList []Position, antennaMatrix [][]string, antinodesMap map[Position]bool) int {
	total := 0
	for i:=0; i< len(posList); i++ {
		pos := posList[i]
		for _, pos2 := range posList[i+1:]{
			distance := Position{pos.X-pos2.X, pos.Y-pos2.Y}

			if distance.X != 0 && distance.Y != 0{
				tempPos1 := pos
				
				for isPositionInMatrixAndTrack(tempPos1, antennaMatrix, antinodesMap){
					if (!antinodesMap[tempPos1]) {
						antinodesMap[tempPos1] = true
						total++
					}
					tempPos1.X = tempPos1.X+distance.X
					tempPos1.Y = tempPos1.Y+distance.Y
				}
				tempPos2 := pos2
				
				for isPositionInMatrixAndTrack(tempPos2, antennaMatrix, antinodesMap){
					if (!antinodesMap[tempPos2]) {
						antinodesMap[tempPos2] = true
						total++
					}
					tempPos2.X = tempPos2.X-distance.X
					tempPos2.Y = tempPos2.Y-distance.Y
				}
			}
		}
	}
	return total
}
func isPositionInMatrixAndTrack(pos Position, antennaMatrix [][]string, antinodesMap map[Position]bool) bool {
	return pos.X < len(antennaMatrix) && pos.X >= 0 && pos.Y >= 0 && pos.Y < len(antennaMatrix[pos.X])
	
}