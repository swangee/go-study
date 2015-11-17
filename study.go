package main

import (
	"bufio"
	"fmt"
	"net"
)

var connections = 0

func handleConnection(conn net.Conn) {
	connections++
	fmt.Printf("New connection established. Connections number is %d\n", connections)

	number := connections

	defer func() {
		fmt.Printf("Connection #%d closed.\n", number)
		conn.Close()
		connections--
	}()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n') // output message received

		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Got bad message: " + string(message))
			}
			break
		}

		if len(message) > 0 {
			conn.Write([]byte("Message Received: " + string(message)))
		}
	}
}

func main() {
	fmt.Println("Launching server...") // listen on all interfaces

	ln, err := net.Listen("tcp", ":8081") // accept connection on port

	if err != nil {
		fmt.Println("Can't start server")
		return
	}

	defer ln.Close()

	for {
		// will listen for message to process ending in newline (\n)
		conn, err := ln.Accept() // run loop forever (or until ctrl-c)

		if err != nil {
			fmt.Println("Connection failed. Continue.")
			continue
		}

		go handleConnection(conn)
	}
}
