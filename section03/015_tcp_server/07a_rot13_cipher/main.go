/*
 * rot13_cipher is a TCP server that
 * receives text via connection, then
 * sends the rot13 translation back.
 * Only rot13's alphabetic text.
 *
 * Run this server, and telnet to it:
 * $ go run main.go &
 * $ telnet localhost 8080
 *   hello
 *   HELLO - URYYB
 *   Ctrl+]
 *   telnet> quit
 * $ fg
 *   Ctrl+C
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToUpper(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n", ln, r)
	}
	defer conn.Close()
}

func rot13(bs []byte) []byte {
	r13 := make([]byte, len(bs))
	for i, v := range bs {
		// ascii 65 - 90
		if v >= 65 && v <= 77 {
			r13[i] = v + 13
		} else if v >= 78 && v <= 90 {
			r13[i] = v - 13
		} else {
			r13[i] = v
		}
	}

	return r13
}
