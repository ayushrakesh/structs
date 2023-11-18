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

	fmt.Printf("%+v", kane)
}
