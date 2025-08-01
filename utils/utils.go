package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		if len(parts) != 2 {
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
func TextToRules(text string) [][]int {
	pattern1 := regexp.MustCompile(`\d{1,2}\|\d{1,2}`)
	pattern2 := regexp.MustCompile(`\d{1,2}`)
	lines := pattern1.FindAllString(text, -1)

	var numbers [][]int
	for _, line := range lines {
		numbersString := pattern2.FindAllString(line, -1)
		number1, err1 := strconv.Atoi(numbersString[0])
		number2, err2 := strconv.Atoi(numbersString[1])
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err1)
		}
		numbers = append(numbers, []int{number1, number2})
	}

	return numbers
}

// TextToListOfNumbers reads an input text and returns a list of numbers separated by commas
func TextToListOfNumbers(text string) [][]int {
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

func FileToIntegerLists(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numberList [][]int
	patternNumber := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var numbers []int
		line := scanner.Text()
		numberString := patternNumber.FindAllString(line, -1)

		for i := range numberString {
			// fmt.Println(numberString[i])
			num, err := strconv.Atoi(numberString[i])
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
		numberList = append(numberList, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numberList, nil
}

func FileToIntList(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {

		line := scanner.Text()

		for i := range line {
			num := int(line[i] - '0')
			numbers = append(numbers, num)
		}
	}
	return numbers, nil
}

func FileToIntMatrix(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numberMatrix [][]int
	for scanner.Scan() {
		var numbers []int

		line := scanner.Text()

		for i := range line {
			num := int(line[i] - '0')
			numbers = append(numbers, num)
		}
		numberMatrix = append(numberMatrix, numbers)
	}
	return numberMatrix, nil
}

func FileToEquations(filePath string) (map[int][][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	/**
	The idea is to return a map where every value is a system of two equations + values of X and Y
	For ex :
	Given the input :
	Button A: X+94, Y+34
	Button B: X+22, Y+67
	Prize: X=8400, Y=5400

	The result would be:
	{1 = [[94 34], [22, 67] [8400,5400]]}
	**/
	equationMap := make(map[int][][]int)
	count := 0 // keep track of each line
	equationSysCount := 1

	var numberList [][]int
	patternNumber := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		if count == 3 {
			equationMap[equationSysCount] = numberList
			numberList = make([][]int, 0)
			equationSysCount++
			count = 0
		}
		var numbers []int
		line := scanner.Text()
		numberString := patternNumber.FindAllString(line, -1)
		if numberString != nil {

			for i := range numberString {
				num, err := strconv.Atoi(numberString[i])
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, num)
			}
			numberList = append(numberList, numbers)
			count++
		}
	}
	equationMap[equationSysCount] = numberList // The last one

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return equationMap, nil

}

func FileToIntegerListsWithNegatives(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numberList [][]int
	patternNumber := regexp.MustCompile(`\d+|-\d+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var numbers []int
		line := scanner.Text()
		numberString := patternNumber.FindAllString(line, -1)

		for i := range numberString {
			//fmt.Println(numberString[i])
			num, err := strconv.Atoi(numberString[i])
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
		numberList = append(numberList, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numberList, nil
}

// For day 15
func FileToMapAndRobotMouvements(filePath string) ([][]string, []string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	var puzzleMap [][]string
	var mouvements []string

	pattern1 := regexp.MustCompile(`([\s\S]+?)\n\n`)
	pattern2 := regexp.MustCompile(`\n\n([\s\S]+)`)

	fileStr := string(file)

	match1 := pattern1.FindAllString(fileStr, -1)
	match2 := pattern2.FindAllString(fileStr, -1)

	lines1 := strings.Split(match1[0], "\n")
	for lineIndex := range lines1 {
		line := make([]string, 0)
		for letterIndex := range lines1[lineIndex] {
			line = append(line, string(lines1[lineIndex][letterIndex]))
		}
		if len(line) > 1 {
			puzzleMap = append(puzzleMap, line)
		}
	}

	for strIndex := range match2[0] {
		if match2[0][strIndex] != 10 {
			mouvements = append(mouvements, string(match2[0][strIndex]))
		}
	}

	return puzzleMap, mouvements, nil
}

// for debugging
func PrintStringMatrix(stringMatrix [][]string) {
	for _, row := range stringMatrix {
		for _, val := range row {
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}
}

// For day 17
func FileToProgramInstructions(filePath string) (map[string]interface{}, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileStr := string(file)

	patternNumber := regexp.MustCompile(`(\d+)`)
	match := patternNumber.FindAllString(fileStr, -1)

	programMap := make(map[string]interface{})

	programMap["A"], err = strconv.Atoi(match[0])
	if err != nil {
		return nil, err
	}

	programMap["B"], err = strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}

	programMap["C"], err = strconv.Atoi(match[2])
	if err != nil {
		return nil, err
	}

	var numbers []int
	for i := 3; i < len(match); i++ {
		//fmt.Println(numberString[i])
		num, err := strconv.Atoi(match[i])
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, num)
	}
	programMap["P"] = numbers

	return programMap, nil

}

// TextToBytePosition reads an input text and return the position of bytes ([3 5], [1 4] for example) for day 18 puzzle
func TextToBytePosition(text string) [][]int {
	pattern1 := regexp.MustCompile(`\d{1,2}\,\d{1,2}`)
	pattern2 := regexp.MustCompile(`\d{1,2}`)
	lines := pattern1.FindAllString(text, -1)

	var numbers [][]int
	for _, line := range lines {
		numbersString := pattern2.FindAllString(line, -1)
		number1, err1 := strconv.Atoi(numbersString[0])
		number2, err2 := strconv.Atoi(numbersString[1])
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err1)
		}
		numbers = append(numbers, []int{number1, number2})
	}

	return numbers
}

// Allows to extract the first line as a slice for each pattern, and the second part as a [][]sting (day 19)
func FileToPatternSlices(filePath string) ([]string, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var patternList []string
	var wordList []string
	//patternNumber := regexp.MustCompile(`\d+|-\d+`)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	// Extract first line
	patternList = strings.Split(lines[0], ", ")

	// Extract the rest of words
	wordList = append(wordList, lines[2:]...)

	return patternList, wordList, nil
}

func FileToStringList(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stringList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringList = append(stringList, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stringList, nil
}

func FileToIntegerList(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	patternNumber := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		numberString := patternNumber.FindAllString(line, -1)

		for i := range numberString {
			// fmt.Println(numberString[i])
			num, err := strconv.Atoi(numberString[i])
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

// Function to read a pair of string of this format kh-tc and return a map with all neighbors
func FileToLanMap(filePath string) (map[string]map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	pairs := make(map[string]map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			tmpMap1, exists1 := pairs[parts[0]]
			if !exists1 {
				tmpMap1 = make(map[string]bool)
			}
			tmpMap1[parts[1]] = true
			pairs[parts[0]] = tmpMap1

			tmpMap2, exists2 := pairs[parts[1]]
			if !exists2 {
				tmpMap2 = make(map[string]bool)
			}
			tmpMap2[parts[0]] = true
			pairs[parts[1]] = tmpMap2
		}
	}
	if err := scanner.Err(); err != nil {
        return nil, err
    }

	return pairs, nil
}

// first get the first part by splitting the (": ")
// so store x,y in a map {x: y}
// then get the second part by rpelacing (" -> ") with ("") and then splitting
// and store the x, op, y, z in a map where {z: {x, op, y}}
func FileToFormulasMaps(filePath string) (map[string]int, map[string][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

    formulas := make(map[string]int)
    operations := make(map[string][]string)
    scanner := bufio.NewScanner(file)
    parsingFormulas := true

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            parsingFormulas = false
            continue
        }

        if parsingFormulas {
            parts := strings.Split(line, ": ")
            if len(parts) == 2 {
                value, err := strconv.Atoi(parts[1])
                if err != nil {
                    return nil, nil, err
                }
                formulas[parts[0]] = value
            }
        } else {
            replacedLine := strings.Replace(line, " -> ", " ", 1)
			parts := strings.Split(replacedLine, " ")
            if len(parts) == 4 {
                operations[parts[3]] = []string{parts[0], parts[1], parts[2]}
            }
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, err
    }

    return formulas, operations, nil
}

// For day 25
func FileToMapOfLocksAndKeysMatrix(filePath string) (map[string][][][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()


	var matrixList [][][]string
	var matrix [][]string
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			matrixList = append(matrixList, matrix)
			matrix = make([][]string, 0)
			row = 0
		} else {
			matrix = append(matrix, []string{})
			for i := range line {
				matrix[row] = append(matrix[row], string(line[i]))
			}
			row++
		}
	}

	// Append the last matrix if the file doesn't end with an empty line
	if len(matrix) > 0 {
		matrixList = append(matrixList, matrix)
	}
	
	matrixMap := make(map[string][][][]string)

	for _, matrix := range matrixList {
		count := 0
		for _, col := range matrix[0] {
			if col == "#" {count++}
		}
		if count == 5 { 
			// found a lock
			// add it to the list of matrix for the lock or create one if it does not exist
			matrixMap["lock"] = append(matrixMap["lock"], matrix)
		} else {
			// found a key
			// add it to the list of matrix for the key or create one if it does not exist
			matrixMap["key"] = append(matrixMap["key"], matrix)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrixMap, nil
}


// PrettyPrintMap prints the map[string][][][]string in a readable format
func PrettyPrintMap(data map[string][][][]string) {
    for key, matrices := range data {
        fmt.Printf("Key: %s\n", key)
        for i, matrix := range matrices {
            fmt.Printf(" Matrix %d:\n", i+1)
            for _, row := range matrix {
                fmt.Println(strings.Join(row, ""))
            }
            fmt.Println()
        }
    }
}