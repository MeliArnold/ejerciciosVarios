package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Crear un string de prueba
	testString := "Hello, World!"

	// Crear un nuevo Reader a través del string
	reader := bufio.NewReader(strings.NewReader(testString))

	// Leer todos los caracteres del Reader y almacenarlos en un string
	readString, _ := reader.ReadString('\n')

	// Imprimir el string leído
	fmt.Print(readString)
}
