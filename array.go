package main

import "fmt"

type Student struct {
	name string
	id   int
}

func p(array []int) {
	fmt.Println(array)
}

func main() {
	v := Student{"kudi", 164114}
	v.name = "Aiba"
	//	fmt.Println(v)
	pntr := &v
	pntr.name = "Keldi" // or like in c++ (*pntr).name
	(*pntr).id = 164221
	//	fmt.Println((*pntr).id, (*pntr).name)
	fmt.Println(*pntr)
	//	fmt.Println(pntr.id, pntr.name)
	pntr2 := pntr
	pntr2.id = 164107
	pntr2.name = "Anchik"
	fmt.Println(v)
	fmt.Println("---------------------------------------------------------------")
	num := []int{1, 2, 5, 6, 7, 8, 9, 8, 10, 11}
	var num1 []int
	num1 = []int{1, 2, 3, 6, 7}
	//simple array
	p(num)
	p(num1)

	a := []Student{{"kudi", 164137}, {"Aiba", 164114}, {"Anchik", 164107}}
	fmt.Println(a)
}
