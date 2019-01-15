package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	alexContact := contactInfo{
		email:   "alex@gmail.com",
		zipCode: 94000,
	}

	alex := person{
		firstName:   "Alex",
		lastName:    "Al",
		contactInfo: alexContact,
	}
	alex.print()

	alex.updateName("Alexander", "Allen-Benedict")
	alex.print()

}

func (pointerToPerson *person) updateName(firstName string, l astName string) {
	//Turn memory address into a value using '*'
	pointerToPerson.firstName = firstName
	pointerToPerson.lastName = lastName

}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v\n", *pointerToPerson)
}
