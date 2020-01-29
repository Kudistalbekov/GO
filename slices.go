package main

import "fmt"

func main() {
	/*
		var q []int = []int{1, 2, 4, 6, 8, 9, 0, 4, 3}
		s := q[1:3]  //2,4 this slice does't store data it's describe section if u change slice array will be changed
		s2 := q[1:4] //2,4,6
		fmt.Println("s = ", s)
		fmt.Println("s2 = ", s2)
		s[0] = 0
		fmt.Println(q) //1 ,0 ,4, 6, 8, 9, 0, 4, 3
		fmt.Println("after s = ", s)
		fmt.Println("after s2 = ", s2)
	*/
	/*
		var a []int = make([]int, 5)
		fmt.Println("a Length", len(a), "cap ", cap(a))
		//store(a)
		fmt.Println(a)

		a1 := [5]int{}
		fmt.Println("a1 Length", len(a1), "cap ", cap(a1))
		fmt.Println(a1)
	*/
	a2 := make([]int, 8, 10)
	store(a2) //[0][1][2][3][4][5][6][7][][]
	/*
		fmt.Println("a2 stored = ", a2)
		printslice(a2)
		fmt.Println()
		a3 := a2[1:6] //[1][2][3][4][5][][][][]
		printslice(a3)
		fmt.Println(a3)
		fmt.Println()
		b := a3[1:4] //[2][3][4][][][][][]
		b[1] = 99    //[2][99][4][][][][][]
		///0][1][2][99][4][5][6][7][][]
		printslice(b)
		fmt.Println(b)
		fmt.Println(a2)
	*/
	/*
		fmt.Println(a2)
		printslice(a2)
		a2 = append(a2, 10, 23, 23)
		printslice(a2)
		fmt.Println(a2)
	*/
	//range a2 <<--- returns two values first index and second copy of the value in that index
	/*
		fmt.Println(a2)
		findEven(a2)
		findOdd(a2)
	*/
}
func findEven(a []int) {
	for _, v := range a {
		if v%2 == 0 {
			fmt.Println("Even ", v)
		}
	}

}
func findOdd(a []int) {
	for i, v := range a {
		if v%2 != 0 {
			fmt.Println("ODD", v, "in index ", i)
		}
	}
}
func store(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
}
func printslice(a []int) {
	fmt.Println("Length", len(a), "cap ", cap(a))
}
