package main

import (
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"os"
	"strconv"
)

func solution(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	var deparsedFileString []string

	fileIdx, emptyIdx := 0, 0
	var isFile bool
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

	var deparsedFileStringConcated string
	for char := range deparsedFileString {
		deparsedFileStringConcated += deparsedFileString[char]
	}

	var deparsedDict = make(map[int][]interface{})
	re := regexp2.MustCompile(`(\d)\1+|\.+|\d`, 0)
	matches := regexp2FindAllString(re, deparsedFileStringConcated)

	fmt.Println(matches)
	//for i, match := range matches {
	//	deparsedDict[i] = []interface{}{match, len(match)}
	//}

	fmt.Println(deparsedDict)

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

	return fileCheckSum, 0

}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

func main() {
	file, _ := os.Open("day_9_Disk_Fragmenter/test_input.txt")
	defer file.Close()

	firstPart, secondPart := solution(file)
	fmt.Printf("Part 1: %d\n", firstPart)
	fmt.Printf("Part 2: %d\n", secondPart)
}
