package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// este es un comentario,
	// un programa simple.

	// Tipos de datos de go

	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)

	// Declaración de variables

	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e int 
	f := "apple" // abreviacion para variables del tipo string
	//g := 3
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(e)

}
