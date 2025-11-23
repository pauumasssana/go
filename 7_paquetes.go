package main

import (
	saludos "conquer-go/saludos" //Podemos poner un alias al paquete importado
	"conquer-go/users/domain"    //Sino tiene el mismo nombre que la carpeta
)

func main() {
	saludos.Saludar("Jaime")
	saludos.Saludar2("Pau")

	domain.Domain()
}
