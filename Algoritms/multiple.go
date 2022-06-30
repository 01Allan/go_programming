package main

import "fmt"

func vals() (int, int){
	return 3, 7
}

func main()  {
	// valores multiples retornados
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// para un subconjunto

	_, c := vals()
	fmt.Println(c)
	
}