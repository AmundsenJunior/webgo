/*
 * execute this tcp client, to write to server
 * run the 02_read server: $ cd ../02_read; go run main.go
 * run this client: $ go run main.go
 */
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic("Couldn't make connection: ", err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
