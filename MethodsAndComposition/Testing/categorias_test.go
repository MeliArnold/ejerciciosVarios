package main

import (
	"testing"
)

// Función para calcular el salario según la categoría del empleado
func CalcularSalario(categoria string) float64 {
	var salario float64

	switch categoria {
	case "A":
		salario = 50000.0 // Salario para categoría "A"
	case "B":
		salario = 40000.0 // Salario para categoría "B"
	case "C":
		salario = 30000.0 // Salario para categoría "C"
	default:
		salario = 0.0 // Categoría no válida
	}

	return salario
}

// Prueba para calcular el salario de la categoría "A"
func TestCalcularSalarioCategoriaA(t *testing.T) {
	categoria := "A"
	salarioEsperado := 50000.0 // Salario esperado para categoría "A"

	salario := CalcularSalario(categoria)

	if salario != salarioEsperado {
		t.Errorf("El salario calculado es %.2f; se esperaba %.2f", salario, salarioEsperado)
	}
}

// Prueba para calcular el salario de la categoría "B"
func TestCalcularSalarioCategoriaB(t *testing.T) {
	categoria := "B"
	salarioEsperado := 40000.0 // Salario esperado para categoría "B"

	salario := CalcularSalario(categoria)

	if salario != salarioEsperado {
		t.Errorf("El salario calculado es %.2f; se esperaba %.2f", salario, salarioEsperado)
	}
}

// Prueba para calcular el salario de la categoría "C"
func TestCalcularSalarioCategoriaC(t *testing.T) {
	categoria := "C"
	salarioEsperado := 30000.0 // Salario esperado para categoría "C"

	salario := CalcularSalario(categoria)

	if salario != salarioEsperado {
		t.Errorf("El salario calculado es %.2f; se esperaba %.2f", salario, salarioEsperado)
	}
}
