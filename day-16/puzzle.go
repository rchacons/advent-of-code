package main

import (
	"container/heap"
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

/****************** HEAP **********************/
// Used for day 16
type Position struct {
	X,Y int
}
type Node struct {
	Pos Position
	Direction string
}

type Path struct {
    Nodes []Node
    score int
}

type Item struct {
	Value Node
	Priority int
    Path []Node
	Index int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {return len(pq)}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}){
	n := len(*pq) //current length of the priority queue
	item := x.(*Item) // asserts that x is of type item
	item.Index = n // allows to later quickly find and update without having to search throught the entire queue
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}
//

/****************** END OF HEAP **********************/



func main() {
	log.Printf("%v days to XMASSS", 24-16)
	input, err := utils.FileToStringMatrix("input.txt")

	// input test 1 -- > 7036  (part two should be 45)
	// input test 2 -- > 11048  (part two should be 64)
	// input test 3 -- > 4013 
	// input tesst 4 --> 5024 
	// inptu test 5 --> 1004

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	visitedMap, _, endNode := getLowerMazeScoreWithDijkstra(input)
    totalPart1, _ := getMinEndNote(visitedMap,endNode)
    
    //totalPart2 := getNumberOfTilesForBestPath(input, visitedMap, startNode, endNodeWithDirection, totalPart1)

    totalPart2 := getLowerMazeScoreWithDijkstraPartTwo(input, totalPart1)

	fmt.Println("Total part 1 : ", totalPart1)
	fmt.Println("Total part 2 : ",totalPart2)

}


func getNeighborsImproved(maze [][]string, node Node) map[Node]int {
    neighbors := make(map[Node]int)

    directions := []struct {
        dx, dy int
        dir    string
    }{
        {-1, 0, "^"},
        {1, 0, "v"},
        {0, -1, "<"},
        {0, 1, ">"},
    }

    for _, d := range directions {
        x, y := node.Pos.X+d.dx, node.Pos.Y+d.dy
        if x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] != "#" {
            if node.Direction == d.dir {
                score := 1
                neighbors[Node{Position{x, y}, d.dir}] = score
            } else {
                score := 1000
                neighbors[Node{Position{node.Pos.X, node.Pos.Y}, d.dir}] = score
            }
            
        }
    }
    return neighbors
}


func getLowerMazeScoreWithDijkstra(maze [][]string) (map [Node]int, Node, Node){
    start := findNode(maze, "S")
    start.Direction = ">"
    end := findNode(maze, "E")

    visitedMap := make(map[Node]int)

    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    item1 := Item{
        Value:    start,
        Priority: 0,
        Path: []Node{start},
    }

    heap.Push(&pq, &item1)
    visitedMap[start] = 0

    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item) // item with the lowest cost
        currentNode := item.Value
        cost := item.Priority
        path := item.Path

        // Skip if we have already processed a better cost for this state
        if visitedMap[currentNode] < cost {
            continue
        } 

        for neighbor, costToReach := range getNeighborsImproved(maze, currentNode) {
            newCost := cost + costToReach
            if oldCost, ok := visitedMap[neighbor]; !ok || newCost < oldCost {
                visitedMap[neighbor] = newCost
                newPath := append([]Node{}, path...)
                newPath = append(newPath, neighbor)
                heap.Push(&pq, &Item{
                    Value:    Node{neighbor.Pos, neighbor.Direction},
                    Priority: newCost,
                    Path: newPath,
                })
//                fmt.Printf("Adding neighbor at position (%d, %d) with direction %s and new cost %d\n", neighbor.Pos.X, neighbor.Pos.Y, neighbor.Direction, newCost)
            } 
        }
    }

    fmt.Printf("VISITED MAP HAS %d POSITIONS", len(visitedMap))

    // fmt.Printf("Final cost to reach end at position (%d, %d) is %d\n", end.Pos.X, end.Pos.Y, minCost)
    return visitedMap, start, end
}


func findNode(maze [][]string, mark string) Node {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == mark {
				return Node{Position{i,j}, ">"}
			}
		}
	}
	return Node{}
}

func getMinEndNote(visitedMap map[Node]int, end Node) (int,Node){
    // Find the lowest cost among all possible end nodes
    possibleEndNodes := []Node{
        {end.Pos, "<"},
        {end.Pos, "^"},
        {end.Pos, ">"},
        {end.Pos, "v"},
    }

    finalEndNode := Node{}
    minCost := int(^uint(0) >> 1) // Initialize to max int
    for _, endNode := range possibleEndNodes {
        if cost, exists := visitedMap[endNode]; exists && cost < minCost {
            minCost = cost
            finalEndNode = endNode
        }
    }

    fmt.Printf("Final cost to reach end at position (%d, %d) is %d\n", end.Pos.X, end.Pos.Y, minCost)
    return minCost, finalEndNode
}

func removeNode(nodes []Node, node Node) []Node {
    index := -1
    for i, p := range nodes {
        if p == node {
            index = i
            break
        }
    }

    if index == -1 {
        return nodes
    }

    return append(nodes[:index], nodes[index+1:]...)
}



// instead of stopping evaluating a path if it sees there's already been another path through the same point and direction, stop evaluating a path only if its score is higher
func getLowerMazeScoreWithDijkstraPartTwo(maze [][]string, maxCost int) int {
    start := findNode(maze, "S")
    start.Direction = ">"
    end := findNode(maze, "E")
    visitedMap := make(map[Node]int)
    paths := make([][]Node, 0)

    pq := make(PriorityQueue, 0)
    heap.Init(&pq)

    item1 := Item{
        Value:    start,
        Priority: 0,
        Path:     []Node{start},
    }

    heap.Push(&pq, &item1)
    visitedMap[start] = 0

    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item) // item with the lowest cost
        currentNode := item.Value
        cost := item.Priority
        path := item.Path

        if cost > maxCost {
            continue
        }
        // Skip if we have already processed a better cost for this state
        if visitedMap[currentNode] < cost {
            continue
        }

        if currentNode.Pos == end.Pos {
            paths = append(paths, path)
        } 

        for neighbor, costToReach := range getNeighborsImproved(maze, currentNode) {
            newCost := cost + costToReach
            if oldCost, ok := visitedMap[neighbor]; !ok || newCost <= oldCost  && currentNode != end  {
                visitedMap[neighbor] = newCost
                newPath := append([]Node{}, path...)
                newPath = append(newPath, neighbor)
                heap.Push(&pq, &Item{
                    Value:    Node{neighbor.Pos, neighbor.Direction},
                    Priority: newCost,
                    Path:     newPath,
                })
            }
        }
    }


    finalMap := make(map[Node]bool)
    for _, path := range paths {
        for _, node := range path {
            finalMap[Node{node.Pos, ""}] = true
        }
    }

    return len(finalMap)
}

