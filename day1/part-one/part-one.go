package main

import (
	"sort"
	"log"
	"fmt"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	
	list1, list2, err := utils.FileToLists("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var distance int
	for i := range list1 {
		distance += absDiffInt(list1[i],list2[i])
	}
	fmt.Println(distance)

}

func absDiffInt(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

