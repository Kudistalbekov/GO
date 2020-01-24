//INITIALIZATION

package main

import "fmt"

func main() {
	//var i int = 2

	//var i = 2

	var i int
	i = 2

	i2 := 3 //short but it cant be used outside a function

	fmt.Println("i = ", i, "\ni2 = ", i2)
	//-------------------------------------------------//

	//var j, k, l int = 1, 3, 4

	//var j, k, l = 1, "kudi", false

	//j, k, l := 1, "stalbekov", true

	var (
		j = 3
		k = "kudi"
		l = false
	)

	fmt.Println("\nj = ", j, "\nk = ", k, "\nl = ", l)

	//-------------------------------------------------//
}
