/*package main

import (
	"fmt"
	"sync"
	"time"
)

//Concurrencia en Go se refiere a la capacidad de ejecutar múltiples tareas simultáneamente.
//Go proporciona primitivas de concurrencia integradas, como goroutines y canales, que facilitan
//la escritura de programas concurrentes.

//Goroutines -> son funciones o métodos que se ejecutan de manera concurrente con otras goroutines.
//Se inician con la palabra clave 'go' seguida de una llamada a función.

//Canales -> son conducciones que permiten la comunicación segura entre goroutines.
//Se utilizan para enviar y recibir valores entre goroutines.

//Ejemplo básico de goroutine y canales:

// Función que se ejecuta como una goroutine
func decirHola(canal chan<- string) {
	time.Sleep(2 * time.Second)         // Simula una tarea que toma tiempo
	canal <- "Hola desde la goroutine!" // Envía un mensaje al canal
}

// Función que recibe un mensaje desde un canal
func ImprimorMensaje(canal <-chan string) {
	mensaje := <-canal // Recibe un mensaje del canal
	fmt.Println(mensaje)
}

//Mirar doonde apunta la flecha
// canal chan <- string  = solo envia
// canal <- chan string  = solo recibe

func main() {
	/*
		canal := make(chan string) // Crea un canal de tipo string

		go decirHola(canal)    // Inicia la goroutine
		ImprimorMensaje(canal) // Llama a la función que recibe el mensaje

		canal2 := make(chan int)
		go func() {
			for i := 0; i < 5; i++ {
				canal2 <- i // Envía valores al canal
				fmt.Printf("Numero enviado al canal2: %d\n", i)
				time.Sleep(2000 * time.Millisecond) //Esperamos dos segundos antes de enviar el siguiente
			}
			close(canal2) // Cierra el canal cuando se termina de enviar
		}() //Usamos () para definir e invocar la función anónima inmediatamente

		//Funciona en paralelo con la goroutine anterior
		for num := range canal2 { // Recibe valores del canal hasta que esté cerrado
			time.Sleep(1000 * time.Millisecond) //Esperamos un segundo antes de printear
			fmt.Printf("Numero recibido del canal2: %d\n", num)
		}
	*/
	/*
	//sync.RWMutex para evitar condiciones de carrera en acceso a variables compartidas
	//Readers-Writers
	contador := 0
	var mu sync.RWMutex

	//Reader goroutine
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock() // Bloquea para lectura
				fmt.Printf("Reader %d lee contador: %d\n", id, contador)
				time.Sleep(3 * time.Second)
				mu.RUnlock() // Desbloquea después de leer

			}

		}(i)

	}

	//Writer goroutine
	go func() {

		for i := 0; i < 5; i++ {
			mu.Lock() // Bloquea para escritura
			time.Sleep(2 * time.Second)
			contador++ // Modifica la variable compartida
			fmt.Printf("Writer incrementa contador a: %d\n", contador)

			mu.Unlock() // Desbloquea después de escribir
		}
	}()

	//adjunto foto en readme para explicar mejor sync.RWMutex

	time.Sleep(30 * time.Second) // Espera para que todas las goroutines terminen
}

