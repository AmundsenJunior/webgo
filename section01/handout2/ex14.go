/*
Adding onto this code (https://play.golang.org/p/CrlJPB3ocb):
You will now learn about CONVERSION.
This is called “CASTING” in a lot of other languages.
Since “g1” is of type “gator” but its underlying type is an “int”, we can use “CONVERSION” to convert the value to an int.
Here is how you do it:

https://play.golang.org/p/zet-WRGZIF
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

	x = int(g1) // type conversion, where able, in order to assign g1 to x, where gator is an int
	fmt.Println(x)
}