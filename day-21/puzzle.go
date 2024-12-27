package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/rchacons/advent-of-code/utils"
)

type Position struct {
	row, col int
}

type KeyPair struct {
	key1, key2 string
}

type QueueElement struct {
	position Position
	moves    string
}

func main() {
	log.Printf("%v days to XMASSS", 24-20)
	input, err := utils.FileToStringList("input.txt")

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// ******** FIRST AND SECOND PART **********
	total := getComplexityOfCodes(input)
	fmt.Println(total)
}

func getComplexityOfCodes(codes []string) int {

	total := 0

	numKeypad := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"none", "0", "A"},
	}

	// First compute the sequence of moves needed to get from one num to another in the num keypad
	numsSeqs := computeSequence(numKeypad)

	dir_keypad := [][]string{
		{"none", "^", "A"},
		{"<", "v", ">"},
	}
	
	// Then compute the sequence of moves needed to get from one direction to another in the dir keypad
	dirSeqs := computeSequence(dir_keypad)

	// Compute the length of the sequences for the dir keypad
	dirLength := make(map[KeyPair]int)
	for key, seq := range dirSeqs {
		dirLength[key] = len(seq[0])
	}

	for _, code := range codes {
		// solve with the code and the num seqs
		inputs := solveCode(code, numsSeqs)

		// For each input, compute the length of the sequence
		lengths := make([]int, len(inputs))
		for i, input := range inputs {
			lengths[i] = computeLength(make(map[string]int), dirSeqs, dirLength, input, 25)
		}

		minLength := 1000000000000000000
    	for _, length := range lengths {
        if length < minLength {
            minLength = length
        }
    }


		total += minLength * stringToInt(code[:len(code)-1])

	}
	return total
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// Compute the sequence of moves on a keypad to get from one key to another
func computeSequence(keypad [][]string) map[KeyPair][]string {
	pos := make(map[string]Position)
	for i := range keypad {
		for j := range keypad[i] {
			if keypad[i][j] != "none" {
				pos[keypad[i][j]] = Position{i, j}
			}
		}
	}

	// The needed sequences to get from one key to another
	seqs := make(map[KeyPair][]string)
	for key1 := range pos {
		for key2 := range pos {
			if key1 == key2 {
				// No need to move, just press A
				seqs[KeyPair{key1, key2}] = []string{"A"}
				continue
			}
			// Do BFS to get the sequence of moves
			seqs[KeyPair{key1, key2}] = bfs(pos[key1], pos[key2], keypad)
		}
	}
	return seqs
}

func bfs(start, end Position, keypad [][]string) []string {
	var possibilities []string
	queue := []QueueElement{{start, ""}}
	shortestPathLength := 1000000

outerLoop:
	for len(queue) > 0 {
		var current QueueElement
		current, queue = popleft(queue)

		r, c := current.position.row, current.position.col
		moves := current.moves

		possibleMoves := []struct {
			newRow, newCol int
			newMove        string
		}{
			{r - 1, c, "^"},
			{r + 1, c, "v"},
			{r, c - 1, "<"},
			{r, c + 1, ">"},
		}

		for _, move := range possibleMoves {
			newRow, newCol := move.newRow, move.newCol
			if newRow < 0 || newRow >= len(keypad) || newCol < 0 || newCol >= len(keypad[0]) {
				continue
			}
			if keypad[newRow][newCol] == "none" {
				continue
			}
			// Found the end position
			if newRow == end.row && newCol == end.col {
				// We already have a shorter path
				if shortestPathLength < len(moves)+1 {
					break outerLoop
				}

				shortestPathLength = len(moves) + 1
				// Add the move to the possibilities with A as we need to press
				possibilities = append(possibilities, moves+move.newMove+"A")
			} else {
				// Enqueue the new position
				queue = append(queue, QueueElement{Position{newRow, newCol}, moves + move.newMove})
			}
		}

	}

	return possibilities
}

// popleft removes the first element from a slice and returns the remaining slice
func popleft[T any](slice []T) (T, []T) {
	var zeroValue T
	if len(slice) == 0 {
		return zeroValue, slice
	}
	return slice[0], slice[1:]
}

// Generate all possible sequences by combining the sequences associated with each key pair
func solveCode(code string, numSeqs map[KeyPair][]string) []string {
	// Create a list of options based on the input code
	var options [][]string
	for i, c := range code {
		var keyPair KeyPair
		if i == 0 {
			keyPair = KeyPair{"A", string(c)}
		} else {
			keyPair = KeyPair{string(code[i-1]), string(c)}
		}
		if seq, ok := numSeqs[keyPair]; ok {
			options = append(options, seq)
		}
	}

	// Get all possible combinations of the sequences
	combinations := cartesianProduct(options)
	var results []string
	for _, comb := range combinations {
		results = append(results, strings.Join(comb, ""))
	}
	return results
}

func cartesianProduct(arrays [][]string) [][]string {
	if len(arrays) == 0 {
		return [][]string{{}}
	}
	result := [][]string{}
	for _, v := range arrays[0] {
		for _, prod := range cartesianProduct(arrays[1:]) {
			result = append(result, append([]string{v}, prod...))
		}
	}
	return result
}
func computeLength(cache map[string]int, dirSeqs map[KeyPair][]string, dirLengths map[KeyPair]int, seq string, depth int) int {
    cacheKey := seq + string(depth)
    if cachedResult, found := cache[cacheKey]; found {
        return cachedResult
    }

    if depth == 1 {
        length := 0
        for i, c := range seq {
            var keyPair KeyPair
            if i == 0 {
                keyPair = KeyPair{"A", string(c)}
            } else {
                keyPair = KeyPair{string(seq[i-1]), string(c)}
            }
            length += dirLengths[keyPair]
        }
        cache[cacheKey] = length
        return length
    }

    length := 0
    for i, c := range seq {
        var keyPair KeyPair
        if i == 0 {
            keyPair = KeyPair{"A", string(c)}
        } else {
            keyPair = KeyPair{string(seq[i-1]), string(c)}
        }
        minLength := 1000000000000000000
        for _, subseq := range dirSeqs[keyPair] {
            subLength := computeLength(cache, dirSeqs, dirLengths, subseq, depth-1)
            if subLength < minLength {
                minLength = subLength
            }
        }
        length += minLength
    }

    cache[cacheKey] = length
    return length
}
