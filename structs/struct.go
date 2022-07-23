package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func structs() {
	justin := person{
		firstName: "justin",
		lastName:  "Wallace",
		contact: contact{
			email: "jpwallace22@gmail.com",
			zip:   92103,
		},
	}

	justinPointer := &justin
	justinPointer.updateName("Bob")
	justin.updateName("but whty")
	justin.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
