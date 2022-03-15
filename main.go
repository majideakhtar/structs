package main

import "fmt"

type contactInfo struct {
	email string
	pin   int
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
			email: "jim@mail.com",
			pin:   831004,
		},
	}
	fmt.Println(jim)
	fmt.Printf("%+v", jim)

}
