package saludos

import "fmt"

func Saludar(nombre string) {
	fmt.Printf("Hola %s, desde el paquete saludos!\n", nombre)
}
