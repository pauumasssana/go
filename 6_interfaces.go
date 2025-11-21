package main

import (
	"fmt"
	"math"
)

//Interfaces en Go funcionan como contratos que definen un conjunto de métodos que
//un tipo debe implementar para cumplir con esa interfaz.

// Definición de una interfaz Forma con un método Area
type Forma interface {
	Area() float64
}

//Para que una struct implemente una interfaz, debe definir todos los métodos de esa interfaz

// Definición de una estructura Circulo que implementa la interfaz Forma
type Circulo struct {
	radio float64
}

// Implementación del método Area para la estructura Circulo
func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

// Definición de una estructura Rectangulo que también implementa la interfaz Forma
type Rectangulo struct {
	ancho, alto float64
}

// Implementación del método Area para la estructura Rectangulo
func (r Rectangulo) Area() float64 {
	return r.ancho * r.alto
}

// Función que acepta cualquier tipo que implemente la interfaz Forma
func imprimirArear(f Forma) {
	fmt.Printf("El area es: %.2f\n", f.Area())
}

func main() {
	circulo := Circulo{radio: 5} // Creación de una instancia de Circulo
	imprimirArear(circulo)       // Llamada a la función que acepta la interfaz Forma

	rectangulo := Rectangulo{ancho: 4, alto: 6} // Creación de una instancia de Rectangulo
	imprimirArear(rectangulo)                   // Llamada a la función que acepta la interfaz Forma

	var cualquiera interface{} = "Esto es un string" // interface{} puede contener cualquier tipo
	fmt.Println("Interface vacia: ", cualquiera)

}
