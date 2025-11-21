package main

import "fmt"

func sumar(a, b int) int {
	a = 2
	b = 2
	return a + b
}

func sumar_punt(a, b *int) int { // a y b son punteros a enteros
	*a = 2 // Desreferenciación para obtener el valor apuntado por a
	*b = 2
	// No se puede hacer a = 2 porque a es un puntero, no un entero
	return *a + *b
}

//Usaremos punteros cuando:
// - Queramos modificar el valor de una variable con una función
// - Queramos evitar copiar grandes estructuras de datos para mejorar el rendimiento

// &var -> obtiene la dirección de memoria de la variable var
// *var -> obtiene el valor almacenado en la dirección de memoria var

//Tambien en structs

type Rectangulo struct {
	ancho, alto int
}

func (r Rectangulo) area() int {
	return r.ancho * r.alto
}

func (r *Rectangulo) setAncho(ancho int) {
	r.ancho = ancho
	// Tambien se puede usar (*r).ancho = ancho
	// porque Go desreferencia automáticamente los punteros a structs
}

func main() {
	a := 10
	b := 10

	c := 10
	d := 10

	suma_punt := sumar_punt(&a, &b)
	suma := sumar(c, d)

	fmt.Println("\n-Probamos punteros\n")
	fmt.Printf("La suma_punt (a y b) de %d y %d es: %d\n", a, b, suma_punt)
	fmt.Printf("La suma (c y d) de %d y %d es: %d\n", c, d, suma)

	fmt.Println("\nVemos que los valores de 'a' e 'b' han cambiado, miramos tambien el resultado de la suma\n")
	fmt.Printf("El valor de a es: %d\n", a)
	fmt.Printf("El valor de b es: %d\n", b)

	fmt.Println("\nVemos que los valores de 'c' y 'd' NO han cambiado, miramos tambien el resultado de la \nsuma se hace con los valores asignados dentro la funcion c=2 y d=2\n")
	fmt.Printf("El valor de c es: %d\n", c)
	fmt.Printf("El valor de d es: %d\n", d)

	// new() -> asigna memoria para un tipo y devuelve un puntero a esa memoria
	puntero := new(int) // puntero es de tipo *int que apunta a un entero se inicializa a 0
	fmt.Println("\n-Probamos new()")
	fmt.Println("\nCreamos un puntero a un entero con: puntero := new(int)")
	fmt.Printf("El valor del puntero es: %d\n", puntero)
	fmt.Printf("El valor del *puntero es: %d\n", *puntero)
	fmt.Printf("El valor del &puntero es: %d\n", &puntero)

	*puntero = 42 // Asignamos un valor al entero apuntado por puntero
	fmt.Println("\nAsignamos un valor al puntero con *puntero = 42")
	fmt.Printf("El valor del puntero es: %d\n", puntero)
	fmt.Printf("El valor del *puntero es: %d\n", *puntero)
	fmt.Printf("El valor del &puntero es: %d\n", &puntero)

	// Uso de punteros con structs
	fmt.Println("\n-Uso de punteros con structs")
	rect := Rectangulo{ancho: 10, alto: 5}
	rect.setAncho(20)
	fmt.Println("\nCreamos un rectángulo con ancho 10 y alto 5")
	fmt.Println("Usamos el método rect.setAncho(20) para cambiar el ancho a 20")
	fmt.Printf("El área del rectángulo es: %d\n", rect.area())
	fmt.Printf("El ancho del rectángulo es: %d\n", rect.ancho)
}
