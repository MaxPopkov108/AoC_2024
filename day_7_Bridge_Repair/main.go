package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solution(file *os.File) (int, int) {

	var equations [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempArr []int
		line := scanner.Text()
		nums := regexp.MustCompile("[0-9]+").FindAllString(line, -1)
		for _, char := range nums {
			num, _ := strconv.Atoi(char)
			tempArr = append(tempArr, num)
		}
		equations = append(equations, tempArr)
	}

	solutionSum := 0
	concatSum := 0

	var deconcatCandidates [][]int

	for _, equation := range equations {
		if solve(equation[0], equation, len(equation)-1, false) {
			solutionSum += equation[0]
			concatSum += equation[0]
		} else {
			deconcatCandidates = append(deconcatCandidates, equation)
		}
	}

	for _, equation := range deconcatCandidates {
		if solve(equation[0], equation, len(equation)-1, true) {
			concatSum += equation[0]
		}
	}

	return solutionSum, concatSum
}

func solve(bigNum int, nums []int, index int, deconcat bool) bool {
	if index == 0 {
		return false
	}

	currentNum := nums[index]

	if currentNum != 0 && bigNum%currentNum == 0 {
		divResult := bigNum / currentNum
		if divResult == 1 && index == 1 {
			return true
		}
		if solve(divResult, nums, index-1, deconcat) {
			return true
		}
	}

	if bigNum-currentNum >= 0 {
		subResult := bigNum - currentNum
		if subResult == 0 && index == 1 {
			return true
		}
		if solve(subResult, nums, index-1, deconcat) {
			return true
		}
	}

	if deconcat {
		bigNumStr := strconv.Itoa(bigNum)
		currentNumStr := strconv.Itoa(currentNum)

		if len(bigNumStr) > len(currentNumStr) && bigNumStr[len(bigNumStr)-len(currentNumStr):] == currentNumStr {
			remainingStr := bigNumStr[:len(bigNumStr)-len(currentNumStr)]
			remainingNum, _ := strconv.Atoi(remainingStr)
			if solve(remainingNum, nums, index-1, deconcat) {
				return true
			}
		}
	}

	return false
}

func main() {
	file, _ := os.Open("day_7_Bridge_Repair/input.txt")
	defer file.Close()

	firstPart, secondPart := solution(file)
	fmt.Printf("Part 1: %d\n", firstPart)
	fmt.Printf("Part 2: %d\n", secondPart)
}
