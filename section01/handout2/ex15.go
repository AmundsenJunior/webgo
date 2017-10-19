/*
Now add a method to type gator with this signature ...
	greeting()
… and have it print “Hello, I am a gator”.
Create a value of type gator.
Call the greeting() method from that value.

https://play.golang.org/p/7Yk5MIlmNb
 */

package main

import (
	"fmt"
	"reflect"
)

type gator int

func (g gator) greeting() {
	fmt.Printf("Hello, I am a %s.\n", reflect.TypeOf(g).Name())
}

func main() {
	var g1 gator

	g1 = 10

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	g1.greeting()
}
