/*
Initialize a SLICE of int using a composite literal;
print out the slice;
range over the slice printing out just the index;
range over the slice printing out both the index and the value

https://play.golang.org/p/cvbhCv5npm
*/

package main

import "fmt"

func main() {
	slice := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

	fmt.Println(slice)

	for idx, _ := range slice {
		fmt.Println(idx)
	}

	for idx, val := range slice {
		fmt.Println(idx, val)
	}
}
