package solutions2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolveDay1(inputLines []string) {
	leftList, rightList, err := parseInputColumns(inputLines)
	if err != nil {
		fmt.Printf("error parsing input lists: %v\n", err)
		return
	}

	part1 := solvePart1(leftList, rightList)
	fmt.Println("Day 1 - Part 1:", part1)

	part2 := solvePart2(inputLines)
	fmt.Println("Day 1 - Part 2:", part2)
}

// Parse day 1 input, which consists of two columns per line.
func parseInputColumns(lines []string) ([]int, []int, error) {
	var leftList, rightList []int

	for _, line := range lines {
		columns := strings.Fields(line)
		if len(columns) < 2 {
			return nil, nil, fmt.Errorf("unexpected input format, expected two columns but got: %q", line)
		}

		leftValue, err := strconv.Atoi(columns[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in left column: %s", columns[0])
		}

		rightValue, err := strconv.Atoi(columns[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in right column: %s", columns[1])
		}

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	return leftList, rightList, nil
}

// calculateDistance calculates the difference between two integers.
func calculateDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func solvePart1(leftList []int, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += calculateDistance(leftList[i], rightList[i])
	}

	return totalDistance
}

func solvePart2(lines []string) int {
	return 0
}
