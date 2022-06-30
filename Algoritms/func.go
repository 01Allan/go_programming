package main 
import "fmt"

func plus(a int, b int)  int{
	return a + b
}

func plusplus(a, b, c int)  int{
	return a + b + c
}

func main()  {

	suma1 := plus(2, 3)
	fmt.Println("suma1: ", suma1)

	suma2 := plusplus(2, 3, 4)
	fmt.Println("suma2: ", suma2)
	
}


