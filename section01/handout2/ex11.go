/*
Create a new type called “gator”.
The underlying type of “gator” is an int.
Using var, declare an identifier “g1” as type gator (var g1 gator).
Assign a value to “g1”.
Print out “g1”.
Print the type of “g1” using fmt.Printf(“%T\n”, g1)

https://play.golang.org/p/qLoAWNmhHL
 */

package main

import "fmt"

type gator int

var g1 gator

func main() {
	g1 = 10

	fmt.Println(g1)

	fmt.Printf("%T\n", g1)
}