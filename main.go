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
	// jimpointer := &jim // Actual process
	// jimpointer.updateName("Jimmy") //Actual process

	// Shortcut to pointers
	// Go automatically understands that if the receiver function is taking a pointer, we can simply give the object of same type, here we have jim
	jim.updateName("Jimmy")
	jim.print()
}

// Structs with receiver functions
// For both process, the below function will be same
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	// pointerToPerson.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
