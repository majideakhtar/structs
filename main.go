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
	jimpointer := &jim
	jimpointer.updateName("Jimmy")
	jim.print()
}

// Structs with receiver functions
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// pointerToPerson.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
