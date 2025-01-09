package core

import (
	"flag"
	"fmt"
)

// ParseFlags parses commandline flags
func ParseFlags() (int, int, error) {
	dayFlag := flag.Int("d", 0, "Specify the day within 1-25 range")
	yearFlag := flag.Int("y", 0, "Specify the year")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  -d <day>   Specify the day within 1-25 range (e.g. -d 1)")
		fmt.Println("  -y <year>  Specify the year (e.g. -y 2024)")
		fmt.Println("\nExamples:")
		fmt.Println("  go run cmd/aoc/main.go -d 1 -y 2024")
	}

	flag.Parse()

	day, year, err := validateFlags(*dayFlag, *yearFlag)
	if err != nil {
		return 0, 0, err
	}
	return day, year, nil
}

// validateFlags validates commandline flags.
func validateFlags(day, year int) (int, int, error) {
	if day == 0 || year == 0 {
		return 0, 0, fmt.Errorf("both -d and -y must be set")
	}
	if day < 1 || day > 25 {
		return 0, 0, fmt.Errorf("invalid day: %d (must be between 1 and 25)", day)
	}
	// GetSupportedYears is imported from solve.go.
	supportedYears := GetSupportedYears()
	if !sliceContainsInt(supportedYears, year) {
		return 0, 0, fmt.Errorf("invalid year: %d (supported years must be one of: %v)", year, supportedYears)
	}

	return day, year, nil
}

// ContainsInt checks if a given searchValue exists in a slices with integers.
func sliceContainsInt(slice []int, searchValue int) bool {
	for _, value := range slice {
		if value == searchValue {
			return true
		}
	}
	return false
}
