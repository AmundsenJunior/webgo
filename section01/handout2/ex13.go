/*
Adding onto this code (https://play.golang.org/p/g4Pcffi1P4):
Can you assign “g1” to “x”?
Why or why not?

https://play.golang.org/p/CrlJPB3ocb
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

	x = g1 // assignment fails as x is not of type 'gator'
}

