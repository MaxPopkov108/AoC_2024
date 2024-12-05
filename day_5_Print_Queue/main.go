package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solution(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	correctUpdates := 0
	inCorrectUpdates := 0

	var isNextChunk bool
	var printQueue [][]int
	orderDict := make(map[int][]int)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isNextChunk = true
			continue
		}

		if isNextChunk {
			var tempArr []int
			nums := strings.Split(line, ",")
			for _, strNum := range nums {
				num, _ := strconv.Atoi(strNum)
				tempArr = append(tempArr, num)
			}
			printQueue = append(printQueue, tempArr)
		} else {
			twoPageChunk := strings.Split(line, "|")
			primePage, _ := strconv.Atoi(twoPageChunk[0])
			afterPage, _ := strconv.Atoi(twoPageChunk[1])
			orderDict[primePage] = append(orderDict[primePage], afterPage)
		}
	}

	var unSafeUpdates [][]int

	for _, arr := range printQueue {
		var isSafe bool
		for i := 1; i < len(arr); i++ {
			if slices.Contains(orderDict[arr[i-1]], arr[i]) {
				isSafe = true
			} else {
				isSafe = false
				unSafeUpdates = append(unSafeUpdates, arr)
				break
			}
		}

		if isSafe {
			correctUpdates += arr[len(arr)/2]
		}
	}

	for _, arr := range unSafeUpdates {
		for i := 1; i < len(arr); i++ {
			if !slices.Contains(orderDict[arr[i-1]], arr[i]) {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				i = 0
			}
		}

		inCorrectUpdates += arr[len(arr)/2]
	}

	return correctUpdates, inCorrectUpdates
}

func main() {
	file, err := os.Open("day_5_Print_Queue/test_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	firstPart, secondPart := solution(file)

	fmt.Println("First part:", firstPart)
	fmt.Println("Second part:", secondPart)
}
