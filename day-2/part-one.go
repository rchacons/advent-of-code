package main

import (
	"fmt"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	reports, err := utils.FileToReportsLists("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var safeCount int
	for _, report := range reports {
		var safe bool

		if (report[0] < report[1]){
			safe = isReportSafe(report, true)
		} else {
			safe = isReportSafe(report, false)
		}
		
		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

// Part one : isReportSafe checks if the report is safe
func isReportSafe(report []int, increasing bool) bool {

	for i := 1; i < len(report); i++ {
		if(!isDiffValid(report[i], report[i-1], increasing)){
			return false
		}
	
	}
	return true
}



func isDiffValid(num1 int, num2 int, incr bool) bool {
	difference := num1 - num2
	if !incr {
		difference *= -1
	}
	if difference < 1 || difference > 3{
		return false
	}
	return true
}
