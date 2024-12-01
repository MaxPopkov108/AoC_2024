package main

import (
	"AoC_2024/file_manager"
	"fmt"
	"math"
	"sort"
)

var inputs = file_manager.ReadByColumns("input.txt")
var col1 = inputs[0]
var col2 = inputs[1]

func firstPart() int {

	sort.Ints(col1)
	sort.Ints(col2)

	sumDistances := 0

	for i := 0; i < len(col1); i++ {
		sumDistances += int(math.Abs(float64(col1[i] - col2[i])))
	}

	return sumDistances
}

func secondPart() int {

	sumInstances := 0

	col2dict := make(map[int]int)

	for i := 0; i < len(col2); i++ {
		if _, exists := col2dict[col2[i]]; exists {
			col2dict[col2[i]]++
		} else {
			col2dict[col2[i]] = 1
		}
	}

	for i := 0; i < len(col1); i++ {
		sumInstances += col1[i] * col2dict[col1[i]]
	}

	return sumInstances
}

func main() {
	fmt.Println("1st part solution: ", firstPart())

	fmt.Println("2nd part solution: ", secondPart())
}
