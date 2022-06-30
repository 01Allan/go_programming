package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	// otra forma de declarar arrays

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// en varias dimenciones dimenciones

	var twoD [10][10]int

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
