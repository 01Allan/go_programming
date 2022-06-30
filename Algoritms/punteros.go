package main

import "fmt"

func zeroval(ival int)  {
	ival = 0	
}

func zeroptr(iptr *int)  {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("initial =", i)

	zeroval(i)
	fmt.Println("zeroval =", i)
	//  cambia el valor de la variable
	zeroptr(&i)
	fmt.Println("zeroptr =", i)
	// imprime la direccion de memoria.
	fmt.Println("pointer =", &i)
}

// repaso de punteros en go.
