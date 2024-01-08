package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	salary float64
	tax    float64
}

func calculateTax(salary float64) float64 {
	var taxRate float64

	switch {
	case salary <= 50000:
		taxRate = 0.1
	case salary <= 150000:
		taxRate = 0.2
	default:
		taxRate = 0.3
	}

	return salary * taxRate
}

func TestCalculateTax(t *testing.T) {
	testCases := []testCase{
		{salary: 40000, tax: 4000},
		{salary: 60000, tax: 12000},
		{salary: 200000, tax: 60000},
	}

	for _, tc := range testCases {
		if tax := calculateTax(tc.salary); tax != tc.tax {
			t.Errorf("Expected tax for salary %v to be %v, but got %v", tc.salary, tc.tax, tax)
		}
	}
}

func main() {
	testCases := []struct {
		salary      float64
		expectedTax float64
	}{
		{salary: 40000, expectedTax: 4000},
		{salary: 60000, expectedTax: 12000},
		{salary: 200000, expectedTax: 60000},
	}

	for _, tc := range testCases {
		actualTax := calculateTax(tc.salary)
		if actualTax != tc.expectedTax {
			fmt.Printf("Test Failed. Salary: %.2f, Expected Tax: %.2f, Got: %.2f\n", tc.salary, tc.expectedTax, actualTax)
		} else {
			fmt.Printf("Test Passed. Salary: %.2f, Tax: %.2f\n", tc.salary, actualTax)
		}
	}
}
