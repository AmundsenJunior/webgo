/*
Using the short declaration operator,
  create a variable with the identifier “s” and assign “i'm sorry dave i can't do that” to “s”.
Print “s”.
Print “s” converted to a slice of byte.
Print “s” converted to a slice of byte and then converted back to a string.
Using slicing, print just “i’m sorry dave”
Using slicing, print just “dave i can’t”
Using slicing, print just “can’t do that”
print every letter of “s” with one rune (character) on each line

https://play.golang.org/p/IziLF0M9tE
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "i'm sorry dave i can't do that"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	chars := []byte(s)
	fmt.Println(chars)

	restring := string(chars)
	fmt.Println(restring)

	sliced := strings.Split(s, " ")

	var daveIdx int
	var cantIdx int
	for idx, val := range sliced {
		switch val {
		case "dave":
			daveIdx = idx
		case "can't":
			cantIdx = idx
		}
	}

	fmt.Println(strings.Join(sliced[:daveIdx+1], " "))
	fmt.Println(strings.Join(sliced[daveIdx:cantIdx+1], " "))
	fmt.Println(strings.Join(sliced[cantIdx:], " "))

	for _, char := range chars {
		fmt.Println(string(char))
	}
}
