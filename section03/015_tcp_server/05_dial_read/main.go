/*
 * execute a tcp client, to read from server
 * run the 01_write server: $ cd ../01_write; go run main.go
 * run this client in a second terminal: $ go run main.go
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic("Error dialing connection: ", err)
	}
	defer conn.Close()

	response, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(response))
}
