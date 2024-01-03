package main

import "fmt"

// Estructura Product
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Slice global de Product
var Products []Product = []Product{
	{ID: 1, Name: "Producto 1", Price: 19.99, Description: "Descripción del producto 1", Category: "Categoria 1"},
	{ID: 2, Name: "Producto 2", Price: 29.99, Description: "Descripción del producto 2", Category: "Categoria 2"},
}

// Método Save asociado a la estructura Product
func (p *Product) Save() {
	Products = append(Products, *p)
}

// Método GetAll asociado a la estructura Product
func (p Product) GetAll() {
	fmt.Println("Listado de Productos:")
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

// Función GetById
func GetById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func main() {

	// Ejecutar el método GetAll
	emptyProduct := Product{}
	emptyProduct.GetAll()

	// Crear un nuevo producto y guardar usando el método Save
	newProduct := Product{ID: 3, Name: "Producto 3", Price: 39.99, Description: "Descripción del producto 3", Category: "Categoria 3"}
	newProduct.Save()

	// Ejecutar el método GetAll nuevamente para mostrar todos los productos, incluyendo el nuevo
	newProduct.GetAll()

	// Ejecutar la función getById
	idToSearch := 2
	foundProduct := GetById(idToSearch)
	if foundProduct != nil {
		fmt.Printf("\nProducto encontrado por ID (%d):\n", idToSearch)
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", foundProduct.ID, foundProduct.Name, foundProduct.Price, foundProduct.Description, foundProduct.Category)
	} else {
		fmt.Printf("\nNo se encontró ningún producto con el ID: %d\n", idToSearch)
	}

}
