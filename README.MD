# ESEP Grade Calculator

A Go library that calculates a student's final letter grade using weighted assignment, exam, and essay scores.

## Requirements

- Go 1.24.4 or compatible Go 1.24 release

## Project Structure

- `README.MD` - Project documentation
- `grade-calculator/go.mod` - Go module configuration
- `grade-calculator/grade_calculator.go` - Grade calculator implementation
- `grade-calculator/grade_calculator_test.go` - Unit tests

## Grade Weights

| Category | Weight |
|---|---|
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

## Letter Grade Scale

| Grade | Numerical Score |
|---|---|
| A | 90-100 |
| B | 80-89 |
| C | 70-79 |
| D | 60-69 |
| F | Below 60 |

A grade of C or higher is considered passing.

## Running Tests

Navigate into the Go module:

```bash
cd grade-calculator
```

Run the unit tests:

```bash
go test -v ./...
```

## Code Coverage

Run tests with coverage:

```bash
go test -coverprofile=coverage.out ./...
```

View total coverage:

```bash
go tool cover -func=coverage.out
```

View the coverage report in a browser:

```bash
go tool cover -html=coverage.out
```

## Usage

Create a calculator and add grades:

```go
calculator := NewGradeCalculator()

calculator.AddGrade("Assignment 1", 95, Assignment)
calculator.AddGrade("Exam 1", 85, Exam)
calculator.AddGrade("Essay 1", 90, Essay)

finalGrade := calculator.GetFinalGrade()
```

The calculator returns the final letter grade as a string.

This example assumes the calculator is used from within the same package.

## Notes

Grades are combined using the weighted averages of their respective categories. Any changes to the implementation should be verified using the automated tests.
