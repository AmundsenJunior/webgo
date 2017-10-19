/*
Adding onto this code(https://play.golang.org/p/7Yk5MIlmNb): create another type called “flamingo”.
Make the underlying type of “flamingo” bool.
Give “flamingo” a method with this signature …
	greeting()
… and have it print “Hello, I am pink and beautiful and wonderful.”
Now create a new type “swampCreature”.
The underlying type of “swapCreature” is interface.
The behavior which the “swampCreature” interface defines is such that any type which has this method …
	greeting()
… will implicitly implement the “swampCreature” interface.
Create a func called “bayou” which takes a value of type “swampCreature” as an argument.
Have this func print out the greeting for whatever “swampCreature” might be passed in.

https://play.golang.org/p/bHqC7nzyvW
 */

package main

import (
	"fmt"
	"reflect"
)

type swampCreature interface {
	greeting()
}

type gator int

type flamingo bool

func (g gator) greeting() {
	fmt.Printf("Hello, I am a %s.\n", reflect.TypeOf(g).Name())
}

func (f flamingo) greeting() {
	fmt.Printf("Hello, I am a %s, and I am pink and beautiful and wonderful.\n", reflect.TypeOf(f).Name())
}

func bayou(c swampCreature) {
	c.greeting()
}

func main() {
	var g1 gator
	var f1 flamingo

	bayou(g1)
	bayou(f1)
}
