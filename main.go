package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type deck []string

func main() {

	// david := person{
	// 	"David", "Warner",
	// }

	// var steve person

	// steve.firstName = "Steve"
	// steve.lastName = "Smith"

	// fmt.Printf("%+v", steve)
	// fmt.Printf("%+v", david)
	// fmt.Printf("%+v", kane)

	kane := person{
		firstName: "Kane",
		lastName:  "Williamson",
		contact: contactInfo{
			email:   "kane@gmail.com",
			zipCode: 123456,
		},
	}

	kane.updateName("kane2")
	kane.print()

	cards := deck{"King", "Queen", "Joker", "Ace"}

	cards.updateDeck("Two")
	cards.print()

}

// function to print person
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// function to update person first name -> need pointer to update struct because it is a value type
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// function to print a deck
func (d deck) print() {
	fmt.Printf("%+v\n", d)
}

// function to update deck -> no pointers needed because slice is a reference types
func (d deck) updateDeck(newCard string) {
	d[0] = newCard
}
