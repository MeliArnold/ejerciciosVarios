package main

import (
	"testing"
)

// Función para calcular impuestos
func CalcularImpuesto(salario float64) float64 {
	var impuesto float64

	// Reglas de cálculo de impuestos
	if salario < 50000 {
		impuesto = salario * 0.1 // Impuesto del 10% para salarios menores a $50,000
	} else if salario >= 50000 && salario <= 150000 {
		impuesto = salario * 0.2 // Impuesto del 20% para salarios entre $50,000 y $150,000
	} else {
		impuesto = salario * 0.3 // Impuesto del 30% para salarios mayores a $150,000
	}

	return impuesto
}

// Prueba para calcular impuesto en caso de que el empleado gane por debajo de $50,000
func TestImpuestoMenor50k(t *testing.T) {
	salario := 40000.0
	resultadoEsperado := 4000.0 // Impuesto esperado para un salario de $40,000

	resultado := CalcularImpuesto(salario)

	if resultado != resultadoEsperado {
		t.Errorf("Para un salario de %.2f, el impuesto calculado es %.2f; se esperaba %.2f", salario, resultado, resultadoEsperado)
	}
}

// Prueba para calcular impuesto en caso de que el empleado gane entre $50,000 y $150,000
func TestImpuestoEntre50kY150k(t *testing.T) {
	salario := 100000.0
	resultadoEsperado := 20000.0 // Impuesto esperado para un salario de $100,000

	resultado := CalcularImpuesto(salario)

	if resultado != resultadoEsperado {
		t.Errorf("Para un salario de %.2f, el impuesto calculado es %.2f; se esperaba %.2f", salario, resultado, resultadoEsperado)
	}
}

// Prueba para calcular impuesto en caso de que el empleado gane por encima de $150,000
func TestImpuestoMayor150k(t *testing.T) {
	salario := 200000.0
	resultadoEsperado := 60000.0 // Impuesto esperado para un salario de $200,000

	resultado := CalcularImpuesto(salario)

	if resultado != resultadoEsperado {
		t.Errorf("Para un salario de %.2f, el impuesto calculado es %.2f; se esperaba %.2f", salario, resultado, resultadoEsperado)
	}
}
