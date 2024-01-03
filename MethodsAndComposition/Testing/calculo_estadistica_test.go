package main

import (
	"testing"
)

// Función para calcular el mínimo de calificaciones
func CalcularMinimo(calificaciones []float64) float64 {
	if len(calificaciones) == 0 {
		return 0.0 // Si no hay calificaciones, se devuelve 0
	}

	min := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion < min {
			min = calificacion
		}
	}
	return min
}

// Función para calcular el máximo de calificaciones
func CalcularMaximo(calificaciones []float64) float64 {
	if len(calificaciones) == 0 {
		return 0.0 // Si no hay calificaciones, se devuelve 0
	}

	max := calificaciones[0]
	for _, calificacion := range calificaciones {
		if calificacion > max {
			max = calificacion
		}
	}
	return max
}

// Función para calcular el promedio de calificaciones
func CalcularPromedios(calificaciones []float64) float64 {
	if len(calificaciones) == 0 {
		return 0.0 // Si no hay calificaciones, se devuelve 0
	}

	total := 0.0
	for _, calificacion := range calificaciones {
		total += calificacion
	}
	return total / float64(len(calificaciones))
}

// Prueba para calcular el mínimo de calificaciones
func TestCalcularMinimo(t *testing.T) {
	calificaciones := []float64{7.5, 8.0, 6.5, 9.0, 7.0} // Calificaciones
	minimoEsperado := 6.5                                // Mínimo esperado

	minimo := CalcularMinimo(calificaciones)

	if minimo != minimoEsperado {
		t.Errorf("El mínimo calculado es %.2f; se esperaba %.2f", minimo, minimoEsperado)
	}
}

// Prueba para calcular el máximo de calificaciones
func TestCalcularMaximo(t *testing.T) {
	calificaciones := []float64{7.5, 8.0, 6.5, 9.0, 7.0} // Calificaciones
	maximoEsperado := 9.0                                // Máximo esperado

	maximo := CalcularMaximo(calificaciones)

	if maximo != maximoEsperado {
		t.Errorf("El máximo calculado es %.2f; se esperaba %.2f", maximo, maximoEsperado)
	}
}

// Prueba para calcular el promedio de calificaciones
func TestCalcularPromedios(t *testing.T) {
	calificaciones := []float64{7.5, 8.0, 6.5, 9.0, 7.0} // Calificaciones
	promedioEsperado := 7.4                              // Promedio esperado

	promedio := CalcularPromedio(calificaciones)

	if promedio != promedioEsperado {
		t.Errorf("El promedio calculado es %.2f; se esperaba %.2f", promedio, promedioEsperado)
	}
}
