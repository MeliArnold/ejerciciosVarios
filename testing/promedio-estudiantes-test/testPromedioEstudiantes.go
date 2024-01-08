package main

import (
	"fmt"
)

func calculateAverage(grades []float64) float64 {
	sum := 0.0
	if len(grades) == 0 {
		return sum
	}
	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}

func main() {
	testCases := []struct {
		grades  []float64
		average float64
	}{
		{grades: []float64{90, 80, 100}, average: 90},
		{grades: []float64{70, 80, 90, 100}, average: 85},
		{grades: []float64{60, 80, 100}, average: 80},
	}

	for i, tc := range testCases {
		average := calculateAverage(tc.grades)
		if average != tc.average {
			fmt.Printf("Test case %d failed: For grades %v expected average %.2f but got %.2f\n", i+1, tc.grades, tc.average, average)
		} else {
			fmt.Printf("Test case %d passed: For grades %v got average %.2f\n", i+1, tc.grades, tc.average)
		}
	}
}
