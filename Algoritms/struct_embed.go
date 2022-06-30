package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main()  {
	
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co = {num: %v, str: %v}\n",  co.num, co.str)
	fmt.Println("also num: ", co.base.num)
	fmt.Println("Describe: ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer:", d.describe())
	
}





