/*
 *
 * Run this server, and telnet to it:
 * $ go run main.go &
 * $ telnet localhost 8080
 *
 *   Ctrl+]
 *   telnet> quit
 * $ fg
 *   Ctrl+C
 */
package main

import (
	"bufio"
	"fmt"
	"io"
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

	io.WriteString(conn,
		"\nIn-Memory Database\n\n"+
			"Usage:\n"+
			"SET key value\n"+
			"GET key1 [key2...]\n"+
			"DEL key\n"+
			"QUIT\n\n"+
			"Examples:\n"+
			"SET flavor chocolate\n"+
			"GET flavor\n\n\n"+
			"> ")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			for _, k := range fs[1:] {
				if v, ok := data[k]; ok {
					fmt.Fprintf(conn, "%s: %s\n", k, v)
				} else {
					fmt.Fprintf(conn, "Key %s not found.\n", k)
				}
			}
			fmt.Fprintf(conn, "> ")
		case "SET":
			k := fs[1]
			v := fs[2]
			data[k] = v
			fmt.Fprintf(conn, "> ")
		case "DEL":
			k := fs[1]
			delete(data, k)
			fmt.Fprintf(conn, "> ")
		case "QUIT":
			fmt.Fprintf(conn, "Adios.\n")
			conn.Close()
		default:
			fmt.Fprintf(conn, "Invalid command: %s\n", fs[0])
			fmt.Fprintf(conn, "> ")
			continue
		}
	}
}
