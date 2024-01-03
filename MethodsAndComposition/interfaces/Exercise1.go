package main

import (
	"fmt"
	"time"
)

// Definición de la interfaz DetalleAlumno
type DetalleAlumno interface {
	Detalle() string
}

// Definición de la estructura Alumno
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    time.Time
}

// Implementación del método Detalle para la estructura Alumno
func (a Alumno) Detalle() string {
	return fmt.Sprintf("Name: %s\nApellido: %s\nDNI: %s\nFecha: %s", a.Nombre, a.Apellido, a.DNI, a.Fecha.Format("2006-01-02"))
}
