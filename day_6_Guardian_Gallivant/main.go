package main

import (
	"bufio"
	"fmt"
	"os"
)

var uniqLcs = make(map[[2]int]bool)

var stepBack = map[[2]int][2]int{
	{-1, 0}: {1, 0},
	{0, 1}:  {0, -1},
	{1, 0}:  {-1, 0},
	{0, -1}: {0, 1},
}

var directionsArr = [][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func solution(file *os.File) (int, int) {
	var matrix [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempRow []string
		line := scanner.Text()
		for _, char := range line {
			tempRow = append(tempRow, string(char))
		}

		matrix = append(matrix, tempRow)
	}

	guardRow, guardCol := findGuardPos(matrix)

	uniqLcs[[2]int{guardRow, guardCol}] = true

	move(matrix, guardRow, guardCol)

	return len(uniqLcs), moveAndPlace(matrix, guardRow, guardCol)
}

func findGuardPos(matrix [][]string) (int, int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			symbol := matrix[row][col]
			if symbol == "^" {
				return row, col
			}
		}
	}

	return -1, -1
}

func move(matrix [][]string, row, col int) {
	currentDirectionIdx := 0
	for {
		row += directionsArr[currentDirectionIdx][0]
		col += directionsArr[currentDirectionIdx][1]

		if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
			break
		}

		if matrix[row][col] == "#" {
			key, _ := stepBack[directionsArr[currentDirectionIdx]]
			row += key[0]
			col += key[1]
			currentDirectionIdx = (currentDirectionIdx + 1) % len(directionsArr)

			continue
		}

		uniqLcs[[2]int{row, col}] = true
	}
}

func moveAndPlace(matrix [][]string, grdRow, grdCol int) int {
	zaLoopCounter := 0

	for posDir := range uniqLcs {

		visited := make(map[[3]int]bool)
		currentDirectionIdx := 0
		row, col := grdRow, grdCol
		obstRow, obstCol := posDir[0], posDir[1]
		matrix[obstRow][obstCol] = "#"

		for {

			row += directionsArr[currentDirectionIdx][0]
			col += directionsArr[currentDirectionIdx][1]

			if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
				matrix[obstRow][obstCol] = "."
				break
			}

			if visited[[3]int{row, col, currentDirectionIdx}] {
				zaLoopCounter++
				matrix[obstRow][obstCol] = "."
				break
			}

			if matrix[row][col] == "#" {
				key, _ := stepBack[directionsArr[currentDirectionIdx]]
				row += key[0]
				col += key[1]

				currentDirectionIdx = (currentDirectionIdx + 1) % len(directionsArr)

				continue
			}

			visited[[3]int{row, col, currentDirectionIdx}] = true
		}

	}

	return zaLoopCounter
}

func main() {
	file, err := os.Open("day_6_Guardian_Gallivant/input.txt")
	if err != nil {
		fmt.Println("Cannot open file:", err)
	}
	defer file.Close()

	firstPart, secondPart := solution(file)

	fmt.Printf("Part 1: %d \nPart 2: %d\n", firstPart, secondPart)

}
