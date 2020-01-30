package main

import "fmt"

type Person struct {
	surname string
	name    string
}

/*
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m:=make(map[string]int)
	sl:=strings.Fields(s)
	for _,v:=range sl{
		_,o:=m[v]
		if o==true{
			m[v]=2
		}else{
		    m[v]=1
		}

	}

	return m
}

func main() {
	wc.Test(WordCount)
}
*/

func main() {
	num := make(map[int]string)
	num[0] = "zero"
	num[1] = "one"
	num[2] = "two"
	num[3] = "three"
	num[4] = "four"
	num[5] = "five"
	for i := 0; i < 6; i++ {
		fmt.Println("num[", i, "]", num[i])
	}

	m := make(map[int]Person)

	v := map[int]Person{
		164105: {"Alya", "Almatova"},
		184150: {"Beksulatn", "Turatbekov"},
	}

	m[164137] = Person{"Stalbekov", "Kudaibergen"}
	m[164114] = Person{"Aibar", "Ulanov"}

	fmt.Println(m[164137])
	fmt.Println(v[184150])
	fmt.Println(v[164105])
}
