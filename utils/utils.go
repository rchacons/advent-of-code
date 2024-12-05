package utils

import (
	"bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
	"regexp"
	"log"
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

// FileToText reads a file and returns its content as a string
func FileToText(filePath string) (string, error) {
	
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}

// FileToText reads a file and returns its content as a string
func FileToTextWithSpaces(filePath string) (string, error) {
	
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var text string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}

// FileToStringMatrix reads a file and returns a matrix of strings 
func FileToStringMatrix(filePath string) ([][]string, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var matrix [][]string
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []string{})
		for i := range line {
			matrix[row] = append(matrix[row], string(line[i]))
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

// TextToRules reads an input text and return the rule ([3 5], [1 4] for example) for day 5 puzzle
func TextToRules(text string) [][]int{
	pattern1 := regexp.MustCompile(`\d{1,2}\|\d{1,2}`)
	pattern2 := regexp.MustCompile(`\d{1,2}`)
	lines := pattern1.FindAllString(text, -1)

	var numbers [][]int
	for _, line := range lines {
		numbersString := pattern2.FindAllString(line, -1)
		number1,err1 := strconv.Atoi(numbersString[0])
		number2,err2 := strconv.Atoi(numbersString[1])
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err1)
		}
		numbers = append(numbers, []int{number1,number2})
	}

	return numbers
}

// TextToListOfNumbers reads an input text and returns a list of numbers separated by commas
func TextToListOfNumbers(text string) [][]int{
	patternNotToMatch := regexp.MustCompile(`\d{1,2}\|\d{1,2}`)
	patternComma := regexp.MustCompile(`\d+`)
	lines := strings.Split(text, "\n")

	var numbers [][]int
	for _, line := range lines {
		
		if line != "" && 
		!patternNotToMatch.MatchString(line) {
			numberString := patternComma.FindAllString(line, -1)

			var lineNumbers []int
			for _, numStr := range numberString {
				number, err := strconv.Atoi(numStr)		
				if err != nil {
					log.Fatal(err)
				}
				lineNumbers = append(lineNumbers, number)
			}
			numbers = append(numbers, lineNumbers)
		}
		
	}
	return numbers
}
