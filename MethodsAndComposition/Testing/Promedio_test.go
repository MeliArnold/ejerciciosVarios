package main

import (
	"testing"
)

// Función para calcular el promedio de un slice de notas
func CalcularPromedio(notas []float64) float64 {
	total := 0.0
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

// Prueba para calcular el promedio de las notas de los alumnos
func TestCalcularPromedio(t *testing.T) {
	notas := []float64{8.5, 9.0, 7.5, 6.0, 9.5} // Notas de los alumnos
	promedioEsperado := 8.1                     // Promedio esperado para estas notas

	promedio := CalcularPromedio(notas)

	if promedio != promedioEsperado {
		t.Errorf("El promedio calculado es %.2f; se esperaba %.2f", promedio, promedioEsperado)
	}
}
