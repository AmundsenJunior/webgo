/*
Add a method to type “person” with the identifier “walk”.
Have the func return this string: <person’s first name> is walking.
Remember, you make a func a method by giving the func a receiver.
	func (r receiver) identifier(parameters) (returns) {
		<code>
	}
To return a string, use fmt.Sprintln.
Call the method assigning the returned string to a variable with the identifier “s”.
Print out “s”.

https://play.golang.org/p/bMBK93Tr4l
 */

package main

import "fmt"

type person struct {
	fName string
	lName string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintf("%s is walkin', heah!\n", p.fName)
}

func main() {
	p1 := person{
		"Scott",
		"Russell",
		[]string{
			"chicken pot pie",
			"Greek salad",
			"wheat bread",
			"mac & cheese",
		},
	}

	s := p1.walk()

	fmt.Println(s)
}
