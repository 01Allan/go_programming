package main

import (
	"fmt"
	"time"
)

//  Una goroutine es un hilo ligero de ejecución.

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("gorutine")

	go func (msg string) {
		fmt.Println(msg)
	}("going") // se imprime el mensaje o el valor de la llamada a la funcion anonima.

	time.Sleep(time.Second)
	fmt.Println("done")
}


