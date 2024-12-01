package file_manager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadByColumns reads a file and separates columns into slices
func ReadByColumns(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil
	}

	var columns [][]int
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line) // Split line into fields based on whitespace

		// Ensure the columns slice can hold the required number of columns
		if len(columns) < len(fields) {
			for i := len(columns); i < len(fields); i++ {
				columns = append(columns, []int{})
			}
		}

		// Convert each field to an integer and append it to the corresponding column
		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Error parsing field '%v' in line '%v': %v\n", field, line, err)
				continue
			}
			columns[i] = append(columns[i], num)
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	return columns
}
