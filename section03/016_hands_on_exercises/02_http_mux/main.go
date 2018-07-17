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

type Request struct {
	Method string
	Target string
}

type Response struct {
	Body   string
	Code   string
	Reason string
}

const (
	index = `<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Index</title>
				</head>
				<body>
					<p><strong>Hello World</strong></p>
					<p><a href="/about">About</a></p>
					<p><a href="/contact">Contact</a></p>
					<form method="POST" action="/ping">
						<input type="submit" value="ping">
					</form>
				</body>
				</html>`
	about    = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>About</title></head><body><p>It's me.</p><p>I'm here.</p></body></html>`
	contact  = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Contact</title></head><body><a href="https://twitter.com/sergeography/">Get at me.</a></body></html>`
	pong     = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Pong</title></head><body><strong>PONG</strong></body></html>`
	notfound = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Not Found</title></head><body><p><strong>404</strong> Not found</p></body></html>'`
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
	request := receive(conn)

	// route the request
	mux := mux(request)

	// send a response
	respond(conn, mux)
}

func receive(conn net.Conn) Request {
	var request Request

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		fs := strings.Fields(ln)
		if i == 0 {
			request.Method = fs[0]
			request.Target = fs[1]
		}
		if ln == "" {
			break
		}
		i++
	}

	return request
}

func mux(request Request) Response {
	method := request.Method
	target := request.Target
	var response Response
	switch method {
	case "GET":
		switch target {
		case "/", "/index", "/index.html":
			response.Body = index
			response.Code = "200"
			response.Reason = "OK"
		case "/about", "/about.html":
			response.Body = about
			response.Code = "200"
			response.Reason = "OK"
		case "/contact", "/contact.html":
			response.Body = contact
			response.Code = "200"
			response.Reason = "OK"
		default:
			response.Body = notfound
			response.Code = "404"
			response.Reason = "Not Found"
		}
	case "POST":
		switch target {
		case "/ping":
			response.Body = pong
			response.Code = "200"
			response.Reason = "OK"
		default:
			response.Body = index
			response.Code = "200"
			response.Reason = "OK"
		}
	}

	return response
}

func respond(conn net.Conn, response Response) {
	fmt.Fprintf(conn, "HTTP/1.1 %s %s\r\n", response.Code, response.Reason)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(response.Body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, response.Body)
}

