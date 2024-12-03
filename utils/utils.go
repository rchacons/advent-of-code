package utils

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func PrintHello() {
	fmt.Println("Hello")
}

// FileToLists reads a file and returns two lists of integers, each corresponding to a column
func FileToLists(filePath string) ([]int, []int, error) {
	
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if(len(parts) != 2) {
			return nil, nil, fmt.Errorf("invalid line: %s", line)
		}
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil

}

// FileTOReportsLists reads a file and returns a list of reports (each report is a list of integers)
func FileToReportsLists(filePath string) ([][]int, error) {
	
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		line := scanner.Text()
		parts := strings.Fields(line)
		
		for i := range parts {
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				return nil, err
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}