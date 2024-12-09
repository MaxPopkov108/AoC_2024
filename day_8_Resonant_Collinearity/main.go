package main

import (
	"bufio"
	"fmt"
	"os"
)

var uniqFreqs = make(map[[2]int]bool)
var uniqFreqsHarmonized = make(map[[2]int]bool)
var grid [][]string

func solution(file *os.File) (int, int) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempArr []string
		line := scanner.Text()
		for _, char := range line {
			tempArr = append(tempArr, string(char))
		}
		grid = append(grid, tempArr)
	}

	antennas := make(map[string][][2]int)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != "." {
				antennas[grid[row][col]] = append(antennas[grid[row][col]], [2]int{row, col})
			}
		}
	}

	for _, antenna := range antennas {
		antinodaradar(antenna, len(grid), len(grid[0]))
		antinodaradarHarmonized(antenna, len(grid), len(grid[0]))
	}

	return len(uniqFreqs), len(uniqFreqsHarmonized)
}

func antinodaradar(antennas [][2]int, rowLen, colLen int) {
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			rowDif := antennas[i][0] - antennas[j][0]
			colDif := antennas[i][1] - antennas[j][1]

			antinodaRow := antennas[i][0] + rowDif
			antinodaCol := antennas[i][1] + colDif

			if antinodaRow >= 0 && antinodaRow < rowLen && antinodaCol >= 0 && antinodaCol < colLen {
				uniqFreqs[[2]int{antinodaRow, antinodaCol}] = true
			}

			antinodaRow = antennas[j][0] - rowDif
			antinodaCol = antennas[j][1] - colDif

			if antinodaRow >= 0 && antinodaRow < rowLen && antinodaCol >= 0 && antinodaCol < colLen {
				uniqFreqs[[2]int{antinodaRow, antinodaCol}] = true
			}
		}
	}
}

func antinodaradarHarmonized(antennas [][2]int, rowLen, colLen int) {
	for i := 0; i < len(antennas); i++ {
		for j := i + 1; j < len(antennas); j++ {
			rowDif := antennas[i][0] - antennas[j][0]
			colDif := antennas[i][1] - antennas[j][1]

			antinodaRow := antennas[i][0]
			antinodaCol := antennas[i][1]

			for {

				if antinodaRow >= 0 && antinodaRow < rowLen && antinodaCol >= 0 && antinodaCol < colLen {
					uniqFreqsHarmonized[[2]int{antinodaRow, antinodaCol}] = true
				} else {
					break
				}

				antinodaRow += rowDif
				antinodaCol += colDif
			}

			antinodaRow = antennas[j][0]
			antinodaCol = antennas[j][1]

			for {

				if antinodaRow >= 0 && antinodaRow < rowLen && antinodaCol >= 0 && antinodaCol < colLen {
					uniqFreqsHarmonized[[2]int{antinodaRow, antinodaCol}] = true
				} else {
					break
				}

				antinodaRow -= rowDif
				antinodaCol -= colDif
			}
		}
	}
}

func main() {
	file, _ := os.Open("day_8_Resonant_Collinearity/input.txt")
	defer file.Close()

	firstPart, secondPart := solution(file)
	fmt.Printf("Part 1: %d\n", firstPart)
	fmt.Printf("Part 2: %d\n", secondPart)
}
