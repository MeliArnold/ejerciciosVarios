package main

import (
	"fmt"
	"time"
)

func main() {
	// Crear un alumno
	//Excercise1
	fmt.Println("Excercise1")
	alumno := Alumno{
		Nombre:   "Juan",
		Apellido: "Perez",
		DNI:      "12345678",
		Fecha:    time.Now(),
	}

	// Llamar al método Detalle de la estructura Alumno
	detalle := alumno.Detalle()

	// Imprimir el detalle del alumno
	fmt.Println("Detalle del alumno:")
	fmt.Println(detalle)

	//Excercise2
	fmt.Println("Excercise2")
	// Crear instancias de productos
	small := SmallProduct{Precio: 100}
	medium := MediumProduct{Precio: 200}
	large := LargeProduct{Precio: 300}

	// Calcular el precio total de cada producto
	fmt.Println("Precio Total Small:", small.PrecioTotal())
	fmt.Println("Precio Total Medium:", medium.PrecioTotal())
	fmt.Println("Precio Total Large:", large.PrecioTotal())
}
