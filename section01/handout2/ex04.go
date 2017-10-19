/*
Using the STRUCT “person”,
  using a composite literal create a value of type person and assign it to a variable with the identifier “p1”;
print out “p1”;
print out just the field fName for “p1”

https://play.golang.org/p/g_LDj9k3IM
 */

package main

import "fmt"

type person struct {
	fName string
	lName string
}

func main() {
	p1 := person {
		"Scott",
		"Russell",
	}

	fmt.Println(p1)
	fmt.Println(p1.fName)
}