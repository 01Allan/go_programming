package main

import "fmt"

func main()  {

	// son como los diccionarios en python
	m:= make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	// se le puede asignar a una variable una key

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
// El integrado devuelve el número de pares clave/valor cuando se llama en un mapa.len

	fmt.Println("len: ", len(m))

	// eliminar elementos de un map

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k1"]

	fmt.Println("prs: ", prs)

	// otra forma de declarar mapas

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

}
