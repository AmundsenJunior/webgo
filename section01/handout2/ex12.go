/*
Adding onto this code (https://play.golang.org/p/qLoAWNmhHL):
Using var, declare an identifier “x” as type int (var x int).
Print out “x”.
Print the type of “x” using fmt.Printf(“%T\n”, x)

https://play.golang.org/p/g4Pcffi1P4
*/

package main

import "fmt"

type gator int

func main() {
	var g1 gator

	g1 = 10

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
