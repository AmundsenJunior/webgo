/*
- create an interface type that both person and secretAgent implement
- declare a func with a parameter of the interface’s type
- call that func in main and pass in a value of type person
- call that func in main and pass in a value of type secretAgent
*/

package main

import "fmt"

type mi6 interface {
	bondifyName()
}

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) bondifyName() {
	fmt.Printf("My name is %s, %s %s.\n", p.lname, p.fname, p.lname)
}

func (sa secretAgent) bondifyName() {
	fmt.Printf("My name is %s, %s %s. I have a license to kill.\n", sa.lname, sa.fname, sa.lname)
}

func sayHello(m mi6) {
	m.bondifyName()
}

func main() {
	p2 := person{
		"Alfred",
		"Pennyworth",
	}

	sa2 := secretAgent{
		person{
			"Bruce",
			"Wayne",
		},
		true,
	}

	sayHello(p2)
	sayHello(sa2)
}
