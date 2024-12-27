package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

type Equation struct {
	a int
	b int
	c int
}

func main() {
	log.Printf("%v days to XMASSS", 24-14)

	input, err := utils.FileToEquations("input.txt")
	// input test -> 480
	// input test 1 -> 14

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	total := getTotalTokensToWin(input)
	
	fmt.Println(total)
}

func getTotalTokensToWin(equationMap map[int][][]int) int {
	total := 0

	for _, eqList := range equationMap{
		total += getTokensForPrize(eqList)
	}
	
	return total
}

func getTokensForPrize(equationList [][]int) int {
	pushesOnA := 0
	pushesOnB := 0

	eq1 := Equation{equationList[0][0], equationList[1][0], equationList[2][0]}
	eq2 := Equation{equationList[0][1], equationList[1][1], equationList[2][1]}

	pushesOnA, pushesOnB = solveEquationsAndGetXAndYPartTwo(eq1, eq2)

	return (pushesOnA*3) + (pushesOnB*1)
}

// Part 1
func solveEquationsAndGetXAndY(eq1 Equation, eq2 Equation) (int, int){
	x := 0
	y := 0

    // Using Cramer's rule: https://en.wikipedia.org/wiki/Cramer%27s_rule
    numerator1 := float64(eq1.c*eq2.b - eq1.b*eq2.c)
    denominator1 := float64(eq1.a*eq2.b - eq1.b*eq2.a)

    numerator2 := float64(eq1.a*eq2.c - eq1.c*eq2.a)
    denominator2 := float64(eq1.a*eq2.b - eq1.b*eq2.a)

    xFloat := numerator1 / denominator1
    yFloat := numerator2 / denominator2

	// Only int works
	if xFloat == float64(int(xFloat)) && yFloat == float64(int(yFloat)) {
		x = int(xFloat)
		y = int(yFloat)
	} 

    if x < 0 || x > 100 || y < 0 || y > 100 {
        return 0, 0
    }

    return x, y
}

// Part 2
func solveEquationsAndGetXAndYPartTwo(eq1 Equation, eq2 Equation) (int, int){
	x := 0
	y := 0

    // Using Cramer's rule: https://en.wikipedia.org/wiki/Cramer%27s_rule
    numerator1 := float64((eq1.c+10000000000000)*eq2.b - eq1.b*(eq2.c+10000000000000))
    denominator1 := float64(eq1.a*eq2.b - eq1.b*eq2.a)

    numerator2 := float64(eq1.a*(eq2.c+10000000000000) - (eq1.c+10000000000000)*eq2.a)
    denominator2 := float64(eq1.a*eq2.b - eq1.b*eq2.a)

    xFloat := numerator1 / denominator1
    yFloat := numerator2 / denominator2

	// Only int works
	if xFloat == float64(int(xFloat)) && yFloat == float64(int(yFloat)) {
		x = int(xFloat)
		y = int(yFloat)
	} 

    return x, y
}
