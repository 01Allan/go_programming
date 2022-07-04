package main 

import (
	"fmt"
	"sync"
	"time"
)

// funcion para cada gorutine
// se usa sleep para simular una tarea costosa:

func worker(id int) {
	fmt.Printf("Worker %d starting \n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}


func main()  {
	
	// Este WaitGroup se utiliza para esperar a que terminen todos los goroutines lanzados aquí. Nota: si un WaitGroup se pasa explícitamente a funciones, debe hacerse mediante puntero.

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}

