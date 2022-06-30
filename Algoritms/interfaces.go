//  interfaces se denominan colecciones de firmas de metodos. 

package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radio float64 
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c circle) perim() float64 {
	return 2*math.Pi*c.radio
}

// se llaman a los metodos en una sola interfaz o un una funcion generica para todos los casos. 

func measure(g geometry) {
	fmt.Println(g) // datos ingresados
	fmt.Println(g.area()) 
	fmt.Println(g.perim())
}

func main()  {
	r1 := rect{width: 2, height: 4}
	c1 := circle{radio: 3}

	measure(r1)
	measure(c1)
}


