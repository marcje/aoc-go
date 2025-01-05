package utils

import (
	"fmt"
	"strconv"
)

var supportedYears = []int{2024}

const minDay = 1
const maxDay = 25

func ValidateArgs(args []string) (int, int, error) {
	if len(args) != 3 {
		return 0, 0, fmt.Errorf("invalid number of arguments")
	}

	year, err := strconv.Atoi(args[1])
	if err != nil || !validYear(year) {
		return 0, 0, fmt.Errorf("invalid year: %s (supported years: %v)", args[1], supportedYears)
	}

	day, err := strconv.Atoi(args[2])
	if err != nil || day < minDay || day > maxDay {
		return 0, 0, fmt.Errorf("invalid day: %s (must be between %d and %d)", args[2], minDay, maxDay)
	}

	return year, day, nil
}

func validYear(year int) bool {
	for _, supportedYear := range supportedYears {
		if year == supportedYear {
			return true
		}
	}
	return false
}
