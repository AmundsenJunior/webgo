/*
Give a method to both the “truck” and “sedan” types with the following signature
	transportationDevice() string
Have each func return a string saying what they do.
Create a value of type truck and populate the fields.
Create a value of type sedan and populate the fields.
Call the method for each value.

https://play.golang.org/p/iHmZg9ucE2
 */

package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println(t1.transportationDevice())
	fmt.Println(s1.transportationDevice())
}
