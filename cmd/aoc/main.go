package main

import (
	"aoc-go/pkg/utils"
	"fmt"
	"os"
)

func main() {
	year, day, err := utils.ValidateArgs(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Usage: go run main.go <year> <day>")
		return
	}

	dailyInput, err := utils.GetDailyInput(year, day)
	if err != nil {
		fmt.Printf("Error retrieving input file: %v\n", err)
		return
	}

	solutionFunction, err := utils.GetSolutionFunction(year, day)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	solutionFunction(dailyInput)
}
