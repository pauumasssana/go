package main

import "fmt"

// Función que suma dos números y devuelve el resultado
func sumar(a int, b int) int {
	return a + b
}

// Función que divide dos números y devuelve el cociente y el residuo
func dividir(a, b float64) (float64, float64) {
	cociente := a / b
	residuo := float64(int(a) % int(b))
	return cociente, residuo
}

// Función que imprime una lista de nombres usando parámetros variádicos
func imprimirNombres(nombres ...string) {
	for _, nombre := range nombres {
		fmt.Println("Nombre:", nombre)
	}
}

// Función que devuelve otra función que cuenta cuántas veces ha sido llamada (Clausura)
// Lo de dento puede accedar a lo de afuera pero no al revés
func contador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// En Go no existen las clases, pero se pueden definir métodos sobre tipos definidos por el usuario
type Rectangulo struct {
	ancho, alto float64
	//No se pueden definir métodos sobre tipos predefinidos como int, string, etc.
}

func (r Rectangulo) Area() float64 {
	return r.ancho * r.alto
}

func main() {
	fmt.Println("\nFunción que suma dos números y devuelve el resultado")
	result := sumar(5, 7)                                  // Llamada a la función sumar
	fmt.Printf("El resultado de la suma es: %d\n", result) // Impresión del resultado

	fmt.Println("\nFunción que divide dos números y devuelve el cociente y el residuo")
	cociente, residuo := dividir(10, 3)                                           // Llamada a la función dividir
	fmt.Printf("El cociente es: %.2f y el residuo es: %.2f\n", cociente, residuo) // Impresión de los resultados

	fmt.Println("\nFunción que imprime una lista de nombres usando parámetros variádicos")
	imprimirNombres("Ana", "Luis", "Carlos") // Llamada con varios argumentos

	fmt.Println("\nFunción que devuelve otra función que cuenta cuántas veces ha sido llamada (Clausura)")
	contar := contador()               // Asignación de la función devuelta a una variable
	fmt.Println("Contador:", contar()) // Llamada 1
	fmt.Println("Contador:", contar()) // Llamada 2
	fmt.Println("Contador:", contar()) // Llamada 3

	fmt.Println("\nDefinición de un método sobre un tipo definido por el usuario")
	rect := Rectangulo{ancho: 5, alto: 10}                       // Creación de una instancia de Rectangulo
	fmt.Printf("El área del rectángulo es: %.2f\n", rect.Area()) // Llamada al método Area del tipo Rectangulo
}
