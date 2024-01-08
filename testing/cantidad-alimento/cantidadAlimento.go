package main

import "fmt"

func calculateFood(animalType string, quantity int) int {
	switch animalType {
	case "dog":
		return quantity * 10
	case "cat":
		return quantity * 5
	case "hamster":
		return quantity * 2
	case "tarantula":
		return quantity * 1
	default:
		return 0
	}
}

func main() {
	testCases := []struct {
		animalType   string
		quantity     int
		expectedFood int
	}{
		{animalType: "dog", quantity: 10, expectedFood: 100},
		{animalType: "cat", quantity: 10, expectedFood: 50},
		{animalType: "hamster", quantity: 10, expectedFood: 20},
		{animalType: "tarantula", quantity: 10, expectedFood: 10},
	}

	for _, tc := range testCases {
		actualFood := calculateFood(tc.animalType, tc.quantity)
		if actualFood != tc.expectedFood {
			fmt.Printf(
				"Test Failed. Animal Type: %s, Quantity: %d, Expected Food: %d, Got: %d\n",
				tc.animalType, tc.quantity, tc.expectedFood, actualFood,
			)
		} else {
			fmt.Printf(
				"Test Passed. Animal Type: %s, Quantity: %d, Food: %d\n",
				tc.animalType, tc.quantity, actualFood,
			)
		}
	}
}
