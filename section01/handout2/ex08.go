/*
Using the vehicle, truck, and sedan structs:
using a composite literal, create a value of type truck and assign values to the fields;
using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values.
Print out a single field from each of these values.

https://play.golang.org/p/-VBq3j_Vs4
 */

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck {
		vehicle {
			2,
			"blue",
		},
		true,
	}

	s1 := sedan {
		vehicle {
			4,
			"red",
		},
		false,
	}

	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.vehicle.color)
	fmt.Println(s1.vehicle.doors)
}
