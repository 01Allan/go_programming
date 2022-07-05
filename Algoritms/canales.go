package main 

import "fmt"

func main()  {
	
	messages := make(chan string)

	go func ()  {
		messages <- "Me amo mucho"
	}()

	msg := <- messages
	fmt.Println(msg)
}