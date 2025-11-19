package main

import "fmt"

var numerito int = 10 // Declaración y asignación de una variable global

func main() {
	fmt.Println("Hello, World!")

	numerito = 20 // Alternativa para reasignar el valor de una variable ya declarada
	fmt.Println("El valor de numerito es:", numerito)
	numerito2 := 30 // Declaración y asignación en una sola línea
	fmt.Println("El valor de numerito2 es:", numerito2)

	i, j := 132, 23566 // Declaración y asignación múltiple
	fmt.Println("El valor de i es:", i)
	fmt.Println("El valor de j es:", j)

}
