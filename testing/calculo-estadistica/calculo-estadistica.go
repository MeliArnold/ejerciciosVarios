package main

import (
	"fmt"
	"sort"
)

func main() {
	testCases := []struct {
		grades      []float64
		expectedMin float64
		expectedMax float64
		expectedAve float64
	}{
		{grades: []float64{70, 100, 80}, expectedMin: 70, expectedMax: 100, expectedAve: 83.33},
		{grades: []float64{60, 80, 70}, expectedMin: 60, expectedMax: 80, expectedAve: 70},
		{grades: []float64{80, 90, 100}, expectedMin: 80, expectedMax: 100, expectedAve: 90},
	}

	for i, tc := range testCases {
		actualMin := calculateMin(tc.grades)
		actualMax := calculateMax(tc.grades)
		actualAverage := calculateAverage(tc.grades)
		if actualMin != tc.expectedMin {
			fmt.Printf("Test case %d failed: Expected min %.2f but got %.2f\n", i+1, tc.expectedMin, actualMin)
		} else {
			fmt.Printf("Test case %d passed: Min grade %.2f\n", i+1, actualMin)
		}
		if actualMax != tc.expectedMax {
			fmt.Printf("Test case %d failed: Expected max %.2f but got %.2f\n", i+1, tc.expectedMax, actualMax)
		} else {
			fmt.Printf("Test case %d passed: Max grade %.2f\n", i+1, actualMax)
		}
		if fmt.Sprintf("%.2f", actualAverage) != fmt.Sprintf("%.2f", tc.expectedAve) {
			fmt.Printf("Test case %d failed: Expected average %.2f but got %.2f\n", i+1, tc.expectedAve, actualAverage)
		} else {
			fmt.Printf("Test case %d passed: Average grade %.2f\n", i+1, actualAverage)
		}
		fmt.Println("----------")
	}
}

func calculateMin(grades []float64) float64 {
	if len(grades) == 0 {
		return 0.0
	}
	sort.Float64s(grades)
	return grades[0]
}

func calculateMax(grades []float64) float64 {
	if len(grades) == 0 {
		return 0.0
	}
	sort.Float64s(grades)
	return grades[len(grades)-1]
}

func calculateAverage(grades []float64) float64 {
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	if len(grades) == 0 {
		return 0.0
	}
	return sum / float64(len(grades))
}
