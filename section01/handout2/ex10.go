/*
Create a new type called “transportation”.
The underlying type of this new type is interface.
An interface defines functionality.
Said another way, an interface defines behavior.
For this interface, any other type that has a method with this signature …
	transportationDevice() string
… will automatically, implicitly implement this interface.
Create a func called “report” that takes a value of type “transportation” as an argument.
The func should print the string returned by “transportationDevice()”
  being called for any type that implements the “transportation” interface.
Call “report” passing in a value of type truck.
Call “report” passing in a value of type sedan.

https://play.golang.org/p/i0p-0TOxlA
 */

package main

import (
	"fmt"
	"reflect"
)

type transportation interface {
	transportationDevice() string
}

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

func (t truck) transportationDevice() string {
	return fmt.Sprintf("I am a %s %s, and I carry goods.", t.vehicle.color, reflect.TypeOf(t).Name())
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintf("I am a %s %s, and I carry people.", s.vehicle.color, reflect.TypeOf(s).Name())
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
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

	report(t1)
	report(s1)
}

