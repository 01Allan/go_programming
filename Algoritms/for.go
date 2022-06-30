package main

import "fmt"

func main() {
	// primera forma
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// segunda forma

	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	// Tercera forma

	for {
		fmt.Println("Loop")
		break
	}

	// cuarta forma

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
