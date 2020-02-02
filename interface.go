package main

import "fmt"

type person interface {
	changeName(string)
	changeSurname(string)
	getId() int
	changeId(int)
	getName() string
	getsurName() string
}

type student struct {
	name    string
	surname string
	id      int
}

func (s student) getName() string {
	return s.name
}
func (s student) getsurName() string {
	return s.surname
}
func (s *student) changeName(n string) {
	s.name = n
}
func (s *student) changeId(i int) {
	s.id = i
}
func (s *student) changeSurname(n string) {
	s.surname = n
}
func (s student) getId() int {
	return s.id
}

type instructor struct {
	name    string
	surname string
	id      int
}

func (s instructor) getName() string {
	return s.name
}
func (s instructor) getsurName() string {
	return s.surname
}
func (s *instructor) changeName(n string) {
	s.name = n
}
func (s *instructor) changeId(i int) {
	s.id = i
}
func (s *instructor) changeSurname(n string) {
	s.surname = n
}
func (s instructor) getId() int {
	return s.id
}

func printValue(p person) {
	//p.changeName("Aiba")
	fmt.Println("Id : ", p.getId())
	fmt.Println("Name : ", p.getName())
	fmt.Println("surname : ", p.getsurName())
}

func main() {
	/*

		var s student = student{"kudi", "stalbekov", 164137}
		i := instructor{"Cam", "Ferhun", 123232}
		s.changeName("Kudaibergen")
		i.changeSurname("Ciogklu")
		fmt.Println(&s)
		printValue(&s)
		fmt.Println(i)
		printValue(&i)

	*/

	/*
		var p person
		var s student = student{"kudi", "stalbekov", 164137}
		i := instructor{"Cam", "Ferhun", 123232}
		p = &s
		p.changeName("Kudaibergen")
		fmt.Println(p)
		printValue(p)
		p = &i
		p.changeSurname("Ciogklu")
		fmt.Println(p)
		printValue(p)
	*/
}
