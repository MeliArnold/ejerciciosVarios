package main

// Definición de la interfaz Producto
type Producto interface {
	PrecioTotal() float64
}

// Definición de la estructura SmallProduct
type SmallProduct struct {
	Precio float64
}

// Método PrecioTotal para SmallProduct
func (s SmallProduct) PrecioTotal() float64 {
	return s.Precio
}

// Definición de la estructura MediumProduct
type MediumProduct struct {
	Precio float64
}

// Método PrecioTotal para MediumProduct
func (m MediumProduct) PrecioTotal() float64 {
	return m.Precio * 1.03 // Aumenta el precio en un 3%
}

// Definición de la estructura LargeProduct
type LargeProduct struct {
	Precio float64
}

// Método PrecioTotal para LargeProduct
func (l LargeProduct) PrecioTotal() float64 {
	return (l.Precio * 1.06) + 2500 // Aumenta el precio en un 6% y suma $2500
}
