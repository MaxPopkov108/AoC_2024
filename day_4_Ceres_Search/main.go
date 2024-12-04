package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solution(file *os.File) (int, int) {
	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	rowLen := len(matrix[0])
	colLen := len(matrix)

	xmasCount := 0
	crossMasCount := 0

	for col := 0; col < colLen; col++ {
		for row := 0; row < rowLen; row++ {

			if matrix[row][col] == 'X' {

				if col > 2 {
					xmasCount += lookLeft(matrix, row, col)
				}

				if col < colLen-3 {
					xmasCount += lookRight(matrix, row, col)
				}

				if row > 2 {
					xmasCount += lookUp(matrix, row, col)
				}

				if row < rowLen-3 {
					xmasCount += lookDown(matrix, row, col)
				}

				if col > 2 && row < rowLen-3 {
					xmasCount += lookLeftDown(matrix, row, col)
				}

				if col < colLen-3 && row < rowLen-3 {
					xmasCount += lookRightDown(matrix, row, col)
				}

				if col > 2 && row > 2 {
					xmasCount += lookLeftUp(matrix, row, col)
				}

				if col < colLen-3 && row > 2 {
					xmasCount += lookRightUp(matrix, row, col)
				}
			}

			if matrix[row][col] == 'A' && row > 0 && row < rowLen-1 && col > 0 && col < colLen-1 {
				crossMasCount += lookForX(matrix, row, col)
			}

		}
	}
	return xmasCount, crossMasCount
}

func lookRight(matrix [][]rune, row int, col int) int {
	if matrix[row][col+1] == 'M' && matrix[row][col+2] == 'A' && matrix[row][col+3] == 'S' {
		return 1
	}
	return 0
}

func lookLeft(matrix [][]rune, row int, col int) int {
	if matrix[row][col-1] == 'M' && matrix[row][col-2] == 'A' && matrix[row][col-3] == 'S' {
		return 1
	}
	return 0
}

func lookDown(matrix [][]rune, row int, col int) int {
	if matrix[row+1][col] == 'M' && matrix[row+2][col] == 'A' && matrix[row+3][col] == 'S' {
		return 1
	}
	return 0
}

func lookUp(matrix [][]rune, row int, col int) int {
	if matrix[row-1][col] == 'M' && matrix[row-2][col] == 'A' && matrix[row-3][col] == 'S' {
		return 1
	}
	return 0
}

func lookRightUp(matrix [][]rune, row int, col int) int {
	if matrix[row-1][col+1] == 'M' && matrix[row-2][col+2] == 'A' && matrix[row-3][col+3] == 'S' {
		return 1
	}
	return 0
}

func lookLeftUp(matrix [][]rune, row int, col int) int {
	if matrix[row-1][col-1] == 'M' && matrix[row-2][col-2] == 'A' && matrix[row-3][col-3] == 'S' {
		return 1
	}
	return 0
}

func lookRightDown(matrix [][]rune, row int, col int) int {
	if matrix[row+1][col+1] == 'M' && matrix[row+2][col+2] == 'A' && matrix[row+3][col+3] == 'S' {
		return 1
	}
	return 0
}

func lookLeftDown(matrix [][]rune, row int, col int) int {
	if matrix[row+1][col-1] == 'M' && matrix[row+2][col-2] == 'A' && matrix[row+3][col-3] == 'S' {
		return 1
	}
	return 0
}

func lookForX(matrix [][]rune, row int, col int) int {
	//	M - S
	//	- A -
	//	M - S
	if matrix[row-1][col-1] == 'M' && matrix[row+1][col-1] == 'M' {
		if matrix[row+1][col+1] == 'S' && matrix[row-1][col+1] == 'S' {
			return 1
		}
	}

	//	S - M
	//	- A -
	//	S - M
	if matrix[row+1][col+1] == 'M' && matrix[row-1][col+1] == 'M' {
		if matrix[row-1][col-1] == 'S' && matrix[row+1][col-1] == 'S' {
			return 1
		}
	}

	//	S - S
	//	- A -
	//	M - M
	if matrix[row+1][col-1] == 'M' && matrix[row+1][col+1] == 'M' {
		if matrix[row-1][col-1] == 'S' && matrix[row-1][col+1] == 'S' {
			return 1
		}
	}

	//	M - M
	//	- A -
	//	S - S
	if matrix[row-1][col-1] == 'M' && matrix[row-1][col+1] == 'M' {
		if matrix[row+1][col-1] == 'S' && matrix[row+1][col+1] == 'S' {
			return 1
		}
	}

	return 0
}

func main() {
	file, err := os.Open("day_4_Ceres_Search/input.txt")
	if err != nil {
		log.Println("open file failed, err:", err)
	}
	defer file.Close()

	firstPart, secondPart := solution(file)

	fmt.Println("First Part:", firstPart)
	fmt.Println("Second Part:", secondPart)
}
