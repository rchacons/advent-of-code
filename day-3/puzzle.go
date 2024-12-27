package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/rchacons/advent-of-code/utils"
)

func main() {
	text, err := utils.FileToText("input.txt")
	// text_part_one := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	text_part_two := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	if err != nil {
		fmt.Println(err)
		fmt.Println(text_part_two)
	}
	// Part one
	mulInString := extractMul(text)
	totalPartOne := extractNumbersAndMultiply(mulInString)

	// Part two
	totalPartTwo := extractDoAndDont(text)  	
	fmt.Println("Total part one : ", totalPartOne)
	fmt.Println("Total part two : ", totalPartTwo)


}

func extractMul(input string) []string {
	pattern := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	return pattern.FindAllString(input, -1)
}

func extractNumbersAndMultiply(input []string) int {
	pattern := regexp.MustCompile(`\d{1,3}`)
	var total int
	for _, mul := range input {
		numbersString := pattern.FindAllString(mul,-1)
		number1,err1 := strconv.Atoi(numbersString[0])
		number2,err2 := strconv.Atoi(numbersString[1])
		if err1 != nil {
			log.Fatal(err1)
		}
		if err2 != nil {
			log.Fatal(err1)
		}
		total += number1*number2
	}
	return total
}

func extractDoAndDont(input string) int {
	var total int
	pattern1 := regexp.MustCompile(`(do\(\))`)
	pattern2 := regexp.MustCompile(`(don't\(\))`)

	indexListDo := pattern1.FindAllStringIndex(input, -1)
	indexListDo = append([][]int{{0, 0}}, indexListDo...)

	indexListDont := pattern2.FindAllStringIndex(input, -1)

    p1 := 0
    p2 := 0
    lastIndex := 0

    for p1 < len(indexListDo) && p2 < len(indexListDont) {
        index1 := indexListDo[p1][1]
        index2 := indexListDont[p2][0]


        if index1 < index2 {
            for _, num := range indexListDo[p1:] {					
                if num[0] > index2 {
                    subStr := input[lastIndex:index2]
                    mulInStr := extractMul(subStr)
                    total += extractNumbersAndMultiply(mulInStr)
                    lastIndex = indexListDo[p1][0]
                    break
                } else {
                    p1++
                }
            }
        } else {
            p2++
        }
    }

    if p1 < len(indexListDo) {
        subStr := input[indexListDo[p1][0]:]
        mulInStr := extractMul(subStr)
        total += extractNumbersAndMultiply(mulInStr)
    }
	
	return total
}