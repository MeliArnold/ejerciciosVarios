package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open the file in read mode
	file, err := os.Open("./panic/leerArchivo/customer.txt")
	if err != nil {
		// If there is an error opening the file, throw a panic
		panic(err)
	}

	// Defer the file close until the main function has finished
	// This is important to prevent resource leak.
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}

		fmt.Println("Execution finished and file closed correctly.")
	}()

	// Read the content of the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Convert byte slice into string and print it
	fmt.Println(string(data))
}
