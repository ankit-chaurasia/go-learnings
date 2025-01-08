package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)       // {Alex Anderson}
	// fmt.Printf("%+v", alex) // {firstName:Alex lastName:Anderson}
	// var alex2 person
	// fmt.Println(alex2) // {}
	// alex.firstName = "Updated name"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim", lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com", zipCode: 12345,
		},
	}
	jim.print()
	jim.updateName("New FirstName")
	jim.print()

	jimPointer := &jim
	jimPointer.updateNameWithPointer(("pointer jimmy"))
	jimPointer.print()

}

// Pass by pointer reference
func (pointerToPerson *person) updateNameWithPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Pass by value
func (p person) updateName(newFirstName string) {
	// It will not update the original person struct passed
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
