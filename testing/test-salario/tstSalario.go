package main

import "fmt"

func calculateSalary(category string) float64 {
	switch category {
	case "A":
		return 50000.00
	case "B":
		return 70000.00
	case "C":
		return 100000.00
	default:
		return 0.00
	}
}

func main() {
	testCases := []struct {
		category       string
		expectedSalary float64
	}{
		{category: "A", expectedSalary: 50000.00},
		{category: "B", expectedSalary: 70000.00},
		{category: "C", expectedSalary: 100000.00},
	}

	for _, tc := range testCases {
		actualSalary := calculateSalary(tc.category)
		if actualSalary != tc.expectedSalary {
			fmt.Printf(
				"Test Fallido: %s, Expected Salary: %.2f, Got: %.2f\n",
				tc.category, tc.expectedSalary, actualSalary,
			)
		} else {
			fmt.Printf(
				"Test Passed. Category: %s, Salary: %.2f\n",
				tc.category, actualSalary,
			)
		}
	}
}
