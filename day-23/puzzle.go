package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	log.Printf("XMAS is over but we still have some work to do...")
	input, err := utils.FileToLanMap("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// ******** FIRST PART **********
	total := getInterconnectedSets(input)
	fmt.Printf("Total part one: %d\n", total)

	// ******** SECOND PART **********
	bronKerboschWrapper(input)


}

func getInterconnectedSets(lanMap map[string]map[string]bool) int {
	total := 0
	reducedMap := make(map[string]map[string]bool)
	
	for k, v := range lanMap {
		if strings.Index(k, "t") == 0{
			reducedMap[k] = v
		}
	}

	visitedSet := make(map[string]bool)
	for k,v := range reducedMap{
		for neighbor, _ := range v {
			for neighbor2, _ := range v{
				set := []string{k, neighbor, neighbor2}
				sort.Strings(set)
				setKey := strings.Join(set, ",")
				if !visitedSet[setKey] {
					if lanMap[neighbor][neighbor2]{
						total++
						// fmt.Printf("Set found: %s, %s, %s\n", k, neighbor, neighbor2)
					}
					visitedSet[setKey] = true
				}
			}
		}
	}
	return total
}

// Part two : Bron-Kerbosch algorithm

func bronKerboschWrapper(lanMap map[string]map[string]bool) {
    vertices := []string{}
    for vertex := range lanMap {
        vertices = append(vertices, vertex)
    }
	totalR := [][]string{}
    bronKerbosch([]string{}, vertices, []string{}, lanMap, &totalR)
	
	// Extract the largest clique
	largest, index := 0, 0
	for i, v := range totalR{
		if len(v) > largest{
			largest = len(v)
			index = i
		}
	}

	largestR := removeDuplicatesAndSort(totalR[index])

	fmt.Println("Largest Clique:", strings.Join(largestR, ","))
}

func removeDuplicatesAndSort(slice []string) []string {
    uniqueMap := make(map[string]bool)
    for _, item := range slice {
        uniqueMap[item] = true
    }

    uniqueSlice := make([]string, 0, len(uniqueMap))
    for item := range uniqueMap {
        uniqueSlice = append(uniqueSlice, item)
    }

    sort.Strings(uniqueSlice)
    return uniqueSlice
}

func bronKerbosch(R, P, X []string, graph map[string]map[string]bool, totalR *[][]string) {
    if len(P) == 0 && len(X) == 0 {
        *totalR = append(*totalR, append([]string{}, R...)) // Make a copy of R
        return
    }
    PCopy := append([]string{}, P...) // Copy of P to iterate over
    for _, v := range PCopy {
        newR := append(R, v) // Add v to R

        // Compute P ⋂ N(v) and X ⋂ N(v)
        neighborsV := neighbors(v, graph)
        newP := intersect(P, neighborsV)
        newX := intersect(X, neighborsV)

        bronKerbosch(newR, newP, newX, graph, totalR)

        P = remove(P, v)
        X = append(X, v)
    }
}

func neighbors(vertex string, graph map[string]map[string]bool) []string {
    neighbors := []string{}
    for neighbor := range graph[vertex] {
        neighbors = append(neighbors, neighbor)
    }
    return neighbors
}

func intersect(slice1, slice2 []string) []string {
    set := make(map[string]bool)
    for _, v := range slice1 {
        set[v] = true
    }
    result := []string{}
    for _, v := range slice2 {
        if set[v] {
            result = append(result, v)
        }
    }
    return result
}

func remove(slice []string, elem string) []string {
    index := -1
    for i, v := range slice {
        if v == elem {
            index = i
            break
        }
    }
    if index == -1 {
        return slice // Element not found; return original slice
    }
    return append(slice[:index], slice[index+1:]...)
}
