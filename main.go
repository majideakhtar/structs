package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		lastName:  "Jim",
		firstName: "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			pinCode: 831004,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

// Structs with receiver functions
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
	p.print()
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v\n", p)
}
