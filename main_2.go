package main

import "fmt"

func main() {

	defer fmt.Println("Esto se ejecuta cuando termina todo\n") // Defer para ejecutar al final

	//Estructura condicional if-else
	fmt.Println("\nEstructura condicional if-else")

	edad := 16      // Declaración y asignación de una variable local
	if edad >= 18 { // Estructura condicional if-else
		fmt.Println("Eres mayor de edad.")
	} else {
		fmt.Println("Eres menor de edad.")
	}

	//Tambien se puede hacer el assertive o negative programming
	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
		//return
	}
	fmt.Println("Eres menor de edad.")
	//Solo controla los casos positivos-> assertive programming
	//El negative programming se usa para controlar los casos negativos y salir temprano de la función
	//Lo mismo pero al revés

	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
		//return
	} else if edad <= 19 {
		fmt.Println("Tienes 18 o 19 años.")
		//return
	}

	//Bucle clasico for
	fmt.Println("\nBucle clasico for")
	for i := 0; i < 5; i++ {
		fmt.Printf("El valor de i es: %d\n", i)
	}

	//Bucle clasico for con array
	fmt.Println("\nBucle clasico for con array")
	array := []int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
		fmt.Printf("El valor de i es: %d en el array = %d\n", i, array[i])
	}

	//Bucle infinito con break
	fmt.Println("\nBucle infinito con break")
	count := 0
	for {
		if count >= 5 {
			break // Sale del bucle cuando count alcanza 5
		}
		fmt.Printf("El valor de count es: %d\n", count)
		count++
	}

	//Bucle infinito con continue
	fmt.Println("\nBucle infinito con continue")
	n := 0
	for {
		if (n % 2) == 0 {
			n++      // Incrementa n para evitar un bucle infinito
			continue // Salta a la siguiente iteración si n es par sin pasar por el print ni el if
		}
		fmt.Printf("El valor impar de n es: %d\n", n)
		if n >= 10 {
			break // Sale del bucle cuando n alcanza 10
		}
		n++ // Incrementa n
	}

	//Bucle while
	fmt.Println("\nBucle while")
	num := 0
	for num < 5 {
		fmt.Printf("El valor de num es: %d\n", num)
		num++
	}

	//Range en slices y arrays
	fmt.Println("\nRange en slices y arrays")
	slice := []string{"Go", "Python", "Java", "C++", "JavaScript"}
	for index, value := range slice {
		fmt.Printf("El índice es: %d y el valor es: %s\n", index, value)
	}

	//Range en mapas
	fmt.Println("\nRange en mapas")
	m := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for key, value := range m {
		fmt.Printf("La clave es: %s y el valor es: %d\n", key, value)
	}

	//Switch
	fmt.Println("\nSwitch")
	day := 3
	switch day {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Día inválido")
	}
}
