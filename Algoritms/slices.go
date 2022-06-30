package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// se pueden crear slice con la longitud de otra y copiar sus elementos

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l1 := s[:5]
	fmt.Println("sl2:", l1)

	l3 := s[2:]
	fmt.Println("sl3:", l3)

	// Declarar slices con una sola linea

	t := []string{"g", "D", "i"}
	fmt.Println("dcl: ", t)

	// Las divisiones se pueden componer en estructuras de datos multidimensionales. La longitud de las rodajas internas puede variar, a diferencia de las matrices multidimensionales

	twod := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twod[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twod[i][j] = i + j
		}
	}

	fmt.Println("matrix: ", twod)

}
