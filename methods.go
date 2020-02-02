package main

import "fmt"

type Person struct {
	name    string
	surname string
}

func (v *Person) changeName(name string) {
	v.name = name
}

func printMap(m map[int]*Person) {
	for i, v := range m {
		fmt.Println("Key[", i, "] value = ", v)
	}
}

//////methods can be used in another way as well
///but only with initialized type like bellow
type MyInt int

func (a *MyInt) add(b int) {

	*a = *a + MyInt(b)
}

func main() {
	v := map[int]*Person{
		164137: {"Kudaibergen", "Stalbekov"},
		164114: {"Aibar", "Ulanov"},
		164105: {"ALiya", "Almatova"},
		164123: {"Beksultan", "Bekovich"},
	}
	v[164137].changeName("Kudi")
	printMap(v)
	/////////////////////////////
	var num MyInt = 5
	num.add(6)
	fmt.Println(num)

}
