package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rchacons/advent-of-code/utils"
)

type Blocks struct {
	idNumber int
	fileBlock int
	freeSpace int
	isTreated bool
	changed bool
}

func main() {
	log.Printf("%v days to XMASSS", 24-11)

	input, err := utils.FileToIntList("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalFilesystemChecksum(input)

	fmt.Println(total)
}

func getTotalFilesystemChecksum(numbers []int) int {
	// First convert numbers into the blocks
	//blocks := convertNumbersIntoBlocks(numbers)  // PART ONE
	blockList, blockMap := convertNumbersIntoBlocksPartTwo(numbers) // PART TWO

	// Then move the fileblocks
	//movedBlockList := moveFileBlocks(blocks) // PART ONE
	movedBlockList := moveFileBlocksPartTwo(blockList, blockMap) // PART TWO

	// Finally calculate checksum
	//return calculateCheckSum(movedBlockList)
	return calculateCheckSumPartTwo(movedBlockList)
}

// Part one
func convertNumbersIntoBlocks(numbers []int) []string {
	var blockList []string

	manualIndex := 0
	for i := range numbers{
		if i%2 == 0 {
			for range numbers[i]{
				blockList = append(blockList, strconv.Itoa(manualIndex))
			}
			manualIndex++
		} else {
			for range numbers[i]{
				blockList = append(blockList, ".")
			}
		}
	}
	return blockList
}

func moveFileBlocks(blocks []string)[]string {
	movedBlockList := make([]string, len(blocks))
	copy(movedBlockList, blocks)

	p1 := 0
	p2 := len(blocks)-1

	for p1 < p2{

		if blocks[p1] == "." && blocks[p2] != "." {
			tmp := movedBlockList[p2]
			movedBlockList[p2] = "."
			movedBlockList[p1] = tmp
			p1++
			p2--
		} else if blocks[p1] != "." {
			p1++
		} else if blocks[p2] == "." {
			p2--
		} else {
			p1++
			p2--
		}
	}

	return movedBlockList
}



func calculateCheckSum(movedBlockList []string) int {
	total := 0
	for i := range movedBlockList{
		if movedBlockList[i] != "."{
			num, err := strconv.Atoi(movedBlockList[i])
			if err != nil {
				log.Fatalf("Error converting number : %v", err)
			}
			total += i*num
		}
	}
	return total
}

// Part two
func convertNumbersIntoBlocksPartTwo(numbers []int) ([]interface{}, map[int]Blocks) {
	var blockList []interface{}
	blockMap := make(map[int]Blocks)

	manualIndex := 0
	for i := range numbers{
		if i%2 == 0 {
			block := Blocks{manualIndex, numbers[i], 0, false, false}
			blockMap[manualIndex] = block
			for range numbers[i]{
				blockList = append(blockList, int(manualIndex))
			}
			manualIndex++
		} else {
			block := blockMap[manualIndex-1]
			blockMap[manualIndex-1] = Blocks{block.idNumber, block.fileBlock, numbers[i], false, false}
			for range numbers[i]{
				blockList = append(blockList, ".")
			}
		}
	}
	return blockList, blockMap
}

func moveFileBlocksPartTwo(blocks []interface{}, blockMap map[int]Blocks)[]interface{} {
	movedBlockList := make([]interface{}, len(blocks))
	copy(movedBlockList, blocks)
	changed := false

	for i:= len(blocks)-1; i >= 0; i--{

		if blocks[i] != "." {
			num := movedBlockList[i].(int)
			block := blockMap[num]
			if block.isTreated && block.changed {
				movedBlockList[i] = "."
			} else if !block.isTreated {
				movedBlockList, changed = getMovedTopLeftArray(movedBlockList, block, i, blockMap)
				if changed {
					movedBlockList[i] = "."
				}
				block := blockMap[num]
				blockMap[num] = Blocks{block.idNumber, block.fileBlock, block.freeSpace, true, block.changed}
			}
		}
	}
	return movedBlockList
}

func getMovedTopLeftArray(blockList []interface{}, block Blocks, index int, blockMap map[int]Blocks) ([]interface{},bool){
	movedBlockListCopy := make([]interface{}, len(blockList))
	copy(movedBlockListCopy, blockList)
	beginIndex := 0
	counter := 0
	for i := 0; i<=index; i++ {
		if blockList[i] == "."{
			if beginIndex == 0 {
				beginIndex = i
			}
			counter++
		} else if beginIndex != 0 {

			if counter >= block.fileBlock {
				for j := beginIndex; j < beginIndex + block.fileBlock ; j++{
					movedBlockListCopy[j] = blockList[index]
					blockMap[block.idNumber] = Blocks{block.idNumber, block.fileBlock, block.freeSpace, true ,true}
				}
				return movedBlockListCopy, true
			} else {
				beginIndex = 0
				counter = 0
			}
		}
	}
	return movedBlockListCopy, false
}



func calculateCheckSumPartTwo(movedBlockList []interface{}) int {
	total := 0
	for i := range movedBlockList{
		if movedBlockList[i] != "."{
			num := movedBlockList[i].(int)
			total += int(i) * num
		}
	}
	return total
}
