package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	
	list1, list2, err := utils.FileToLists("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Store each number in list2 in a map with its ocurrences (bucket sort)
	m := make(map[int]int)
	for _, num := range list2 {
		m[num] = m[num] + 1
	}
	
	var similarity int
	// Then just retrieve the ocurrences of each number of list1 in the map
	for _, num := range list1 {
		if ocurrences, exists := m[num]; exists {
			similarity += num*ocurrences
		}
	}
	
	fmt.Println(similarity)

}