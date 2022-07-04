package main

import (
	"fmt"
)

//  se generaliza una funcion para cualquier tipo de dato. 

// func divide_by_two[T any] (a T) {
	// return a / 2
// }

type int_or_float interface {
	int | float64
}

func divide_by_two[T int_or_float] (a T) T {
	return a/2
} 

func main()  {
	c := divide_by_two(5.0)
	fmt.Println(c)
}

