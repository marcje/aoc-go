package main

import (
	"aoc-go/internal/core"
	"aoc-go/pkg/utils"
	"flag"
	"fmt"
)

func main() {
	day, year, err := core.ParseFlags()
	if err != nil {
		fmt.Println("Error:", err)
		flag.CommandLine.Usage()
		return
	}

	dailyInput, err := utils.GetDailyInput(year, day)
	if err != nil {
		fmt.Printf("Error retrieving input file: %v\n", err)
		return
	}

	solutionFunction, err := core.GetSolutionFunction(year, day)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	solutionFunction(dailyInput)
}
