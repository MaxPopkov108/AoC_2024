package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{0, -1}, // left
	{1, 0},  //down
	{0, 1},  // right
	{-1, 0}, // up
}

var visited = make(map[[3]int]bool)

var trailMap = make(map[[3]int][][3]int)

func solution(file *os.File) (int, int) {
	var hights [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempArr []int
		line := scanner.Text()
		for _, char := range line {
			num := int(char - '0')
			tempArr = append(tempArr, num)
		}
		hights = append(hights, tempArr)
	}

	for row := 0; row < len(hights); row++ {
		for col := 0; col < len(hights[row]); col++ {
			lookAround(hights, row, col)
		}
	}

	trailCountSum := 0
	trailCountSum2 := 0
	for row := 0; row < len(hights); row++ {
		for col := 0; col < len(hights[row]); col++ {
			if hights[row][col] == 0 {
				visited = make(map[[3]int]bool)
				trailCountSum += countTrails(row, col, 0)
				trailCountSum2 += countTrails2(row, col, 0)
			}
		}
	}

	return trailCountSum, trailCountSum2
}

func lookAround(hights [][]int, row, col int) {
	for _, direction := range directions {
		currentPos := hights[row][col]
		newRow, newCol := row+direction[0], col+direction[1]
		if newRow < 0 || newRow >= len(hights) || newCol < 0 || newCol >= len(hights[row]) {
			continue
		}
		nextPos := hights[row+direction[0]][col+direction[1]]
		if currentPos < nextPos && nextPos-currentPos == 1 {
			trailMap[[3]int{row, col, currentPos}] = append(
				trailMap[[3]int{row, col, currentPos}], [3]int{row + direction[0], col + direction[1], nextPos},
			)
		}
	}
}

func countTrails(row, col, hight int) int {
	key := [3]int{row, col, hight}

	if visited[key] {
		return 0
	}

	visited[key] = true

	if hight == 9 {
		return 1
	}

	trails, exists := trailMap[[3]int{row, col, hight}]
	if !exists {
		visited[key] = false
		return 0
	}

	trailCount := 0
	for _, trail := range trails {
		trailCount += countTrails(trail[0], trail[1], trail[2])
	}

	return trailCount
}

func countTrails2(row, col, hight int) int {
	if hight == 9 {
		return 1
	}

	trails, exists := trailMap[[3]int{row, col, hight}]
	if !exists {
		return 0
	}

	trailCount := 0
	for _, trail := range trails {
		trailCount += countTrails2(trail[0], trail[1], trail[2])
	}

	return trailCount
}

func main() {
	file, _ := os.Open("day_10_Hoof_It/input.txt")
	defer file.Close()

	firstPart, secondPart := solution(file)
	fmt.Printf("Part 1: %d\n", firstPart)
	fmt.Printf("Part 2: %d\n", secondPart)
}
