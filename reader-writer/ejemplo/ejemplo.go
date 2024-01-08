package main

import (
	"fmt"
	"log"
	"os"
)

// readFile es una función que toma un path de archivo como string y lee el contenido del archivo.
func readFile(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("contenido del archivo : %s", content)
}

func main() {
	readFile("./reader-writer/ejemplo/ejemplo.txt")
}
