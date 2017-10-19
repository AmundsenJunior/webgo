/*
- create a struct that holds person fields
- create a struct that holds secret agent fields and embeds person type
- attach a method to person: pSpeak
- attach a method to secret agent: saSpeak
- create a variable of type person
- create a variable of type secret agent
- print a field from person
- run pSpeak attached to the variable of type person
- print a field from secret agent
- run saSpeak attached to the variable of type secret agent
- run pSpeak attached to the variable of type secret agent
*/
package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("I am", p.fname, p.lname)
}

func (sa secretAgent) saSpeak() {
	fmt.Println("My name is", sa.lname, sa.fname, sa.lname)
}

func main() {
	p1 := person{
		"Miss",
		"Moneypenny",
	}
	p1.pSpeak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	if sa1.licenseToKill {
		sa1.person.pSpeak()
		sa1.saSpeak()
	}
}
