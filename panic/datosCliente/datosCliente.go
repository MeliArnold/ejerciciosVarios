package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("ejecución finalizada")

	data, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	// Imprimir datos si no hubo error:
	fmt.Println(string(data))
}
