package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{lastName: "Lima", firstName: "Alex"}
	// fmt.Println(alex)
	// alex.lastName = "Dickson"
	// fmt.Println(alex)
	Another way to assogn structs
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Lima"
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // {firstName: lastName:}

}
