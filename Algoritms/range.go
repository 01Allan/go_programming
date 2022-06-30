package main

import "fmt"

func main()  {
	nums  := []int{2, 3, 4} // array
	sum := 0
	
	//sumar los elementos de un array.
	// se ignora el indice con el identificador en blanco _,

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	// No se ignora el identificador. 

	for i, num := range nums {
		//fmt.Println(i)
		num = num + 1
		fmt.Println(num)
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range en el mapa itera sobra pares clave/valor

	kvs := map[string]string{"a":"apple", "b": "banana"}
	for k, v := range kvs{
		//fmt.Printf("%s -> %s\n", k, v)
		fmt.Println(k, "->", v)
	}

	// range tambien puede iterar solo sobre las claves de un mapa
	
	for k := range kvs{
		fmt.Println("key: ", k)
	}

	for i, c := range "go"{
		fmt.Println(i, c)
	}

	str := "hola"
	for _, letra := range str {
		fmt.Print(letra)
	}

}