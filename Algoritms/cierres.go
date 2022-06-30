// En go se admiten funciones anonimas que pueden formar cierres. Las funciones anonimas son utiles cuando se desea definir una funcion en linea sin tener que nombrarla. 

package main 

import "fmt"

// La función devuelta se cierra sobre la variable para formar un cierre. 

func intseq() func() int{

	i := 0
	return func() int {
		i++
		return i
	}
	
}


func main() {
	nextInt := intseq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intseq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

}