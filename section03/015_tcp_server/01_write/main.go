/*
 * Run to start TCP server
 * Connect via telnet: $ telnet localhost 8080
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Problem accepting connection: ", err)
		}

		io.WriteString(conn, "\nHello from the TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope so!\n")

		conn.Close()
	}
}
