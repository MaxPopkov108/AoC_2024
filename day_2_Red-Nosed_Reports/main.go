package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solution(file *os.File) (int, int) {
	safeLevelsFirstpart := 0
	safeLevelsSecondpart := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stringNums := strings.Fields(line)

		var nums []int

		for _, strNum := range stringNums {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Printf("Error parsing field '%v' in line '%v': %v\n", strNum, line, err)
				continue
			}
			nums = append(nums, num)
		}

		if checkStableLevelChange(nums, false)+checkLevelDifference(nums, false) == 0 {
			safeLevelsFirstpart++
		}

		if checkStableLevelChange(nums, true)+checkLevelDifference(nums, true) == 0 {
			safeLevelsSecondpart++
		}
	}

	return safeLevelsFirstpart, safeLevelsSecondpart
}

func checkStableLevelChange(levels []int, damper bool) int {
	prevMarker := levels[0] > levels[1]
	errorCounter := 0
	for i := 1; i < len(levels); i++ {
		if i+1 < len(levels) {
			marker := levels[i] > levels[i+1]
			if prevMarker != marker {
				errorCounter++

				if damper == true {
					damper = false

					damperStability1 := 0
					damperStability2 := 0
					damperStability3 := 0

					levels1 := removeElement(levels, i-1)
					damperStability1 += checkStableLevelChange(levels1, false) + checkLevelDifference(levels1, false)

					levels2 := removeElement(levels, i)
					damperStability2 += checkStableLevelChange(levels2, false) + checkLevelDifference(levels2, false)

					levels3 := removeElement(levels, i+1)
					damperStability3 += checkStableLevelChange(levels3, false) + checkLevelDifference(levels3, false)

					if damperStability1*damperStability2*damperStability3 != 0 {
						return 100
					}

					return 0
				}
			}
		}
	}
	return errorCounter
}

func checkLevelDifference(levels []int, damper bool) int {
	errorCounter := 0
	for i := 0; i < len(levels); i++ {
		if i+1 < len(levels) {
			levelChange := int(math.Abs(float64(levels[i] - levels[i+1])))
			if levelChange < 1 || levelChange > 3 {
				errorCounter++
				if damper == true {
					damper = false

					damperStability1 := 0
					damperStability2 := 0
					damperStability3 := 0

					levels1 := removeElement(levels, i-1)
					damperStability1 += checkStableLevelChange(levels1, false) + checkLevelDifference(levels1, false)

					levels2 := removeElement(levels, i)
					damperStability2 += checkStableLevelChange(levels2, false) + checkLevelDifference(levels2, false)

					levels3 := removeElement(levels, i+1)
					damperStability3 += checkStableLevelChange(levels3, false) + checkLevelDifference(levels3, false)

					if damperStability1*damperStability2*damperStability3 != 0 {
						return 100
					}

					return 0
				}
			}
		}
	}
	return errorCounter
}

func removeElement(levels []int, index int) []int {
	if index < 0 || index >= len(levels) {
		return levels
	}
	result := make([]int, 0, len(levels)-1)
	result = append(result, levels[:index]...)
	return append(result, levels[index+1:]...)
}

func main() {
	file, err := os.Open("day_2_Red-Nosed_Reports/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	firstPart, secondPart := solution(file)

	fmt.Println("1st part result:", firstPart)
	fmt.Println("2nd part result:", secondPart)
}
