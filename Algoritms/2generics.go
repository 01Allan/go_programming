package main

import (
	"fmt"
)

//  se generaliza una funcion para cualquier tipo de dato. 

func sum[T any](v1, v2 T) T {
	return v1 + v2
}

