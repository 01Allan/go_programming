package main

import "fmt"

// go no es un lenguaje orientado a objetos

type rect struct {
	width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
	t := rect{width: 10, height: 5}

	fmt.Println("area =", t.area())
	fmt.Println("perimetro =", t.perim())
	
	tp := &t
	fmt.Println("area =", tp.area())
	fmt.Println("perim =", tp.perim())

}

