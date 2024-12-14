package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	// part 1

	deparsedFileString := deparsedStringBuilder(line)
	deparsedFileStringBlocks := splitToBlocks(deparsedFileString)

	for i := 0; i < len(deparsedFileString); i++ {
		if deparsedFileString[i] == "." {
			for j := len(deparsedFileString) - 1; j >= 0; j-- {
				if i >= j {
					break
				}
				if deparsedFileString[j] != "." {
					deparsedFileString[i], deparsedFileString[j] = deparsedFileString[j], deparsedFileString[i]
					break
				}
			}
		}
	}

	fileCheckSum := 0

	for i := 0; i < len(deparsedFileString); i++ {
		if deparsedFileString[i] == "." {
			break
		}
		num, _ := strconv.Atoi(deparsedFileString[i])
		fileCheckSum += num * i
	}

	// part 2 TODO: works, but doesn't work

	for i := len(deparsedFileStringBlocks) - 1; i >= 0; i-- {
		if strings.ContainsAny(deparsedFileStringBlocks[i][0], "0123456789") {
			numBlock := deparsedFileStringBlocks[i]

			numLen := len(numBlock)

			for j := 0; j < len(deparsedFileStringBlocks); j++ {

				if j > i {
					break
				}

				if deparsedFileStringBlocks[j][0] == "." {
					dotBlock := deparsedFileStringBlocks[j]

					if len(dotBlock) >= numLen {
						for k := 0; k < numLen; k++ {
							dotBlock[k], numBlock[k] = numBlock[k], dotBlock[k]
						}

						deparsedFileStringBlocks = normalizeBlocks(deparsedFileStringBlocks)
						j = 0
						break
					}
				}
			}
		}
	}

	finalArray := []string{}
	for _, block := range deparsedFileStringBlocks {
		finalArray = append(finalArray, block...)
	}

	arrayFileCheckSum := 0
	for i := 0; i < len(finalArray); i++ {
		if finalArray[i] != "." {
			num, _ := strconv.Atoi(finalArray[i])
			arrayFileCheckSum += num * i
		}
	}

	return fileCheckSum, arrayFileCheckSum

}

func splitToBlocks(deparsedFileString []string) [][]string {
	var result [][]string
	var currentBlock []string

	for i, char := range deparsedFileString {
		if len(currentBlock) == 0 {
			currentBlock = append(currentBlock, char)
		} else {
			if char == "." && currentBlock[0] == "." {
				currentBlock = append(currentBlock, char)
			} else if char != "." && currentBlock[0] != "." && char == currentBlock[len(currentBlock)-1] {
				currentBlock = append(currentBlock, char)
			} else {
				result = append(result, currentBlock)
				currentBlock = []string{char}
			}
		}

		// Append the last block if at the end of the input
		if i == len(deparsedFileString)-1 {
			result = append(result, currentBlock)
		}
	}

	return result
}

func normalizeBlocks(blocks [][]string) [][]string {
	var normalized [][]string
	var currentDots []string

	for _, block := range blocks {
		if block[0] == "." {
			currentDots = append(currentDots, block...)
		} else {
			if len(currentDots) > 0 {
				normalized = append(normalized, currentDots)
				currentDots = nil
			}

			var currentNumbers []string
			for _, char := range block {
				if char == "." {
					if len(currentNumbers) > 0 {
						normalized = append(normalized, currentNumbers)
						currentNumbers = nil
					}
					currentDots = append(currentDots, char)
				} else {
					if len(currentDots) > 0 {
						normalized = append(normalized, currentDots)
						currentDots = nil
					}
					currentNumbers = append(currentNumbers, char)
				}
			}

			if len(currentNumbers) > 0 {
				normalized = append(normalized, currentNumbers)
			}
		}
	}

	if len(currentDots) > 0 {
		normalized = append(normalized, currentDots)
	}

	return normalized
}

func deparsedStringBuilder(line string) []string {
	var deparsedFileString []string
	var isFile bool
	fileIdx, emptyIdx := 0, 0

	for i, char := range line {
		idxToChar := 0

		if i%2 != 0 {
			isFile = false
			idxToChar = emptyIdx
			emptyIdx++
		} else {
			isFile = true
			idxToChar = fileIdx
			fileIdx++
		}

		num, _ := strconv.Atoi(string(char))
		idxChar := strconv.Itoa(idxToChar)
		for num != 0 {

			if isFile {
				deparsedFileString = append(deparsedFileString, idxChar)
			} else {
				deparsedFileString = append(deparsedFileString, ".")
			}

			num--
		}
	}

	return deparsedFileString
}

func main() {
	file, _ := os.Open("day_9_Disk_Fragmenter/input.txt")
	defer file.Close()

	firstPart, secondPart := solution(file)
	fmt.Printf("Part 1: %d\n", firstPart)
	fmt.Printf("Part 2: %d\n", secondPart)
}

// 116987784736, 116916073215 -> too low
// 6514433664797 -> too high
