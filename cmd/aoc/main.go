package main

import (
	"aoc-go/internal/core"
	"aoc-go/pkg/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	os.Exit(cliRunner())
}

func cliRunner() int {
	day, year, err := core.ParseFlags()
	if err != nil {
		fmt.Println("Error:", err)
		flag.CommandLine.Usage()
		return 1
	}

	dailyInput, err := utils.GetDailyInput(year, day)
	if err != nil {
		fmt.Printf("Error retrieving input file: %v\n", err)
		return 1
	}

	solutionFunction, err := core.GetSolutionFunction(year, day)
	if err != nil {
		fmt.Println("Error:", err)
		return 1
	}

	part1, part2, err := solutionFunction(dailyInput)
	if err != nil {
		fmt.Println("Error: err")
		return 1
	}
	fmt.Println(part1)
	fmt.Println(part2)
	return 0
}
