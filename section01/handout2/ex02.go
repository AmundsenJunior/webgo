/*
Initialize a MAP using a composite literal where the key is a string and the value is an int;
print out the map;
range over the map printing out just the key;
range over the map printing out both the key and the value

https://play.golang.org/p/wLynuAPIAb
 */

package main

import "fmt"

func main() {
	imdb := map[string]int {
		"terminator": 7,
		"terminator 2": 10,
		"terminator 3": 5,
	}

	fmt.Println(imdb)

	for key, _ := range imdb {
		fmt.Println(key)
	}

	for key, val := range imdb {
		fmt.Printf("I rate %s as a(n) %d.\n", key, val)
	}
}
