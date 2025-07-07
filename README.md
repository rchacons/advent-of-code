# 🎄 Advent of Code 2024 in Go

This repository contains my solutions to the [Advent of Code 2024](https://adventofcode.com/) challenges, implemented in Go. I used this opportunity to learn the Go programming language while solving the daily puzzles.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Advent of Code](https://img.shields.io/badge/Advent%20of%20Code-2024-brightgreen)

## 📚 About

Advent of Code is an annual set of Christmas-themed programming puzzles that can be solved in any programming language. Each puzzle has two parts, with the second part building upon the first but typically requiring a more optimized or extended solution.

I decided to use the Advent of Code 2024 challenges as a practical way to learn and explore Go, focusing on:
- Go syntax and idioms
- Efficient data structures
- Algorithms implementation
- Problem-solving in Go

## 🗂️ Repository Structure

The repository is organized by days, with each day containing one or two parts of the puzzle:

```
advent-of-code/
├── day-01/
│   ├── part-one/
│   │   └── part-one.go
│   └── part-two/
│       └── part-two.go
├── day-02/
│   └── ...
...
├── day-25/
│   └── puzzle.go
└── utils/
    └── utils.go
```

Some days have a single file `puzzle.go` that contains solutions for both parts, while others have separate directories for each part.

> **Note**: Input files (`input.txt`) are not committed to the repository in accordance with Advent of Code's request not to share inputs publicly.

## 🚀 How to Run

To run any solution:

1. Clone the repository:
   ```bash
   git clone https://github.com/rchacons/advent-of-code.git
   cd advent-of-code
   ```

2. Make sure you have Go installed (version 1.23.3 or later):
   ```bash
   go version
   ```

3. Navigate to a specific day's directory:
   ```bash
   cd day-1/part-one
   ```

4. Create an `input.txt` file with your puzzle input from the Advent of Code website

5. Run the solution:
   ```bash
   go run part-one.go
   ```

   Or for days with a single file:
   ```bash
   cd day-5
   go run puzzle.go
   ```

## 🛠️ Common Utilities

I've created a `utils` package with reusable functions that help with:
- Reading and parsing input files in various formats
- Converting between data types
- Common algorithms and data structures
- Helper functions for specific puzzle types

This approach allowed me to focus on solving the puzzles rather than rewriting boilerplate code for each day.

##  Learning Outcomes

Throughout this challenge, I gained experience with:

### Go Language Fundamentals
- Mastering Go's type system and syntax
- Error handling patterns and best practices
- Working with slices, maps, and custom data structures
- File I/O and text parsing in Go
- Efficient string manipulation techniques

### Algorithms and Data Structures
- Graph traversal (DFS, BFS) for path finding and connectivity analysis
- Dynamic programming and recursive problem solving
- Two-pointer and sliding window techniques
- Hash-based data structures for efficient lookups and tracking
- Matrix traversal and manipulation algorithms
- Vector and geometric calculations

### Problem-Solving Approaches
- Breaking down complex problems into manageable components
- Pattern recognition in sequence and grid-based problems
- Implementing state machines for simulation problems
- Optimization techniques for reducing time and space complexity
- Backtracking and constraint satisfaction algorithms

### Software Engineering Practices
- Building reusable utility functions and packages
- Test-driven development for algorithmic solutions
- Documentation and code organization
- Performance profiling and optimization
- Regular expression pattern matching and parsing

## 📊 Progress

- ✅ Day 1: **Sorting & Map-based Frequency Counting** - Part 1: Sorting lists and calculating Manhattan distance (sum of absolute differences) between corresponding elements. Part 2: Using hash maps for efficient frequency counting and weighted sum calculation.
- ✅ Day 2: **Report Validation Logic** - Validating reports based on increasing/decreasing sequences with different tolerance rules for variations.
- ✅ Day 3: **Regular Expression Pattern Extraction** - Using regex for extracting and processing nested patterns like "mul(x,y)", with tokenization and parsing to handle context-dependent evaluation regions.
- ✅ Day 4: **Multi-directional Matrix Traversal** - Implementing eight-directional search algorithm (horizontal, vertical, diagonal) to detect "XMAS" pattern sequences within a character matrix.
- ✅ Day 5: **Recursive Ordering Validation with Map Lookups** - Using maps for efficient rule storage and recursive divide-and-conquer approach to validate ordering constraints between elements, with special pivot-based middle element processing.
- ✅ Day 6: **Directional Matrix Traversal with State Machine** - Implementing a state-based guard movement simulation with collision detection and direction change rules, tracking visited positions in a matrix.
- ✅ Day 7: **Recursive Expression Evaluation with Set Building** - Using recursive computation to build sets of possible values, tracking all possible numerical expressions through different operations (addition, multiplication, concatenation) with hash maps.
- ✅ Day 8: **Vector-based Geometric Calculation** - Finding antinode positions for antennas using vector arithmetic and position tracking with hash maps to avoid duplicate counting.
- ✅ Day 9: **Two-Pointer Block Rearrangement** - Implementing a two-pointer algorithm for efficient file block movement, optimizing space utilization by shifting blocks in-place based on specific rules.
- ✅ Day 10: **Recursive Path Traversal with Backtracking** - Using depth-first recursive traversal to follow numbered paths from trailhead positions (zeros), with position tracking to calculate path scores and ratings.
- ✅ Day 11: **Recursive Stone Blinking Simulation** - Implementing recursive and iterative approaches with caching to simulate "blinking" stones and their splitting rules.
- ✅ Day 12: **DFS Garden Region Analysis** - Using Depth-First Search algorithm to identify connected plant regions in a garden and calculate their fencing costs.
- ✅ Day 13: **Linear Equation Solving with Cramer's Rule** - Solving systems of linear equations using Cramer's rule to find optimal token distributions.
- ✅ Day 14: **Robot Movement Simulation with Wraparound** - Simulating multiple robots' movements with wraparound boundaries and calculating quadrant-based safety factors.
- ✅ Day 15: **Sokoban-style Box Pushing Simulation** - Implementing a recursive box-pushing logic similar to the Sokoban puzzle game, calculating GPS coordinates of final box positions.
- ✅ Day 16: **Dijkstra's Algorithm with Priority Queue** - Implementing Dijkstra's shortest path algorithm with a custom priority queue to navigate a maze with directional constraints.
- ✅ Day 17: **Virtual Machine & Reverse Engineering** - Building a virtual machine with custom instruction set and reverse engineering program behavior using recursive search.
- ✅ Day 18: **BFS Path Finding with Obstacle Detection** - Using Breadth-First Search algorithm to find paths through a memory map with fallen bytes and identify blocking positions.
- ✅ Day 19: **Dynamic Programming with Memoization** - Applying dynamic programming with memoization to efficiently solve pattern matching problems and count possible design combinations.
- ✅ Day 20: **Graph Traversal & Manhattan Distance** - Combining BFS traversal with Manhattan distance calculations to find shortcuts ("cheats") in a race map.
- ✅ Day 21: **Keypad Navigation with BFS & DP** - Using Breadth-First Search to find shortest paths between keypad numbers, then applying dynamic programming with memoization to compute optimal movement sequences.
- ✅ Day 22: **Pattern Recognition & Bitwise Operations** - Computing "secret numbers" through bitwise operations and detecting patterns in sequences of digits to find the most valuable combinations.
- ✅ Day 23: **Bron-Kerbosch Maximal Clique Algorithm** - Implementing Bron-Kerbosch algorithm to find maximal cliques in an undirected graph, working with interconnected sets to identify the largest complete subgraph.
- ✅ Day 24: **Circuit Simulation & Logical Gate Analysis** - Simulating a ripple carry adder circuit with recursive evaluation of wire values through logical gates (AND, OR, XOR) to identify incorrect wire configurations.
- ✅ Day 25: **Geometrical Constraints & Combinatorial Counting** - Computing heights of lock and key matrices and counting valid combinations where combined heights don't exceed specified thresholds.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Note that while this code is licensed under MIT, the [Advent of Code](https://adventofcode.com/) puzzles themselves are copyrighted by Eric Wastl. This repository contains only my solutions to these puzzles.

## 🔗 Links

- [Advent of Code](https://adventofcode.com/)
- [Go Programming Language](https://golang.org/)