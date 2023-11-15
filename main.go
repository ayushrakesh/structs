package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	kane := person{
		firstName: "Kane", lastName: "Williamson",
	}

	david := person{
		"David", "Warner",
	}

	var steve person

	steve.firstName = "Steve"
	steve.lastName = "Smith"

	fmt.Printf("%+v", steve)
	fmt.Printf("%+v", david)
	fmt.Printf("%+v", kane)
}
