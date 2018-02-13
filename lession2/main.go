package main

import (
	"fmt"
)

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
	var alexon person

	alexon.firstName = "Kanokrat"
	alexon.lastName = "Hongkhao"
	alexon.contactInfo.email = "aofiee666@gmail.com"
	alexon.contactInfo.zip = 10210

	//alexonPointer := &alexon
	//alexonPointer.updateName("Karnwara")
	alexon.updateName("Karnwara")
	alexon.print()

	man := person{
		firstName: "Khomkrid",
		lastName:  "Lerdprasert",
		contactInfo: contactInfo{
			email: "aofiee666@gmail.com",
			zip:   10210,
		},
	}

	man.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}
