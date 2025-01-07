package core

import (
	solutions2024 "aoc-go/internal/solutions/2024"
	"fmt"
)

// Dynamically map years and days to implemented functions.
type SolutionFunc func([]string)

// solutionsMap maps implemented years and day numbers to their corresponding solution functions.
var solutionsMap = map[int]map[int]SolutionFunc{
	2024: {
		1: solutions2024.SolveDay1,
	},
}

// GetSupportedYears returns a slice of years that are implemented.
func GetSupportedYears() []int {
	var years []int
	for year := range solutionsMap {
		years = append(years, year)
	}
	return years
}

// Check if a solution exists for the given year and day, and run it if available.
func GetSolutionFunction(year, day int) (SolutionFunc, error) {
	daySolutions, ok := solutionsMap[year]
	if !ok {
		return nil, fmt.Errorf("year %d not implemented", year)
	}

	solutionFunction, ok := daySolutions[day]
	if !ok {
		return nil, fmt.Errorf("day %d not implemented for year %d", day, year)
	}

	return solutionFunction, nil
}
