/*
 *
 * Run this server, and connect via browser:
 * $ go run main.go &
 * Chrome: http://localhost:8080
 * Ctrl+C
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// take a request
	receive(conn)

	// send a response
	respond(conn)
}

func receive(conn net.Conn) {
	var target, host string

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		fs := strings.Fields(ln)
		if i == 0 {
			target = fs[1]
		}
		if ln != "" && strings.HasPrefix(fs[0], "Host:") {
			host = fs[1]
		}
		if ln == "" {
			fmt.Printf("Full request URL: %s%s\r\n\r\n", host, target)
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
