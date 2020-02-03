package main

import "fmt"

type person interface {
	changename(string)
	getname() string
}

type student struct {
	name string
}

func (s student) getname() string {
	return s.name
}

func (s *student) changename(n string) {
	s.name = n
}

type instructor struct {
	name string
}

func (i *instructor) changename(n string) {
	i.name = n
}
func (i instructor) getname() string {
	return i.name
}
func getname(p person) string {
	fmt.Printf("%v, %T\n", p, p)
	if p == nil {
		return "person in empty"
	}
	return p.getname()
}

func main() {
	var p person
	fmt.Println(getname(p))
	i := instructor{"Ferhun"}
	s := student{"Aibar"}
	p = &i
	fmt.Println(getname(p)) //p.getname() -- works as well
	p = &s
	p.changename("Kudi")
	fmt.Println(getname(p))
}
