package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	alex := person{
		firstName: "Alec",
		lastName:  "Krafzik",
		contactInfo: contactInfo{
			email: "aleckfzk@gmail.com",
			zip:   54321,
		},
	}

	alex.print()
	alex.updateName("Alex")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
