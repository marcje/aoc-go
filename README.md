# aoc-go

A Go application that implements solutions for [Advent of Code (AoC)](https://adventofcode.com/) challenges.

## Table of Contents

- [Advent of Code - Go](#aoc-go)
- [Table of Contents](#table-of-contents)
- [Description](#description)
  - [Features](#features)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Implement challenges](#implement-challenges)
- [Future ideas](#future-ideas)

## Description

This application helps focussing on solving the Advent of Code (AoC) challenges by automating [the boring parts](#features).

This project serves mostly as a way to familiarize myself with Go. Developing the application proved to be a fun way to dive into Go. By occasionally improving the code and solving some puzzles I aim to strengthen my knowledge around Go.

### Features

- Easily implement [a new day or year](#implement-challenges)
- Solutions are dynamically mapped
- AoC input files are retrieved automatically
- Reusable utilities for input validation and parsing

## prerequisites

To run this application, ensure you have the following installed:

- **Go 1.20 or higher**: Install from [https://go.dev/dl/](https://go.dev/dl/).

Get your `SESSION_COOKIE` from AoC and setup in your environment:

```bash
export AOC_SESSION_COOKIE=your_cookie_value
```

## Usage

```bash
go run cmd/aoc/main.go <year> <day>
```

## Implement challenges

Follow these steps to implement a new challenge solution:

- Create `day<x>.go` in `internal/solutions/<year>/`
- Implement at least `SolveDay<x>` in `day<x>.go`
- Import the year package in `solve.go`
- Add day (and year if not already present) to `solutionsMap` in `solve.go`

## Future ideas

- Replace CLI arguments with CLI flags, allowing `-y` for year and `-d` for day
- Prompt user input if year and / or day are not provided
- Implement logging and use it for error handling
- Use interfaces to ensure daily solutions implement `SolveDay<x>` and `solvePart<1,2>`
- Add unit tests and allow for test input from AoC challenges
- Setup pre-commit hooks and GitHub actions to automatically run tests
- Containerize application
