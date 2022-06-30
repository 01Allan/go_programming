
// Las estructuras de Go son colecciones mecanografiadas de campos. Son útiles para agrupar datos para formar registros.

package main

import "fmt"

type person struct {
	name string 
	age int 
}

func newperson(name string) *person{

	p := person{name: name}
	p.age = 42
	return &p
}


func main()  {
	c := person{"bob", 20}
	fmt.Println(c)

	fmt.Println(person{name: "Alice", age: 50})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newperson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// Las estructuras son mutables.
	sp.age = 43
	fmt.Println(sp.age)

}
