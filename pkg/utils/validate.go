package utils

import (
	"fmt"
	"strconv"
)

func ValidateArgs(args []string, supportedYears []int) (int, int, error) {
	if len(args) != 3 {
		return 0, 0, fmt.Errorf("invalid number of arguments")
	}

	year, err := strconv.Atoi(args[1])
	if err != nil || !validYear(year, supportedYears) {
		return 0, 0, fmt.Errorf("invalid year: %s (supported years: %v)", args[1], supportedYears)
	}

	day, err := strconv.Atoi(args[2])
	if err != nil || day < 1 || day > 25 {
		return 0, 0, fmt.Errorf("invalid day: %s (must be between 1 and 25)", args[2])
	}

	return year, day, nil
}

func validYear(year int, supportedYears []int) bool {
	for _, supportedYear := range supportedYears {
		if year == supportedYear {
			return true
		}
	}
	return false
}
