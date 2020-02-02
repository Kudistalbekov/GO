package main

import "fmt"

type Person struct {
	name    string
	surname string
}

func main() {
	/*
		v := map[int]*Person{
			164137: {"Kudaibergen", "Stalbekov"},
			164114: {"Aibar", "Ulanov"},
			164105: {"ALiya", "Almatova"},
			164123: {"Beksultan", "Bekovich"},
		}
		fmt.Println(v)

		v2 := make(map[int]Person)
		v2[123445] = Person{"qweqweq", "dsfsdfas"}
		v2[123442] = Person{"q423weq", "dwerwdfas"}
		fmt.Println(v2)
	*/

	//v3 := *Person{"with", "pointer"} error
	v := Person{"without", "pointer"}
	v2 := &Person{"with", "pointer"}
	fmt.Println(v)
	fmt.Println(v2)
}
