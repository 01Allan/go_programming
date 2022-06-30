package main

import "fmt"

///  Una funcion variadica recibe un nimero indefinido de elementos 

func sum(nums ... int){

	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println("su suma:", total)
}

func main()  {

	// se hace el llamado de la funcion
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums ...)
	
}