package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solution(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)
	multResult1 := 0
	multResult2 := 0
	var longline string
	for scanner.Scan() {
		line := scanner.Text()
		longline += line
		multResult1 += parseAndMult(line)
	}

	matches := regexp.MustCompile(`(?s)`+regexp.QuoteMeta("don't()")+`(.*?)(`+regexp.QuoteMeta("do()")+`|$)`).FindAllString(longline, -1)
	for _, match := range matches {
		longline = strings.Replace(longline, match, "", 1)
	}
	multResult2 += parseAndMult(longline)
	return multResult1, multResult2
}

func parseAndMult(line string) int {
	multResult := 0
	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		multResult += num1 * num2
	}
	return multResult
}

func main() {
	file, err := os.Open("day_3_Mull_It_Over/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	firstPart, secondPart := solution(file)

	fmt.Println("1st part:", firstPart)
	fmt.Println("2nd part:", secondPart)
}
