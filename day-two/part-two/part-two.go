package main

import (
	"fmt"
	"log"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	
	reports, err := utils.FileToReportsLists("../input.txt")
	reports_test := [][]int{
        {1, 4, 2, 3, 4},
		{5, 7, 6, 8, 9},
		{2, 4, 3, 2, 1},
		{3, 4, 2, 1},
		{2, 4, 3, 5, 6},
		{5, 8, 15, 11, 12},
		{17, 25, 26, 27},
		{17, 25, 18, 19},
		 {5, 8, 11, 14, 16, 19, 20, 26},
		{-1, 4, 5, 6, 9},
		{-6, -7, -8, -9, -10},
		{-6, -7, 15, -9, 12},
        {7, 6, 4, 2, 1},  // Safe without removing any level.
        {1, 2, 7, 8, 9},  // Unsafe regardless of which level is removed.
        {9, 7, 6, 2, 1},  // Unsafe regardless of which level is removed.
        {1, 3, 2, 4, 5},  // Safe by removing the second level, 3.
        {8, 6, 4, 4, 1},  // Safe by removing the third level, 4.
        {1, 3, 6, 7, 9},  // Safe without removing any level.
	}

    if err != nil {
        log.Println(err)
		log.Println(reports_test)
        return
	}
    


	var safeCount int
	for _, report := range reports {
		var safe bool

		if (isReportInIncOrder(report)){
			safe = isReportSafeWithSingleBadTolerance(report, true)
		} else {
			safe = isReportSafeWithSingleBadTolerance(report, false)
		}
		
		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isReportInIncOrder(report []int) bool {
	posCount := 0
	negCount := 0

	for i := 1; i < len(report); i++ {
		if report[i-1] < report[i] {
			posCount++
		} else{
			negCount++
		}
	}

	return posCount > negCount
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

// Part two : isReportSafeWithSingleBadTolerance checks if the report is safe having one single bad tolerance (i brutforced it)
func isReportSafeWithSingleBadTolerance(report []int, incr bool) bool {
    p1 := 0
    p2 := 1

    for p2 < len(report) {
        if !isDiffValid(report[p2], report[p1], incr) {
            // brutforce :(

            report1 := append([]int(nil), report[:p1]...)
            report1 = append(report1, report[p1+1:]...)

            report2 := append([]int(nil), report[:p2]...)
            report2 = append(report2, report[p2+1:]...)

            if isReportSafe(report1, incr) || isReportSafe(report2, incr) {
                return true
            } else {
                return false
            }
        }
        p1++
        p2++
    }
    return true
}


func isDiffValid(num1 int, num2 int, incr bool) bool {
    difference := num1 - num2
    if !incr {
        difference *= -1
    }
    if difference < 1 || difference > 3 {
        log.Printf("Difference %d is out of valid range (1-3)\n", difference)
        return false
    }
    log.Printf("Difference %d is within valid range (1-3)\n", difference)
    return true
}
