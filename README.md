# Getting Started with Go

[![CI](https://github.com/DevOpsDerek/go-beginners/actions/workflows/ci.yml/badge.svg)](https://github.com/DevOpsDerek/go-beginners/actions/workflows/ci.yml)

A beginner-friendly Go course built as small, testable lesson packages. Each lesson teaches one core topic through exported functions, detailed comments, and unit tests written with Go's built-in `testing` package.

## Course Overview

This course is designed for developers who are new to Go and want hands-on practice with idiomatic basics before moving on to larger programs.

You will learn how to:
- read and write simple Go packages
- use variables, constants, and basic types
- work with operators, conditionals, and loops
- build and test functions
- handle slices, maps, closures, and errors
- create structs, methods, and interfaces

## Prerequisites

- Go 1.22+
- A terminal
- VS Code with the Go extension (recommended)
- `golangci-lint` for local linting

## Install Go and golangci-lint

Install Go from the official downloads page:
- https://go.dev/dl/

Install `golangci-lint`:

```bash
brew install golangci-lint
```

Or see the official install guide:
- https://golangci-lint.run/welcome/install/

## Project Layout

```text
.
├── README.md
├── go.mod
├── .golangci.yml
├── .github/workflows/ci.yml
└── lessons/
```

Each lesson lives in its own directory under `lessons/` and uses a package name like `lesson01`, `lesson02`, and so on.

## How to Explore a Lesson

These lessons are packages, not standalone executables, so there is no `main()` function inside the lesson folders.

### Run tests for all lessons

```bash
go test ./...
```

### Run tests for one lesson

```bash
go test ./lessons/01-hello-world/
```

### Run the linter

```bash
golangci-lint run
```

### Try a lesson in a scratch file

Create a small `main.go` in the repository root or another folder:

```go
package main

import (
    "fmt"

    lesson01 "github.com/DevOpsDerek/go-beginners/lessons/01-hello-world"
)

func main() {
    fmt.Println(lesson01.Greet("Go learner"))
}
```

Then run:

```bash
go run .
```

## Lesson Summary

| Lesson | Topic | Key Functions |
| --- | --- | --- |
| 01 | Hello World & Packages | `Greet`, `FullName` |
| 02 | Variables & Types | `CelsiusToFahrenheit`, `TypeName`, `AbsoluteValue` |
| 03 | Operators | `Calculate`, `IsEven`, `IsPalindrome` |
| 04 | Conditionals | `GetSeason`, `GetLetterGrade`, `Classify` |
| 05 | Loops | `Factorial`, `FizzBuzz`, `SumDigits` |
| 06 | Slices & Maps | `SliceStats`, `WordCount`, `MergeMaps` |
| 07 | Functions Basics | `CircleArea`, `MinMax`, `Apply` |
| 08 | Functions Advanced | `MakeCounter`, `MakeAdder`, `RunWithDefer` |
| 09 | Error Handling | `SafeDivide`, `ValidateAge`, `SafeRun` |
| 10 | Capstone: Structs & Interfaces | `NewStudent`, `ClassAverage`, `TopStudents`, `Student.Grade` |

## Beginner Tips for Go

- **Zero values matter.** Variables in Go always start with a useful default like `0`, `false`, `""`, or `nil`.
- **Use `:=` inside functions** when the type is obvious, and `var` when you want the zero value or need a specific type.
- **Errors are values.** In Go, functions often return `(value, error)` and callers handle problems explicitly.
- **Keep packages small.** These lessons are intentionally tiny so you can focus on one concept at a time.
- **Run tests often.** `go test ./...` is one of the fastest ways to build confidence while learning.

## Learning Objectives

By the end of this course, you should be able to:
- read package declarations and imports confidently
- write functions with clear return values
- use slices, maps, and structs effectively
- test Go code with the standard library
- understand how Go encourages simple, explicit error handling
