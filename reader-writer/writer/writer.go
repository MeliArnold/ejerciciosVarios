package main

import (
	"log"
	"os"
)

func main() {
	// Abre el archivo en modo append. Crea el archivo si no existe.
	file, err := os.OpenFile("./reader-writer/writer/ejemplo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo: %v", err)
	}

	defer file.Close()

	// Escribe en el archivo
	if _, err := file.WriteString("Hola mundo con Go!\n"); err != nil {
		log.Fatalf("No se pudo escribir en el archivo: %v", err)
	}
}
