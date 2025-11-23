package domain

import (
	adapters "conquer-go/users/adapters" //Podemos poner un alias al paquete importado
	utilities "conquer-go/users/utilities"
	"fmt" //Sino tiene el mismo nombre que la carpeta
)

var name string = "Gentleman"

func Domain() {
	adaptado := adapters.UsersAdapter(name)
	fmt.Println(utilities.ToUppercase(adaptado))
}
